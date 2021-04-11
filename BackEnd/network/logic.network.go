package network

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"strconv"
	"strings"
)

const (
	// See http://man7.org/linux/man-pages/man8/route.8.html
	file = "/proc/net/route"
)

type linuxRouteStruct struct {
	Destination string
	Gateway     string
}

func GetPersonalIp() (string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}

	for _, addr := range addrs {
		s := strings.Split(addr.String(), "/")
		if s[1] == "24" {
			return s[0], nil
		}
	}

	return "", errors.New("There is no valid IP")
}

func parseToLinuxRouteStruct(output []byte) (linuxRouteStruct, error) {

	const (
		sep              = "\t" // field separator
		destinationField = 1    // field containing hex destination address
		gatewayField     = 2    // field containing hex gateway address
	)
	scanner := bufio.NewScanner(bytes.NewReader(output))

	// Skip header line
	if !scanner.Scan() {
		return linuxRouteStruct{}, errors.New("Invalid linux route file")
	}

	for scanner.Scan() {
		row := scanner.Text()
		tokens := strings.Split(row, sep)
		if len(tokens) < 11 {
			return linuxRouteStruct{}, fmt.Errorf("invalid row '%s' in route file: doesn't have 11 fields", row)
		}

		// Cast hex destination address to int
		destinationHex := "0x" + tokens[destinationField]
		destination, err := strconv.ParseInt(destinationHex, 0, 64)
		if err != nil {
			return linuxRouteStruct{}, fmt.Errorf(
				"parsing destination field hex '%s' in row '%s': %w",
				destinationHex,
				row,
				err,
			)
		}

		// The default interface is the one that's 0
		if destination != 0 {
			continue
		}

		return linuxRouteStruct{
			Destination: tokens[1],
			Gateway:     tokens[2],
		}, nil
	}
	return linuxRouteStruct{}, errors.New("interface with default destination not found")
}

func parseLinuxGatewayIP(output []byte) (string, error) {

	parsedStruct, err := parseToLinuxRouteStruct(output)
	if err != nil {
		return "", err
	}

	destinationHex := "0x" + parsedStruct.Destination
	gatewayHex := "0x" + parsedStruct.Gateway

	// cast hex address to uint32
	d, err := strconv.ParseInt(gatewayHex, 0, 64)
	if err != nil {
		return "", fmt.Errorf(
			"parsing default interface address field hex '%s': %w",
			destinationHex,
			err,
		)
	}
	// make net.IP address from uint32
	ipd32 := make(net.IP, 4)
	binary.LittleEndian.PutUint32(ipd32, uint32(d))

	// format net.IP to dotted ipV4 string
	return net.IP(ipd32).String(), nil
}

func GetGatewayAddress() (string, error) {
	f, err := os.Open(file)
	if err != nil {
		return "", fmt.Errorf("Can't access %s", file)
	}
	defer f.Close()

	bytes, err := ioutil.ReadAll(f)
	if err != nil {
		return "", fmt.Errorf("Can't read %s", file)
	}
	return parseLinuxGatewayIP(bytes)
}

package controllers

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"serverHome/resources"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var ADMINUSER resources.UserResource

func createHash(pass string) string {
	hasher := md5.New()
	hasher.Write([]byte(pass))
	return hex.EncodeToString(hasher.Sum(nil))
}

func initAdminUser() {

	passwordMD5 := createHash("password")

	ADMINUSER = resources.UserResource{
		Id:       1,
		UserName: "cusadmin",
		Password: passwordMD5,
	}
}

func InitAdminUser() {
	initAdminUser()
}

func checkUser(user string, password string) (bool, error) {

	if user != ADMINUSER.UserName {
		return false, errors.New("UserName not found")
	}

	passwordMD5 := createHash(password)

	if passwordMD5 != ADMINUSER.Password {

		return false, errors.New("Password error")
	}

	return true, nil

}

func generateToken(userId int) (string, error) {
	var err error

	os.Setenv("ACCESS_SECRET", "jdnfksdmfksd")
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = userId
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return "", err
	}
	return token, nil

}

func Login(w http.ResponseWriter, r *http.Request) {

	setCORS(&w, r)

	body, err := ioutil.ReadAll(r.Body)

	fmt.Printf("%s\n", body)

	if err != nil {
		http.Error(w, "There is an error with body: "+err.Error(), http.StatusBadRequest)
		return
	}

	user := resources.UserResource{}

	err = json.Unmarshal(body, &user)

	if err != nil {
		http.Error(w, "There is an error: "+err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Printf("Username : => %s\nPassword: => %s", user.UserName, user.Password)

	isUser, err := checkUser(user.UserName, user.Password)

	if err != nil {
		http.Error(w, "There is an error: "+err.Error(), http.StatusNotFound)
		return
	}
	tokenSecret, err := generateToken(ADMINUSER.Id)

	if err != nil {
		println("Token couldn't be generated")
		return
	}

	loginRes := resources.LoginResource{
		Logged: isUser,
		Token:  tokenSecret,
	}

	json, err := json.Marshal(loginRes)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	sendResponse(w, json)

}

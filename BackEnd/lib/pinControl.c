
#include <errno.h>
#include <fcntl.h>
#include <stdio.h>
#include <stdlib.h>
#include <sys/stat.h>
#include <sys/types.h>
#include <unistd.h>
#include <string.h>
#include "../include/pinControl.h"


int unExportPin(char *pin)
{
    int fd = open("/sys/class/gpio/unexport", O_WRONLY);
    if (fd == -1) {
        perror("Unable to open /sys/class/gpio/unexport");
        exit(1);
    }
    if (write(fd, pin, strlen(pin)) != 2) {
        perror("Error writing to /sys/class/gpio/unexport");
        exit(1);
    }
    close(fd);
}

int pinMode(char *pin, bool mode)
{
    int fd = open("/sys/class/gpio/export", O_WRONLY);
    if (fd == -1)
    {
        perror("Unable to open /sys/class/gpio/export");
        exit(1);
    }
    if (write(fd, pin, strlen(pin)) != 2)
    {
        perror("Error writing to /sys/class/gpio/export");
        exit(1);
    }
    close(fd);
    char mPath[32];
    sprintf(mPath, "/sys/class/gpio/gpio%s/direction", pin);
    printf("%s \n", mPath);
    fd = open(mPath, O_WRONLY);
    if (fd == -1)
    {
        perror("Unable to open direction file");
        exit(1);
    }
    if (mode)
    {
        if (write(fd, "out", 3) != 3)
        {
            perror("Error writing to direction file");
            exit(1);
        }
        else
        {
            if (write(fd, "in", 2) != 3)
            {
                perror("Error writing to direction file");
                exit(1);
            }
        }
        close(fd);
        return 0;
    }
}

int digitalWrite(char *pin, char *value)
{
    char pathValue[28];
    sprintf(pathValue, "/sys/class/gpio/gpio%s/value", pin); 
    printf("%s \n", pathValue);
    int fd = open(pathValue, O_WRONLY);
    if (fd == -1)
    {
        perror("Unable to open value file");
        exit(1);
    }
    if (write(fd, value, 1) != 1) {
            perror("Error writing to /sys/class/gpio/gpio24/value");
            exit(1);
    }
    close(fd);

    //unExportPin(pin);
    return 0;
}

char digitalRead(char *pin)
{
    char pathValue[28];
    sprintf(pathValue, "/sys/class/gpio/gpio%s/value", pin); 
    printf("%s \n", pathValue);
    int fd = open(pathValue, O_RDONLY);
    if (fd == -1)
    {
        perror("Unable to open value file");
        exit(1);
    }
    char *value;
    if (read(fd, value, 1) != 1) {
            perror("Error writing to value file");
            exit(1);
    }
    close(fd);
    //unExportPin(pin);
    return *value;
}

int blinkFun(char *pin, int freq, int duration)
{
    pinMode(pin, true);
    char pathValue[28];
    sprintf(pathValue, "/sys/class/gpio/gpio%s/value", pin); 
    printf("%s \n", pathValue);
    int fd = open(pathValue, O_WRONLY);
    if (fd == -1)
    {
        perror("Unable to open value file");
        exit(1);
    }
    for (size_t i = 0; i < duration/freq; i++)
    {
        if (write(fd, "1", 1) != 1) {
            perror("Error writing to /sys/class/gpio/gpio24/value");
            exit(1);
        }
        sleep(1/(2*freq));
        if (write(fd, "0", 1) != 1) {
            perror("Error writing to /sys/class/gpio/gpio24/value");
            exit(1);
        }
        sleep(1/(2*freq));
    }
    close(fd);
    return 0;
}



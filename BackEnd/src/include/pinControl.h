#ifndef PINCONTROL_H_
#define PINCONTROL_H_

#include <errno.h>
#include <fcntl.h>
#include <stdio.h>
#include <stdlib.h>
#include <sys/stat.h>
#include <sys/types.h>
#include <unistd.h>
#include <string.h>
#include <stdbool.h>

int pinMode(char *pin, bool mode);
int digitalWrite(char *pin, char *value);
int digitalRead(char *pin);
int blinkFun(char *pin, int freq, int duration);
int unExportPin(char *pin);
char *getPhoto(long filelen);
long takePhoto();
int writePhoto(char *photo, int a);

#endif

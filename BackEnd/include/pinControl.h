

#ifndef PINCONTROL_H_  
#define PINCONTROL_H_

#include <stdbool.h>

int pinMode(char *pin, bool mode);
int digitalWrite(char *pin, char *value);
char digitalRead(char *pin);
int blinkFun(char *pin, int freq, int duration);

#endif

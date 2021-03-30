
#include "../include/pinControl.h"
#include <stdio.h>

int main()
{
    pinMode(2, true);
    digitalWrite(2, '1');
    pinMode(3, false);
    printf("%c \n", digitalRead(3));
}
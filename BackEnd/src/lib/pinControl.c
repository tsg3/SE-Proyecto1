#include "pinControl.h"

char *readPathValue;

long takePhoto()
{
    int status = system("fswebcam /dev/video0 photo.jpeg");

    FILE *fileptr;
    FILE *filedet;
    long filelen;
    char *buff;

    fileptr = fopen("photo.jpeg", "rb");
    fseek(fileptr, 0, SEEK_END);
    filelen = ftell(fileptr);
    rewind(fileptr);
    fclose(fileptr);

    return filelen;
}

char *getPhoto(long filelen)
{
    FILE *fileptr;
    FILE *filedet;
    char *buff;

    buff = (char *)malloc(filelen);

    fileptr = fopen("photo.jpeg", "rb");
    // Enough memory for the file
    fread(buff, filelen, 1, fileptr);
    fclose(fileptr);

    printf("IN C: ::::: : %s\n", buff);

    return buff;
}

int unExportPin(char *pin)
{
    int fd = open("/sys/class/gpio/unexport", O_WRONLY);
    if (fd == -1)
    {
        perror("Unable to open /sys/class/gpio/unexport");
        exit(1);
    }
    if (write(fd, pin, strlen(pin)) < 0)
    {
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
    if (write(fd, pin, strlen(pin)) <= 0)
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
    }
    else
    {
        if (write(fd, "in", 2) != 2)
        {
            perror("Error writing to direction file");
            exit(1);
        }
    }
    close(fd);
    return 0;
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
    if (write(fd, value, 1) < 0)
    {
        perror("Error writing to /sys/class/gpio/gpio24/value");
        exit(1);
    }
    close(fd);
    memset(pathValue, 0, 28);
    return 0;
}

int digitalRead(char *pin)
{
    readPathValue = (char *)malloc(28 * sizeof(char));
    memset(readPathValue, 0, 28);
    sprintf(readPathValue, "/sys/class/gpio/gpio%s/value", pin);
    printf("%s \n", readPathValue);
    int fd = open(readPathValue, O_RDONLY);
    if (fd == -1)
    {
        perror("Unable to open value file");
        exit(1);
    }
    int *value;

    int result = read(fd, value, 1);
    printf("Read result %d from %s and value %d\n", result, pin, *value);
    if (result < 0)
    {
        perror("Error reading to value file");
        exit(1);
    }
    close(fd);
    free(readPathValue);
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
    for (size_t i = 0; i < duration / freq; i++)
    {
        if (write(fd, "1", 1) < 0)
        {
            perror("Error writing to /sys/class/gpio/gpio24/value");
            exit(1);
        }
        sleep(1 / (2 * freq));
        if (write(fd, "0", 1) < 0)
        {
            perror("Error writing to /sys/class/gpio/gpio24/value");
            exit(1);
        }
        sleep(1 / (2 * freq));
    }
    close(fd);
    memset(pathValue, 0, 28);
    return 0;
}

int writePhoto(char *photo, int a)
{
    FILE *filedet;
    filedet = fopen("photo-test.jpeg", "w");
    fwrite(photo, 1, a, filedet);

    fclose(filedet);

    return 0;
}
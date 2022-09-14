#include <stdlib.h>
#include <unistd.h>
int main() {
    system("/bin/sh -c '/usr/bin/cat /opt/script/files/input.txt > /opt/script/files/output.txt'");
    sleep(60);
}

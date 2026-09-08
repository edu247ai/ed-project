#include <stdio.h>
#include <stdint.h>
#include "chat_server.h"

int main(void) {
    const char *host = "0.0.0.0";
    uint16_t port = 9000;

    printf("EDU247CHAT server starting on %s:%u\n", host, port);

    if (chat_server_start(host, port) != 0) {
        fprintf(stderr, "Failed to start chat server\n");
        return 1;
    }

    return 0;
}

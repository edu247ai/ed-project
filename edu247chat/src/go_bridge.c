#include "go_bridge.h"

#include <arpa/inet.h>
#include <netinet/in.h>
#include <sys/socket.h>
#include <unistd.h>
#include <string.h>
#include <stdio.h>
#include <stdint.h>

int go_bridge_request_ai(const char *host, uint16_t port,
                         const char *payload,
                         char *out, size_t out_size)
{
    // 1. Create TCP socket
    int fd = socket(AF_INET, SOCK_STREAM, 0);
    if (fd < 0) {
        perror("socket");
        return -1;
    }

    // 2. Prepare address
    struct sockaddr_in addr;
    memset(&addr, 0, sizeof(addr));
    addr.sin_family = AF_INET;
    addr.sin_port = htons(port);
    addr.sin_addr.s_addr = inet_addr(host);

    // 3. Connect to Go bridge
    if (connect(fd, (struct sockaddr *)&addr, sizeof(addr)) < 0) {
        perror("connect");
        close(fd);
        return -1;
    }

    // 4. Build request frame
    uint8_t header[5];
    header[0] = 0x03; // AI request opcode

    uint32_t n = htonl((uint32_t)strlen(payload));
    memcpy(&header[1], &n, 4);

    // 5. Send header
    if (write(fd, header, 5) != 5) {
        perror("write header");
        close(fd);
        return -1;
    }

    // 6. Send payload
    if (write(fd, payload, strlen(payload)) != (ssize_t)strlen(payload)) {
        perror("write payload");
        close(fd);
        return -1;
    }

    // 7. Read reply header
    uint8_t rheader[5];
    ssize_t r = read(fd, rheader, 5);
    if (r != 5) {
        perror("read reply header");
        close(fd);
        return -1;
    }

    uint8_t rop = rheader[0];
    uint32_t rlen;
    memcpy(&rlen, &rheader[1], 4);
    rlen = ntohl(rlen);

    if (rop != 0x03) {
        fprintf(stderr, "unexpected opcode %u\n", rop);
        close(fd);
        return -1;
    }

    // 8. Read reply payload
    if (rlen >= out_size)
        rlen = out_size - 1;

    ssize_t got = read(fd, out, rlen);
    if (got <= 0) {
        perror("read reply payload");
        close(fd);
        return -1;
    }

    out[got] = '\0';

    // 9. Close connection
    close(fd);
    return 0;
}

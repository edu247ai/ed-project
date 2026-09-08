#include "client.h"
#include "protocol.h"
#include "go_bridge.h"

#include <stdio.h>
#include <string.h>
#include <unistd.h>

void client_handle(client_t *cli) {
    char buf[MSG_MAX_LEN];
    ssize_t n = read(cli->fd, buf, sizeof(buf));
    if (n <= 0) return;

    char msg[MSG_MAX_LEN];
    if (protocol_parse(buf, (size_t)n, msg, sizeof(msg)) != 0) {
        const char *err = "Invalid message\n";
        write(cli->fd, err, strlen(err));
        return;
    }

    printf("Received: %s\n", msg);

    if (strncmp(msg, "/ai ", 4) == 0) {
        const char *payload = msg + 4; // strip "/ai "
        char ai_reply[MSG_MAX_LEN];

        if (go_bridge_request_ai("127.0.0.1", 9100, payload,
                                 ai_reply, sizeof(ai_reply)) == 0) {
            write(cli->fd, ai_reply, strlen(ai_reply));
            write(cli->fd, "\n", 1);
        } else {
            const char *err = "[AI error]\n";
            write(cli->fd, err, strlen(err));
        }
    } else {
        write(cli->fd, msg, strlen(msg));
    }
}

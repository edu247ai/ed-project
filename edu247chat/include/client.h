#ifndef CLIENT_H
#define CLIENT_H

#include <netinet/in.h>

typedef struct {
    int fd;
    struct sockaddr_in addr;
} client_t;

void client_handle(client_t *cli);

#endif

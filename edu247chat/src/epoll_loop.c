#define _GNU_SOURCE
#include "epoll_loop.h"
#include "client.h"

#include <sys/epoll.h>
#include <unistd.h>
#include <stdio.h>
#include <stdlib.h>
#include <sys/socket.h>
#include <netinet/in.h>


#define MAX_EVENTS 64

int epoll_loop(int server_fd) {
    int epfd = epoll_create1(0);
    if (epfd < 0) {
        perror("epoll_create1");
        return -1;
    }

    struct epoll_event ev = {0};
    ev.events = EPOLLIN;
    ev.data.fd = server_fd;

    epoll_ctl(epfd, EPOLL_CTL_ADD, server_fd, &ev);

    struct epoll_event events[MAX_EVENTS];

    while (1) {
        int n = epoll_wait(epfd, events, MAX_EVENTS, -1);

        for (int i = 0; i < n; i++) {
            if (events[i].data.fd == server_fd) {
                client_t cli;
                socklen_t len = sizeof(cli.addr);
                cli.fd = accept4(server_fd, (struct sockaddr *)&cli.addr, &len, SOCK_NONBLOCK);

                if (cli.fd >= 0) {
                    struct epoll_event cev = {0};
                    cev.events = EPOLLIN;
                    cev.data.fd = cli.fd;
                    epoll_ctl(epfd, EPOLL_CTL_ADD, cli.fd, &cev);
                }
            } else {
                client_t cli = { .fd = events[i].data.fd };
                client_handle(&cli);
            }
        }
    }

    close(epfd);
    return 0;
}

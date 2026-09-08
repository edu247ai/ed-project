#ifndef GO_BRIDGE_H
#define GO_BRIDGE_H

#include <stddef.h>
#include <stdint.h>

// Send framed AI request and receive framed AI reply
int go_bridge_request_ai(const char *host, uint16_t port,
                         const char *payload,
                         char *out, size_t out_size);

#endif

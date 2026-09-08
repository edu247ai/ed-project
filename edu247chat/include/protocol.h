#ifndef PROTOCOL_H
#define PROTOCOL_H

#include <stddef.h>

#define MSG_MAX_LEN 2048

#define OP_JOIN_ROOM  0x01
#define OP_SEND_MSG   0x02
#define OP_AI_REQUEST 0x03

int protocol_parse(const char *buf, size_t len, char *out, size_t out_size);

#endif

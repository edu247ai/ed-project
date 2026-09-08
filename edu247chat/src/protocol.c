#include "protocol.h"
#include <string.h>

int protocol_parse(const char *buf, size_t len, char *out, size_t out_size) {
    if (len == 0 || out_size == 0) {
        return -1;
    }

    size_t copy_len = len < out_size - 1 ? len : out_size - 1;
    memcpy(out, buf, copy_len);
    out[copy_len] = '\0';

    return 0;
}

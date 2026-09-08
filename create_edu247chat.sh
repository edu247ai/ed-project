#!/bin/bash

BASE="edu247/edu247chat"

echo "Creating EDU247CHAT C project..."

# Create folders
mkdir -p $BASE/include
mkdir -p $BASE/src
mkdir -p $BASE/build

##############################################
# Makefile
##############################################
cat << 'EOF' > $BASE/Makefile
CC = gcc
CFLAGS = -Wall -Wextra -O2 -std=c11
INCLUDES = -Iinclude
SRC = src/main.c src/chat_server.c src/client.c src/protocol.c src/epoll_loop.c
OBJ = $(SRC:.c=.o)
BIN = build/edu247chat

.PHONY: all run clean

all: $(BIN)

$(BIN): $(OBJ)
    mkdir -p build
    $(CC) $(CFLAGS) $(INCLUDES) -o $(BIN) $(OBJ)

%.o: %.c
    $(CC) $(CFLAGS) $(INCLUDES) -c $< -o $@

run: all
    ./$(BIN)

clean:
    rm -rf $(OBJ) $(BIN)
EOF

##############################################
# README.md
##############################################
cat << 'EOF' > $BASE/README.md
# EDU247CHAT (C)

Ultra-fast chat transport layer for EDU247.

Features:
- TCP server
- epoll event loop
- zero-copy buffers
- binary protocol
- room/channel fan-out
- ready for Go + Python AI integration

## Build & Run

```bash
make run

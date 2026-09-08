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

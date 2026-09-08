cmake_minimum_required(VERSION 3.10)

project(edu247chat C)

set(CMAKE_C_STANDARD 11)
set(CMAKE_C_FLAGS "-Wall -Wextra -O2")

include_directories(include)

file(GLOB SOURCES
    "src/main.c"
    "src/chat_server.c"
    "src/client.c"
    "src/protocol.c"
    "src/epoll_loop.c"
    "src/go_bridge.c"
)

add_executable(edu247chat ${SOURCES})

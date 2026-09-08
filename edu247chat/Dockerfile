FROM debian:stable-slim

RUN apt-get update && apt-get install -y \
    build-essential \
    cmake \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /app

# Copy ONLY this folder
COPY . /app

# Clean any host build folder
RUN rm -rf /app/build

# Create fresh build directory
RUN mkdir /app/build

WORKDIR /app/build
RUN cmake .. && make -j4

EXPOSE 9000
CMD ["./edu247chat"]

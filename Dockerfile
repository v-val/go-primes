FROM golang:alpine
LABEL authors="v"
ARG ws_port=18700
ENV PRIMES_WS_PORT ${ws_port}
RUN mkdir /opt/primes
WORKDIR /opt/primes
COPY . .
RUN go build -ldflags "-s -w" -o /usr/local/bin/primes-ws
WORKDIR /run
ENTRYPOINT /usr/local/bin/primes-ws -D -p${PRIMES_WS_PORT}

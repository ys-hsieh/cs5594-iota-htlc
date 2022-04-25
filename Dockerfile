ARG GOLANG_IMAGE_TAG=1.17-buster


# Build stage
FROM golang:${GOLANG_IMAGE_TAG}
ARG BUILD_TAGS="rocksdb,builtin_static"
ARG BUILD_LD_FLAGS=""
ARG BUILD_TARGET="/wasp"
WORKDIR /wasp

# Make sure that modules only get pulled when the module file has changed
COPY go.mod go.sum ./
RUN go mod download
RUN go mod verify

# Project build stage

COPY . .
RUN make build

EXPOSE 7000/tcp
EXPOSE 9090/tcp
EXPOSE 5550/tcp
EXPOSE 4000/udp

COPY docker_config.json /etc/wasp_config.json

RUN wget https://github.com/tinygo-org/tinygo/releases/download/v0.22.0/tinygo_0.22.0_amd64.deb
RUN dpkg -i tinygo_0.22.0_amd64.deb

RUN mkdir smart-contracts
RUN apt upgrade && apt update
RUN apt install -y vim

CMD ["./wasp", "-c", "/etc/wasp_config.json"]
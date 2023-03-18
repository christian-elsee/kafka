#################################################
FROM docker.io/golang:1.19-bullseye as build

LABEL stage=build
WORKDIR /opt/main

COPY go.mod go.sum ./
RUN  go mod download && go mod verify

COPY . ./
RUN go build -v -o ./target main.go

#################################################
FROM debian:bullseye-slim
WORKDIR /opt/main

COPY --from=build /opt/main/target /usr/local/bin/kafka
COPY --from=build /opt/main/sh     /opt/main/

RUN apt update && \
  apt install -y --no-install-recommends \
    netcat \
    dnsutils

ENTRYPOINT [ "/usr/local/bin/kafka" ]
CMD [ "-h" ]

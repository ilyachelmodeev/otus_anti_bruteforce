FROM golang:1.22 as build

ENV BIN_FILE /opt/anti-bruteforce/anti-bruteforce
ENV CODE_DIR /go/src/

WORKDIR ${CODE_DIR}

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . ${CODE_DIR}

ARG LDFLAGS
RUN CGO_ENABLED=0 go build \
        -ldflags "$LDFLAGS" \
        -o ${BIN_FILE} cmd/anti-bruteforce/*

FROM alpine:3.9

ENV BIN_FILE "/opt/anti-bruteforce/anti-bruteforce"
COPY --from=build ${BIN_FILE} ${BIN_FILE}

CMD ${BIN_FILE}
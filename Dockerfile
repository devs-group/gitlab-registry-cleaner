FROM golang:alpine AS base

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /app

COPY go.mod .
RUN go mod download

COPY . .

FROM golang:alpine as build

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /build

RUN apk add -U --no-cache ca-certificates

COPY --from=base /app /build

RUN go build -o main .

FROM scratch

COPY --from=build /build/main /
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

ENTRYPOINT ["/main"]

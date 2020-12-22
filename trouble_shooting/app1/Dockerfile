FROM golang:1.14.0-alpine as builder
RUN apk add alpine-sdk
WORKDIR /go/src
COPY . .
RUN go install ./cmd

FROM alpine:3.11
RUN apk add --no-cache tzdata ca-certificates
COPY --from=builder /go/bin/cmd /web
CMD ["/web"]

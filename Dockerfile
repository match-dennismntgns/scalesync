FROM golang:1.16.5 as builder
MAINTAINER dennis@bluecanary.be

WORKDIR /app/
COPY . /app/

RUN set -x && \

  go get -d -v . && \
  CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ScaleSync

FROM alpine:latest
WORKDIR /app/

COPY --from=builder /app/ /app/
CMD ./ScaleSync
FROM golang:1.16-alpine

# installs GCC, libc-dev, etc
RUN apk add build-base

# makes working with alpine-linux a little easier
RUN apk add --no-cache shadow

# Create a non-privileged user for running the go app
RUN groupadd -r dockeruser && useradd -r -g dockeruser dockeruser

WORKDIR /home/dockeruser

ADD . .

RUN go test -v
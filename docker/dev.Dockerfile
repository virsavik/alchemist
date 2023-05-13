# syntax=docker/dockerfile:1

FROM golang:1.20.3-alpine3.17

WORKDIR /app

RUN apk add git

# Copy go mod
COPY go.mod .
COPY go.sum .

RUN go mod download && go mod verify

# Install sqlboiler
# RUN go install github.com/virsavik/sqlboiler/v4@latest
RUN go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-psql@latest

# Use custom template
RUN cd /opt && \
  git clone https://github.com/virsavik/sqlboiler && \
  cd sqlboiler &&  \
  go build -o sqlboiler && \
  cp sqlboiler $GOPATH/bin



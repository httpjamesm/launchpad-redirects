FROM golang:1.20-alpine

RUN apk add musl-dev
RUN apk add libc-dev
RUN apk add gcc

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY ./ /app

RUN go build -o /launchpad

EXPOSE 8080

CMD ["/anonymousoverflow"]
# BACKEND-CONTAINER
FROM golang

WORKDIR /backend

COPY ./ .

RUN go mod tidy
RUN go build -o ./backend
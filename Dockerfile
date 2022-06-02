FROM golang:1.17-alpine

ARG USER_ID=1001
ARG PROJECT_DIR=/home/app

ENV USER=mobydick
ENV GROUP=app_users
ENV GROUP_ID=1000
ENV HOME $PROJECT_DIR

WORKDIR $HOME

COPY go.mod go.sum main.go $HOME
RUN go mod download
RUN go build -o app-server

EXPOSE $PORT

CMD go run .
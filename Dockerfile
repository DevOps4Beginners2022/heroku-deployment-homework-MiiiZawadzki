FROM golang:1.17-alpine

ARG USER_ID=1001
ARG PROJECT_DIR=/home/app

ENV USER=mobydick
ENV GROUP=app_users
ENV GROUP_ID=1000
ENV HOME $PROJECT_DIR
ENV PORT=8000
WORKDIR $HOME

COPY go.mod .
COPY go.sum .
COPY main.go .

RUN go mod download
RUN go build -o app-server

EXPOSE $PORT

CMD go run main.go
FROM golang:1.17-alpine

ARG PROJECT_DIR=/home/app
ARG PORT=8090

ENV HOME $PROJECT_DIR
ENV PORT $PORT
WORKDIR $HOME

COPY go.mod .
COPY go.sum .
COPY main.go .

RUN go mod download
RUN go build -o app-server

EXPOSE $PORT

CMD go run main.go
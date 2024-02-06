FROM golang:1.21-alpine
WORKDIR /var/bench

COPY src/ src/
COPY Taskfile.yaml .

RUN go install github.com/go-task/task/cmd/task@latest

CMD task bench
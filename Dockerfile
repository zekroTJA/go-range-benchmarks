FROM golang:alpine
WORKDIR /var/bench

COPY src/ src/
COPY Taskfile.yaml .

RUN go install github.com/go-task/task/cmd/task@latest

CMD task bench
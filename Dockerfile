FROM golang:1.21-alpine
WORKDIR /var/bench

COPY src/ src/
COPY Taskfile.yaml .

RUN go install github.com/go-task/task/v3/cmd/task@latest

RUN apk add curl gcc musl-dev
RUN curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs | ash -s - -y

ENV PATH="/root/.cargo/bin:${PATH}"
RUN rustup default nightly
RUN cargo install --git https://github.com/zekroTJA/testanalyzer

CMD task bench
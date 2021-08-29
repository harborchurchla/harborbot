FROM golang:1.16.4 as builder
WORKDIR /opt/app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY cmd ./cmd
#COPY internal ./internal
RUN CGO_ENABLED=0 go build -o /bin/harborbot ./cmd/harborbot.go

FROM target/flottbot
COPY --from=builder /bin/harborbot ./harborbot
COPY .config ./config

ENTRYPOINT ./harborbot

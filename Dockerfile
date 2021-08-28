FROM golang:1.16.4 as builder
WORKDIR /opt/app
RUN go get -u github.com/target/flottbot/cmd/flottbot

FROM python:3.7.2-slim
WORKDIR /opt/app

# Need ca-certificates to make https requests from container
RUN apt-get update
RUN apt-get install -y ca-certificates

COPY --from=builder go/bin/flottbot ./flottbot
COPY .config ./config

RUN pip install -r ./config/scripts/requirements.txt

ENTRYPOINT ./flottbot

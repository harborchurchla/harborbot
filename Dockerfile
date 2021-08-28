#FROM golang:1.16.4 as builder
#WORKDIR /opt/app
#RUN go get -u github.com/target/flottbot/cmd/flottbot
#
#FROM python:3.7.2-slim
#WORKDIR /opt/app
#
## Need ca-certificates to make https requests from container
#RUN apt-get update
#RUN apt-get install -y ca-certificates
#ENV USERNAME=flottbot
#ENV GROUP=flottbot
#ENV UID=900
#ENV GID=900
#RUN addgroup -gid "$GID" -S "$GROUP" && adduser -S -u "$UID" -G "$GROUP" "$USERNAME"
#
#COPY --from=builder go/bin/flottbot ./flottbot
#COPY .config ./config
#
#RUN pip install -r ./config/scripts/requirements.txt
#
#USER ${USERNAME}
#ENTRYPOINT ./flottbot

FROM golang:1.16.4 as builder
COPY main.go ./main.go
RUN go build -o /bin/main ./main.go

FROM target/flottbot:python
COPY --from=builder /bin/main ./main
COPY .config ./config
RUN pip install -r ./config/scripts/requirements.txt

CMD ["/main"]
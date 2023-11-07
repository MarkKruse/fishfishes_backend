FROM alpine:latest

COPY ./target/ff-linux-amd64 /
COPY ./.env /

ENTRYPOINT ./ff-linux-amd64 \

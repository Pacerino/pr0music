#build stage
FROM golang:bullseye AS builder
RUN apt update
RUN apt install build-essential -y
WORKDIR /go/src/app
COPY . .
RUN go get -d -v ./...
RUN go build -o /go/bin/app/ -v ./...

#final stage
FROM debian:bullseye-slim
ENV TZ="Europe/Berlin"
RUN apt update
RUN apt install ffmpeg -y
COPY --from=builder /go/bin/app/pr0music /app/pr0music
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /go/src/app/acrcloud/libacrcloud_extr_tool.so /usr/lib/libacrcloud_extr_tool.so
RUN chmod 755 /usr/lib/libacrcloud_extr_tool.so
ENTRYPOINT /app/pr0music
LABEL Name=pr0music Version=0.0.1
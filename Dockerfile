#build stage
FROM golang:alpine AS builder
RUN apk add --no-cache git
WORKDIR /go/src/app
COPY . .
RUN go get -d -v ./...
RUN go build -o /go/bin/app/ -v ./...

#final stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates
RUN apk add --no-cache ffmpeg
COPY --from=builder /go/bin/app/pr0music /app/pr0music
ENTRYPOINT /app/pr0music
LABEL Name=pr0music Version=0.0.1
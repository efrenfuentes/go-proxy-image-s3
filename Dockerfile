FROM golang:1.17-alpine as builder
RUN apk update && apk add --no-cache ca-certificates && rm -rf /var/cache/apk/*
RUN update-ca-certificates
RUN mkdir /app
ADD . /app/
WORKDIR /app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o go_proxy_image_s3 .
EXPOSE 8080
CMD ["./go_proxy_image_s3", "web"]

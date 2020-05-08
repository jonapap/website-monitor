FROM golang:alpine AS builder
WORKDIR /build
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN ./build.sh

FROM alpine:latest
ENV InsideDockerContainer=True
RUN apk add ca-certificates && rm -rf /var/cache/apk/*
COPY check_websites_cron /etc/periodic/daily/ 
RUN  chmod +x /etc/periodic/daily/check_websites_cron
WORKDIR /dist
COPY --from=builder /build/bin/* ./
CMD ["crond", "-f"]

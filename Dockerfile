FROM golang:1.8.1-alpine
WORKDIR /root/
COPY healthcheck /root/healthcheck
ENTRYPOINT ["./healthcheck"]

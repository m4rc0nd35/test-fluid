FROM golang:1.18-alpine3.15 AS builder
WORKDIR /app
COPY . .
# This will download all certificates (ca-certificates) and builds it in a
# single file under /etc/ssl/certs/ca-certificates.crt (update-ca-certificates)
# I also add git so that we can download with `go mod download` and
# tzdata to configure timezone in final image
RUN apk --update add --no-cache ca-certificates openssl git tzdata && \
update-ca-certificates
RUN GO111MODULE="on" CGO_ENABLED=0 GOOS=linux go build -o app -ldflags="-w -s" .

FROM scratch
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
# This line will copy all certificates to final image
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /app/app /app 
ENTRYPOINT ["/app"]
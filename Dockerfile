FROM golang:1.22.1-alpine as builder

WORKDIR /app
RUN apk add --no-cache make upx
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN make build && upx -9 -k /app/bin/satontdev


FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/bin/satontdev /app/satontdev
ENTRYPOINT ["/app/satontdev"]
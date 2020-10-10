FROM golang:1.15.2-alpine3.12 as builder

RUN apk add git curl

WORKDIR /go/src/app
COPY ./go.* ./*.go ./
RUN go build -o /app main.go

FROM alpine:3.12
COPY --from=builder /app .
ENTRYPOINT ["/app"]
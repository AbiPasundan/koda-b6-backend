FROM golang:1.25.0-alpine

WORKDIR /workspace

COPY *go* .

RUN go tidy

RUN go build -o test main.go

ENTRYPOINT ["/workspace/test"]

FROM alpine:latest

WORKDIR /app

COPY --from=build /workspace/test /workspace/test /app/

ENTRYPOINT ["/app/test"]


# kapai ./
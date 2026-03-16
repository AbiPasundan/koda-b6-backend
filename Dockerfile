FROM golang:1.25.0-alpine AS build

WORKDIR /workspace

COPY . .

RUN go mod tidy

RUN go build -o test ./cmd/main.go

ENTRYPOINT ["/workspace/test"]

FROM golang:1.25.0-alpine

WORKDIR /app

COPY --from=build /workspace/test /workspace/test /app/

ENTRYPOINT ["/app/test"]

# docker run --rm --network=server-name migrate/migrate:latest -source github:// link/ghcr/project/migration -database postgresql://postgres:1@localhost:5432?sslmode=disable down 1

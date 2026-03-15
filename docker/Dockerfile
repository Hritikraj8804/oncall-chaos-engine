# -------- BUILD STAGE --------

FROM golang:1.23.4-alpine AS builder

WORKDIR /app

COPY app/go.mod app/go.sum ./

RUN go mod download

COPY app/ .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o chaos-engine

# -------- RUNTIME STAGE --------

# -------- RUNTIME STAGE --------
FROM gcr.io/distroless/base-debian12

WORKDIR /app

COPY --from=builder /app/chaos-engine .
COPY --from=builder /app/data ./data

EXPOSE 5000

USER nonroot:nonroot

CMD ["./chaos-engine"]
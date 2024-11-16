# Stage 1: Build the Go application
FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN ls -la
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main


FROM scratch
LABEL org.opencontainers.image.source="https://github.com/mari0theminer/Hackathon_2024_HN_TEAM"
EXPOSE 3000
ENV SECRET =""

WORKDIR /app

COPY --from=builder /app/main .

ENTRYPOINT ["./main"]
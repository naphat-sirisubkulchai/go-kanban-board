# syntax=docker/dockerfile:1

# Build stage
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Install dependencies for Go and build tools
RUN apk add --no-cache gcc musl-dev

# Copy go mod files and download dependencies
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# Copy all source files
COPY . ./

# Build the Go application to a specific path
RUN go build -o /app/kanban cmd/main.go

# ตรวจสอบว่ามีไฟล์ /app/kanban ใน container หรือไม่
RUN ls -l /app/kanban

# Final stage with alpine base image
FROM frolvlad/alpine-glibc:alpine-3.15

WORKDIR /

# Copy the built Go application from the builder stage
COPY --from=builder /app/kanban .

# ให้สิทธิ์ execute กับไฟล์
RUN chmod +x /kanban

# Expose port 3000
EXPOSE 3000

# Run the application
ENTRYPOINT ["/kanban"]

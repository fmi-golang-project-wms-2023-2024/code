# Stage 1: Build protobufs
FROM bufbuild/buf AS buf_stage

# Set the working directory to /app
WORKDIR /app

# Copy the entire backend directory to the working directory
COPY ./backend .

# Run buf build and buf generate to generate protobufs
RUN buf build
RUN buf generate
RUN buf format -w
RUN buf lint

# Stage 2: Build WMS binary
FROM golang:1.21 AS go_build_stage

# Set the working directory to /app
WORKDIR /app

# Copy the generated protobuf files from the buf_stage
COPY --from=buf_stage /app ./

# Download dependencies
RUN go mod download

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /main

# Stage 3: Create a lightweight final image
FROM alpine:latest

# Set the working directory to /app
WORKDIR /app

# Copy the binary from the go_build_stage
COPY --from=go_build_stage /main .

# Expose ports 8080 and 8090 (gRPC & gateway)
EXPOSE 8080 8090

# Run the binary
CMD ["./main"]
# syntax=docker/dockerfile:1

FROM golang:1.16-alpine

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod .
RUN go mod download

# Copy the source code
COPY *.go ./

# Build
RUN go build -o /docker-cmentarz-app

# Run
CMD [ "/docker-cmentarz-app" ]

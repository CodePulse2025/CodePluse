# Use Golang Alpine base image
FROM golang:1.22-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod ./
COPY go.sum ./

# Copy the vendor folder containing the dependencies
COPY vendor/ ./vendor/
# ENV GOPROXY=direct

# RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the application
RUN go build -o main ./server

# Expose port 8080
EXPOSE 8080

# Run the application
CMD ["./main"]

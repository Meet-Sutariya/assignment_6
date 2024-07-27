# Use the official Go image as a base
FROM golang:1.18-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the Go application code into the container
COPY . .

# Build the Go application
RUN go build -o main .

# Start the application
CMD ["./main"]

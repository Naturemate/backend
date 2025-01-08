# Use the official Go image
FROM golang:1.20

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files first, for dependency caching
COPY go.mod ./
RUN go mod download

# Copy the rest of the application code
COPY . .

# Clean up unused dependencies
RUN go mod tidy

# Build the Go application
RUN go build -o /godocker

# Expose the port your application will run on
EXPOSE 8080

# Command to run the application
CMD ["/godocker"]

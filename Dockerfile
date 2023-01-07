# Use a lightweight base image
FROM golang:alpine

# Set the working directory
WORKDIR /app

# Copy the go.mod and go.sum files
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the application
RUN go build -o URLshortener

# Expose the application on port 8080
EXPOSE 8080

# Run the built binary
CMD ["./URLshortener"]


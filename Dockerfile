# The base go-image
FROM golang:1.14-alpine
 
# Create a directory for the app
RUN mkdir /app

# Copy all files from the current directory to the app directory
COPY . /app

# Set working directory
WORKDIR /app

# Download all dependencies. Dependencies will be cached if the go.mod and the go.sum files are not changed
COPY go.mod go.sum ./
RUN go mod download
COPY . .
 
# Run command as described:
# go build will build an executable file named server in the current directory
RUN go build -o server . 
 
# Run the server executable
CMD [ "./server" ]
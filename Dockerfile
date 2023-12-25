FROM golang:1.21.4 
 
# Creates an app directory to hold your app’s source code
WORKDIR /app
 
# Copies everything from your root directory into /app
COPY . .
 
# Installs Go dependencies
RUN go mod download
 
# Builds your app with optional configuration
RUN go build main.go
 
# Tells Docker which network port your container listens on
EXPOSE 3000
 
# Specifies the executable command that runs when the container starts
CMD [ "./main" ]

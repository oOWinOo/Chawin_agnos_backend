FROM golang:1.21.4 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main .

# #nginx
# FROM nginx:latest

# COPY ./index.html /usr/share/nginx/html/index.html

# EXPOSE 8080

# CMD ["nginx", "-g", "daemon off;"]
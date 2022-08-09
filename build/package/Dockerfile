# from alpine image
FROM golang:1.18-alpine

# in app directory
WORKDIR /../../

# copy go mod and go sum
COPY go.mod ./
COPY go.sum ./

# running go mod download
RUN go mod download

# copy all files
COPY * ./

# building our application
RUN go build -o ./internal/cmd/darolfonoon/main

# expose on port 8080
EXPOSE 8080

# executing the .sh file
CMD [ "/main" ]


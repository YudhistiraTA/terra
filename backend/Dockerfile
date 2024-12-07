# syntax=docker/dockerfile:1

FROM golang:1.23-alpine

RUN apk add --no-cache bash

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code. Note the slash at the end, as explained in
# https://docs.docker.com/engine/reference/builder/#copy
COPY . .
RUN chmod +x ./wait-for-it.sh

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /terra-docker ./app/main.go

# To bind to a TCP port, runtime parameters must be supplied to the docker command.
# But we can (optionally) document in the Dockerfile what ports
# the application is going to listen on by default.
# https://docs.docker.com/engine/reference/builder/#expose
EXPOSE 8000

# Run
# CMD [ "/terra-docker" ]
CMD ["./wait-for-it.sh", "postgres:5432", "--", "/terra-docker"]

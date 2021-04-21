#FROM golang:1.12-alpine
#FROM golang:1.12 AS builder

FROM golang:alpine AS builder

LABEL stage=builder
# RUN apk add --no-cache gcc libc-dev tzdata
RUN apk add --no-cache gcc tcptraceroute

# Set the Current Working Directory inside the container
WORKDIR /apigoevermos/

# We want to populate the module cache based on the go.{mod,sum} files.
COPY go.mod go.sum ./
# COPY go.sum .

RUN go mod download

COPY . .

COPY src/apievermos/main.go .

USER 0:0
# Build the Go app
RUN go build -o src/apievermos/main .
RUN chmod -R 777 src/

#second stage

#FROM golang:1.12
FROM alpine AS final

COPY --from=builder /apigoevermos .

# Run the binary program produced by `go install`
CMD ["./src/apievermos/main"]
# CMD ./src/apibafgate/main
EXPOSE 2021

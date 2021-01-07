#https://hub.docker.com/r/supinf/go-swagger/dockerfile

FROM golang:1.14.3-alpine AS build_base

RUN apk add --no-cache git

# Set the Current Working Directory inside the container
WORKDIR /app

# We want to populate the module cache based on the go.{mod,sum} files.
COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

# # Build the Go app
#RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o ./out . #/app/out
RUN go build -o /app/out

WORKDIR /app
RUN go build -o /app/out/apic

# # Start fresh from a smaller image
FROM alpine:3.12.3

WORKDIR /

COPY --from=build_base /app/out /

# Run the binary program produced by `go install`
CMD ["/app"]
# Build stage
FROM golang:1.12.7-alpine3.10 AS build
RUN apk add --no-cache git
WORKDIR /app/
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o beagle main.go

# Final stage
FROM alpine:3.10.1
LABEL maintainer="danielkvist@protonmail.com"
RUN apk add --no-cache ca-certificates
COPY --from=build /app/beagle /app/beagle
COPY --from=build /app/urls.csv /app/urls.csv
WORKDIR /app/
ENTRYPOINT ["./beagle"]
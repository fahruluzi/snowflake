FROM golang:1.20-alpine as builder

RUN apk update && apk add --no-cache git && apk --no-cache add tzdata

WORKDIR /build
COPY ./go.mod ./go.sum ./

COPY . .

RUN go mod download

# Build Go App
RUN CGO_ENABLED=0 GOOS=linux go build -o main

FROM scratch

WORKDIR /app

COPY --from=builder /build/main ./main
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo

# Expose port
EXPOSE 8080

# Command to run the executeable
ENTRYPOINT ["./main"]


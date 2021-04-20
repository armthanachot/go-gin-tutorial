#build stage
FROM golang:alpine AS builder
RUN apk add --no-cache git
WORKDIR /go/src/app
COPY . .
RUN go get -d -v ./...
RUN go build -o /go/bin/app -v ./...

#final stage
FROM alpine:latest
RUN go mod init GO_API
RUN go mod tidy 
RUN go mod vendor
RUN go get -u github.com/gin-gonic/gin
RUN apk --no-cache add ca-certificates
COPY --from=builder /go/bin/app /app
ENTRYPOINT /app
LABEL Name=goapi Version=0.0.1
EXPOSE 8080

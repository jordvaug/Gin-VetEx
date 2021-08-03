#Using Multistage build to create smallest container possible
FROM golang:latest AS builder

#docker run -d --env-file .env -p 8080:8080/tcp --name gin gin-vetex 
#docker build --tag gin-vetex -q .

#Copy all files and folders not enumerated in .dockerignore
COPY . /app
WORKDIR /app

#Install dependencies
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -a -o /main .

#Get latest certs 
FROM alpine:latest as certs
RUN apk --update add ca-certificates

#scratch is an empty docker image for use with languages capable of producing statically linked binaries
FROM scratch
COPY --from=certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /main ./
ENTRYPOINT ["./main"]
EXPOSE 8080
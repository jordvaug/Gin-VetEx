FROM golang:latest AS builder

#docker run gin-vetex --env-file .env -p 127.0.0.1:8080:8080/tcp
#Create dir in image to use
WORKDIR /app

#Copy all files and folders not enumerated in .dockerignore
COPY . ./

#Install dependencies
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o /main .

ENTRYPOINT ["./main"]
EXPOSE 8080

#scratch is a docker image for Go
#FROM scratch
#COPY --from=builder /main ./
#ENTRYPOINT ["./main"]
#EXPOSE 8080
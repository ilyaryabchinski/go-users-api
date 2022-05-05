FROM golang:1.18-alpine


WORKDIR /go/users-api

COPY . /go/users-api


RUN go build /go/users-api/main.go

EXPOSE 8080

ENTRYPOINT ["./main"]
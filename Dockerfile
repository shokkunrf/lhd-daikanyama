FROM golang:1.13.5-buster

RUN go get "github.com/gin-gonic/gin"
RUN go get "github.com/go-sql-driver/mysql"
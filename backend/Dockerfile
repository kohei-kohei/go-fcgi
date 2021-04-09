FROM golang:latest

RUN mkdir /backend
WORKDIR /backend

RUN go get github.com/gin-contrib/cors \
  && go get github.com/gin-gonic/gin \
  && go get github.com/go-sql-driver/mysql

FROM russellwmy/golang

MAINTAINER Russell Wong <russellwmy@gmail.com>

RUN go get -v gopkg.in/gin-gonic/gin.v1 && \
    go get -v github.com/olahol/melody && \
    go get -u -v github.com/jinzhu/gorm && \
    go get -v github.com/gin-gonic/contrib/sessions && \
    go get -v github.com/go-sql-driver/mysql

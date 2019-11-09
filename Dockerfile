FROM golang:1.13.3

WORKDIR /home/go-csd-api
COPY . .

RUN go get -u github.com/gorilla/mux
RUN go get -u github.com/go-sql-driver/mysql

EXPOSE 80

RUN go build

CMD ["./go-csd-api"]

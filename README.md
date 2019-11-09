# Golang CSD live API


**Dependecies:**

```
go get -u github.com/gorilla/mux
```
```
go get -u github.com/go-sql-driver/mysql
```



**Running container on Server with port 8080:**

```
docker build -t csd_go_server .
```
```
docker run -d -p 8080:80 csd_go_server
```

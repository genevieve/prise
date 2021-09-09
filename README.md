# prise

*prise* - French for socket

> First, a Unix socket is represented as a special file.
Your server listens on the file, accepts connections, and
reads data from those connections. Your client dials the
file to create a connection and then writes data to the connection.
[source](https://johnrefior.com/gobits/read?post=12)

### run

window 1:
```
 go run cmd/manager/main.go
```

window 2:
```
go run cmd/report/main.go
```

expected output in window 1:
```
hello, world
```

expected output in window 2:
```
wrote 12 bytes
```


### reading

- https://eli.thegreenplace.net/2019/unix-domain-sockets-in-go/
- https://github.com/eliben/code-for-blog/blob/master/2019/unix-domain-sockets-go/local-latency-benchmark.go


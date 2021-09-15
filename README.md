# prise

*prise* - French for socket

> First, a Unix socket is represented as a special file.
Your server listens on the file, accepts connections, and
reads data from those connections. Your client dials the
file to create a connection and then writes data to the connection.
[source](https://johnrefior.com/gobits/read?post=12)

### run

window 1:
```bash
go run cmd/manager/main.go
```

window 2:
```bash
go run cmd/report/main.go
```

expected output by report:
```
report connected to manager
Meeting set for:  tomorrow
```


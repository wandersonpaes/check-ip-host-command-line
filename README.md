# check-ip-host-command-line

Simple App to check IP and Server name throuht the command line.

## Search for IPs

To find an IP you can run the code below:

```
go run main.go ip --host <URL>
```

Example - Finding Amazon Brazil IP:

```
go run main.go ip --host amazon.com.br
```

## Search for Server Name

To find a Server Name you can run the code below:

```
go run main.go servers --host <URL>
```

Example - Finding Amazon Brazil Server Name:

```
go run main.go servers --host amazon.com.br
```

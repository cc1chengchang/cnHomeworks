
how to run

server
export GOVERSION="go1.19.1"; go run main.go

client
(base) chchang@chchang-a01 ~ % curl http://localhost:8090/
hello [stranger]
===================Details of the http request header:============
User-Agent=[curl/7.78.0]
Accept=[*/*]
===================Server go version:============
Version=go1.19.1

(base) chchang@chchang-a01 ~ % curl http://localhost:8090/healthz
200
(base) chchang@chchang-a01 ~ %

http://127.0.0.1:8090/healthz
 200



server access log output:
(base) chchang@chchang-a01 httpserver % export GOVERSION="go1.19.1"; go run main.go
Starting http server...
entering logging
entering root handler
2022/10/08 21:38:28 GET		/		127.0.0.1:53237		200
entering logging
2022/10/08 21:38:51 GET		/healthz		127.0.0.1:53253		200




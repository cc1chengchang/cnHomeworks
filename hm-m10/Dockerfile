FROM ubuntu
ENV MY_SERVICE_PORT=8080
LABEL multi.label1="value1" multi.label2="value2" other="value3"
ADD bin/amd64/httpserver /httpserver
EXPOSE 8080
ENTRYPOINT /httpserver

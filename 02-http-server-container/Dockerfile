FROM ubuntu
ENV MY_SERVICE_PORT=80
ADD bin/amd64/02-http-server-container /httpserver
EXPOSE 80
ENTRYPOINT /httpserver
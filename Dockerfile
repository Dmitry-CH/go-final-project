FROM ubuntu:latest

WORKDIR /usr/my-todo

COPY app .

COPY web/ ./web/

EXPOSE 3000

CMD ["/usr/my-todo/app"]

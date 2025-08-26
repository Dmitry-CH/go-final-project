FROM alpine:latest

ENV TODO_PORT=7540

WORKDIR /usr/my-todo

COPY app .

COPY web/ ./web/

EXPOSE $TODO_PORT

CMD ["/usr/my-todo/app"]

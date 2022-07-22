FROM golang
COPY ./example/hello /app/hello
WORKDIR /app
ENTRYPOINT [ "/app/hello" ]

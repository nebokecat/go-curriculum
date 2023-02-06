FROM golang:latest

ENV APP_ROOT /app
WORKDIR $APP_ROOT
COPY ./ $APP_ROOT

CMD ["go", "mod", "download"]
CMD ["go", "mod", "tidy"]

EXPOSE 8080
CMD ["go", "run", "server.go"]
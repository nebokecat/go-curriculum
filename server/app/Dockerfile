FROM golang:latest

ENV APP_ROOT /app
WORKDIR $APP_ROOT
COPY ./ $APP_ROOT

CMD ["go", "mod", "download"]
CMD ["go", "mod", "tidy"]

RUN go install github.com/cosmtrek/air@v1.49.0 && \
  go install github.com/volatiletech/sqlboiler/v4@v4.11.0 && \
  go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-psql@v4.11.0

EXPOSE 8080
CMD ["go", "run", "server.go"]

FROM golang:1.18

ENV APP_HOME /go/src/form3-client
RUN mkdir -p "$APP_HOME"
WORKDIR "$APP_HOME"

COPY ./src .

RUN go mod download

RUN go install -v ./...

CMD ["go", "test", "-v"]
FROM golang:1.18

ENV APP_HOME /go/src/form3-client
RUN mkdir -p "$APP_HOME"
WORKDIR "$APP_HOME"

COPY ./src .

RUN go get -d -v ./...

RUN go install -v ./...

EXPOSE 8085

CMD ["form3-client"]
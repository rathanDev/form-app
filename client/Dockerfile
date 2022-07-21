FROM golang:1.10

ENV APP_HOME /go/src/form-app
RUN mkdir -p "$APP_HOME"
WORKDIR "$APP_HOME"

COPY . .

RUN go get -d -v ./...

RUN go install -v ./...

CMD ["form-app"]
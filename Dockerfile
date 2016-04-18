FROM golang:1.6.1

RUN mkdir -p /go/src/app
WORKDIR /go/src/app

COPY . /go/src/app

RUN go-wrapper download
RUN go-wrapper install

ENV PORT 3000

EXPOSE 3000

CMD ["app"]

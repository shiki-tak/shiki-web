FROM golang:1.13.7

RUN go get github.com/labstack/echo/...

WORKDIR /shiki-web
ADD . /shiki-web

CMD ["go", "run", "app/main.go"]

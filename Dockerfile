FROM golang:1.13.4
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go clean --modcache
RUN go build -o ./src/server .
CMD["app/server"]

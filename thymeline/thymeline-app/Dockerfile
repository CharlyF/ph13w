FROM golang:latest 
RUN mkdir /app 
ADD . /app/ 
WORKDIR /app 
ADD vendor /go/src
RUN go build -o main . 
CMD ["/app/main"]

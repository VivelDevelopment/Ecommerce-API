FROM golang:1.16-alpine

WORKDIR /app
ADD . /app
EXPOSE 8080
RUN go build -o main .
CMD [ "./main" ]
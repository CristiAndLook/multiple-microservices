FROM  golang:1.18.1-alpine 
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go build -o main .
CMD ["/app/main"]

#docker run -it --rm -p 8080:8081 --name go-service go-microservice
FROM golang:1.23-alpine

COPY . /app 

WORKDIR /app 

RUN go mod tidy 

EXPOSE 8090 

ENTRYPOINT ["go", "run", "./cmd"]

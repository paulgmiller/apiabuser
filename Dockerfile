FROM golang:alpine
RUN mkdir /app 
COPY go.mod go.sum /app/
WORKDIR /app 
RUN go mod download
ADD . /app/
RUN go build -o apiabuser .
RUN adduser -S -D -H -h /app appuser
USER appuser
CMD ["./apiabuser"]
FROM golang:la

WORKDIR /go/src/app
COPY . .

RUN go-wrapper download
RUN go-wrapper install    
CMD ["go-wrapper", "run"]
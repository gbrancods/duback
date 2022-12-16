FROM golang:1.19

WORKDIR /app

COPY ./ ./

RUN go mod download

RUN go build -o /duback

RUN apt-get update
RUN apt-get install -y postgresql-client-13

CMD [ "/duback" ]
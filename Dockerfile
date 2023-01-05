FROM golang:1.19

WORKDIR /app

COPY ./ ./

RUN sh -c 'echo "deb http://apt.postgresql.org/pub/repos/apt $(lsb_release -cs)-pgdg main" > /etc/apt/sources.list.d/pgdg.list'
RUN wget -qO- https://www.postgresql.org/media/keys/ACCC4CF8.asc | tee /etc/apt/trusted.gpg.d/pgdg.asc &>/dev/null
RUN apt update
RUN apt install postgresql postgresql-client -y

RUN go mod download
RUN go build -o /duback

CMD [ "/duback" ]

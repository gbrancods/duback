# stage de build
FROM golang:1.19 AS build

WORKDIR /app

COPY . /app

RUN CGO_ENABLED=0 GOOS=linux go build -o duback cmd/main.go

# stage imagem final
FROM ubuntu:latest

RUN apt update

RUN apt install lsb-core -y && apt install wget -y

RUN sh -c 'echo "deb http://apt.postgresql.org/pub/repos/apt $(lsb_release -cs)-pgdg main" > /etc/apt/sources.list.d/pgdg.list'

RUN wget -qO- https://www.postgresql.org/media/keys/ACCC4CF8.asc | tee /etc/apt/trusted.gpg.d/pgdg.asc &>/dev/null

RUN apt-key adv --keyserver keyserver.ubuntu.com --recv-keys 7FCC7D46ACCC4CF8 

RUN apt update -y

RUN apt install postgresql-client -y

WORKDIR /app

COPY --from=build /app/duback ./

RUN mkdir scripts 
COPY --from=build /app/scripts/backup.sh ./scripts/backup.sh

RUN mkdir backups

CMD [ "./duback" ]

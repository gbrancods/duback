# --Build stage
FROM golang:1.19 AS build

WORKDIR /app

COPY . /app

RUN CGO_ENABLED=0 GOOS=linux go build -o duback cmd/main.go

# --Final stage
FROM ubuntu:latest

# Install dependences and postgres client v15
RUN apt update -y
RUN apt install lsb-core -y && apt install wget -y
RUN sh -c 'echo "deb http://apt.postgresql.org/pub/repos/apt $(lsb_release -cs)-pgdg main" > /etc/apt/sources.list.d/pgdg.list'
RUN wget -qO- https://www.postgresql.org/media/keys/ACCC4CF8.asc | tee /etc/apt/trusted.gpg.d/pgdg.asc &>/dev/null
RUN apt-key adv --keyserver keyserver.ubuntu.com --recv-keys 7FCC7D46ACCC4CF8 
RUN apt update -y
RUN apt install postgresql-client -y

WORKDIR /app

# Copy app bin
COPY --from=build /app/duback ./

# Copy backup shell script
RUN mkdir scripts 
COPY --from=build /app/scripts/backup.sh ./scripts/backup.sh

# Create folder to receive backups
RUN mkdir backups

CMD [ "./duback" ]

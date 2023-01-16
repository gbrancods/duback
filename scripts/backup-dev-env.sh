#!/bin/bash

backup_date=$(date +'%d-%m-%Y')
backup_hour=$(date +'%Hh%Mm')

mkdir backups/dev-env

mkdir backups/dev-env/${backup_date}-${backup_hour}

pg_dump --dbname=postgresql://postgres:000DHPdhpDHP000@192.168.55.250:5433/postgres > backups/dev-env/${backup_date}-${backup_hour}/backup-dbpay-${backup_date}-${backup_hour}

pg_dump --dbname=postgresql://postgres:000DHPdhpDHP000@192.168.55.250:5432/postgres > backups/dev-env/${backup_date}-${backup_hour}/backup-dbeco-${backup_date}-${backup_hour}

pg_dump --dbname=postgresql://postgres:000DHPdhpDHP000@192.168.55.250:5431/postgres > backups/dev-env/${backup_date}-${backup_hour}/backup-dbchat-${backup_date}-${backup_hour}

#!/bin/bash

backup_date=$(date +'%d-%m-%Y')

pg_dump --dbname=postgresql://postgres:000DHPdhpDHP000@192.168.55.250:5433/postgres > backups/backup-dbpay-${backup_date}
pg_dump --dbname=postgresql://postgres:000DHPdhpDHP000@192.168.55.250:5432/postgres > backups/backup-dbecommerce-${backup_date}
pg_dump --dbname=postgresql://postgres:000DHPdhpDHP000@192.168.55.250:5431/postgres > backups/backup-dbchat-${backup_date}

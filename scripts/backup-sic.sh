#!/bin/bash

backup_date=$(date +'%d-%m-%Y')
backup_hour=$(date +'%Hh%Mm')

mkdir backups/SIC

mkdir backups/SIC/${backup_date}

gbak -user SYSDBA -pas masterkey -b 192.168.0.241:C:/SupraSys/SIC/Arq01/ARQSIST.FDB backups/SIC/${backup_date}/sic-${backup_date}-${backup_hour}.fbk

gzip backups/SIC/${backup_date}/sic-${backup_date}-${backup_hour}.fbk

rm -r backups/SIC/${backup_date}/sic-${backup_date}-${backup_hour}.fbk
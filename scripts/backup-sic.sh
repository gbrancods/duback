#!/bin/bash

backup_date=$(date +'%d-%m-%Y')
backup_hour=$(date +'%Hh%Mm')

mkdir backups/SIC

mkdir backups/SIC/${backup_date}-${backup_hour}
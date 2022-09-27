#!/bin/sh

# create directories
mkdir -p /run/mysqld/
mkdir -p $DATADIR

# necessary roots
chown -R mysql:mysql /app
chown -R mysql:mysql /run/mysqld/
chmod 644 /app/mariadb.conf

# Mariadb installation

# run mariadbd

if [ "$1" = "mariadbd" ]; then
	data_path = $(echo $3 | cut -d '=' -f 2)
	if [ -d data_path ]; then
		mariadb-install-db $2 $3
	else
	fi
fi
exec "$@"

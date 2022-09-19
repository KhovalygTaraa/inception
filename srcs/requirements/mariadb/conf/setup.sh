#!/bin/sh

# create directories
mkdir -p /run/mysqld/

# necessary roots
chown -R mysql:mysql /app
chown -R mysql:mysql /run/mysqld/

# Mariadb installation
mariadb-install-db $1

# run mariadbd
exec "$@"

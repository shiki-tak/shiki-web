#!/bin/sh
chmod 0775 docker-entrypoint-initdb.d/init-mysql.sh
./docker-entrypoint-initdb.d/init-database.sh
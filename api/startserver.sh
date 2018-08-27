#! /bin/bash

set -e

until mysqladmin ping -h shokubunka_db --silent; do
  echo "Waiting for mysqld to be connectable..."
  sleep 1
done

dep ensure

go build
./api

exec "$@"
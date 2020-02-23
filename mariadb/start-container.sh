#!/bin/sh

# should mount volumes later and make persistent, and need a way to populate db tables on startup
docker run --rm -d --env-file .env mariadb:latest

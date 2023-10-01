#!/bin/bash

if [ "$(docker ps -a -q -f name=mysql_db)" ]; then
  docker stop mysql_db
  docker rm -f mysql_db
fi

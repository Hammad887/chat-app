#!/bin/bash

if [ "$(docker ps -a -q -f name=mysql_dev)" ]; then
  docker rm -f mysql_dev
fi

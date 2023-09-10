#!/bin/bash

if [ "$(docker ps -a -q -f name=mysql_dev)" ]; then
  docker stop mysql_dev
  docker rm -f mysql_dev
fi
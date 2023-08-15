#!/bin/bash

if [ ! -d "docs" ]; then
  mkdir docs 
fi

swagger generate server -t docs --exclude-main
swagger generate client -t docs

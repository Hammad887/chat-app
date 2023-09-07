#!/bin/bash

# Directory check and creation
if [ ! -d "docs" ]; then
  mkdir docs 
fi

swagger docserate server -t docs --exclude-main
swagger docserate client -t docs

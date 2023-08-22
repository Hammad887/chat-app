#!/bin/bash

if [ ! -d "gen" ]; then
  mkdir gen 
fi

swagger generate server -t gen --exclude-main
swagger generate client -t gen

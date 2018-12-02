#!/bin/bash

if [ "$1" = "" ]; then
  echo "Usage: $0 dir"
  exit 1
fi

cp -R template "$1"

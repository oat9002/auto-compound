#! /bin/bash

if grep -q Alive "/tmp/auto-compound-healthcheck.txt"; then
  exit 0
fi

exit 1
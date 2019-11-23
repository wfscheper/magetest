#!/bin/bash

set -e

rc=0
while test ${rc} -eq 0; do
    go run mage/mage.go -clean
    go run mage/mage.go -l -debug
    rc=$?
done

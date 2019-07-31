#!/usr/bin/env bash

cat input | sed -e 's/[^0-9-]/\n/g' | grep . | awk '{sum+=$1}END{print sum}'


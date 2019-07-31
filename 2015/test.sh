#!/usr/bin/env bash

for i in $(seq 1 25); do
  echo "Testing output of day $i"
  cd "$i"

  echo -n "-- Part1......"
  START=$(date +%s%N | cut -b1-13)
  diff <(make -s 1) output1 && echo -n "ok" || echo -n "FAIL"
  END=$(date +%s%N | cut -b1-13)
  DIFF=$(echo "$END - $START" | bc)
  echo "     time: ${DIFF}ms"

  echo -n "-- Part2......"
  START=$(date +%s%N | cut -b1-13)
  diff <(make -s 2) output2 && echo -n "ok" || echo -n "FAIL"
  END=$(date +%s%N | cut -b1-13)
  DIFF=$(echo "$END - $START" | bc)
  echo "     time: ${DIFF}ms"

  cd ..
done 2>/dev/null

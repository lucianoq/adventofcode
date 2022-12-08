#!/usr/bin/env bash

YEAR="$1"

printf "%11sDay   Part   Result   Duration\n" ""

for i in $(seq 1 25); do
  cd "$YEAR/$i" || exit

  make -s main1
  make -s main2

  for part in 1 2; do

    [[ $part -eq 1 ]] && printf 'Checking%5s' "$i" || printf '%13s' ""
    printf '%7d' "$part"

    START=$(date +%s%N | cut -b1-13)
    OUT=$(make -s run${part})
    END=$(date +%s%N | cut -b1-13)

    printf "    "
    if diff <(echo "$OUT") "output$part"; then
      printf "  ok  "
    else
      printf " FAIL "
    fi

    ELAPSED=$(echo "$END - $START" | bc)
    printf '%8.4sms\n' "${ELAPSED}"

  done

  make -s clean

  cd ../..
done 2>/dev/null

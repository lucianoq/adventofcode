#!/usr/bin/env bash

COOKIE_TOKEN="$1"
YEAR="$2"
DAY="$3"

#create dir
mkdir -p "$YEAR/$DAY"

# download input files
http "https://adventofcode.com/$YEAR/day/$DAY/input" "Cookie:session=$COOKIE_TOKEN;" >"$YEAR/$DAY/input"

# download assignment
http "https://adventofcode.com/$YEAR/day/$DAY" "Cookie:session=$COOKIE_TOKEN;" | pup 'article.day-desc' >"$YEAR/$DAY/tmp.html"
lynx -dump "$YEAR/$DAY/tmp.html" -width 80 >"$YEAR/$DAY/assignment"
rm -f "$YEAR/$DAY/tmp.html"

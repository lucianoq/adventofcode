#!/usr/bin/env bash

YEAR=2017
COOKIE_TOKEN="$2"
DAY="$1"

#create dir
mkdir "$DAY"

# download input files
http "https://adventofcode.com/$YEAR/day/$DAY/input" "Cookie:session=$COOKIE_TOKEN;" >"$DAY/input"

# download assignment
http "https://adventofcode.com/$YEAR/day/$DAY" "Cookie:session=$COOKIE_TOKEN;" | pup 'article.day-desc' >"$DAY/tmp.html"
lynx -dump "$DAY/tmp.html" -width 80 >"$DAY/assignment"
rm -f "$DAY/tmp.html"

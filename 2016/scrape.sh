#!/usr/bin/env bash

COOKIE_TOKEN="$2"
DAY="$1"

#create dir
mkdir "$DAY"

# download input files
http "https://adventofcode.com/2016/day/$DAY/input" "Cookie:session=$COOKIE_TOKEN;" >"$DAY/input"

# download assignment
http "https://adventofcode.com/2016/day/$DAY" "Cookie:session=$COOKIE_TOKEN;" | pup 'article.day-desc' >"$DAY/tmp.html"
lynx -dump "$DAY/tmp.html" -width 80 >"$DAY/assignment"
rm -f "$DAY/tmp.html"

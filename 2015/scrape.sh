#!/usr/bin/env bash

export COOKIE_TOKEN="$1"

for i in $(seq 1 25); do

  #create dir
  mkdir "$i"

  # download input files
  http "https://adventofcode.com/2015/day/$i/input" "Cookie:session=$COOKIE_TOKEN;" >"$i/input"

  # download assignment
  http "https://adventofcode.com/2015/day/$i" "Cookie:session=$COOKIE_TOKEN;" | pup 'article.day-desc' >"$i/tmp.html"
  lynx -dump "$i/tmp.html" -width 80 >"$i/assignment"
  rm -f "$i/tmp.html"

done

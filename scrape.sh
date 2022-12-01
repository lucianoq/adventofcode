#!/usr/bin/env bash

YEAR="$1"
DAY="$2"

#create dir
mkdir -p "$YEAR/$DAY"
cd "$YEAR/$DAY" || exit

touch output1 output2

GO_TEMPLATE="package main\n\nfunc main() {\n    \n}\n"
echo -en "$GO_TEMPLATE" | tee main1.go main2.go >/dev/null


cat >Makefile <<EOF
1:
	go run main1.go <input

2:
	go run main2.go <input
EOF

# download input files
http "https://adventofcode.com/$YEAR/day/$DAY/input" "Cookie:session=$AOC_SESSION;" >input

# download assignment
http "https://adventofcode.com/$YEAR/day/$DAY" "Cookie:session=$AOC_SESSION;" | pup 'article.day-desc' >tmp.html
lynx -dump tmp.html -width 80 >assignment
rm -f tmp.html

#!/usr/bin/env bash

YEAR="$1"
DAY="$2"

#create dir
mkdir -p "$YEAR/$DAY"
cd "$YEAR/$DAY" || exit

touch output1 output2

GO_TEMPLATE="package main\n\nfunc main() {\n    \n}\n"
if [ ! -f main1.go ]; then echo -en "$GO_TEMPLATE" > main1.go; fi
if [ ! -f main2.go ]; then echo -en "$GO_TEMPLATE" > main2.go; fi

if [ ! -f Makefile ]; then cat >Makefile <<EOF
1:
	go run main1.go <input

2:
	go run main2.go <input
EOF
fi

# download input files
http "https://adventofcode.com/$YEAR/day/$DAY/input" "Cookie:session=$AOC_SESSION;" >input

# download assignment
http "https://adventofcode.com/$YEAR/day/$DAY" "Cookie:session=$AOC_SESSION;" | pup 'article.day-desc' >tmp.html
lynx -dump tmp.html -width 80 >assignment
rm -f tmp.html

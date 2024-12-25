input:
	http "https://adventofcode.com/2024/day/25/input" "Cookie:session=${AOC_SESSION};" >input

main1:
	go build -o main1 main1.go

.PHONY: run1 clean

run1: main1 input
	./main1 <input

clean:
	rm -f main1 input

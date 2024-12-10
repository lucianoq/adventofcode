input:
	http "https://adventofcode.com/2024/day/9/input" "Cookie:session=${AOC_SESSION};" >input

main1:
	go build -o main1 main1.go

main2:
	go build -o main2 main2.go

.PHONY: run1 run2 clean

run1: main1 input
	./main1 <input

run2: main2 input
	./main2 <input

clean:
	rm -f main1 main2 input


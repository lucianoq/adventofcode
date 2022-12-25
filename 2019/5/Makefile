main1:
	go build -o main1 main.go vm.go

main2:
	go build -o main2 main.go vm.go

.PHONY: run1 run2 clean

run1: main1
	./main1 input <part1.input | tail -n 1

run2: main2
	./main2 input <part2.input

clean:
	rm -f main1 main2


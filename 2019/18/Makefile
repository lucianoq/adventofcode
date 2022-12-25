main1:
	go build -o main1 main1.go graph.go common.go

main2:
	go build -o main2 main2.go graph.go common.go

.PHONY: run1 run2 clean

run1: main1
	./main1 <input1

run2: main2
	./main2 <input2

clean:
	rm -f main1 main2


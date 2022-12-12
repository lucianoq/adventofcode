main1:
	go build -o main1 main1.go common.go bfs.go

main2:
	go build -o main2 main2.go common.go bfs.go

.PHONY: run1 run2 clean

run1: main1
	./main1 <input

run2: main2
	./main2 <input

clean:
	rm -f main1 main2


main1:
	go build -o main1 main1.go common.go

main2:
	go build -o main2 main2.go common.go

.PHONY: run1 run2 clean draw

run1: main1
	./main1 <input

run2: main2
	./main2 <input

clean:
	rm -f main1 main2

draw:
	go run main2_draw.go common.go <input

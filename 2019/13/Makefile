main1:
	go build -o main1 main1.go common.go cmd.go vm.go

main2:
	go build -o main2 main2.go common.go cmd.go vm.go

.PHONY: run1 run2 clean play

run1: main1
	./main1 <input

run2: main2
	./main2 <input

clean:
	rm -f main1 main2

play:
	go run main2_play.go cmd.go vm.go common.go
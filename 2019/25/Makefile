main1:
	go build -o main1 main1.go vm.go cmd.go

main2:


.PHONY: run1 run2 clean

run1: main1
	./main1 <input | sed -En 's/^.*typing ([0-9]+) on.*$$/\1/p'

run2: main2
	true

clean:
	rm -f main1 main2

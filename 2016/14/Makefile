1:
	go run main1.go hashes.go common.go

2:
	go run main2.go hashes.go common.go

test:
	go test common.go common_test.go -v
	go test hashes.go hashes_test.go -v

benchmark:
	go test hashes.go hashes_benchmark_test.go --bench=. -v

1:
	go run main1.go hashes.go common.go

2:
	go run main2.go hashes.go common.go

test:
	go test -v common*
	go test -v hashes*

benchmark:
	go test hashes* --bench=. -v

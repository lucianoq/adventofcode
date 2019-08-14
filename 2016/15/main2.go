package main

import "fmt"

func main() {
	disks := []Disk{{0, 7}, {0, 13}, {2, 3}, {2, 5}, {0, 17}, {7, 19}, {0, 11}}

	fmt.Println(run(disks))
}

package main

type Disk struct {
	InitialState int
	Period       int
}

func run(disks []Disk) int {
time:
	for t := 0; ; t++ {
		for i, d := range disks {
			delay := t + i + 1
			if (d.InitialState+delay)%d.Period != 0 {
				continue time
			}
		}

		return t
	}
}

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
	"time"
)

type eventType int

const (
	begin = iota
	fallsAsleep
	wakesUp
)

type event struct {
	when  time.Time
	guard int
	eType eventType
}

func main() {
	logs := getOrderedLogs()
	asleepRecord := make(map[int]time.Duration)

	for i, e := range logs {
		if e.eType == wakesUp {
			asleepRecord[e.guard] += e.when.Sub(logs[i-1].when)
		}
	}

	var sleepyGuard int
	var mostSleep time.Duration
	for id, duration := range asleepRecord {
		if duration > mostSleep {
			sleepyGuard = id
			mostSleep = duration
		}
	}

	var minutes [60]int
	for i, e := range logs {
		if e.guard == sleepyGuard {
			if e.eType == wakesUp {
				for j := logs[i-1].when.Minute(); j < e.when.Minute(); j++ {
					minutes[j]++
				}
			}
		}
	}

	var maxMinute, maxTimes int
	for m, times := range minutes {
		if times > maxTimes {
			maxMinute = m
			maxTimes = times
		}
	}

	fmt.Println(maxMinute * sleepyGuard)
}

func getOrderedLogs() []*event {
	logs := make([]*event, 0)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		var y, m, d, hh, mm, id int
		switch {
		case strings.Contains(line, "begins"):
			n, err := fmt.Sscanf(line, "[%d-%d-%d %d:%d] Guard #%d begins shift", &y, &m, &d, &hh, &mm, &id)
			if err != nil || n != 6 {
				log.Fatal(err)
			}
			logs = append(logs, &event{
				eType: begin,
				guard: id,
				when:  time.Date(y, time.Month(m), d, hh, mm, 0, 0, time.UTC),
			})
		case strings.Contains(line, "asleep"):
			n, err := fmt.Sscanf(line, "[%d-%d-%d %d:%d] falls asleep", &y, &m, &d, &hh, &mm)
			if err != nil || n != 5 {
				log.Fatal(err)
			}
			logs = append(logs, &event{
				eType: fallsAsleep,
				guard: -1,
				when:  time.Date(y, time.Month(m), d, hh, mm, 0, 0, time.UTC),
			})
		case strings.Contains(line, "wakes"):
			n, err := fmt.Sscanf(line, "[%d-%d-%d %d:%d] wakes up", &y, &m, &d, &hh, &mm)
			if err != nil || n != 5 {
				log.Fatal(err)
			}
			logs = append(logs, &event{
				eType: wakesUp,
				guard: -1,
				when:  time.Date(y, time.Month(m), d, hh, mm, 0, 0, time.UTC),
			})
		}
	}

	sort.Slice(logs, func(i, j int) bool {
		return logs[i].when.Before(logs[j].when)
	})

	var tempID int
	for _, e := range logs {
		if e.guard != -1 {
			tempID = e.guard
		} else {
			e.guard = tempID
		}
	}

	return logs
}

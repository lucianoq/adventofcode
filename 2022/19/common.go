package main

import (
	"bufio"
	"fmt"
	"os"
)

type Blueprint struct {
	ID,
	OreBotOreCost, ClayBotOreCost,
	ObsidianBotOreCost, ObsidianBotClayCost,
	GeodeBotOreCost, GeodeBotObsidianCost uint16
}

func parse() []Blueprint {
	list := []Blueprint{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()

		var (
			id,
			oreRobotCost, clayCost, obsidianCostOre,
			obsidianCostClay, geodeCostOre, geodeCostObsidian uint16
		)
		fmt.Sscanf(
			line,
			"Blueprint %d: Each ore robot costs %d ore. Each clay robot costs %d ore. Each obsidian robot costs %d ore and %d clay. Each geode robot costs %d ore and %d obsidian.",
			&id,
			&oreRobotCost, &clayCost, &obsidianCostOre,
			&obsidianCostClay, &geodeCostOre, &geodeCostObsidian,
		)
		list = append(list, Blueprint{id, oreRobotCost, clayCost, obsidianCostOre,
			obsidianCostClay, geodeCostOre, geodeCostObsidian},
		)
	}

	return list
}

type Status struct {
	Ore, Clay, Obsidian, Geode,
	OreBots, ClayBots, ObsidianBots, GeodeBots uint16
}

func (s Status) Gather() Status {
	s.Ore += s.OreBots
	s.Clay += s.ClayBots
	s.Obsidian += s.ObsidianBots
	s.Geode += s.GeodeBots
	return s
}

func (s Status) BuyGeodeBot(b Blueprint) Status {
	s.Ore -= b.GeodeBotOreCost
	s.Obsidian -= b.GeodeBotObsidianCost
	s = s.Gather()
	s.GeodeBots++
	return s
}

func (s Status) BuyObsidianBot(b Blueprint) Status {
	s.Ore -= b.ObsidianBotOreCost
	s.Clay -= b.ObsidianBotClayCost
	s = s.Gather()
	s.ObsidianBots++
	return s
}

func (s Status) BuyClayBot(b Blueprint) Status {
	s.Ore -= b.ClayBotOreCost
	s = s.Gather()
	s.ClayBots++
	return s
}

func (s Status) BuyOreBot(b Blueprint) Status {
	s.Ore -= b.OreBotOreCost
	s = s.Gather()
	s.OreBots++
	return s
}

func bfs(b Blueprint) int {

	var (
		maxNeededOrePerMinute      = max(b.ClayBotOreCost, b.ObsidianBotOreCost, b.GeodeBotOreCost)
		maxNeededClayPerMinute     = b.ObsidianBotClayCost
		maxNeededObsidianPerMinute = b.GeodeBotObsidianCost

		// 1 ore bot only, 0 resources
		start = Status{0, 0, 0, 0, 1, 0, 0, 0}

		toDo    = []Status{start}
		visited = map[Status]struct{}{start: {}}

		maxGeodes uint16
		minutes   uint16
		curr      Status

		appendCandidate = func(next Status) {
			if _, ok := visited[next]; !ok {
				visited[next] = struct{}{}
				toDo = append(toDo, next)
			}
		}
	)

	for len(toDo) > 0 && minutes < MinuteLimit {
		items := len(toDo)
		for i := 0; i < items; i++ {
			curr, toDo = toDo[0], toDo[1:]

			// No need to reach MinuteLimit,
			// There is no point on buying bots in the last minute
			// 'cause they'll have no time to be useful.
			// We can calculate the value the last but one minute.
			if minutes == MinuteLimit-1 {
				curr = curr.Gather()
				if curr.Geode > maxGeodes {
					maxGeodes = curr.Geode
				}
			}

			// if you can build a geode robot,
			if curr.Ore >= b.GeodeBotOreCost && curr.Obsidian >= b.GeodeBotObsidianCost {
				appendCandidate(curr.BuyGeodeBot(b))

				// that's the best path for sure
				// ignore other steps
				continue
			}

			// No need to create more ObsidianBots if we reached the
			// rate needed to create 1 GeodeBot per minute
			if curr.ObsidianBots < maxNeededObsidianPerMinute {
				if curr.Ore >= b.ObsidianBotOreCost && curr.Clay >= b.ObsidianBotClayCost {
					appendCandidate(curr.BuyObsidianBot(b))
				}
			}

			// No need to create more ClayBots if we reached the
			// rate needed to create 1 ClayBot per minute
			if curr.ClayBots < maxNeededClayPerMinute {
				if curr.Ore >= b.ClayBotOreCost {
					appendCandidate(curr.BuyClayBot(b))
				}
			}

			// No need to create more OreBots if we reached the
			// rate needed to create 1 OreBot per minute
			if curr.OreBots < maxNeededOrePerMinute {
				if curr.Ore >= b.OreBotOreCost {
					appendCandidate(curr.BuyOreBot(b))
				}
			}

			// buy nothing and wait
			appendCandidate(curr.Gather())
		}
		minutes++
	}
	return int(maxGeodes)
}

func max(items ...uint16) uint16 {
	var max uint16
	for _, item := range items {
		if item > max {
			max = item
		}
	}
	return max
}

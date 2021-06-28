package main

import (
	"bufio"
	"io"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func parse(reader io.Reader) []*G {
	re := regexp.MustCompile(`(\d+) units each with (\d+) hit points (\(.*\) )?with an attack that does (\d+) (.*) damage at initiative (\d+)`)

	scanner := bufio.NewScanner(reader)

	immune := true
	var groups []*G

	id := 1

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		switch line {
		case "Immune System:", "Infection:":
			// noop
		case "":
			immune = false
			id = 1
		default:
			group := &G{}

			match := re.FindStringSubmatch(line)

			group.ID = id
			group.NumUnits, _ = strconv.Atoi(match[1])
			group.HP, _ = strconv.Atoi(match[2])
			group.AttDamage, _ = strconv.Atoi(match[4])
			group.AttType = parseAttType(match[5])
			group.Initiative, _ = strconv.Atoi(match[6])

			group.Immunities = map[AttackType]struct{}{}
			immunities, weaknesses := parseBonusMalus(match[3])
			for _, i := range immunities {
				group.Immunities[i] = struct{}{}
			}

			group.Weaknesses = map[AttackType]struct{}{}
			for _, w := range weaknesses {
				group.Weaknesses[w] = struct{}{}
			}

			if immune {
				group.Team = ImmuneSystem
			} else {
				group.Team = Infection
			}

			groups = append(groups, group)

			id++
		}
	}

	return groups
}

func parseBonusMalus(s string) ([]AttackType, []AttackType) {
	s = strings.TrimSpace(strings.Trim(s, " ()"))

	if s == "" {
		return []AttackType{}, []AttackType{}
	}

	// examples:
	//   immune to radiation, slashing; weak to cold
	//   weak to fire; immune to slashing

	parseList := func(list string) []AttackType {
		types := make([]AttackType, 0)
		items := strings.Split(list, ",")
		for _, i := range items {
			types = append(types, parseAttType(strings.TrimSpace(i)))
		}
		return types
	}

	parseLine := func(line string) []AttackType {
		if line == "" {
			return make([]AttackType, 0)
		}
		parts := strings.Split(line, " to ")
		return parseList(parts[1])
	}

	var weak, immu string

	idx := strings.Index(s, ";")
	switch {
	case idx != -1:
		if strings.HasPrefix(s, "weak") {
			weak = s[:idx]
			immu = s[idx+1:]
		} else {
			immu = s[:idx]
			weak = s[idx+1:]
		}

	case strings.HasPrefix(s, "weak"):
		weak = s
	case strings.HasPrefix(s, "immune"):
		immu = s
	default:
		log.Fatal("wrong parsing")
	}
	return parseLine(immu), parseLine(weak)
}

func parseAttType(s string) AttackType {
	switch s {
	case "bludgeoning":
		return Bludgeoning
	case "fire":
		return Fire
	case "radiation":
		return Radiation
	case "slashing":
		return Slashing
	case "cold":
		return Cold
	}
	log.Fatal("wrong attack type")
	return -1
}

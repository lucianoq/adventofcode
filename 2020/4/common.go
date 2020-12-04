package main

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Document struct {
	BYR, IYR, EYR, HGT, HCL, ECL, PID, CID string
}

func (d *Document) Present() bool {
	return d.BYR != "" && d.IYR != "" && d.EYR != "" &&
		d.HGT != "" && d.HCL != "" && d.ECL != "" && d.PID != ""
}

func (d *Document) Valid() bool {

	// byr (Birth Year) - four digits; at least 1920 and at most 2002.
	n, err := strconv.Atoi(d.BYR)
	if err != nil || n < 1920 || n > 2002 {
		return false
	}

	// iyr (Issue Year) - four digits; at least 2010 and at most 2020.
	n, err = strconv.Atoi(d.IYR)
	if err != nil || n < 2010 || n > 2020 {
		return false
	}

	// eyr (Expiration Year) - four digits; at least 2020 and at most 2030.
	n, err = strconv.Atoi(d.EYR)
	if err != nil || n < 2020 || n > 2030 {
		return false
	}

	// hgt (Height) - a number followed by either cm or in:
	// If cm, the number must be at least 150 and at most 193.
	// If in, the number must be at least 59 and at most 76.
	l := len(d.HGT)
	if l < 3 {
		return false
	}
	h, unit := d.HGT[:l-2], d.HGT[l-2:]
	switch unit {
	case "in":
		n, err = strconv.Atoi(h)
		if err != nil || n < 59 || n > 76 {
			return false
		}
	case "cm":
		n, err = strconv.Atoi(h)
		if err != nil || n < 150 || n > 193 {
			return false
		}
	default:
		return false
	}

	// hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
	matched, err := regexp.MatchString(`^#[0-9a-f]{6}$`, d.HCL)
	if err != nil || !matched {
		return false
	}

	// ecl (Eye Color) - exactly one of: amb blu brn gry grn hzl oth.
	switch d.ECL {
	case "amb", "blu", "brn", "gry", "grn", "hzl", "oth":
	default:
		return false
	}

	// pid (Passport ID) - a nine-digit number, including leading zeroes.
	matched, err = regexp.MatchString(`^[0-9]{9}$`, d.PID)
	if err != nil || !matched {
		return false
	}

	// cid (Country ID) - ignored, missing or not.

	return true
}

func parse() []*Document {
	scanner := bufio.NewScanner(os.Stdin)

	d := &Document{}
	ds := make([]*Document, 0)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if line == "" {
			ds = append(ds, d)
			d = &Document{}
			continue
		}

		ff := strings.Split(line, " ")
		for _, f := range ff {
			p := strings.Split(f, ":")

			switch p[0] {
			case "byr":
				d.BYR = p[1]
			case "iyr":
				d.IYR = p[1]
			case "eyr":
				d.EYR = p[1]
			case "hgt":
				d.HGT = p[1]
			case "hcl":
				d.HCL = p[1]
			case "ecl":
				d.ECL = p[1]
			case "pid":
				d.PID = p[1]
			case "cid":
				d.CID = p[1]
			}
		}
	}

	ds = append(ds, d)

	return ds
}

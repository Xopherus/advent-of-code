// https://adventofcode.com/2020/day/4
package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"
)

type Passport struct {
	Byr, Iyr, Eyr      int
	Hcl, Ecl, Pid, Cid string

	Hgt struct {
		Value int
		Unit  string
	}
}

// Unmarshal converts bytes into a Passport struct
func (p *Passport) Unmarshal(data []byte) error {
	for _, field := range bytes.Split(data, []byte(" ")) {
		parts := bytes.Split(field, []byte(":"))
		k, v := string(parts[0]), string(parts[1])

		switch k {
		case "byr":
			fmt.Sscanf(v, "%d", &p.Byr)
		case "iyr":
			fmt.Sscanf(v, "%d", &p.Iyr)
		case "eyr":
			fmt.Sscanf(v, "%d", &p.Eyr)
		case "hgt":
			fmt.Sscanf(v, "%d%2s", &p.Hgt.Value, &p.Hgt.Unit)
		case "hcl":
			fmt.Sscanf(v, "%s", &p.Hcl)
		case "ecl":
			fmt.Sscanf(v, "%s", &p.Ecl)
		case "pid":
			fmt.Sscanf(v, "%s", &p.Pid)
		case "cid":
			fmt.Sscanf(v, "%s", &p.Cid)
		}
	}

	return nil
}

func part1Validator(p Passport) bool {
	requiredFields := []string{
		"byr",
		"iyr",
		"eyr",
		"hgt",
		"hcl",
		"ecl",
		"pid",
	}

	passportFields := map[string]bool{}
	v := reflect.ValueOf(p)

	for i := 0; i < v.NumField(); i++ {
		if !v.Field(i).IsZero() {
			passportFields[strings.ToLower(v.Type().Field(i).Name)] = true
		}

	}

	for _, f := range requiredFields {
		if _, ok := passportFields[f]; !ok {
			return false
		}
	}

	return true
}

// for part 2
var (
	errInvalidByr = fmt.Errorf("invalid Birth Year")
	errInvalidIyr = fmt.Errorf("invalid Issue Year")
	errInvalidEyr = fmt.Errorf("invalid Expiration Year")
	errInvalidHgt = fmt.Errorf("invalid Height")
	errInvalidHcl = fmt.Errorf("invalid Hair Color")
	errInvalidEcl = fmt.Errorf("invalid Eye Color")
	errInvalidPid = fmt.Errorf("invalid Passport ID")
)

func (p Passport) validateByr() error {
	if p.Byr < 1920 || p.Byr > 2002 {
		return errInvalidByr
	}

	return nil
}

func (p Passport) validateIyr() error {
	if p.Iyr < 2010 || p.Iyr > 2020 {
		return errInvalidIyr
	}

	return nil
}

func (p Passport) validateEyr() error {
	if p.Eyr < 2020 || p.Eyr > 2030 {
		return errInvalidEyr
	}

	return nil
}

func (p Passport) validateHgt() error {
	switch p.Hgt.Unit {
	case "cm":
		if p.Hgt.Value < 150 || p.Hgt.Value > 193 {
			return errInvalidHgt
		}
	case "in":
		if p.Hgt.Value < 59 || p.Hgt.Value > 76 {
			return errInvalidHgt
		}
	default:
		return errInvalidHgt
	}

	return nil
}

func (p Passport) validateHcl() error {
	if len(p.Hcl) != 7 || !strings.HasPrefix(p.Hcl, "#") {
		return errInvalidHcl
	}

	// try parsing hcl as a hex number (0-9, a-f)
	_, err := strconv.ParseInt(strings.TrimPrefix(p.Hcl, "#"), 16, 0)
	if err != nil {
		return errInvalidHcl
	}

	return nil
}

func (p Passport) validateEcl() error {
	values := []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}

	for _, v := range values {
		if p.Ecl == v {
			return nil
		}
	}

	return errInvalidEcl
}

func (p Passport) validatePid() error {
	if len(p.Pid) != 9 {
		return errInvalidPid
	}

	if _, err := strconv.Atoi(p.Pid); err != nil {
		return errInvalidPid
	}

	return nil
}

// ignored by design!
func (p Passport) validateCid() error {
	return nil
}

func part2Validator(p Passport) bool {
	validatorFuncs := []func() error{
		p.validateByr,
		p.validateIyr,
		p.validateEyr,
		p.validateHgt,
		p.validateHcl,
		p.validateEcl,
		p.validatePid,
		p.validateCid,
	}

	for _, fn := range validatorFuncs {
		if err := fn(); err != nil {
			// log.Printf("func %s errored with %s", runtime.FuncForPC(reflect.ValueOf(fn).Pointer()).Name(), err)
			return false
		}
	}

	return true
}

func checkPassports(validator func(Passport) bool) (int, int) {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("could not open file: %s", err)
	}
	defer f.Close()

	var passports []string

	scanner := bufio.NewScanner(f)
	ppBuilder := strings.Builder{}

	for scanner.Scan() {
		line := scanner.Text()

		if line != "" {
			ppBuilder.WriteString(line)
			ppBuilder.WriteString(" ")
			continue
		}

		passports = append(passports, strings.TrimSpace(ppBuilder.String()))
		ppBuilder.Reset()
	}

	// scanner may exit if file does not end in newline, count this as a passport
	passports = append(passports, strings.TrimSpace(ppBuilder.String()))

	var valid, total int

	for _, passport := range passports {
		p := Passport{}
		if err := p.Unmarshal([]byte(passport)); err != nil {
			log.Fatalf("could not unmarshal passport: %s", err)
		}

		if validator(p) {
			valid++
		}

		total++
	}

	return valid, total
}

func main() {
	valid, total := checkPassports(part1Validator)
	log.Printf("%d / %d passports are valid", valid, total)

	valid, total = checkPassports(part2Validator)
	log.Printf("%d / %d passports are valid", valid, total)
}

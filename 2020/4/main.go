package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Passport struct {
	byr string
	iyr string
	eyr string
	hgt string
	hcl string
	ecl string
	pid string
	cid string
}

func NewPassportFromMap(m map[string]string) Passport {
	return Passport{
		byr: m["byr"],
		iyr: m["iyr"],
		eyr: m["eyr"],
		hgt: m["hgt"],
		hcl: m["hcl"],
		ecl: m["ecl"],
		pid: m["pid"],
		cid: m["cid"],
	}
}

func (p Passport) HasRequiredFields() bool {
	return (p.byr != "" &&
		p.iyr != "" &&
		p.eyr != "" &&
		p.hgt != "" &&
		p.hcl != "" &&
		p.ecl != "" &&
		p.pid != "")
}

func isHexDigit(b byte) bool {
	return ('0' <= b && b <= '9') || ('a' <= b && b <= 'f')
}

func isDigit(b byte) bool {
	return '0' <= b && b <= '9'
}

func validateYear(year string, lower, upper int) bool {
	year_num, err := strconv.Atoi(year)
	if err != nil {
		return false
	}
	if lower > year_num || year_num > upper {
		return false
	}
	return true
}

func validateHeight(height string) bool {
	num_idx := 0
	for i := 0; i < len(height); i++ {
		if isDigit(height[i]) {
			num_idx = i
		} else {
			break
		}
	}
	if num_idx == 0 {
		return false
	}
	height_num, err := strconv.Atoi(height[:num_idx+1])
	if err != nil {
		return false
	}
	unit := height[num_idx+1:]
	if unit == "in" {
		if !(59 <= height_num && height_num <= 76) {
			return false
		}
	} else if unit == "cm" {
		if !(150 <= height_num && height_num <= 193) {
			return false
		}
	} else {
		return false
	}
	return true
}

func validateHairColor(hair_color string) bool {
	if !(len(hair_color) == 7 && hair_color[0] == '#') {
		return false
	}
	for i := 1; i < len(hair_color); i++ {
		if !isHexDigit(hair_color[i]) {
			return false
		}
	}
	return true
}

func validateEyeColor(eye_color string) bool {
	allowed := []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
	for _, v := range allowed {
		if eye_color == v {
			return true
		}
	}
	return false
}

func validatePassportId(id string) bool {
	if len(id) != 9 {
		return false
	}
	for i := 0; i < len(id); i++ {
		if !isDigit(id[i]) {
			return false
		}
	}
	return true
}

func (p Passport) IsValid() bool {
	return (p.HasRequiredFields() &&
		validateYear(p.byr, 1920, 2002) &&
		validateYear(p.iyr, 2010, 2020) &&
		validateYear(p.eyr, 2020, 2030) &&
		validateHeight(p.hgt) &&
		validateHairColor(p.hcl) &&
		validateEyeColor(p.ecl) &&
		validatePassportId(p.pid))
}

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	var passports []Passport
	fields := make(map[string]string)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			passports = append(passports, NewPassportFromMap(fields))
			clear(fields)
		} else {
			line_fields := strings.Fields(line)
			for _, field := range line_fields {
				field_arr := strings.Split(field, ":")
				if len(field_arr) != 2 {
					log.Fatal("Unexpected field format!")
				}
				key, value := field_arr[0], field_arr[1]
				fields[key] = value
			}
		}
	}
	passports = append(passports, NewPassportFromMap(fields))

	rf, vp := 0, 0
	for _, p := range passports {
		if p.HasRequiredFields() {
			rf++
		}
		if p.IsValid() {
			vp++
		}
	}

	fmt.Println("[1] passports with required fields count: ", rf)
	fmt.Println("[2] valid passports count: ", vp)
}

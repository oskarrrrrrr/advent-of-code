package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Rule struct {
	source     string
	contains   []string
	multiplies []int
}

func newRule(s string) Rule {
	fields := strings.Fields(s)
	var source string
	var i int
	for ; ; i++ {
		if fields[i] == "bags" {
			break
		}
		if source != "" {
			source += " "
		}
		source += fields[i]
	}
	if fields[i] != "bags" || fields[i+1] != "contain" {
		panic("Expected 'bags contain' after source color definition.")
	}
	i += 2
	if fields[i] == "no" {
		return Rule{source: source}
	} else {
		var contains []string
		var multiplies []int
		for i < len(fields) && (fields[i] != "bags." && fields[i] != "bag.") {
			m, err := strconv.Atoi(fields[i])
			if err != nil {
				log.Fatal("Failed to parse number!")
			}
			multiplies = append(multiplies, m)
			i++
			var color string
			for ; ; i++ {
				if fields[i] == "bag," ||
					fields[i] == "bags," ||
					fields[i] == "bag." ||
					fields[i] == "bags." {
					break
				} else {
					if color != "" {
						color += " "
					}
					color += fields[i]
				}
			}
			contains = append(contains, color)
			i++
		}
		return Rule{
			source:     source,
			contains:   contains,
			multiplies: multiplies,
		}
	}
}

func can_reach_shiny_gold_iter(rules []Rule, can_reach_shiny_gold map[string]bool) bool {
	any_change := false
	for _, rule := range rules {
		if _, exists := can_reach_shiny_gold[rule.source]; !exists {
			any_change = true
			if rule.source == "shiny gold" {
				can_reach_shiny_gold[rule.source] = true
				continue
			}
			can_reach_shiny_gold[rule.source] = false
		}
		for _, color := range rule.contains {
			if ok, _ := can_reach_shiny_gold[color]; ok {
				can_reach, exists := can_reach_shiny_gold[rule.source]
				if !exists || !can_reach {
					can_reach_shiny_gold[rule.source] = true
					any_change = true
				}
			}
		}
	}
	return any_change
}

func find_rule(rules []Rule, source string) *Rule {
	for _, rule := range rules {
		if rule.source == source {
			return &rule
		}
	}
	log.Fatal("Could not find rule for bag ", source)
	return nil
}

func count_bags(rule *Rule, rules []Rule) int {
	sum := 1
	for i := 0; i < len(rule.contains); i++ {
		color_name := rule.contains[i]
		multiple := rule.multiplies[i]
		rule := find_rule(rules, color_name)
		sum += multiple * count_bags(rule, rules)
	}
	return sum
}

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var rules []Rule
	for scanner.Scan() {
		rules = append(rules, newRule(scanner.Text()))
	}

	can_reach_shiny_gold := make(map[string]bool)
	for can_reach_shiny_gold_iter(rules, can_reach_shiny_gold) {
	}
	count := -1
	for _, can_reach := range can_reach_shiny_gold {
		if can_reach {
			count++
		}
	}
	fmt.Println("[1]", count)

	shiny_gold := find_rule(rules, "shiny gold")
	bags := count_bags(shiny_gold, rules) - 1
	fmt.Println("[2]", bags)
}

package day05

import (
	"aoc/src/utils"
	"sort"
	"strconv"
	"strings"
)

type SafetyManual struct {
	pages [][]string
	rules []Rule
}

type Rule struct {
	first string
	last  string
}

func (s SafetyManual) getIndexes(update []string, page string) []int {
	indexes := []int{}
	for idx, otherPage := range update {
		if page == otherPage {
			indexes = append(indexes, idx)
		}
	}

	return indexes
}

func (s SafetyManual) checkRule(update []string) bool {
	for pageIndex, page := range update {
		for _, rule := range s.rules {
			if page == rule.last {
				indexes := s.getIndexes(update, rule.first)

				for _, otherPageIndex := range indexes {
					if otherPageIndex > pageIndex {
						return false
					}
				}
			}
		}
	}

	return true
}

func (s SafetyManual) checkRules() ([][]string, [][]string) {
	printable := [][]string{}
	rejected := [][]string{}
	for _, update := range s.pages {
		if s.checkRule(update) {
			printable = append(printable, update)
		} else {
			rejected = append(rejected, update)
		}
	}

	return printable, rejected
}

func (s SafetyManual) correctRejected(rejected [][]string) [][]string {
	for _, update := range rejected {
		sort.Slice(update, func(i, j int) bool {
			prev, next := update[i], update[j]

			for _, rule := range s.rules {
				if rule.first == next && rule.last == prev {
					return false
				}
			}

			return true
		})
	}

	return rejected
}

func SolveProblem1() (string, error) {
	input := strings.Split(utils.GetInput(5), "\n\n")

	safetyManual := SafetyManual{pages: [][]string{}, rules: []Rule{}}

	for _, line := range strings.Split(input[0], "\n") {
		entry := strings.Split(line, "|")
		safetyManual.rules = append(safetyManual.rules, Rule{first: entry[0], last: entry[1]})
	}

	for _, update := range strings.Split(input[1], "\n") {
		if update == "" {
			continue
		}

		safetyManual.pages = append(safetyManual.pages, strings.Split(update, ","))
	}

	sum := 0
	printable, _ := safetyManual.checkRules()
	for _, v := range printable {
		mid := v[int(len(v)/2)]
		s, err := strconv.Atoi(mid)
		if err != nil {
			panic(err)
		}

		sum += s
	}

	return strconv.Itoa(sum), nil
}

func SolveProblem2() (string, error) {
	input := strings.Split(utils.GetInput(5), "\n\n")

	safetyManual := SafetyManual{pages: [][]string{}, rules: []Rule{}}

	for _, line := range strings.Split(input[0], "\n") {
		entry := strings.Split(line, "|")
		safetyManual.rules = append(safetyManual.rules, Rule{first: entry[0], last: entry[1]})
	}

	for _, update := range strings.Split(input[1], "\n") {
		if update == "" {
			continue
		}

		safetyManual.pages = append(safetyManual.pages, strings.Split(update, ","))
	}

	_, rejected := safetyManual.checkRules()
	rejected = safetyManual.correctRejected(rejected)

	sum := 0
	for _, v := range rejected {
		mid := v[int(len(v)/2)]
		s, err := strconv.Atoi(mid)
		if err != nil {
			panic(err)
		}

		sum += s
	}

	return strconv.Itoa(sum), nil
}

package main

import (
	_ "embed"
	"fmt"
	"sort"
	"strings"
)

//go:embed input.txt
var Rucksacks string

func separateRucksacks() []string {
	return strings.Split(Rucksacks, "\n")
}

func processRucksacks() int {
	var commonItems string
	rucks := separateRucksacks()

	for _, ruck := range rucks {
		commonItems = commonItems + findCommonItemsInRuck(ruck)
	}

	return tabulateTotalPriority(commonItems)
}

func findBadgeType() string {
	var badgeTypes string

	rucks := separateRucksacks()

	ruckGroups := make(map[int][]string)
	ruckGroupIndex := 0
	for _, ruck := range rucks {
		ruckGroups[ruckGroupIndex] = append(ruckGroups[ruckGroupIndex], ruck)

		if len(ruckGroups[ruckGroupIndex]) == 3 {
			ruckGroupIndex++
		}
	}

	for _, ruckGroup := range ruckGroups {
		badgeTypes = badgeTypes + findCommonItemInRucks(ruckGroup)
	}

	return badgeTypes
}

func findCommonItemInRucks(rucks []string) string {
	var commonItem string

	sort.Slice(rucks, func(i, j int) bool {
		return len(rucks[i]) < len(rucks[j])
	})

	for _, item := range rucks[0] {
		if strings.Contains(rucks[1], string(item)) && strings.Contains(rucks[2], string(item)) {
			commonItem = string(item)
		}
	}

	return commonItem
}

func findCommonItemsInRuck(ruck string) string {
	var items string

	divider := len(ruck) / 2
	compartmentA := ruck[:divider]
	compartmentB := ruck[divider:]

	for _, item := range compartmentB {
		if isPresent := strings.Contains(compartmentA, string(item)); isPresent && !strings.Contains(items, string(item)) {
			items = items + string(item)
		}
	}

	return items
}

func tabulateTotalPriority(itemList string) int {
	var total int

	for _, item := range itemList {
		total += priorityList[string(item)]
	}

	return total
}

func PartOne() int {
	return processRucksacks()
}

func PartTwo() int {
	return tabulateTotalPriority(findBadgeType())
}

func main() {
	fmt.Println("Part One:", PartOne())
	fmt.Println("Part Two:", PartTwo())
}

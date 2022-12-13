package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var input = [][]string{
	{},
	{"Q", "M", "G", "C", "L"},
	{"R", "D", "L", "C", "T", "F", "H", "G"},
	{"V", "J", "F", "N", "M", "T", "W", "R"},
	{"J", "F", "D", "V", "Q", "P"},
	{"N", "F", "M", "S", "L", "B", "T"},
	{"R", "N", "V", "H", "C", "D", "P"},
	{"H", "C", "T"},
	{"G", "S", "J", "V", "Z", "N", "H", "P"},
	{"Z", "F", "H", "G"},
}

type CargoCrane struct {
	Stacks [][]string
	Moves  *os.File
	inTransit
	orders []order
}

type order struct {
	amount, source, destination int
}

type inTransit struct {
	from, to int
	cargo    []string
}

func (c *CargoCrane) convertMovesToOrders() {
	scanner := bufio.NewScanner(c.Moves)

	for scanner.Scan() {
		moveOrder := scanner.Text()
		re := regexp.MustCompile(`\d+`)
		preConvertedOrder := re.FindAllString(moveOrder, -1)

		var convertedOrders []int
		for _, assignment := range preConvertedOrder {
			convertedString, _ := strconv.Atoi(assignment)
			convertedOrders = append(convertedOrders, convertedString)
		}

		c.orders = append(c.orders, order{
			amount:      convertedOrders[0],
			source:      convertedOrders[1],
			destination: convertedOrders[2],
		})
	}
}

func (c *CargoCrane) processOrders() {
	for _, order := range c.orders {
		c.from(order.source).move(order.amount).to(order.destination)
		c.reset()
	}
}

func (c *CargoCrane) from(source int) *CargoCrane {
	c.inTransit.from = source
	return c
}

func (c *CargoCrane) move(amount int) *CargoCrane {
	amt := len(c.Stacks[c.inTransit.from]) - amount
	c.inTransit.cargo = c.Stacks[c.inTransit.from][amt:]
	c.Stacks[c.inTransit.from] = c.Stacks[c.inTransit.from][:amt]

	return c
}

func (c *CargoCrane) to(destination int) *CargoCrane {
	total := len(c.inTransit.cargo)

	for i := total - 1; i >= 0; i-- {
		c.Stacks[destination] = append(c.Stacks[destination], c.inTransit.cargo[i])
	}

	return c
}

func (c *CargoCrane) reset() {
	c.inTransit = inTransit{}
}

func (c *CargoCrane) getPlacements() string {
	var finalPlacements []string

	stacksTotal := len(c.Stacks) - 1

	for i := 1; i <= stacksTotal; i++ {
		topmostCargo := len(c.Stacks[i]) - 1

		finalPlacements = append(finalPlacements, c.Stacks[i][topmostCargo])
	}

	return strings.Join(finalPlacements[:], "")
}

func PartOne(loadingDock CargoCrane) string {
	loadingDock.convertMovesToOrders()
	loadingDock.processOrders()
	return loadingDock.getPlacements()
}

func main() {
	orders, _ := os.Open("input_commands")
	defer orders.Close()

	toSort := CargoCrane{
		Stacks: input,
		Moves:  orders,
	}

	fmt.Println("PartOne:", PartOne(toSort))
}

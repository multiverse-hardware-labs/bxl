package bxlparser

import (
	"strconv"
	"strings"
	"unicode"
)

// CompPin BXL Component pin
type CompPin struct {
	Number       int
	Name         string
	PartNum      int
	SymbolPinNum int
	GateEq       int
	PinEq        int
	PinType      string
	Side         string
	Group        int
	InnerGraphic string
	OuterGraphic string
}

// FindCompPins FindCompPins
func (c *Component) FindCompPins() {
	var a, b int
	for i, l := range c.data {
		if strings.HasPrefix(l, "CompPins") {
			a = i + 1
		}
		if strings.HasPrefix(l, "EndCompPins") {
			b = i
		}
	}
	for _, l := range c.data[a:b] {
		c.parseCompPin(&l)
	}
}

func (c *Component) parseCompPin(l *string) {
	var p CompPin
	f := func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c) || c == '(' || c == ')' || c == '"'
	}

	fields := strings.FieldsFunc(*l, f)
	for i, s := range fields {
		switch s {
		case "CompPin":
			p.Number, _ = strconv.Atoi(fields[i+1])
			p.Name = fields[i+2]
		case "PartNum":
			p.PartNum, _ = strconv.Atoi(fields[i+1])
		case "SymPinNum":
			p.SymbolPinNum, _ = strconv.Atoi(fields[i+1])
		case "GateEq":
			p.GateEq, _ = strconv.Atoi(fields[i+1])
		case "Pineq":
			p.PinEq, _ = strconv.Atoi(fields[i+1])
		case "PinType":
			p.PinType = fields[i+1]
		case "Side":
			p.Side = fields[i+1]
		case "Group":
			p.Group, _ = strconv.Atoi(fields[i+1])
		case "InnerGraphic":
			p.InnerGraphic = fields[i+1]
		case "OuterGraphic":
			p.OuterGraphic = fields[i+1]
		}
	}
}

// FindPinMap Find Pin map
func (c *Component) FindPinMap() {
	var a, b int
	for i, l := range c.data {
		if strings.HasPrefix(l, "PinMap") {
			a = i + 1
		}
		if strings.HasPrefix(l, "EndPinMap") {
			b = i
		}
	}

	f := func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c) || c == '(' || c == ')' || c == '"'
	}

	for _, l := range c.data[a:b] {
		fields := strings.FieldsFunc(l, f)
		padnum, _ := strconv.Atoi(fields[1])
		c.PinMap[padnum], _ = strconv.Atoi(fields[3])
	}
}
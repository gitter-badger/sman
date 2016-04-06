package sman

import (
	//"fmt"
	"github.com/fatih/color"
	"strconv"
	"strings"
)

// Placeholder struct
type Placeholder struct {
	Name, Desc, Input string
	Options           []string
	Patterns          []string
}

func ParseOptions(in string) (options []string) {
	split := strings.Split(in, ",")
	// Loop Through Options, and check for escaped comma
	var toAppend string
	for _, o := range split {
		if o[len(o)-1:] == `\` {
			toAppend += o[:len(o)-1]
			toAppend += ","
			continue
		} else {
			toAppend += o
		}
		options = append(options, toAppend)
		toAppend = ""
	}
	return options
}

func SearchPlaceholder(in []Placeholder, n string) (i int, ok bool) {
	for i, p := range in {
		if p.Name == n {
			return i, true
		}
	}
	return i, false
}

func (p *Placeholder) DisplayName() string {
	c := color.New(color.FgCyan).SprintFunc()
	return c("[", p.Name, "]")
}

func (p *Placeholder) DisplayOptions() string {
	if len(p.Options) > 0 {
		return strings.Join(p.Options, " | ")
	}
	return "----"
}

func (p *Placeholder) DisplayDesc() string {
	if len(p.Desc) > 0 {
		return strings.Title(p.Desc)
	}
	return "----"
}

func (p *Placeholder) AddPattern(pattern string) {
	if !SliceContains(p.Patterns, pattern) {
		p.Patterns = append(p.Patterns, pattern)
	}
}

func (p *Placeholder) SetInput(input string) {
	if len(p.Options) != 0 {
		if len(input) == 0 {
			input = p.Options[0]
		} else if i, err := strconv.Atoi(input); err == nil {
			if (i > 0) && (i <= len(p.Options)) {
				i--
				input = p.Options[i]
			}
		}
	}
	p.Input = input
}
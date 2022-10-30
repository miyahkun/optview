package main

import "miyahkun/optview/sample"

type ParseState struct {
	buffer []rune
	point  int
}

func (p *ParseState) at() rune {
	var ret rune
	l := len(p.buffer)
	if l <= p.point {
		ret = 0
	} else {
		ret = p.buffer[p.point]
	}
	return ret
}

// func (p *ParseState) next() rune {
// 	c := p.at()
// 	if c == 0 {
// 		return c
// 	} else if c == '\n' {

// 	}
// }

// func (p *ParseState) skip() {
// 	for c := p.at(); unicode.IsSpace(c); c = p.at() {
// 		p.next()
// 	}
// }

// func (p *ParseState) Parse(text string) []string {

// }

func NewParseState(buffer string) ParseState {
	sample.Sample()
	return ParseState{
		buffer: []rune(buffer),
	}
}

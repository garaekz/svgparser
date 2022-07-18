package svgparser_test

import (
	"encoding/xml"
	"strings"

	"github.com/garaekz/svgparser"
)

func element(name xml.Name, attrs []xml.Attr) *svgparser.Element {
	return &svgparser.Element{
		Name:       name,
		Attributes: attrs,
		Children:   []*svgparser.Element{},
	}
}

func parse(svg string, validate bool) (*svgparser.Element, error) {
	element, err := svgparser.Parse(strings.NewReader(svg), validate)
	return element, err
}

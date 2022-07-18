package svgparser_test

import (
	"encoding/xml"
	"testing"

	"github.com/garaekz/svgparser"
)

func testElement() *svgparser.Element {
	svg := `
		<svg width="1000" height="600">
			<g id="first">
				<rect width="5" height="3" id="inFirst"/>
				<rect width="5" height="2" id="inFirst"/>
			</g>
			<g id="second">
				<path d="M50 50 Q50 100 100 100"/>
				<rect width="5" height="1"/>
			</g>
		</svg>
	`
	element, _ := parse(svg, false)
	return element
}

func equals(t *testing.T, name string, expected, actual *svgparser.Element) {
	if !(expected == actual || expected.Compare(actual)) {
		t.Errorf("%s: expected %v, actual %v\n", name, expected, actual)
	}
}

func equalSlices(t *testing.T, name string, expected, actual []*svgparser.Element) {
	if len(expected) != len(actual) {
		t.Errorf("%s: expected %v, actual %v\n", name, expected, actual)
		return
	}

	for i, r := range actual {
		equals(t, name, expected[i], r)
	}
}

func TestFindAll(t *testing.T) {
	svgElement := testElement()

	equalSlices(t, "Find", []*svgparser.Element{
		element(xml.Name{Space: "", Local: "rect"}, []xml.Attr{
			{Name: xml.Name{Space: "", Local: "width"}, Value: "5"},
			{Name: xml.Name{Space: "", Local: "height"}, Value: "3"},
			{Name: xml.Name{Space: "", Local: "id"}, Value: "inFirst"},
		}),
		element(xml.Name{Space: "", Local: "rect"}, []xml.Attr{
			{Name: xml.Name{Space: "", Local: "width"}, Value: "5"},
			{Name: xml.Name{Space: "", Local: "height"}, Value: "2"},
			{Name: xml.Name{Space: "", Local: "id"}, Value: "inFirst"},
		}),
		element(xml.Name{Space: "", Local: "rect"}, []xml.Attr{
			{Name: xml.Name{Space: "", Local: "width"}, Value: "5"},
			{Name: xml.Name{Space: "", Local: "height"}, Value: "1"},
		}),
	}, svgElement.FindAll("rect"))

	equalSlices(t, "Find", []*svgparser.Element{}, svgElement.FindAll("circle"))
}

func TestFindID(t *testing.T) {
	svgElement := testElement()

	equals(t, "Find", &svgparser.Element{
		Name: xml.Name{Space: "", Local: "g"},
		Attributes: []xml.Attr{
			{Name: xml.Name{Space: "", Local: "id"}, Value: "second"},
		},
		Children: []*svgparser.Element{
			element(xml.Name{Space: "", Local: "path"}, []xml.Attr{
				{Name: xml.Name{Space: "", Local: "d"}, Value: "M50 50 Q50 100 100 100"},
			}),
			element(xml.Name{Space: "", Local: "rect"}, []xml.Attr{
				{Name: xml.Name{Space: "", Local: "width"}, Value: "5"},
				{Name: xml.Name{Space: "", Local: "height"}, Value: "1"},
			}),
		},
	},
		svgElement.FindID("second"))

	equals(t, "Find",
		element(xml.Name{Space: "", Local: "rect"}, []xml.Attr{
			{Name: xml.Name{Space: "", Local: "width"}, Value: "5"},
			{Name: xml.Name{Space: "", Local: "height"}, Value: "3"},
			{Name: xml.Name{Space: "", Local: "id"}, Value: "inFirst"},
		}),
		svgElement.FindID("inFirst"),
	)

	equals(t, "Find", nil, svgElement.FindID("missing"))
}

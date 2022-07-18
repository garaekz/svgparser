package svgparser_test

import (
	"encoding/xml"
	"testing"

	"github.com/garaekz/svgparser"
)

func TestParser(t *testing.T) {
	var testCases = []struct {
		svg     string
		element svgparser.Element
	}{
		{
			`
		<svg width="100" height="100">
			<circle cx="50" cy="50" r="40" fill="red" />
		</svg>
		`,
			svgparser.Element{
				Name: xml.Name{Space: "", Local: "svg"},
				Attributes: []xml.Attr{
					{Name: xml.Name{Space: "", Local: "width"}, Value: "100"},
					{Name: xml.Name{Space: "", Local: "height"}, Value: "100"},
				},
				Children: []*svgparser.Element{
					element(xml.Name{Space: "", Local: "circle"},
						[]xml.Attr{
							{Name: xml.Name{Space: "", Local: "cx"}, Value: "50"},
							{Name: xml.Name{Space: "", Local: "cy"}, Value: "50"},
							{Name: xml.Name{Space: "", Local: "r"}, Value: "40"},
							{Name: xml.Name{Space: "", Local: "fill"}, Value: "red"},
						},
					),
				},
			},
		},
		{
			`
		<svg height="400" width="450">
			<g stroke="black" stroke-width="3" fill="black">
				<path id="AB" d="M 100 350 L 150 -300" stroke="red" />
				<path id="BC" d="M 250 50 L 150 300" stroke="red" />
				<path d="M 175 200 L 150 0" stroke="green" />
			</g>
		</svg>
		`,
			svgparser.Element{
				Name: xml.Name{Space: "", Local: "svg"},
				Attributes: []xml.Attr{
					{Name: xml.Name{Space: "", Local: "height"}, Value: "400"},
					{Name: xml.Name{Space: "", Local: "width"}, Value: "450"},
				},
				Children: []*svgparser.Element{
					{
						Name: xml.Name{Space: "", Local: "g"},
						Attributes: []xml.Attr{
							{Name: xml.Name{Space: "", Local: "stroke"}, Value: "black"},
							{Name: xml.Name{Space: "", Local: "stroke-width"}, Value: "3"},
							{Name: xml.Name{Space: "", Local: "fill"}, Value: "black"},
						},
						Children: []*svgparser.Element{
							element(xml.Name{Space: "", Local: "path"}, []xml.Attr{
								{Name: xml.Name{Space: "", Local: "id"}, Value: "AB"},
								{Name: xml.Name{Space: "", Local: "d"}, Value: "M 100 350 L 150 -300"},
								{Name: xml.Name{Space: "", Local: "stroke"}, Value: "red"},
							}),
							element(xml.Name{Space: "", Local: "path"}, []xml.Attr{
								{Name: xml.Name{Space: "", Local: "id"}, Value: "BC"},
								{Name: xml.Name{Space: "", Local: "d"}, Value: "M 250 50 L 150 300"},
								{Name: xml.Name{Space: "", Local: "stroke"}, Value: "red"},
							}),
							element(xml.Name{Space: "", Local: "path"}, []xml.Attr{
								{Name: xml.Name{Space: "", Local: "d"}, Value: "M 175 200 L 150 0"},
								{Name: xml.Name{Space: "", Local: "stroke"}, Value: "green"},
							}),
						},
					},
				},
			},
		},
		{
			"",
			svgparser.Element{},
		},
	}

	for _, test := range testCases {
		actual, err := parse(test.svg, false)

		if !(test.element.Compare(actual) && err == nil) {
			t.Errorf("Parse: expected %v, actual %v\n", test.element, actual)
		}
	}
}

func TestValidDocument(t *testing.T) {
	svg := `
		<svg xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" id="svg-root" width="100%" height="100%" viewBox="0 0 480 360">
			<title id="test-title">color-prop-01-b</title>
			<desc id="test-desc">Test that viewer has the basic capability to process the color property</desc>
			<rect id="test-frame" x="1" y="1" width="478" height="358" fill="none" stroke="#000000"/>
		</svg>
		`

	element, err := parse(svg, true)
	if !(element != nil && err == nil) {
		t.Errorf("Validation: expected %v, actual %v\n", nil, err)
	}
}

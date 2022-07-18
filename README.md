# svgparser

Library for parsing and manipulating SVG files. We've made some changes to original repo and soon we will update this, sorry in advance.

### Installation

	go get github.com/garaekz/svgparser

### Features

##### Validation
Checks if the SVG input is valid according to the [W3C Recommendation](https://www.w3.org/TR/SVG/Overview.html).

##### Find functionality
Provides capability to search for SVG elements by id or element name.

##### Path Parser
Parsing the 'd' attribute of a path element into a structure containing all subpaths with their commands and parameters.

##### Style Parser
Parsing the value of a style element.

### Example

	func ExampleParse() {
		svg := `
			<svg width="100" height="100">
				<circle cx="50" cy="50" r="40" fill="red" />
			</svg>
		`
		reader := strings.NewReader(svg)

		element, _ := svgparser.Parse(reader, false)

		fmt.Printf("SVG width: %s", element.GetAttribute("width"))
		fmt.Printf("Circle fill: %s", element.Children[0].GetAttribute("fill"))

		// Output:
		// SVG width: 100
		// Circle fill: red
	}

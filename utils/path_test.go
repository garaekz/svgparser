package utils_test

import (
	"testing"

	"github.com/garaekz/svgparser/utils"
)

func TestPathParser(t *testing.T) {
	var testCases = []struct {
		d        string
		expected *utils.Path
	}{
		{
			"M 10,20 L 30,30 Z",
			&utils.Path{
				[]*utils.Subpath{
					&utils.Subpath{
						Commands: []*utils.Command{
							&utils.Command{Symbol: "M", Params: []float64{10, 20}},
							&utils.Command{Symbol: "L", Params: []float64{30, 30}},
							&utils.Command{Symbol: "Z", Params: []float64{}},
						},
					},
				},
			},
		},
		{
			"M .2.3 L 30,30 Z",
			&utils.Path{
				[]*utils.Subpath{
					&utils.Subpath{
						Commands: []*utils.Command{
							&utils.Command{Symbol: "M", Params: []float64{0.2, 0.3}},
							&utils.Command{Symbol: "L", Params: []float64{30, 30}},
							&utils.Command{Symbol: "Z", Params: []float64{}},
						},
					},
				},
			},
		},
		{
			"M10-20 L30,30 Z",
			&utils.Path{
				[]*utils.Subpath{
					&utils.Subpath{
						Commands: []*utils.Command{
							&utils.Command{Symbol: "M", Params: []float64{10, -20}},
							&utils.Command{Symbol: "L", Params: []float64{30, 30}},
							&utils.Command{Symbol: "Z", Params: []float64{}},
						},
					},
				},
			},
		},
		{
			"M 10-20 L 30,30 L 40,40 Z",
			&utils.Path{
				[]*utils.Subpath{
					&utils.Subpath{
						Commands: []*utils.Command{
							&utils.Command{Symbol: "M", Params: []float64{10, -20}},
							&utils.Command{Symbol: "L", Params: []float64{30, 30}},
							&utils.Command{Symbol: "L", Params: []float64{40, 40}},
							&utils.Command{Symbol: "Z", Params: []float64{}},
						},
					},
				},
			},
		},
		{
			"M10,20 L20,30 L10,20",
			&utils.Path{
				[]*utils.Subpath{
					&utils.Subpath{
						Commands: []*utils.Command{
							&utils.Command{Symbol: "M", Params: []float64{10, 20}},
							&utils.Command{Symbol: "L", Params: []float64{20, 30}},
							&utils.Command{Symbol: "L", Params: []float64{10, 20}},
						},
					},
				},
			},
		},
	}
	for _, test := range testCases {
		path, err := utils.PathParser(test.d)
		if !(test.expected.Compare(path) && err == nil) {
			t.Errorf("Path: expected %v, actual %v\n", test.expected, path)
		}
	}
}

func TestParamNumberInPath(t *testing.T) {
	path, err := utils.PathParser("M 10 20 30 Z")
	expectedError := "Incorrect number of parameters for M"

	if !(path == nil && err.Error() == expectedError) {
		t.Errorf("Path: expected %v, actual %v\n", expectedError, err)
	}
}

func TestMissingZero(t *testing.T) {
	var testCases = []struct {
		d        string
		expected *utils.Path
	}{
		{
			"M 0.2 0.3 L 30,30 Z",
			&utils.Path{
				[]*utils.Subpath{
					&utils.Subpath{
						Commands: []*utils.Command{
							&utils.Command{Symbol: "M", Params: []float64{0.2, 0.3}},
							&utils.Command{Symbol: "L", Params: []float64{30, 30}},
							&utils.Command{Symbol: "Z", Params: []float64{}},
						},
					},
				},
			},
		},
	}

	for _, test := range testCases {
		path, err := utils.PathParser(test.d)
		if !(test.expected.Compare(path) && err == nil) {
			t.Errorf("Path: expected %v, actual %v\n", test.expected, path)
		}
	}
}

func TestTwoSubpaths(t *testing.T) {
	var testCases = []struct {
		d        string
		expected *utils.Path
	}{
		{
			"M25,0 L0,30 L50,50 Z m 10, 10 L50,50 l10,0 Z",
			&utils.Path{
				[]*utils.Subpath{
					&utils.Subpath{
						Commands: []*utils.Command{
							&utils.Command{Symbol: "M", Params: []float64{25, 0}},
							&utils.Command{Symbol: "L", Params: []float64{0, 30}},
							&utils.Command{Symbol: "L", Params: []float64{50, 50}},
							&utils.Command{Symbol: "Z", Params: []float64{}},
						},
					},
					&utils.Subpath{
						Commands: []*utils.Command{
							&utils.Command{Symbol: "m", Params: []float64{10, 10}},
							&utils.Command{Symbol: "L", Params: []float64{50, 50}},
							&utils.Command{Symbol: "l", Params: []float64{10, 0}},
							&utils.Command{Symbol: "Z", Params: []float64{}},
						},
					},
				},
			},
		},
	}

	for _, test := range testCases {
		path, err := utils.PathParser(test.d)
		if len(test.expected.Subpaths) != len(path.Subpaths) {
			t.Errorf("Incorrect number of subpaths found")
		}

		if !(test.expected.Compare(path) && err == nil) {
			t.Errorf("Path: expected %v, actual %v\n", *(test.expected), *path)
		}
	}
}

func TestImplicitLineCommands(t *testing.T) {
	var testCases = []struct {
		d        string
		expected *utils.Path
	}{
		{
			"M 10,20 30,40 Z m 10,20 30,40 Z",
			&utils.Path{
				[]*utils.Subpath{
					&utils.Subpath{
						Commands: []*utils.Command{
							&utils.Command{Symbol: "M", Params: []float64{10, 20}},
							&utils.Command{Symbol: "L", Params: []float64{30, 40}},
							&utils.Command{Symbol: "Z", Params: []float64{}},
						},
					},
					&utils.Subpath{
						Commands: []*utils.Command{
							&utils.Command{Symbol: "m", Params: []float64{10, 20}},
							&utils.Command{Symbol: "l", Params: []float64{30, 40}},
							&utils.Command{Symbol: "Z", Params: []float64{}},
						},
					},
				},
			},
		},
	}

	for _, test := range testCases {
		path, err := utils.PathParser(test.d)
		if !(test.expected.Compare(path) && err == nil) {
			t.Errorf("Path: expected %v, actual %v\n", test.expected, path)
		}
	}
}

package handlers

import (
	"assignment-2/fetchers"
	"errors"
	"fmt"
	"math"
	"strings"
)

func castToCoords(polygonsInterface [][][]interface{}) ([][][][]interface{}, error) {
	/*
		The polygon data is either [][][]float64 or [][][][]float64. This is because each land mass
		is an array inside of an array of all the landmasses. But when there is only one land mass,
		it does not need to be IN an array, since it would be the only element.
		Therefore, the when parsed into [][][]interface{} this interface{} is either a float or
		an array. We will need to convert it into [][][][]interface{}.
	*/

	var shapes [][][][]interface{}

	var err error
	err = nil

	switch polygonsInterface[0][0][0].(type) {

	case float64: // The type is [][][]float64
		// Put it into an array
		shapes = make([][][][]interface{}, 1)
		shapes[0] = polygonsInterface

	case interface{}: // The type is [][][][]float64

		// Copy into 'shapes'
		shapes = make([][][][]interface{}, len(polygonsInterface))
		for i := range polygonsInterface {
			shapes[i] = make([][][]interface{}, len(polygonsInterface[i]))
			for j := range polygonsInterface[i] {
				shapes[i][j] = make([][]interface{}, len(polygonsInterface[i][j]))
				for k := range polygonsInterface[i][j] {
					if coords, ok := polygonsInterface[i][j][k].([]interface{}); ok {
						shapes[i][j][k] = make([]interface{}, len(coords))
						copy(shapes[i][j][k], coords)
					}
				}
			}
		}

	default:
		err = errors.New("failed to read necessary data")
	}

	return shapes, err
}

func scaleCoordinates(width float64, height float64, bbox []float64, coordinates []float64) (int, int) {
	x := int((coordinates[0] - bbox[0]) / math.Abs(bbox[0]-bbox[2]) * width)
	y := int((coordinates[1] - bbox[1]) / math.Abs(bbox[1]-bbox[3]) * height)

	return x, y
}

func AsciiMap(iso3 string) (string, error) {
	// Fetch geometric data
	geoJson, err := fetchers.FetchGeoJson(iso3)
	if err != nil {
		return "", fmt.Errorf("generate ascii map: %v", err)
	}
	polygonData := geoJson.Features[0].Geometry.Coordinates

	// Cast the data
	shapes, err := castToCoords(polygonData)
	if err != nil {
		return "", err
	}

	// Find significant landmasses, smaller ones will be ignored
	polygons := make([][][]interface{}, 0)
	maxPolygon := 0
	mainlandInd := 0
	for i, shape := range shapes {
		amt := len(shape[0])

		if amt > maxPolygon {
			mainlandInd = i
			maxPolygon = amt
		}
		if amt > 100 {
			polygons = append(polygons, shapes[i][0])
		}
	}

	// If no landmasses were big enough, the biggest one is added
	if len(polygons) == 0 {
		polygons = append(polygons, shapes[mainlandInd][0])
	}

	// Bounding box
	bbox := geoJson.Features[0].Properties.Bbox
	bboxHeight := math.Abs(bbox[3] - bbox[1])
	bboxWidth := math.Abs(bbox[2] - bbox[0])

	// Scale properly
	width := 150
	height := int((bboxHeight / bboxWidth) * float64(width))

	// Initialize ascii map
	countryMap := make([][]byte, height+1)
	for i := range countryMap {
		countryMap[i] = []byte(strings.Repeat(" ", width+1))
	}

	// Draw the map
	for _, polygon := range polygons {
		for _, point := range polygon {
			// Convert from []interface{} to []float
			coords := make([]float64, 2)
			coords[0] = point[0].(float64)
			coords[1] = point[1].(float64)

			// Scale coordinates to width*height
			x, y := scaleCoordinates(float64(width), float64(height), bbox, coords)
			realy := math.Abs(float64(height - y))
			countryMap[int(realy)][x] = '#'
		}
	}

	// Turn into ascii map
	result := make([]byte, 0)
	for _, row := range countryMap {
		row[len(row)-1] = '\n'
		result = append(result, row...)
	}

	return string(result), nil
}

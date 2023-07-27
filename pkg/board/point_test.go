package board

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPointBox(t *testing.T) {
	point := Point{Row: 0, Col: 0}
	assert.Equal(t, 0, point.BoxIdx())
	for i, boxCenter := range boxCenters {
		assert.Equal(t, i, boxCenter.BoxIdx())
	}
	point = Point{Row: 8, Col: 8}
	assert.Equal(t, 8, point.BoxIdx())
	point = Point{Row: 4, Col: 4}
	assert.Equal(t, 4, point.BoxIdx())
	point = Point{Row: 4, Col: 5}
	assert.Equal(t, 4, point.BoxIdx())
}

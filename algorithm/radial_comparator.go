package algorithm

import (
	"github.com/twpayne/go-geom"
	"github.com/twpayne/go-geom/algorithm/orientation"
	"github.com/twpayne/go-geom/bigplanar"
	"github.com/twpayne/go-geom/sorting"
	"sort"
)

// NewRadialSorting creates an implementation sort.Interface which will sort the wrapped coordinate array
// radially around the focal point.  The comparison is based on the angle and distance
// from the focal point.
// First the angle is checked.
// Counter clockwise indicates a greater value and clockwise indicates a lesser value
// If co-linear then the coordinate nearer to the focalPoint is considered less.
func NewRadialSorting(layout geom.Layout, coordData []float64, focalPoint geom.Coord) sort.Interface {
	isLess := func(v1, v2 []float64) bool {
		orient := bigplanar.OrientationIndex(focalPoint, v1, v2)

		if orient == orientation.CounterClockwise {
			return false
		}
		if orient == orientation.Clockwise {
			return true
		}

		dxp := v1[0] - focalPoint[0]
		dyp := v1[1] - focalPoint[1]
		dxq := v2[0] - focalPoint[0]
		dyq := v2[1] - focalPoint[1]

		// points are collinear - check distance
		op := dxp*dxp + dyp*dyp
		oq := dxq*dxq + dyq*dyq
		if op < oq {
			return true
		}

		return false
	}
	return sorting.NewFlatCoordSorting(layout, coordData, isLess)
}

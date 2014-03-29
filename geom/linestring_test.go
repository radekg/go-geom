package geom

import (
	"math"

	. "launchpad.net/gocheck"
)

type LineStringSuite struct{}

var _ = Suite(&LineStringSuite{})

func (s *LineStringSuite) TestXY(c *C) {

	coords1 := [][]float64{{1, 2}, {3, 4}}
	ls, err := NewLineString(XY, coords1)
	c.Assert(err, IsNil)
	c.Assert(ls, Not(IsNil))

	c.Check(ls.Coords(), DeepEquals, coords1)
	c.Check(ls.Envelope(), DeepEquals, NewEnvelope(1, 2, 3, 4))
	c.Check(ls.Layout(), Equals, XY)
	c.Check(ls.Length(), Equals, math.Sqrt(8))
	c.Check(ls.Stride(), Equals, 2)

	c.Check(ls.Ends(), IsNil)
	c.Check(ls.Endss(), IsNil)
	c.Check(ls.FlatCoords(), DeepEquals, []float64{1, 2, 3, 4})

}

func (s *LineStringSuite) TestXYZ(c *C) {

	coords1 := [][]float64{{1, 2, 3}, {4, 5, 6}}
	ls, err := NewLineString(XYZ, coords1)
	c.Assert(err, IsNil)
	c.Assert(ls, Not(IsNil))

	c.Check(ls.Coords(), DeepEquals, coords1)
	c.Check(ls.Envelope(), DeepEquals, NewEnvelope(1, 2, 3, 4, 5, 6))
	c.Check(ls.Layout(), Equals, XYZ)
	c.Check(ls.Length(), Equals, math.Sqrt(18))
	c.Check(ls.Stride(), Equals, 3)

	c.Check(ls.Ends(), IsNil)
	c.Check(ls.Endss(), IsNil)
	c.Check(ls.FlatCoords(), DeepEquals, []float64{1, 2, 3, 4, 5, 6})

}

func (s *LineStringSuite) TestClone(c *C) {
	ls1, err := NewLineString(XY, [][]float64{{1, 2}, {3, 4}})
	c.Assert(err, IsNil)
	ls2 := ls1.Clone()
	c.Check(ls2, Not(Equals), ls1)
	c.Check(ls2.Coords(), DeepEquals, ls1.Coords())
	c.Check(ls2.Envelope(), DeepEquals, ls1.Envelope())
	c.Check(ls2.FlatCoords(), Not(Aliases), ls1.FlatCoords())
	c.Check(ls2.Layout(), Equals, ls1.Layout())
	c.Check(ls2.Stride(), Equals, ls1.Stride())
}

func (s *LineStringSuite) TestStrideMismatch(c *C) {
	var ls *LineString
	var err error
	ls, err = NewLineString(XY, [][]float64{{1, 2}, {}})
	c.Check(ls, IsNil)
	c.Check(err, DeepEquals, ErrStrideMismatch{Got: 0, Want: 2})
	ls, err = NewLineString(XY, [][]float64{{1, 2}, {3}})
	c.Check(ls, IsNil)
	c.Check(err, DeepEquals, ErrStrideMismatch{Got: 1, Want: 2})
	ls, err = NewLineString(XY, [][]float64{{1, 2}, {3, 4}})
	c.Check(ls, Not(IsNil))
	c.Check(err, IsNil)
	ls, err = NewLineString(XY, [][]float64{{1, 2}, {3, 4, 5}})
	c.Check(ls, IsNil)
	c.Check(err, DeepEquals, ErrStrideMismatch{Got: 3, Want: 2})
}

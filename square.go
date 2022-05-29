package square

type Point struct {
	x, y int
}

type Square struct {
	start Point
	a     uint
}

func (s Square) End() Point {
	var p Point
	p.x = s.start.x + int(s.a)
	p.y = s.start.y + int(s.a)
	return p
}

func (s Square) Area() uint {
	var area uint
	area = s.a * s.a
	return area
}

func (s Square) Perimeter() uint {
	var per uint
	per = 4 * s.a
	return per
}

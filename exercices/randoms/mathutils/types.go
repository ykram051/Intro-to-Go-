package mathutils

type Point struct{
	X int
	Y int
}

func (p Point) GetX() int {
	return p.X
}

func (p *Point) SetX(x int){
	p.X = x
}
package cmd

type Square struct {
	A int
	B int
}

func (s Square) Accept(visitor Visitor) {
	visitor.VisitForSquare(s)
}

package navigator

type Navigator struct {
	Strategy
}

func (n *Navigator) SetStrategy(strategy Strategy) {
	n.Strategy = strategy
}

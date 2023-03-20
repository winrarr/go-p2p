package blockchain

type tree struct {
	nodes   [][]*node
	current *node
	longest int
}

type node struct {
	Slot         int
	Transactions []transaction

	parent   *node
	children []*node
}

func MakeTree() tree {
	genesis := getGenesis()
	return tree{
		nodes:   [][]*node{{genesis}},
		current: genesis,
		longest: 1,
	}
}

func getGenesis() *node {
	return &node{
		Slot:         0,
		Transactions: []transaction{},

		parent:   nil,
		children: []*node{},
	}
}

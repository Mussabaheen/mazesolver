package queue

type Node struct {
	Key        string
	Val        interface{}
	ParentNode *Node
}

type Queue struct {
	Elements []Node
	Size     int
}

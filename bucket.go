package main

import "fmt"

const (
	ACAP = 8
	BCAP = 5
	CCAP = 3
	ENDVALUE = 4
)

type step struct {
	a int
	b int
	c int
}

type node struct {
	s      step
	childs []*node
	parent *node
}

var root *node

func init() {
	s := step{ACAP, 0, 0}

	root = &node{s, make([]*node, 0), nil}
}

func (n *node) checkEqual(s step) bool {
	if n.s.a == s.a &&
		n.s.b == s.b &&
		n.s.c == s.c {
		return true
	}

	return false
}

func findNode(n *node, s step) bool{
	if n == nil {
		return false
	}

	if n.checkEqual(s) {
		return true
	}

	for _, tmp := range n.childs {
		if findNode(tmp, s) {
			return true
		}
	}

	return false
}

func checkEnd(s step) bool{
	if s.a == ENDVALUE ||
		s.b == ENDVALUE {
		return true
	}
	return false
}

func pour(bucket1, bucket2, cap int) (a int, b int, check bool){
	if bucket1 > 0 &&
		bucket2 < cap {
		if bucket1 >= (cap - bucket2) {
			a = bucket1 - (cap - bucket2)
			b = cap
		} else {
			a = 0
			b = bucket1 + bucket2
		}
		return a, b, true
	}

	return
}

func appendAndEnd(s step, parent *node) (*node, bool) {
	newNode := &node{
		s: s,
		parent: parent,
		childs: make([]*node, 0),
	}

	if !findNode(root, s) {
		//将此step加入到树中
		parent.childs = append(parent.childs, newNode)
		if checkEnd(s){
			printSteps(newNode)
			return newNode, false
		}
		return newNode, false
	}

	return nil, false
}

func printSteps(n *node){
	fmt.Println("---------------------")
	steps := make([]step, 0)
	for n != nil {
		steps = append(steps, n.s)
		n = n.parent
	}

	for i := len(steps) - 1; i >= 0; i-- {
		fmt.Printf("%d, %d, %d\n", steps[i].a, steps[i].b, steps[i].c)
	}
}

func main() {
	pCurHeader := make([]*node, 0)
	pCurHeader = append(pCurHeader, root)

	for pCurHeader != nil && len(pCurHeader) > 0 {
		thisList := make([]*node, 0)
		for _, pCurNode := range pCurHeader {
			if !checkEnd(pCurNode.s) {
				s := pCurNode.s
				//A to B
				if x,y,ok := pour(s.a, s.b, BCAP); ok{
					tmpStep := step{x,y,s.c}
					n,ok := appendAndEnd(tmpStep, pCurNode)
					if n != nil {
						thisList = append(thisList, n)
					}
					if ok {
						goto END
					}
				}
				//A to C
				if x,y,ok := pour(s.a, s.c, CCAP); ok{
					tmpStep := step{x,s.b,y}
					n,ok := appendAndEnd(tmpStep, pCurNode)
					if n != nil {
						thisList = append(thisList, n)
					}
					if ok {
						goto END
					}
				}
				//B to A
				if x,y,ok := pour(s.b, s.a, ACAP); ok{
					tmpStep := step{y,x,s.c}
					n,ok := appendAndEnd(tmpStep, pCurNode)
					if n != nil {
						thisList = append(thisList, n)
					}
					if ok {
						goto END
					}
				}
				//B to C
				if x,y,ok := pour(s.b, s.c, CCAP); ok{
					tmpStep := step{s.a,x,y}
					n,ok := appendAndEnd(tmpStep, pCurNode)
					if n != nil {
						thisList = append(thisList, n)
					}
					if ok {
						goto END
					}
				}
				//C to A
				if x,y,ok := pour(s.c, s.a, ACAP); ok{
					tmpStep := step{y,s.b,x}
					n,ok := appendAndEnd(tmpStep, pCurNode)
					if n != nil {
						thisList = append(thisList, n)
					}
					if ok {
						goto END
					}
				}
				//C to B
				if x,y,ok := pour(s.c, s.b, BCAP); ok{
					tmpStep := step{s.a,y,x}
					n,ok := appendAndEnd(tmpStep, pCurNode)
					if n != nil {
						thisList = append(thisList, n)
					}
					if ok {
						goto END
					}
				}
			}
		}

		pCurHeader = thisList
	}
END:
	fmt.Println("GAME OVER")
}
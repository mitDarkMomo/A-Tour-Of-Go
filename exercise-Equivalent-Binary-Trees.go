package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {

	// 左节点遍历
	if t.Left != nil {
		Walk(t.Left, ch)
	}

	// 输出自己
	ch <- t.Value

	// 右节点遍历
	if t.Right != nil {
		Walk(t.Right, ch)
	}

}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	same := true
	ch1, ch2 := make(chan int), make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	for i := 0; i < 10; i++ {
		if x1, x2 := <-ch1, <-ch2; x1 != x2 {
			same = false
		}
	}

	return same
}

func main() {
	result := Same(tree.New(1), tree.New(1))
	fmt.Printf("两个树的比较结果是：%v", result)
}

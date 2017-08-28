package conctw

//import "fmt"
import "golang.org/x/tour/tree"

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func recurseWalk(t *tree.Tree, ch chan int) {

	if t == nil {
		return
	}

	recurseWalk(t.Left, ch)
	ch <- t.Value
	recurseWalk(t.Right, ch)

	// fmt.Println("returning")
}

func Walk(t *tree.Tree, ch chan int) {
	recurseWalk(t, ch)
	close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
// test comment
func Same(t1, t2 *tree.Tree) bool {
	var ch1, ch2 chan int = make(chan int), make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for {
		x1, ok1 := <-ch1
		x2, ok2 := <-ch2

		if !ok1 && !ok2 {
			return true
		}
		if ok1 && ok2 {
			if x1 != x2 {
				return false
			}
		} else {
			return false
		}
	}
}

/* func main() {
	fmt.Println("Same t1(1), t2(1):", Same(tree.New(1), tree.New(1)))
	fmt.Println("Same t1(1), t2(10):", Same(tree.New(1), tree.New(10)))
	fmt.Println("Same t1(10), t2(10):", Same(tree.New(10), tree.New(10)))
}
*/

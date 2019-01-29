// select max(balance) from user where balanc not in (select max(balance) from user) 
package main
import "fmt"

type LinkedList struct {     // Linked list
    value int
    right *LinkedList
}

// Add data to LinkedList one by one
func (sl *LinkedList) add (x int, right *LinkedList) {
    sl.value = x
    if right.value != 0 {
        sl.right = right
    }
}

// Check LinkedList, whether a loop is included yet
func isLoop (lls []LinkedList) bool {
    var temp map[int]LinkedList

    for _, ll := range lls {
        if _, exist := temp[ll.value]; exist {
            temp[1] = LinkedList{}
        }
        // if ll.right {
        //     return true
        // }
    }

    return false
}

func main() {
    var ll LinkedList
    // ll = LinkedList{1}
    fmt.Println(LinkedList{1, *LinkedList{}})
    // fmt.Println(*ll)
    fmt.Println(&ll)
    // print(ll)
    // var lls []struct
    //lls = struct{{1}, {2}}
    //println(lls)

    // todo: insert
    // lls.add(1, nil)
    // lls.add(1, nil)
    // lls.add(1, nil)
}

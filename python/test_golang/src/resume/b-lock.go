package main
import  "sync"
// import  "fmt"

type UserAges struct {
    ages map[string] int
    sync.Mutex
}

func (u *UserAges) Add(name string,age int)  {
    u.Lock()
    defer u.Unlock()
    if _, ok := u.ages[0]; ok {
        println("ok-ok")
        u.ages[name] = age
    } else {
        println("ok-no")
    }
}

func (u *UserAges) Get(name string) int {
    if age,ok:=u.ages[name];ok{
        return age
    }
    return -1
}

func main() {
	u := UserAges{}
    u.Add("hi", 20)
    // u.Get("hi")
}

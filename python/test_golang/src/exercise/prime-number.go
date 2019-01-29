// Computing <1000 prime numbers
package main

func main() {
    for i:=1; i<10; i++ {
        // println("------------ \n")
        // println("i = ", i)

        if i <= 3 {
            println("prime=", i)
            continue
        }
        if i%2 == 0 || i%3 == 0 || i%5 == 0 {
            continue
        }

        is_prime := true
        for j:=6; j<i; j++ {
            // println("Loop j: ", i, j, i%j)
            if i%j == 0  {
                is_prime = false
                continue
            }
        }

        if is_prime {
            println("prime=", i)
        }
    }
}

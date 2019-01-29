package main

func is_prime(number int) bool {
    if number <= 3 {
        return number >= 2
    }

    if number%2 == 0 || number%3 == 0{
        return false
    }

    for i:=5; i*i<=number; i++ {
        if number%i == 0  {
            return false
        }
    }
    return true
}

func main() {
    counter := 0
    for i:=0; i<=1000; i++ {
        // if isPrime(i){
        if is_prime(i){
            println("prime=", i)
            counter++
        }
    }

    // for i:2: i<=
    println("Total=", counter)
}

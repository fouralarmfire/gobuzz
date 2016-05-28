package fizzbuzz

import "fmt"

func DivisibleByThree(number int) bool {
  return IsDivisible(number, 3)
}

func DivisibleByFive(number int) bool {
  return IsDivisible(number, 5)
}

func DivisibleByFifteen(number int) bool {
  return IsDivisible(number, 15)
}

func IsDivisible(number int, divisor int) bool {
  if number % divisor == 0 {
    return true
  } else {
    return false
  }
}

func FizzbuzzSays(number int) string {
  if DivisibleByFifteen(number) {
    return "fizzbuzz"
  }
  if DivisibleByFive(number) {
    return "buzz"
  }
  if DivisibleByThree(number) {
    return "fizz"
  }
  return fmt.Sprintf("number %v neither fizzed nor buzzed :(", number) 
}

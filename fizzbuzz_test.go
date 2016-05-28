package fizzbuzz

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

func TestDivisibleByThree(t *testing.T) {
  assert.True(t, DivisibleByThree(3))
}

func TestNotDivisibleByThree(t *testing.T) {
  assert.False(t, DivisibleByThree(1))  
}

func TestDivisibleByFive(t *testing.T) {
  assert.True(t, DivisibleByFive(5))
}

func TestNotDivisibleByFive(t *testing.T) {
  assert.False(t, DivisibleByFive(1))
}

func TestDivisibleByFifteen(t *testing.T) {
  assert.True(t, DivisibleByFifteen(15))
}

func TestNotDivisibleByFifteen(t *testing.T) {
  assert.False(t, DivisibleByFifteen(1))
}

func TestFizzbuzzSaysFizz(t *testing.T) {
  assertFizzbuzz(3, "fizz", t)
}

func TestFizzbuzzSaysBuzz(t *testing.T) {
  assertFizzbuzz(5, "buzz", t)
}

func TestFizzbuzzSaysFizzBuzz(t *testing.T) {
  assertFizzbuzz(15, "fizzbuzz", t)
}

func TestFizzbuzzSaysNothing(t *testing.T) {
  assertFizzbuzz(1, "number 1 neither fizzed nor buzzed :(", t)
}

func assertFizzbuzz(number int, expected string, t *testing.T) {
	assert.Equal(t, FizzbuzzSays(number), expected)
}

package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Log("sum(1,2):", sum(1, 2))

	if sum(1, 2) != 3 {
		t.Errorf("SALAH! harusnya 3")
	}
}

func TestMutiply(t *testing.T) {
	t.Log("multiply(1,2):", multiply(1, 2))

	if multiply(1, 2) != 2 {
		t.Errorf("SALAH! harusnya 2")
	}
}

func TestFirstPrime(t *testing.T) {
	t.Log("fistPrime(4):", firstPrime(4))
	expected := []int64{2, 3, 5, 7}

	if !reflect.DeepEqual(expected, firstPrime(4)) {
		t.Errorf("SALAH! harusnya %d", expected)
	}
}

func TestFirstFibonacci(t *testing.T) {
	t.Log("fistFibonnaci(4):", firstFibonacci(4))
	expected := []int{0, 1, 1, 2}

	if !reflect.DeepEqual(expected, firstFibonacci(4)) {
		t.Errorf("SALAH! harusnya %d", expected)
	}
}

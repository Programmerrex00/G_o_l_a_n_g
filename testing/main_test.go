package main

import (
	"testing"
)

func TestSum(t *testing.T) {
	// total := Sum(5, 5)

	// if total != 11 {
	// 	t.Errorf("Sum was incorrect, got %d expected %d", total, 11)
	// }

	tables := []struct {
		a int
		b int
		n int
	}{
		{1, 2, 3},
		{2, 2, 4},
		{25, 26, 51},
	}
	for _, item := range tables {
		total := Sum(item.a, item.b)
		if total != item.n {
			t.Errorf("Sum was incorrect, got %d expected %d", total, item.n)
		}
	}
}

// Agregando un test para GetMax
func TestMax(t *testing.T) {
	table := []struct {
		a int
		b int
		c int
	}{
		{4, 2, 4},
		{3, 2, 3},
		{2, 8, 8},
	}
	for _, item := range table {
		max := GetMax(item.a, item.b)
		if max != item.c {
			t.Errorf("GetMax was incorrect, got %d, expected %d", max, item.c)
		}
	}
}

func TestFib(t *testing.T) {
	table := []struct {
		a int
		n int
	}{
		{1, 1},
		{8, 21},
		{50, 12586269025},
	}
	for _, item := range table {
		fibo := Fibonacci(item.a)
		if fibo != item.n {
			t.Errorf("Fibonacci was incorrect, got %d expected %d", fibo, item.n)
		}
	}
}

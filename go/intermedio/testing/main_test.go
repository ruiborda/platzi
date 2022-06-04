package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1+2=3", args{1, 2}, 3},
		{"2+2=4", args{2, 2}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sum(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("Sum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFibonacci(t *testing.T) {
	type args struct {
		n uint8
	}
	tests := []struct {
		name string
		args args
		want []uint64
	}{
		{"fibonacci of 0", args{0}, []uint64{}},
		{"fibonacci of 1", args{1}, []uint64{0}},
		{"fibonacci of 2", args{2}, []uint64{0, 1}},
		{"fibonacci of 3", args{3}, []uint64{0, 1, 1}},
		{"fibonacci of 4", args{4}, []uint64{0, 1, 1, 2}},
		{"fibonacci of 5", args{5}, []uint64{0, 1, 1, 2, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Fibonacci(tt.args.n); reflect.DeepEqual(got, tt.want) {
				t.Errorf("Fibonacci() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFibonacciFunction(t *testing.T) {
	type args struct {
		n uint8
	}
	tests := []struct {
		name string
		args args
		want uint64
	}{
		{"fibonacci of 0", args{0}, 0},
		{"fibonacci of 1", args{1}, 1},
		{"fibonacci of 2", args{2}, 1},
		{"fibonacci of 3", args{3}, 2},
		{"fibonacci of 4", args{4}, 3},
		{"fibonacci of 5", args{5}, 5},
		{"fibonacci of 30", args{30}, 832040},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FibonacciFunction(tt.args.n); got != tt.want {
				t.Errorf("FibonacciFunction() = %v, want %v", got, tt.want)
			}
		})
	}
}

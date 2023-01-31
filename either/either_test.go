package either

import (
	"testing"
)

func TestLeft(t *testing.T) {
	input := 10
	output := Left[int, int](input)
	if output.left != input {
		t.Errorf("Received %d - expected %d", output.left, input)
	}
}

func TestRight(t *testing.T) {
	input := 10
	output := Right[int, int](input)
	if output.right != input {
		t.Errorf("Received %d - expected %d", output.left, input)
	}
}

func TestIsLeft(t *testing.T) {
	input := 10
	left := Left[int, int](input)
	if output := left.IsLeft(); output == false {
		t.Errorf("Received %t - expected true", output)
	}
	right := Right[int, int](input)
	if output := right.IsLeft(); output == true {
		t.Errorf("Received %t - expected false", output)
	}
}

func TestIsRight(t *testing.T) {
	input := 10
	right := Right[int, int](input)
	if output := right.IsRight(); output == false {
		t.Errorf("Received %t - expected true", output)
	}
	left := Left[int, int](input)
	if output := left.IsRight(); output == true {
		t.Errorf("Received %t - expected false", output)
	}
}

func TestMapLeft(t *testing.T) {
	f := func(x int) int {
		return x * 2
	}
	input := Left[int, int](10)
	output := MapLeft[int, int](f)(input)
	expected := f(input.left)
	if output.left != expected {
		t.Errorf("Received %d - expected %d", output.left, expected)
	}
}

func TestMap(t *testing.T) {
	f := func(x int) int {
		return x * 2
	}
	input := Right[int, int](10)
	output := Map[int, int](f)(input)
	expected := f(input.right)
	if output.right != expected {
		t.Errorf("Received %d - expected %d", output.right, expected)
	}
}

func TestChainLeft(t *testing.T) {
	f := func(x int) *Either[int, int] {
		return Left[int, int](x * 2)
	}
	input := Left[int, int](10)
	output := ChainLeft[int, int](f)(input)
	expected := f(input.left).left
	if output.left != expected {
		t.Errorf("Received %d - expected %d", output.left, expected)
	}
}

func TestChain(t *testing.T) {
	f := func(x int) *Either[int, int] {
		return Right[int, int](x * 2)
	}
	input := Right[int, int](10)
	output := Chain[int, int](f)(input)
	expected := f(input.right).right
	if output.right != expected {
		t.Errorf("Received %d - expected %d", output.right, expected)
	}
}

func TestFold(t *testing.T) {
	var left float32
	var right int32
	onLeft := func(x float32) {
		left = x
	}
	onRight := func(x int32) {
		right = x
	}
	input := Either[float32, int32]{
		left:  10.0,
		right: 10,
	}
	Fold(onLeft, onRight)(input)
	if left != input.left {
		t.Errorf("Received %f - expected %f", left, input.left)
	}
	if right != input.right {
		t.Errorf("Received %d - expected %d", right, input.right)
	}
}

func TestSwap(t *testing.T) {
	input := Either[float32, int32]{
		left:  10.0,
		right: 10,
	}
	output := Swap(input)
	if output.right != input.left {
		t.Errorf("Received %f - expected %f", output.right, input.left)
	}
	if output.left != input.right {
		t.Errorf("Received %d - expected %d", output.left, input.right)
	}
}

package either

import "reflect"

type Either[A, B any] struct {
	left  A
	right B
}

func (e *Either[A, B]) IsLeft() bool {
	return !reflect.ValueOf(e.left).IsZero()
}

func (e *Either[A, B]) IsRight() bool {
	return !reflect.ValueOf(e.right).IsZero()
}

func Left[A, B any](value A) *Either[A, B] {
	return &Either[A, B]{
		left: value,
	}
}

func Right[A, B any](value B) *Either[A, B] {
	return &Either[A, B]{
		right: value,
	}
}

func MapLeft[A, B, C any](f func(A) C) func(*Either[A, B]) *Either[C, B] {
	return func(e *Either[A, B]) *Either[C, B] {
		if e.IsLeft() {
			return Left[C, B](f(e.left))
		}
		return Right[C, B](e.right)
	}
}

func Map[A, B, C any](f func(B) C) func(*Either[A, B]) *Either[A, C] {
	return func(e *Either[A, B]) *Either[A, C] {
		if e.IsRight() {
			return Right[A, C](f(e.right))
		}
		return Left[A, C](e.left)
	}
}

func ChainLeft[A, B, C any](f func(A) *Either[C, B]) func(*Either[A, B]) *Either[C, B] {
	return func(e *Either[A, B]) *Either[C, B] {
		if e.IsLeft() {
			return Left[C, B](f(e.left).left)
		}
		return Right[C, B](e.right)
	}
}

func Chain[A, B, C any](f func(B) *Either[A, C]) func(*Either[A, B]) *Either[A, C] {
	return func(e *Either[A, B]) *Either[A, C] {
		if e.IsRight() {
			return Right[A, C](f(e.right).right)
		}
		return Left[A, C](e.left)
	}
}

func Fold[A, B any](onLeft func(A), onRight func(B)) func(Either[A, B]) {
	return func(e Either[A, B]) {
		onLeft(e.left)
		onRight(e.right)
	}
}

func Swap[A, B any](e Either[A, B]) Either[B, A] {
	return Either[B, A]{
		left:  e.right,
		right: e.left,
	}
}

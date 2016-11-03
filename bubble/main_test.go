package main

import (
	"reflect"
	"testing"
)

func TestSortTwo(t *testing.T) {
	t.Parallel()

	cases := []struct {
		input []int
		wants []int
	}{
		{
			[]int{0, 0},
			[]int{0, 0},
		},
		{
			[]int{1, 2},
			[]int{1, 2},
		},
		{
			[]int{2, 1},
			[]int{1, 2},
		},
		{
			[]int{13, 1},
			[]int{1, 13},
		},
		{
			[]int{-13, 1},
			[]int{-13, 1},
		},
	}

	for _, c := range cases {
		c := c

		t.Run("", func(t *testing.T) {
			t.Parallel()

			got := SwapFirstTwoElementsIfOutOfOrder(c.input)
			if !reflect.DeepEqual(got, c.wants) {
				t.Fatalf("Difference in expected result\nGot: %v\nExpected: %v\n", got, c.wants)
			}
		})
	}
}

func TestMoveFirstElementUntilAllItemsBeforeItAreSmaller(t *testing.T) {
	t.Parallel()

	cases := []struct {
		input []int
		wants []int
	}{
		{
			[]int{0, 0, 0},
			[]int{0, 0, 0},
		},
		{
			[]int{1, 3, 0},
			[]int{1, 3, 0},
		},
		{
			[]int{3, 1, 0},
			[]int{1, 0, 3},
		},
		{
			[]int{13, 1, 0},
			[]int{1, 0, 13},
		},
		{
			[]int{-13, 1, 30},
			[]int{-13, 1, 30},
		},
	}

	for _, c := range cases {
		c := c

		t.Run("", func(t *testing.T) {
			t.Parallel()

			got := MoveFirstElementUntilAllItemsBeforeItAreSmaller(c.input)
			if !reflect.DeepEqual(got, c.wants) {
				t.Fatalf("Difference in expected result\nGot: %v\nExpected: %v\n", got, c.wants)
			}
		})
	}
}

func TestKeepMovingTheFirstElementUntilTheSliceIsSorted(t *testing.T) {
	t.Parallel()

	cases := []struct {
		input []int
		wants []int
	}{
		{
			[]int{0, 0, 0},
			[]int{0, 0, 0},
		},
		{
			[]int{1, 3, 0},
			[]int{0, 1, 3},
		},
		{
			[]int{3, 1, 0},
			[]int{0, 1, 3},
		},
		{
			[]int{13, 1, 0},
			[]int{0, 1, 13},
		},
		{
			[]int{-13, 1, 30},
			[]int{-13, 1, 30},
		},
		{
			[]int{-13, 20, 30, 0, 9, 10000},
			[]int{-13, 0, 9, 20, 30, 10000},
		},
	}

	for _, c := range cases {
		c := c

		t.Run("", func(t *testing.T) {
			t.Parallel()

			got := KeepMovingTheFirstElementUntilTheSliceIsSorted(c.input)
			if !reflect.DeepEqual(got, c.wants) {
				t.Fatalf("Difference in expected result\nGot: %v\nExpected: %v\n", got, c.wants)
			}
		})
	}
}

func TestBubbleSort(t *testing.T) {
	t.Parallel()

	cases := []struct {
		input []int
		wants []int
	}{
		{
			[]int{0, 0, 0},
			[]int{0, 0, 0},
		},
		{
			[]int{1, 3, 0, 10, 39, 2},
			[]int{0, 1, 2, 3, 10, 39},
		},
		{
			[]int{1, 2, -1},
			[]int{-1, 1, 2},
		},
		{
			[]int{13, 1, 0},
			[]int{0, 1, 13},
		},
		{
			[]int{-13, 1, 30},
			[]int{-13, 1, 30},
		},
		{
			[]int{-13, 20, 30, 0, 9, 10000, -1, 0},
			[]int{-13, -1, 0, 0, 9, 20, 30, 10000},
		},
	}

	for _, c := range cases {
		c := c

		t.Run("", func(t *testing.T) {
			t.Parallel()

			got := BubbleSort(c.input)
			if !reflect.DeepEqual(got, c.wants) {
				t.Fatalf("Difference in expected result\nGot: %v\nExpected: %v\n", got, c.wants)
			}
		})
	}
}

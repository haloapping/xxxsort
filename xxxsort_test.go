package xxxsort_test

import (
	"testing"

	"github.com/haloapping/xxxsort"
	"github.com/stretchr/testify/assert"
)

func TestBubbleSort(t *testing.T) {
	t.Run("sort number", func(t *testing.T) {
		t.Run("ascending mode", func(t *testing.T) {
			s := []int{4, 5, 2, 1, 1, 9, 0}
			c, err := xxxsort.BubbelSort(s, "asc")
			assert.Equal(t, []int{0, 1, 1, 2, 4, 5, 9}, c)
			assert.Equal(t, nil, err)
		})

		t.Run("descending mode", func(t *testing.T) {
			s := []int{4, 5, 2, 1, 1, 9, 0}
			c, err := xxxsort.BubbelSort(s, "desc")
			assert.Equal(t, []int{9, 5, 4, 2, 1, 1, 0}, c)
			assert.Equal(t, nil, err)
		})

		t.Run("mode is not available", func(t *testing.T) {
			s := []int{4, 5, 2, 1, 1, 9, 0}
			_, err := xxxsort.BubbelSort(s, "xxx")
			assert.EqualError(t, err, "xxx not available, use instead 'asc' or 'desc'")
		})
	})

	t.Run("sort string", func(t *testing.T) {
		t.Run("ascending mode", func(t *testing.T) {
			s := []string{"apping", "appang", "appeng", "appung", "appong"}
			c, err := xxxsort.BubbelSort(s, "asc")
			assert.Equal(t, []string{"appang", "appeng", "apping", "appong", "appung"}, c)
			assert.Equal(t, nil, err)
		})

		t.Run("descending mode", func(t *testing.T) {
			s := []string{"apping", "appang", "appeng", "appung", "appong"}
			c, err := xxxsort.BubbelSort(s, "desc")
			assert.Equal(t, []string{"appung", "appong", "apping", "appeng", "appang"}, c)
			assert.Equal(t, nil, err)
		})

		t.Run("mode is not available", func(t *testing.T) {
			s := []string{"apping", "appang", "appeng", "appung", "appong"}
			_, err := xxxsort.BubbelSort(s, "xxx")
			assert.EqualError(t, err, "xxx not available, use instead 'asc' or 'desc'")
		})
	})
}

func TestSelectionSort(t *testing.T) {
	t.Run("sort number", func(t *testing.T) {
		t.Run("ascending mode", func(t *testing.T) {
			s := []int{4, 5, 2, 1, 1, 9, 0}
			c, err := xxxsort.SelectionSort(s, "asc")
			assert.Equal(t, []int{0, 1, 1, 2, 4, 5, 9}, c)
			assert.Equal(t, nil, err)
		})

		t.Run("descending mode", func(t *testing.T) {
			s := []int{4, 5, 2, 1, 1, 9, 0}
			c, err := xxxsort.SelectionSort(s, "desc")
			assert.Equal(t, []int{9, 5, 4, 2, 1, 1, 0}, c)
			assert.Equal(t, nil, err)
		})

		t.Run("mode is not available", func(t *testing.T) {
			s := []int{4, 5, 2, 1, 1, 9, 0}
			_, err := xxxsort.SelectionSort(s, "xxx")
			assert.EqualError(t, err, "xxx not available, use instead 'asc' or 'desc'")
		})
	})

	t.Run("sort string", func(t *testing.T) {
		t.Run("ascending mode", func(t *testing.T) {
			s := []string{"apping", "appang", "appeng", "appung", "appong"}
			c, err := xxxsort.SelectionSort(s, "asc")
			assert.Equal(t, []string{"appang", "appeng", "apping", "appong", "appung"}, c)
			assert.Equal(t, nil, err)
		})

		t.Run("descending mode", func(t *testing.T) {
			s := []string{"apping", "appang", "appeng", "appung", "appong"}
			c, err := xxxsort.SelectionSort(s, "desc")
			assert.Equal(t, []string{"appung", "appong", "apping", "appeng", "appang"}, c)
			assert.Equal(t, nil, err)
		})

		t.Run("mode is not available", func(t *testing.T) {
			s := []string{"apping", "appang", "appeng", "appung", "appong"}
			_, err := xxxsort.SelectionSort(s, "xxx")
			assert.EqualError(t, err, "xxx not available, use instead 'asc' or 'desc'")
		})
	})
}

func TestInsertionSort(t *testing.T) {
	t.Run("sort number", func(t *testing.T) {
		t.Run("ascending mode", func(t *testing.T) {
			s := []int{4, 5, 2, 1, 1, 9, 0}
			c, err := xxxsort.InsertionSort(s, "asc")
			assert.Equal(t, []int{0, 1, 1, 2, 4, 5, 9}, c)
			assert.Equal(t, nil, err)
		})

		t.Run("descending mode", func(t *testing.T) {
			s := []int{4, 5, 2, 1, 1, 9, 0}
			c, err := xxxsort.InsertionSort(s, "desc")
			assert.Equal(t, []int{9, 5, 4, 2, 1, 1, 0}, c)
			assert.Equal(t, nil, err)
		})

		t.Run("mode is not available", func(t *testing.T) {
			s := []int{4, 5, 2, 1, 1, 9, 0}
			_, err := xxxsort.InsertionSort(s, "xxx")
			assert.EqualError(t, err, "xxx not available, use instead 'asc' or 'desc'")
		})
	})

	t.Run("sort string", func(t *testing.T) {
		t.Run("ascending mode", func(t *testing.T) {
			s := []string{"apping", "appang", "appeng", "appung", "appong"}
			c, err := xxxsort.InsertionSort(s, "asc")
			assert.Equal(t, []string{"appang", "appeng", "apping", "appong", "appung"}, c)
			assert.Equal(t, nil, err)
		})

		t.Run("descending mode", func(t *testing.T) {
			s := []string{"apping", "appang", "appeng", "appung", "appong"}
			c, err := xxxsort.InsertionSort(s, "desc")
			assert.Equal(t, []string{"appung", "appong", "apping", "appeng", "appang"}, c)
			assert.Equal(t, nil, err)
		})

		t.Run("mode is not available", func(t *testing.T) {
			s := []string{"apping", "appang", "appeng", "appung", "appong"}
			_, err := xxxsort.InsertionSort(s, "xxx")
			assert.EqualError(t, err, "xxx not available, use instead 'asc' or 'desc'")
		})
	})
}

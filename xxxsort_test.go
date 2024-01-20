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
			xxxsort.BubbelSortAsc(s)
			assert.Equal(t, []int{0, 1, 1, 2, 4, 5, 9}, s)
		})

		t.Run("descending mode", func(t *testing.T) {
			s := []int{4, 5, 2, 1, 1, 9, 0}
			xxxsort.BubbelSortDesc(s)
			assert.Equal(t, []int{9, 5, 4, 2, 1, 1, 0}, s)
		})
	})

	t.Run("sort string", func(t *testing.T) {
		t.Run("ascending mode", func(t *testing.T) {
			s := []string{"apping", "appang", "appeng", "appung", "appong"}
			xxxsort.BubbelSortAsc(s)
			assert.Equal(t, []string{"appang", "appeng", "apping", "appong", "appung"}, s)
		})

		t.Run("descending mode", func(t *testing.T) {
			s := []string{"apping", "appang", "appeng", "appung", "appong"}
			xxxsort.BubbelSortDesc(s)
			assert.Equal(t, []string{"appung", "appong", "apping", "appeng", "appang"}, s)
		})
	})
}

func TestSelectionSort(t *testing.T) {
	t.Run("sort number", func(t *testing.T) {
		t.Run("ascending mode", func(t *testing.T) {
			s := []int{4, 5, 2, 1, 1, 9, 0}
			xxxsort.SelectionSortAsc(s)
			assert.Equal(t, []int{0, 1, 1, 2, 4, 5, 9}, s)
		})

		t.Run("descending mode", func(t *testing.T) {
			s := []int{4, 5, 2, 1, 1, 9, 0}
			xxxsort.SelectionSortDesc(s)
			assert.Equal(t, []int{9, 5, 4, 2, 1, 1, 0}, s)
		})
	})

	t.Run("sort string", func(t *testing.T) {
		t.Run("ascending mode", func(t *testing.T) {
			s := []string{"apping", "appang", "appeng", "appung", "appong"}
			xxxsort.SelectionSortAsc(s)
			assert.Equal(t, []string{"appang", "appeng", "apping", "appong", "appung"}, s)
		})

		t.Run("descending mode", func(t *testing.T) {
			s := []string{"apping", "appang", "appeng", "appung", "appong"}
			xxxsort.SelectionSortDesc(s)
			assert.Equal(t, []string{"appung", "appong", "apping", "appeng", "appang"}, s)
		})
	})
}

func TestInsertionSort(t *testing.T) {
	t.Run("sort number", func(t *testing.T) {
		t.Run("ascending mode", func(t *testing.T) {
			s := []int{4, 5, 2, 1, 1, 9, 0}
			xxxsort.InsertionSortAsc(s)
			assert.Equal(t, []int{0, 1, 1, 2, 4, 5, 9}, s)
		})

		t.Run("descending mode", func(t *testing.T) {
			s := []int{4, 5, 2, 1, 1, 9, 0}
			xxxsort.InsertionSortDesc(s)
			assert.Equal(t, []int{9, 5, 4, 2, 1, 1, 0}, s)
		})
	})

	t.Run("sort string", func(t *testing.T) {
		t.Run("ascending mode", func(t *testing.T) {
			s := []string{"apping", "appang", "appeng", "appung", "appong"}
			xxxsort.InsertionSortAsc(s)
			assert.Equal(t, []string{"appang", "appeng", "apping", "appong", "appung"}, s)
		})

		t.Run("descending mode", func(t *testing.T) {
			s := []string{"apping", "appang", "appeng", "appung", "appong"}
			xxxsort.InsertionSortDesc(s)
			assert.Equal(t, []string{"appung", "appong", "apping", "appeng", "appang"}, s)
		})
	})
}

func TestQuickSort(t *testing.T) {
	t.Run("sort number", func(t *testing.T) {
		t.Run("ascending mode", func(t *testing.T) {
			s := []int{4, 5, 2, 1, 1, 9, 0}
			xxxsort.QuickSortAsc(s, 0, len(s)-1)
			assert.Equal(t, []int{0, 1, 1, 2, 4, 5, 9}, s)
		})

		t.Run("descending mode", func(t *testing.T) {
			s := []int{4, 5, 2, 1, 1, 9, 0}
			xxxsort.QuickSortDesc(s, 0, len(s)-1)
			assert.Equal(t, []int{9, 5, 4, 2, 1, 1, 0}, s)
		})
	})

	t.Run("sort string", func(t *testing.T) {
		t.Run("ascending mode", func(t *testing.T) {
			s := []string{"apping", "appang", "appeng", "appung", "appong"}
			xxxsort.QuickSortAsc(s, 0, len(s)-1)
			assert.Equal(t, []string{"appang", "appeng", "apping", "appong", "appung"}, s)
		})

		t.Run("descending mode", func(t *testing.T) {
			s := []string{"apping", "appang", "appeng", "appung", "appong"}
			xxxsort.QuickSortDesc(s, 0, len(s)-1)
			assert.Equal(t, []string{"appung", "appong", "apping", "appeng", "appang"}, s)
		})
	})
}

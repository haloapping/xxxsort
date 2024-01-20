package xxxsort

import (
	"cmp"
)

func BubbelSortAsc[T cmp.Ordered](s []T) {
	for i := 0; i < len(s); i++ {
		is_swap := false
		for j := 0; j < len(s)-i-1; j++ {
			if cmp.Compare(s[j], s[j+1]) == 1 {
				s[j], s[j+1] = s[j+1], s[j]
				is_swap = true
			}
		}

		if !is_swap {
			break
		}
	}
}

func BubbelSortDesc[T cmp.Ordered](s []T) {
	for i := 0; i < len(s); i++ {
		is_swap := false
		for j := 0; j < len(s)-i-1; j++ {
			if cmp.Compare(s[j], s[j+1]) == -1 {
				s[j], s[j+1] = s[j+1], s[j]
				is_swap = true
			}
		}

		if !is_swap {
			break
		}
	}
}

func SelectionSortAsc[T cmp.Ordered](s []T) {
	for i := 0; i < len(s)-1; i++ {
		lowestNumberIdx := i
		for j := i + 1; j < len(s); j++ {
			if cmp.Compare(s[j], s[lowestNumberIdx]) == -1 {
				lowestNumberIdx = j
			}
		}

		if lowestNumberIdx != i {
			s[i], s[lowestNumberIdx] = s[lowestNumberIdx], s[i]
		}
	}
}

func SelectionSortDesc[T cmp.Ordered](s []T) {
	for i := 0; i < len(s)-1; i++ {
		lowestNumberIdx := i
		for j := i + 1; j < len(s); j++ {
			if cmp.Compare(s[j], s[lowestNumberIdx]) == 1 {
				lowestNumberIdx = j
			}
		}

		if lowestNumberIdx != i {
			s[i], s[lowestNumberIdx] = s[lowestNumberIdx], s[i]
		}
	}
}

func InsertionSortAsc[T cmp.Ordered](s []T) {
	for i := 1; i < len(s); i++ {
		tempVal := s[i]
		position := i - 1
		for position >= 0 {
			if cmp.Compare(tempVal, s[position]) == -1 {
				s[position+1] = s[position]
				position--
			} else {
				break
			}
		}

		s[position+1] = tempVal
	}
}

func InsertionSortDesc[T cmp.Ordered](s []T) {
	for i := 1; i < len(s); i++ {
		tempVal := s[i]
		position := i - 1
		for position >= 0 {
			if cmp.Compare(tempVal, s[position]) == 1 {
				s[position+1] = s[position]
				position--
			} else {
				break
			}
		}

		s[position+1] = tempVal
	}
}

func partitionAsc[T cmp.Ordered](s []T, low int, high int) int {
	pivot := s[high]
	i := low - 1

	for j := low; j <= high; j++ {
		if cmp.Compare(s[j], pivot) == -1 {
			i++
			s[i], s[j] = s[j], s[i]
		}
	}

	s[i+1], s[high] = s[high], s[i+1]

	return i + 1
}

func QuickSortAsc[T cmp.Ordered](s []T, low int, high int) {
	if low < high {
		part := partitionAsc(s, low, high)
		QuickSortAsc(s, low, part-1)
		QuickSortAsc(s, part+1, high)
	}
}

func partitionDesc[T cmp.Ordered](s []T, low int, high int) int {
	pivot := s[high]
	i := low - 1

	for j := low; j <= high; j++ {
		if cmp.Compare(s[j], pivot) == 1 {
			i++
			s[i], s[j] = s[j], s[i]
		}
	}

	s[i+1], s[high] = s[high], s[i+1]

	return i + 1
}

func QuickSortDesc[T cmp.Ordered](s []T, low int, high int) {
	if low < high {
		part := partitionDesc(s, low, high)
		QuickSortDesc(s, low, part-1)
		QuickSortDesc(s, part+1, high)
	}
}

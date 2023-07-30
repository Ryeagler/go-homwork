package slice

import "errors"

func Delete[T any](slice []T, index int) ([]T, error) {
	if index < 0 || index >= len(slice) {
		return slice, errors.New("illegal index")
	}

	cursorIndex := 0
	for i, _ := range slice {
		if i == index {
			continue
		}
		slice[cursorIndex] = slice[i]
		cursorIndex += 1
	}

	return slice[:cursorIndex], nil
}

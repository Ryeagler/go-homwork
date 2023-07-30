package slice

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDelete(t *testing.T) {
	tests := []struct {
		slice         []string
		indexToDelete int
		expected      []string
		name          string
	}{
		{
			slice:         []string{"apple", "banana", "strawberry"},
			indexToDelete: 0,
			expected:      []string{"banana", "strawberry"},
			name:          "delete an element with index 0",
		},
		{
			slice:         []string{"apple", "banana", "strawberry"},
			indexToDelete: -1,
			expected:      []string{"apple", "banana", "strawberry"},
			name:          "delete an element with index -1",
		},
		{
			slice:         []string{"apple", "banana", "strawberry"},
			indexToDelete: 3,
			expected:      []string{"apple", "banana", "strawberry"},
			name:          "delete an element with index 10",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res, _ := Delete(test.slice, test.indexToDelete)
			assert.Equal(t, test.expected, res)
		})
	}
}

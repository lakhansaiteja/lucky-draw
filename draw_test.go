package lucky_draw

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRemoveEmptyItems(t *testing.T) {
	items := []string{"", "hello"}
	modifiedItems := RemoveEmptyItems(items)
	assert.Equal(t, 1, len(modifiedItems))
}

func TestPrintTheWinner(t *testing.T) {
	items := []string{"Alpha"}
	PrintTheWinner(0, items)
}

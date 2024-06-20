package memory_test

import (
	"testing"

	"github.com/mraron/njudge/internal/njudge/memory"
	"github.com/stretchr/testify/assert"
)

func TestPaginate(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	res, pdata := memory.Paginate(nums, 1, 10)
	assert.Len(t, res, 10)
	assert.Equal(t, res[0], 1)
	assert.Equal(t, res[9], 10)

	assert.Equal(t, pdata.Count, 15)
	assert.Equal(t, pdata.Page, 1)
	assert.Equal(t, pdata.Pages, 2)
	assert.Equal(t, pdata.PerPage, 10)

	res, pdata = memory.Paginate(nums, 2, 10)
	assert.Len(t, res, 5)
	assert.Equal(t, res[0], 11)
	assert.Equal(t, res[4], 15)

	assert.Equal(t, pdata.Count, 15)
	assert.Equal(t, pdata.Page, 2)
	assert.Equal(t, pdata.Pages, 2)
	assert.Equal(t, pdata.PerPage, 10)

	res, pdata = memory.Paginate(nums, 100, 10)
	assert.Len(t, res, 5)

	res, pdata = memory.Paginate(nums, -100, 10)
	assert.Len(t, res, 10)
}

package testdata

import "net/http"

var CategoryTests = []struct {
	In       string
	Expected int
}{
	{
		In:       "热血",
		Expected: http.StatusOK,
	},
}

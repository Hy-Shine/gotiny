package ginx

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

// QueryArray returns the query values associated with the given key.
// It returns nil if the are no values associated with the key.
// The returned slice is NOT a view into the original query values
// and can be modified without affecting the original.
func QueryArray(ctx *gin.Context, key string) []string {
	if array := ctx.QueryArray(key); len(array) > 0 {
		values := make([]string, 0, len(array))
		for _, v := range array {
			values = append(values, v)
		}
		return values
	}

	m := ctx.QueryMap(key)
	if len(m) == 0 {
		return nil
	}

	values := make([]string, 0, len(m))
	mLen := len(m)
	for i := 0; i < mLen; i++ {
		v, _ := m[strconv.Itoa(i)]
		values = append(values, v)
	}
	return values
}

package ginx

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestQueryArray(t *testing.T) {
	// Test for a key with multiple associated values in ctx.QueryArray
	ctx1, _ := gin.CreateTestContext(httptest.NewRecorder())
	ctx1.Request, _ = http.NewRequest("GET", "/?key=value1&key=value2", nil)
	assert.Equal(t, []string{"value1", "value2"}, QueryArray(ctx1, "key"))

	// Test for a key with no associated values in ctx.QueryArray
	ctx2, _ := gin.CreateTestContext(httptest.NewRecorder())
	ctx2.Request, _ = http.NewRequest("GET", "/", nil)
	assert.Nil(t, QueryArray(ctx2, "key"))

	// Test for a key with associated values in ctx.QueryMap
	ctx3, _ := gin.CreateTestContext(httptest.NewRecorder())
	ctx3.Request, _ = http.NewRequest("GET", "/?key=value1&key=", nil)
	assert.Equal(t, []string{"value1", ""}, QueryArray(ctx3, "key"))

	// Test for a key with no associated values in ctx.QueryMap
	ctx4, _ := gin.CreateTestContext(httptest.NewRecorder())
	ctx4.Request, _ = http.NewRequest("GET", "/?key[0]=value1&key[1]=value2", nil)
	assert.Equal(t, []string{"value1", "value2"}, QueryArray(ctx4, "key"))

	// Test for a key with both associated values in ctx.QueryArray and ctx.QueryMap
	ctx5, _ := gin.CreateTestContext(httptest.NewRecorder())
	ctx5.Request, _ = http.NewRequest("GET", "/?key[0]=value1&key[1]=", nil)
	assert.Equal(t, []string{"value1", ""}, QueryArray(ctx5, "key"))
}

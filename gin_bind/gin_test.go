package gin_bind

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/zhwei820/gconv"
)

// ================================================================= for test
type TransferType uint8 // for test

const (
	TransferTypeNormalOrder   TransferType = 0
	TransferTypeTransferOrder TransferType = 1
)

func TestBindQuery(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = &http.Request{}
	c.Request.URL = &url.URL{RawQuery: "k=&k1="}
	req := struct {
		K  *TransferType `json:"k" form:"k"`
		K2 *TransferType `json:"k2" form:"k2"`
		K1 TransferType  `json:"k1" form:"k1"`
	}{}

	BindQuery(context.Background(), c, &req)
	fmt.Println("req", gconv.Export(req))
	assert.Nil(t, req.K)
	assert.Equal(t, req.K1, TransferType(0))
	// BindQuery(context.Background(), c, req)
}

// PageFilter .
type PageFilter struct {
	Offset int32 `json:"offset" form:"offset" validate:"omitempty,gte=0"` // 偏移量
	Limit  int32 `json:"limit" form:"limit" validate:"required"`          // 限制数量
}

type ListUserQuotaRecordReq struct {
	// Biz      shared.BizType `json:"biz" form:"biz"`
	UserId   string `json:"user_id" form:"user_id"`
	Currency string `json:"currency" form:"currency"`
	// PageFilter
}

func TestBindQuery2(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = &http.Request{}
	c.Request.URL = &url.URL{RawQuery: "user_id=1"}
	req := ListUserQuotaRecordReq{}

	BindQuery(context.Background(), c, &req)
	fmt.Println("req", gconv.Export(req))
	// BindQuery(context.Background(), c, req)
}

// =================================================================

type AdminGetProductsReq struct {
	*ProductFilter
}

type ProductFilter struct {
	*TimeFilter
}

type TimeFilter struct {
	ValueDay *int64 `json:"value_day" form:"value_day"` //
	CloseDay *int64 `json:"close_day" form:"close_day"` //
}

func TestBindQuery3(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = &http.Request{}
	c.Request.URL = &url.URL{RawQuery: "value_day=1&close_day="}
	req := &AdminGetProductsReq{ProductFilter: &ProductFilter{TimeFilter: &TimeFilter{}}}

	fmt.Println("req", gconv.Export(req))

	BindQuery(context.Background(), c, req)
	fmt.Println("req", gconv.Export(req))
	// BindQuery(context.Background(), c, req)
}

func TestBindQuery4(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = &http.Request{}
	c.Request.URL = &url.URL{RawQuery: "value_day=1&close_day="}
	req := &AdminGetProductsReq{}

	fmt.Println("req", gconv.Export(req))

	BindQuery(context.Background(), c, req)
	fmt.Println("req", gconv.Export(req))
	// BindQuery(context.Background(), c, req)
}

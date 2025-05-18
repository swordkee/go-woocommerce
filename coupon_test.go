package woocommerce

import (
	"errors"
	"strings"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/hiscaler/gox/jsonx"
	"github.com/hiscaler/gox/randx"
	"github.com/stretchr/testify/assert"
	"github.com/swordkee/woocommerce-go/entity"
)

func TestCouponService_All(t *testing.T) {
	params := CouponsQueryParams{Search: ""}
	params.PerPage = 100
	params.Order = SortDesc
	items, _, _, _, err := wooClient.Services.Coupon.All(params)
	if err != nil {
		t.Errorf("wooClient.Services.Coupon.All error: %s", err.Error())
	} else {
		t.Logf("items = %s", jsonx.ToPrettyJson(items))
	}
}

func TestCouponService_One(t *testing.T) {
	couponId := 4
	item, err := wooClient.Services.Coupon.One(couponId)
	if err != nil {
		t.Errorf("wooClient.Services.Coupon.One error: %s", err.Error())
	} else {
		assert.Equal(t, couponId, item.ID, "coupon id")
	}
}

func TestCouponService_Create(t *testing.T) {
	code := strings.ToLower(randx.Letter(8, false))
	req := CreateCouponRequest{
		Code:               code,
		DiscountType:       "percent",
		Amount:             1,
		IndividualUse:      false,
		ExcludeSaleItems:   false,
		MinimumAmount:      2,
		UseageCount:        0,
		UseageLimit:        1,
		UseageLimitPerUser: 1,
		FreeShipping:       true,
		Description:        "TEST",
		DateExpires:        time.UTC.String(),
		EmailRestrictions:  []string{"*@gmail.com"},
	}
	item, err := wooClient.Services.Coupon.Create(req)
	if err != nil {
		t.Fatalf("wooClient.Services.Coupon.Create error: %s", err.Error())
	} else {
		assert.Equal(t, code, item.Code, "code")
	}
}

func TestCouponService_CreateUpdateDelete(t *testing.T) {
	code := gofakeit.LetterN(8)
	req := CreateCouponRequest{
		Code:             code,
		DiscountType:     "percent",
		Amount:           1,
		IndividualUse:    false,
		ExcludeSaleItems: false,
		MinimumAmount:    2,
	}
	var oldItem, newItem entity.Coupon
	var err error
	oldItem, err = wooClient.Services.Coupon.Create(req)
	if err != nil {
		t.Fatalf("wooClient.Services.Coupon.Create error: %s", err.Error())
	} else {
		assert.Equal(t, code, oldItem.Code, "code")
	}

	newItem, err = wooClient.Services.Coupon.One(oldItem.ID)
	if err != nil {
		t.Errorf("wooClient.Services.Customer.One(%d) error: %s", oldItem.ID, err.Error())
	} else {
		updateReq := UpdateCouponRequest{
			Amount:           11,
			IndividualUse:    true,
			ExcludeSaleItems: true,
			MinimumAmount:    22,
		}
		newItem, err = wooClient.Services.Coupon.Update(oldItem.ID, updateReq)
		if err != nil {
			t.Fatalf("wooClient.Services.Coupon.Update error: %s", err.Error())
		} else {
			assert.Equal(t, oldItem.Code, newItem.Code, "code")
			assert.Equal(t, 11.0, newItem.Amount, "Amount")
			assert.Equal(t, true, newItem.IndividualUse, "IndividualUse")
			assert.Equal(t, true, newItem.ExcludeSaleItems, "ExcludeSaleItems")
			assert.Equal(t, 22.0, newItem.MinimumAmount, "MinimumAmount")
		}

		// Only change amount
		updateReq = UpdateCouponRequest{Amount: 11.23}
		newItem, err = wooClient.Services.Coupon.Update(oldItem.ID, updateReq)
		if err != nil {
			t.Fatalf("wooClient.Services.Coupon.Update error: %s", err.Error())
		} else {
			assert.Equal(t, 11.23, newItem.Amount, "Amount")
			assert.Equal(t, true, newItem.IndividualUse, "IndividualUse")
			assert.Equal(t, true, newItem.ExcludeSaleItems, "ExcludeSaleItems")
			assert.Equal(t, 22.0, newItem.MinimumAmount, "MinimumAmount")
		}

		_, err = wooClient.Services.Coupon.Delete(oldItem.ID, true)
		if err != nil {
			t.Fatalf("wooClient.Services.Coupon.Delete(%d) error: %s", oldItem.ID, err.Error())
		} else {
			_, err = wooClient.Services.Coupon.One(oldItem.ID)
			if !errors.Is(err, ErrNotFound) {
				t.Fatalf("wooClient.Services.Coupon.Delete(%d) failed", oldItem.ID)
			}
		}
	}
}

func TestCouponService_Batch(t *testing.T) {
	n := 3
	createRequests := make([]BatchCouponsCreateItem, n)
	codes := make([]string, n)
	for i := 0; i < n; i++ {
		code := strings.ToLower(gofakeit.LetterN(8))
		req := BatchCouponsCreateItem{
			Code:             code,
			DiscountType:     "percent",
			Amount:           float64(i),
			IndividualUse:    false,
			ExcludeSaleItems: false,
			MinimumAmount:    float64(i),
		}
		createRequests[i] = req
		codes[i] = req.Code
	}
	batchReq := BatchCouponsRequest{
		Create: createRequests,
	}
	result, err := wooClient.Services.Coupon.Batch(batchReq)
	if err != nil {
		t.Fatalf("wooClient.Services.Coupon.Batch() error: %s", err.Error())
	}
	assert.Equal(t, n, len(result.Create), "Batch create return len")
	returnCodes := make([]string, 0)
	for _, d := range result.Create {
		returnCodes = append(returnCodes, d.Code)
	}
	assert.Equal(t, codes, returnCodes, "check codes is equal")
}

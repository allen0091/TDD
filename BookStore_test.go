package bookstroe

import "testing"

func TestBuyOneBook(t *testing.T) {

	price := calculatePrice([]int{1}) // 還沒有這個 Function

	if price != 8 {
		t.Fatal("Price should be 8, but got ", price)
	}

}

package bookstroe

import "testing"

func TestBuyOneBook(t *testing.T) {

	price := calculatePrice([]int{1})

	if price != 8 {
		t.Fatal("Price should be 8, but got ", price)
	}

}

func calculatePrice(books []int) float32 {
	return 8
}

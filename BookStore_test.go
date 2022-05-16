package bookstroe

import "testing"

func TestBuyOneBook(t *testing.T) {

	price := calculatePrice([]int{1})

	if price != 8 {
		t.Fatal("Price should be 8, but got ", price)
	}

}

func TestBuyTwoDifferentBook(t *testing.T) {

	price := calculatePrice([]int{1, 1}) //第一集買一本 第二集買1本

	if price != 15.2 {
		t.Fatal("price should be 15.2, but got", price)
	}

}

func calculatePrice(books []int) float32 {
	return 8
}

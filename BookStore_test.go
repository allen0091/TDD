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

func Test_calculatePrice(t *testing.T) {
	type args struct {
		books []int
	}
	tests := []struct {
		name string // test name
		args args
		want float32 // expected value
	}{
		{
			name: "Basic : zero book",
			args: args{books: []int{1}},
			want: 8,
		},
		{
			name: "Basic : Ep1",
			args: args{books: []int{1}},
			want: 8,
		},
		{
			name: "SimpleDiscounts : two different book",
			args: args{books: []int{1, 2}},
			want: 2 * 8 * 0.95,
		},
		{
			name: "SimpleDiscounts : three different book",
			args: args{books: []int{1, 2, 4}},
			want: 3 * 8 * 0.95,
		},
		{
			name: "SimpleDiscounts : five different book",
			args: args{books: []int{1, 2, 3, 4, 5}},
			want: 5 * 8 * 0.95,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculatePrice(tt.args.books); got != tt.want {
				t.Errorf("calculatePrice() should got %v, but got %v", tt.want, got)
			}
		})
	}
}

func calculatePrice(books []int) float32 {
	count := len(books)
	if count > 1 {
		return float32(count) * 8 * 0.95
	}
	return 8
}

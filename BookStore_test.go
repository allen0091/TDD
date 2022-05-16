package bookstroe

import (
	"testing"
)

func Test_calculatePrice(t *testing.T) {
	type args struct {
		books []int
	}
	tests := []struct {
		name string // test name
		args args
		want float64 // expected value
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
			want: 3 * 8 * 0.9,
		},
		{
			name: "SimpleDiscounts : five different book",
			args: args{books: []int{1, 2, 3, 4, 5}},
			want: 5 * 8 * 0.75,
		},
		{
			name: "SeveralDiscounts : 2_EP1&1_EP2",
			args: args{books: []int{1, 1, 2}},
			want: 8 + (8 * 2 * 0.95),
		},
		{
			name: "SeveralDiscounts : 2_EP1&1_EP2&2_EP3&1_EP4",
			args: args{books: []int{1, 1, 2, 3, 3, 4}},
			want: (8 * 4 * 0.8) + (8 * 2 * 0.95),
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

func calculatePrice(books []int) float64 {
	epMap := make(map[int]int)
	for i := 0; i < len(books); i++ {
		epMap[books[i]]++
	}

	total := float64(0)
	mapLen := len(epMap)
	for mapLen > 0 {
		temp := 0
		for key, element := range epMap {
			if element > 0 {
				temp += 1
				temp_element := element - 1
				epMap[key] = temp_element
				if temp_element == 0 {
					mapLen--
				}
			}
		}
		total += calculateDiscount(temp)
	}
	return total
}

func calculateDiscount(len int) float64 {
	switch len {
	case 1:
		return 8
	case 2:
		return float64(len) * 8 * 0.95
	case 3:
		return float64(len) * 8 * 0.9
	case 4:
		return float64(len) * 8 * 0.8
	case 5:
		return float64(len) * 8 * 0.75
	default:
		return 0
	}
}

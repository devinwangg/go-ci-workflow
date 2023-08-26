package service

import "testing"

func Test_calculateDiscount(t *testing.T) {
	type args struct {
		price    float64
		discount float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "Test 10% discount on 100",
			args: args{price: 100.0, discount: 10.0},
			want: 90.0,
		},
		{
			name: "Test 20% discount on 200",
			args: args{price: 200.0, discount: 20.0},
			want: 160.0,
		},
		{
			name: "Test 0% discount on 100",
			args: args{price: 100.0, discount: 0.0},
			want: 100.0,
		},
		{
			name: "Test 100% discount on 100",
			args: args{price: 100.0, discount: 100.0},
			want: 0.0,
		},
		{
			name: "Test invalid discount -10%",
			args: args{price: 100.0, discount: -10.0},
			want: -1.0,
		},
		{
			name: "Test invalid discount 110%",
			args: args{price: 100.0, discount: 110.0},
			want: -1.0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalculateDiscount(tt.args.price, tt.args.discount); got != tt.want {
				t.Errorf("calculateDiscount() = %v, want %v", got, tt.want)
			}
		})
	}
}

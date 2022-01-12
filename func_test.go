package main

import (
	"math/big"
	"reflect"
	"testing"
)

func Test_fibbonaci(t *testing.T) {
	huge_num, _ := new(big.Int).SetString("222232244629420445529739893461909967206666939096499764990979600", 10)

	type args struct {
		n *big.Int
	}
	tests := []struct {
		name string
		args args
		want *big.Int
	}{
		{"boundry_0", args{big.NewInt(0)}, big.NewInt(0)},
		{"boundry_1", args{big.NewInt(1)}, big.NewInt(1)},
		{"boundry_2", args{big.NewInt(2)}, big.NewInt(1)},
		{"small_number_5", args{big.NewInt(5)}, big.NewInt(5)},
		{"big_number_300", args{big.NewInt(300)}, huge_num},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fibbonaci(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fibbonaci() = %v, want %v", got, tt.want)
			}
		})
	}
}

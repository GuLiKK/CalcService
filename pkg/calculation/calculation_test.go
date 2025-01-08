package calculation_test

import (
	"testing"

	"github.com/GuLiKK/CalcService/pkg/calculation"
)

func TestCalc(t *testing.T) {
	cases := []struct {
		expr string
		want float64
		err  bool
	}{
		{"-5+3", -2, false},
		{"--5", 0, true},
		{"2*-3", -6, false},
		{"(-3)*(-2)", 6, false},
		{"-3+(-2)", -5, false},
		{"-3*(-2)", 6, false},
		{"-3/-1", 3, false},
		{"3+-5", -2, false},
		{"1+1", 2, false},
		{"2*3", 6, false},
		{"10-5", 5, false},
		{"10-15", -5, false},
		{"0*100", 0, false},
		{"2/0", 0, true},
		{"1/2", 0.5, false},
		{"3+4", 7, false},
		{"3.5+4.7", 8.2, false},
		{"3.5/0", 0, true},
		{"30/3", 10, false},
		{"30/5", 6, false},
		{"(5+5)*2", 20, false},
		{"5+(5*2)", 15, false},
		{"((5+5)*2)", 20, false},
		{"5*(5+2", 0, true},
		{"2*2+2*2", 8, false},
		{"0/0", 0, true},
		{"(2+2)(3+3)", 0, true},
		{"(2+2)*(3+3)", 24, false},
		{"100*0.5", 50, false},
		{"100.5*2", 201, false},
		{"9999999999999999+1", 1.0e16, false},
		{"3*-5", -15, false},
		{"3/2", 1.5, false},
		{"8/4", 2, false},
		{"(1+1)+(2*2)", 6, false},
		{"((1+1)+(2*2))", 6, false},
		{"((1+1)+(2*2)", 0, true},
		{"1.1+2.2", 3.3, false},
		{"1.1+2.2.3", 0, true},
		{"((1.1+2.2)*3.3)", 10.89, false},
		{"2^3", 0, true},
		{"2--3", 0, true},
		{"((3-2)", 0, true},
		{"3----2", 0, true},
		{"2+2*2", 6, false},
		{"100/3", 33.3333333333, false}, // float rounding can vary
		{"3*(5)-4*6/(10-8)", 3, false},
		{"(3*(5)-4*6)/(10-8)", -4.5, false},
		{"(3*((5)-4)*6)/(10-8)", 9, false},
		{"(((2+3)*4)-5)/2", 7.5, false},
		{"((1+2)*3)/((4-2))", 4.5, false},
		{"40000000000/2", 20000000000, false},
		{"(2+3)+(3+4)-(2*3)/(4-2)", 9, false},
	}
	for _, c := range cases {
		got, e := calculation.Calc(c.expr)
		if (e != nil) != c.err {
			t.Fatalf("expr %q error = %v, wantErr=%v", c.expr, e, c.err)
		}
		if !c.err && got != c.want {
			t.Fatalf("expr %q => %f, want %f", c.expr, got, c.want)
		}
	}
}

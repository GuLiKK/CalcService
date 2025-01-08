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
        {"1+1", 2, false},
        {"2+2*2", 6, false},
        {"2/0", 0, true},
        {"(2+3", 0, true},
    }
    for _, c := range cases {
        got, e := calculation.Calc(c.expr)
        if (e != nil) != c.err {
            t.Fatalf("expr %q error = %v, wantErr=%v", c.expr, e, c.err)
        }
        if !c.err && got != c.want {
            t.Fatalf("%q => %f, want %f", c.expr, got, c.want)
        }
    }
}

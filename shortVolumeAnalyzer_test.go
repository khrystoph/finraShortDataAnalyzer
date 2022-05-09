package main

import (
	"reflect"
	"testing"
)

func Test_tickerCalculations(t *testing.T) {
	type args struct {
		stockTickerData ShortData
	}
	tests := []struct {
		name            string
		args            args
		wantTickerFinal ShortData
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotTickerFinal := tickerCalculations(tt.args.stockTickerData); !reflect.DeepEqual(gotTickerFinal, tt.wantTickerFinal) {
				t.Errorf("tickerCalculations() = %v, want %v", gotTickerFinal, tt.wantTickerFinal)
			}
		})
	}
}

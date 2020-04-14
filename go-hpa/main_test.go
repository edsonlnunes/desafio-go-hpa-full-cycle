package main

import "testing"

//2499999390081494.000000

func TestMakeCalc(t *testing.T){
	got := makeCalc();
	want := 2499999390081494.000000;

	if got != want{
		t.Errorf("makeCalc() \n got: %f \n want: %f", got, want);
	}
}

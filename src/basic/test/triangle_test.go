package main

import "testing"

func TestTriangle(t *testing.T){

	tests:= []struct{
		a,b,c int
	}{
		{3,4,7},
		{5,12,17},
		{12,35,47},
	}

	for _, tt:=range tests  {
		if actual := Add(tt.a,tt.b); actual!=tt.c{
			t.Errorf("sum (%d,+ %d)= %d",tt.a,tt.b,tt.c)
		}
	}
}



func BenchmarkSubstr(b *testing.B){
	aa:= 9
	bb:=10
	cc:=19


	for i:=0;i<b.N;i++  {
		if actual := Add(aa,bb); actual!=cc{
			b.Errorf("sum (%d,+ %d)= %d",aa,bb,cc)
		}
	}

}
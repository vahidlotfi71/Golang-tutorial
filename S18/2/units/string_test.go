package units

import "testing"

func BenchmarkConcatStringWithoutBuilder(b *testing.B){
	for i := 0 ; i <b.N; i++ {
		ConcatStringWithoutBuilder("test_1","test_2")
	}
}


func BenchmarkConcatStringWithBuilder(b *testing.B){
	for i:= 0 ; i< b.N ; i++ {
		ConcatStringWithBuilder("test_1","test_2",)
	}
}

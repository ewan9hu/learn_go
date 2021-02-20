package iteration

import (
	"fmt"
	"testing" )


func TestRepeat(t *testing.T){
	repeated := Repeat("a", 4)
	expected := "aaaa"
	if repeated != expected{
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func Repeat(character string, length int) string{
	var repeated  string
	for i := 0; i < length; i++ {
		repeated = repeated + character
	}
	return repeated
}

func BenchmarkRepeat(b *testing.B){
	for i := 0; i < b.N; i++ {
		Repeat("a", 4)
	}
}

func ExampleRepeat(){
	repeated := Repeat("c", 3)
	fmt.Printf(repeated)
	// Output: ccc
}
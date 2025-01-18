package integers

import "testing"
import "fmt"

func TestAdder(t *testing.T) {
	got := Add(34, 66)
	want := 100

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func ExampleAdd()  {
  sum := Add(1, 4)
  fmt.Println(sum)
  //Output: 5
}

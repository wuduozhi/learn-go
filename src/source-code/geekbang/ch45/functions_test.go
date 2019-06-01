package testing

import(
	"testing"
	"fmt"
)

func TestSquare(t *testing.T){
	inputs := [...]int{1,2,3}
	expected := [...]int{1,4,9}

	for i := 0;i<len(inputs);i++{
		ret := square(inputs[i])
		if ret != expected[i] {
			t.Error("Error",ret)
		}
	}
}

func TestErrorInCode(t *testing.T){
	fmt.Println("Start")
	t.Error("Error")
	fmt.Println("Finish")
}

func TestFailInCode(t *testing.T){
	fmt.Println("Start")
	t.Fatal("Error")
	fmt.Println("Finish")
}
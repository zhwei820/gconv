package deepcopy

import (
	"fmt"
	"testing"

	"github.com/zhwei820/gconv"
	"github.com/zhwei820/gconv/empty"
)

type Person struct {
	Name string
	Age  int
}
type Person2 struct {
	Name string
	Age  int
}
type Person3 struct {
	Name int
	Age  int
}

func TestCopyMapToStruct(t *testing.T) {

	p := &Person{"John", 30}
	p2 := &Person{"", 40}

	SimpleCopyStruct(p, p2, true)

	fmt.Println("p", gconv.Export(p))

	// =====================================
	SimpleCopyStruct(p, p2)

	fmt.Println("\np", gconv.Export(p))
}

func TestCopyMapToStruct2(t *testing.T) {

	p := &Person{"John", 30}
	p2 := &Person2{"", 40}

	SimpleCopyStruct(p, p2, true)

	fmt.Println("p", gconv.Export(p))

	// =====================================
	SimpleCopyStruct(p, p2)

	fmt.Println("\np", gconv.Export(p))
}
func TestCopyMapToStruct3(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("should not panic: ", err)
		}
	}()

	p := &Person{"John", 30}
	p2 := &Person3{0, 40}

	SimpleCopyStruct(p, p2, true)

	fmt.Println("p", gconv.Export(p))

	// =====================================
	SimpleCopyStruct(p, p2)

	fmt.Println("\np", gconv.Export(p))
}

func Test_IsEmpty(t *testing.T) {
	var a *string
	var b *int
	var c *Person
	p := &Person{"John", 30}

	fmt.Println(empty.IsEmpty(""))
	fmt.Println(empty.IsEmpty("b"))
	fmt.Println(empty.IsEmpty('a')) // char to int
	fmt.Println(empty.IsEmpty(a))
	fmt.Println(empty.IsEmpty(b))
	fmt.Println(empty.IsEmpty(c))
	fmt.Println(empty.IsEmpty(p))
	fmt.Println(empty.IsEmpty("aaa"))

}

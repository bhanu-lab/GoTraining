package jedilevel1

import "fmt"

type dosa int

var mas dosa
var pls int

func main() {
	fmt.Println(mas)
	fmt.Printf("%T\n", mas)
	mas = 99
	fmt.Println(mas)

	pls = int(mas)
	fmt.Println(pls)
	fmt.Printf("%T", pls)
}

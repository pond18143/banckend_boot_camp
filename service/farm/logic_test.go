package farm

import (
	"fmt"
	"testing"
)

func Test_cryptoRandom(t *testing.T) {
	fmt.Println("*****************Test_cryptoRandom**************")

	var White, Blue, Green, Red, Orange, Pink, Black, Rainbow int

	n := 1000
	for i := 1; i < n; i++ {
		randomResult, _ := cryptoRandom(100)
		fmt.Println(randomResult)
		if randomResult <= int64(ApricotWhite) {
			White++
		} else if randomResult <= int64(ApricotBlue) {
			Blue++
		} else if randomResult <= int64(ApricotGreen) {
			Green++
		} else if randomResult <= int64(ApricotRed) {
			Red++
		} else if randomResult <= int64(ApricotOrange) {
			Orange++
		} else if randomResult <= int64(ApricotPink) {
			Pink++
		} else if randomResult <= int64(ApricotBlack) {
			Black++
		} else {
			Rainbow++
		}
	}
	fmt.Println("random : ", n)
	fmt.Println("White : ", White)
	fmt.Println("Blue : ", Blue)
	fmt.Println("Green : ", Green)
	fmt.Println("Red : ", Red)
	fmt.Println("Orange : ", Orange)
	fmt.Println("Pink : ", Pink)
	fmt.Println("Black : ", Black)
	fmt.Println("Rainbow : ", Rainbow)
}

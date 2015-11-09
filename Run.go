package CalculatorZFU
//anıl demirkol
import (
	"fmt"
	"bufio"
	"os"

)

func main() {

	result := 0.0
	var number float64
	var number2 float64
	readerStart := "d"
	option := "1"
	for readerStart == "d" {

		fmt.Println("İşlem yapılacak ilk sayı:")
		// old version fmt.Scanf("%f", &number)
		fmt.Scanf("%f\n", &number)

		fmt.Println("İşlem yapılacak ikinci sayı:")
		fmt.Scanf("%f\n", &number2)

		fmt.Print("İşlem Tipi Seçiniz:  ")
		fmt.Print("Topla(+) ")
		fmt.Print("Çıkar(-) ")
		fmt.Print("Böl(/) ")
		fmt.Println("Çarp(*) ")

		optionTemp := bufio.NewReader(os.Stdin)
		d, _, _ := optionTemp.ReadLine()
		option = string(d)

		switch option {
		case "+": result = sum(number, number2)
		case "-": result = sub(number, number2)
		case "*": result = div(number, number2)
		case "/": result = mul(number, number2)
		default:
			fmt.Println("Hatalı Seçim")
			readerStart = "Çıkış"
		}

		fmt.Println("Sonuç: ", result)
		fmt.Println("Devam etmek için - d - yazın")

		readerResume := bufio.NewReader(os.Stdin)
		readerTemp, _, _ := readerResume.ReadLine()
		readerStart = string(readerTemp)

	}
}

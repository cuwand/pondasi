package phoneNumberHelper

import (
	"fmt"
	"testing"
)

func TestReformatNumber(t *testing.T) {

	fmt.Println(fmt.Sprintf("From %s | To %s", "+6281395957596", ReformatNumber("+6281395957596")))
	fmt.Println(fmt.Sprintf("From %s | To %s", "6281395957596", ReformatNumber("6281395957596")))
	fmt.Println(fmt.Sprintf("From %s | To %s", "+62081395957596", ReformatNumber("+62081395957596")))
	fmt.Println(fmt.Sprintf("From %s | To %s", "62081395957596", ReformatNumber("62081395957596")))
	fmt.Println(fmt.Sprintf("From %s | To %s", "81395957596", ReformatNumber("81395957596")))
	fmt.Println(fmt.Sprintf("From %s | To %s", "081395957596", ReformatNumber("081395957596")))
	//fmt.Println(ReformatNumber("6281395957596"))
	//fmt.Println(ReformatNumber("+62081395957596"))
	//fmt.Println(ReformatNumber("62081395957596"))
	//fmt.Println(ReformatNumber("81395957596"))
	//fmt.Println(ReformatNumber("081395957596"))

}

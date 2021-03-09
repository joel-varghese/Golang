package main

import "fmt"

type account struct {
	amt float32
	bal float32
}
type bank struct {
	accnt []account
}

func (bk *bank) addinterest(acc account, ir float32) {
	bk.accnt = append(bk.accnt, acc)
	for _, v := range bk.accnt {
		v.bal = v.bal * ir
		fmt.Println("%v\n", v.amt-v.bal)
	}
}
func main() {
	accs := bank{}
	acc1 := account{amt: 2000, bal: 50}
	accs.addinterest(acc1, 0.05)
	fmt.Printf("%v", accs)
}

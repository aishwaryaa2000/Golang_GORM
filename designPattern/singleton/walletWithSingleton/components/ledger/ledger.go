package ledger

import (
	"fmt"
	"strconv"
)

type Ledger struct {
	entries []string
}

func (l*Ledger) MakeEntry(accName string, txnType string, amount int) {
	amountStr := strconv.Itoa(amount)
	str := accName + " done a transaction of " + txnType +" Rs" +amountStr
	l.entries = append(l.entries, str)
}

func (l*Ledger)GetAllEntries() {

	if len(l.entries)==0{
		fmt.Println("No entries inside the ledger")
		return
	}
	for i,entry := range l.entries{
		fmt.Println("\n Entry ",i+1,"\n",entry)
	}
}
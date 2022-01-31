package cell

import (
	"testing"
)



func TestNew(t *testing.T){  //Test___()
	newCell := New()
	var actual Mark = newCell.cellMark
	var expected Mark = NoMark

	if actual != expected{
		t.Error("Actual is ",actual,"but expected is ",expected)
	}
}

func TestGetMark(t *testing.T){
	newCell := New()	
	err := newCell.SetMark(XMark)
	var actual Mark = newCell.GetMark()
	var expected Mark=XMark

	if actual!=expected{
		t.Error("Actual is ",actual,"but expected is ",expected,"\nError while setting mark is :",err)
	}

}

func TestSetMark(t *testing.T){
	newCell := New()	
	newCell.SetMark(XMark)
	actualError:=newCell.SetMark(OMark) //Error should come as this cell is already marked
	if actualError==nil{
		t.Error("Error is expected as the cell is already marked")
	}

}


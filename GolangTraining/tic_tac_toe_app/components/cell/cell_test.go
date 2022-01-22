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
	cell,_ :=newCell.setMark(XMark)
	var actual Mark = cell.getMark()
	var expected Mark=XMark

	if actual!=expected{
		t.Error("Actual is ",actual,"but expected is ",expected)
	}


}

func TestSetMark(t *testing.T){
	newCell := New()	
	c,_ :=newCell.setMark(XMark)
	_,actual:=c.setMark(OMark) //Error should come as this cell is already marked
	if actual==nil{
		t.Error("Error is expected as the cell is already marked")
	}

}


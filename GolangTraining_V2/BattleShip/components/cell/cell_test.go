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
	 newCell.SetMark(BattleShip)
	var actual Mark = newCell.GetMark()
	var expected Mark=BattleShip

	if actual!=expected{
		t.Error("Actual is ",actual,"but expected is ",expected)
	}

}

func TestGetSetMarkBulk(t *testing.T){

	var listTest = []struct{ //slice of struct
		input Mark
		expected Mark
	}{
		{
			BattleShip,BattleShip,
		},
		
		{
			Hit,Hit,
		},
		{
			Miss,Miss,
		},
		{
			NoMark,NoMark,
		},
	}

	for _,structAtiIndex := range listTest{
		newCell := New()	

		newCell.SetMark(structAtiIndex.input)
		actual := newCell.GetMark()
		if actual !=structAtiIndex.expected{
			t.Error("Actual status is ",actual,"but expected is ",structAtiIndex.expected)
		}
	}
}


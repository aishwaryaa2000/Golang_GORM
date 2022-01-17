package maths
import "testing"

func TestAdd(t *testing.T){  //Test___()
	actual := add(20,30)
	var expected int64 = 50

	if actual != expected{
		t.Error("Actual is ",actual,"but expected is ",expected)
	}
}

func TestAddBulk(t *testing.T){
	var list = []struct{ //slice of struct
		firstNo int
		secondNo int
		expected int64
	}{
		{
			20,30,50,
		},
		
		{
			-20,30,10,
		},
		{
			-20,-30,-50,
		},
		{
			20,-30,-10,
		},
	}

	for _,str := range list{
		actual := add(str.firstNo,str.secondNo) 
		if actual !=str.expected{
			t.Error("Actual is ",actual,"but expected is ",str.expected)
		}
	}
}
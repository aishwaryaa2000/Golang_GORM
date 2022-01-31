package result
import(
	"testing"
)

// func TestGetStatus(t *testing.T){  //Test___()
// 	actualStatus := GetStatus("XXXOXOOXO")
// 	expectedStatus := Win

// 	if actualStatus != expectedStatus{
// 		t.Error("Actual is ",actualStatus,"but expected is ",expectedStatus)
// 	}
// }

func TestGetStatusBulk(t *testing.T){
	var listTest = []struct{ //slice of struct
		strBoard string
		expected Status
	}{
		{
			"XXXOXOXOX",Win,
		},
		
		{
			"X-----O--",Running,
		},
		{
			"XOXXOOOXX",Tie,
		},
		{
			"XOXOXOXOXOXOX---",Win,
		},
	}

	for _,iIndex := range listTest{
		actual := GetStatus(iIndex.strBoard) 
		if actual !=iIndex.expected{
			t.Error("Actual status is ",actual,"but expected is ",iIndex.expected)
		}
	}
}

func TestRowWinBulk(t *testing.T){
	var listTest = []struct{ //slice of struct
		strBoard string
		expected bool
	}{
		{
			"XXXOXOXOX",true,
		},
		
		{
			"X-----O--",false,
		},
		{
			"XOXXOOOXX",false,
		},
		{
			"XOXOXOXOXOXOX---",false,
		},
		{
			"X-XOOO-X-",true,
		},
	}

	for _,iIndex := range listTest{
		actual := checkRowWin(iIndex.strBoard) 
		if actual !=iIndex.expected{
			t.Error("Actual status is ",actual,"but expected is ",iIndex.expected)
		}
	}
}
 
func TestColWinBulk(t *testing.T){
	var listTest = []struct{ //slice of struct
		strBoard string
		expected bool
	}{
		{
			"XXOOXOXXO",true,
		},
		
		{
			"X-----O--",false,
		},
		{
			"XOXXOOOXX",false,
		},
		{
			"XOXOXO--XO--X---",true,
		},	
	}

	for _,iIndex := range listTest{
		actual := checkColWin(iIndex.strBoard) 
		if actual !=iIndex.expected{
			t.Error("Actual status is ",actual,"but expected is ",iIndex.expected)
		}
	}
}
package player
import(
	"testing"
	"tic_tac_toe_app/components/cell"

)

func TestNew(t *testing.T){  //Test___()
	newPlayer := New(cell.XMark,"Aishwarya Anand")

	var actualMark cell.Mark = newPlayer.Mark
	var expectedMark cell.Mark = cell.XMark

	if actualMark != expectedMark{
		t.Error("Actual is ",actualMark,"but expected is ",expectedMark)
	}

	actualName := newPlayer.Name
	expectedName := "Aishwarya Anand"
	if actualName != expectedName{
		t.Error("Actual is ",actualName,"but expected is ",expectedName)
	}
}

package player
import(
	"testing"

)

func TestNew(t *testing.T){  //Test___()
	newPlayer := New("Aishwarya Anand")
	actualName := newPlayer.Name
	expectedHit := 0
	expectedMiss := 0
	actualHit := newPlayer.Hit
	actualName := newPlayer.Name
	expectedName := "Aishwarya Anand"
	if actualName != expectedName{
		t.Error("Actual is ",actualName,"but expected is ",expectedName)
	}
}

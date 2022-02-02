package board

import (
	"testing"
	"tic_tac_toe_app/components/cell"
)



func TestNew(t *testing.T){  //Test___()
	newBoard := New(4)
	actual  := newBoard.Size
	expected := 4

	// size is proper
	if int(actual) != expected{
		t.Error("Actual is ",actual,"but expected is ",expected)
	}

	// all cells have no mark
	for _,icell := range newBoard.NCells{
		actualMark := icell.GetMark()
		expectedMark := cell.NoMark
		if (actualMark) != expectedMark{
			t.Error("Actual is ",actual,"but expected is ",expected)
		}
	}

	noOfCellsActual := len(newBoard.NCells)
	noOfCellsExpected := 16
	if noOfCellsActual != noOfCellsExpected{
		t.Error("Actual is ",actual,"but expected is ",expected)
	}
}

func TestGetBoard(t *testing.T){
	newBoard := New(4)

	for index:=0;index<=14;index=index+2{
	icell := newBoard.NCells[index]
	icell.SetMark(cell.XMark)
	}
	strBoardActual := newBoard.GetBoard()
	strBoardExpected := "X-X-X-X-X-X-X-X-"
	if strBoardActual != strBoardExpected{
		t.Error("Actual is ",strBoardActual,"but expected is ",strBoardExpected)
	}
}
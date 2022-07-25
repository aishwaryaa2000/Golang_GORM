package command

import (
	"commandDp/receiver"
)
/* 
your commands implement the same interface. Usually it has just a single execution method that takes no parameters. 
This interface lets you use various commands with the same request sender, without coupling it to concrete classes of commands.
*/
//Command interface
type Command interface {
    Execute()
}


//Concrete command
type LikeCommand struct {
    LikeButtonVar receiver.LikeButton
}

func (l *LikeCommand) Execute() {
    l.LikeButtonVar.Like()
}

type DislikeCommand struct {
    LikeButtonVar receiver.LikeButton
}

func (d *DislikeCommand) Execute() {
    d.LikeButtonVar.Dislike()
}

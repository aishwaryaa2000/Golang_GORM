

/* 
Instagram like button when tapped will like the photo
and when tapped again,it will dislike the photo
*/

package main

import(
	"commandDp/command"
	"commandDp/receiver"

)

type Button struct {
    command command.Command
}

func (b *Button) tap() {
    b.command.Execute()
}


func main() {
    instaLike := &receiver.InstaDoubleTapButton{}

    likeCommand := &command.LikeCommand{
        LikeButtonVar: instaLike,
    }

    dislikeCommand := &command.DislikeCommand{
        LikeButtonVar: instaLike,
    }

    likeButton := &Button{
        command: likeCommand,
    }

	//Invoker object looks for the appropriate object which can handle this command 
	//and passes the command to the corresponding object which executes the command.

    likeButton.tap()

    dislikeButton := &Button{
        command: dislikeCommand,
    }
    dislikeButton.tap()
}


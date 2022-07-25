package receiver

import(
	"fmt"
)

//Receiver interface
type LikeButton interface {
    Like()
    Dislike()
}




//concrete receiver
type InstaDoubleTapButton struct {
    beingLiked bool
}

func (i *InstaDoubleTapButton) Like() {
	i.beingLiked = true
    fmt.Println("You have liked this pic")
}

func (i *InstaDoubleTapButton) Dislike() {
    i.beingLiked = false
    fmt.Println("You have disliked this pic")
}

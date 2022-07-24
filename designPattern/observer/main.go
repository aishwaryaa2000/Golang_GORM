package main

import(
	"observerDP/observer"
	"observerDP/subject"

)

func main() {

    nodeA := subject.NewNode("nodeAsender")

    observerFirst := &observer.ObserverNode{Name: "nodeBlistener"}
    observerSecond := &observer.ObserverNode{Name: "nodeClistener"}

    nodeA.Register(observerFirst)
    nodeA.Register(observerSecond)

    nodeA.Updateload(3)
	//nodeBlistener and nodeClistener are notified

	nodeA.Deregister(observerSecond)
	//if observerSecond has failed so remove it from the network

	nodeA.Updateload(4)
	//now only nodeBlistener is notified


}
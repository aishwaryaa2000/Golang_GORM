/* 
A particular node in distributed computing system-concrete subject-Node
If a particular node ka load changes then notification is sent to all the subscribers
*/
package subject

import (
    "fmt"
    "observerDP/observer"
)
type Subject interface {
    //Subject, the instance which publishes an event when anything happens.
    Register(observer observer.Observer)
    Reregister(observer observer.Observer)
    notifyAll()
}

type Node struct {
    // Concrete subject
    observerList []observer.Observer
    name         string //name of node
    load      int
}

func NewNode(name string) *Node {
    return &Node{
        name: name,
		load: 0,
    }
}
func (n *Node) Updateload(newLoad int) {
    str := "Load of Node" + n.name + " has now changed " + string(n.load)
    fmt.Println(str)
    n.load = newLoad
    n.notifyAll()
}
func (n *Node) Register(o observer.Observer) {
    n.observerList = append(n.observerList, o)
}

func (n *Node) Deregister(o observer.Observer) {
    n.observerList = removeFromslice(n.observerList, o)
}

func (n *Node) notifyAll() {
    for _, observer := range n.observerList {
        observer.Update(n.name)
    }
}

func removeFromslice(observerList []observer.Observer, observerToRemove observer.Observer) []observer.Observer {
    observerListLength := len(observerList)
    for i, observer := range observerList {
        if observerToRemove.GetName() == observer.GetName() {
            observerList[observerListLength-1], observerList[i] = observerList[i], observerList[observerListLength-1]
            return observerList[:observerListLength-1]
        }
    }
    return observerList
}
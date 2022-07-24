package observer

import (
    "fmt"
)

type Observer interface {
    Update(string)
    GetName() string
}

//Observer, which subscribes to the subject events and gets notified when they happen


type ObserverNode struct {
    Name string
}

func (o *ObserverNode) Update(subjectNode string) {
    fmt.Printf("Sending message to node %s for load change in  %s node\n", o.Name, subjectNode)
}

func (o *ObserverNode) GetName() string {
    return o.Name
}

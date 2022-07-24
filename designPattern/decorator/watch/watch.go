package watch

type SmatWatch interface {
	//component interface
    GetPrice() int
}

type BoatSmartWatch struct {
	//Concrete component
}

func (b *BoatSmartWatch) GetPrice() int {
    return 4000
}
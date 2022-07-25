package product

type Food interface {
   
    GetName() string
    GetCalorie() int
}


type FastFood struct {
    Name  string
    Calorie int
}


func (f *FastFood) GetName() string {
    return f.Name
}

func (f *FastFood) GetCalorie() int {
    return f.Calorie
}

type Waffle struct {
    FastFood
}

func NewWaffle() Food {
    return &Waffle{
        FastFood: FastFood{
            Name:  "BelgianWaffle",
            Calorie: 50,
        },
    }
}

type Sandwich struct {
    FastFood
}

func NewSandwich() Food {
    return &Sandwich{
        FastFood: FastFood{
            Name:  "CheeseSandwich",
            Calorie: 40,
        },
    }
}
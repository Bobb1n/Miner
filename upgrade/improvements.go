package upgrade

type Upgrade interface {
	BuyUpgrade()
}

type Pickaxe struct {
	name  string
	price int
	count int
}

func NewPickaxe() *Pickaxe {
	return &Pickaxe{
		name:  "Кирка",
		price: 3000,
		count: 0,
	}
}
func (m *Pickaxe) BuyUpgrade() {
	m.count++
}

type Ventilation struct {
	name  string
	price int
	count int
}

func NewVentilation() *Ventilation {
	return &Ventilation{
		name:  "Вентиляция в шахту",
		price: 15000,
		count: 0,
	}
}
func (m *Ventilation) BuyUpgrade() {
	m.count++
}

type Trolley struct {
	name  string
	price int
	count int
}

func NewTrolley() *Trolley {
	return &Trolley{
		name:  "Вентиляция в шахту",
		price: 15000,
		count: 0,
	}
}
func (m *Trolley) BuyUpgrade() {
	m.count++
}

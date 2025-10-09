package upgrade

type Upgrade interface {
	BuyUpgrade() int
	GetName() string
	GetCount() int
	GetPrice() int
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
func (m *Pickaxe) BuyUpgrade() int {
	m.count++
	return 3000
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
func (m *Ventilation) BuyUpgrade() int {
	m.count++
	return 15000
}

type Trolley struct {
	name  string
	price int
	count int
}

func NewTrolley() *Trolley {
	return &Trolley{
		name:  "Вагонетка",
		price: 50000,
		count: 0,
	}
}
func (m *Trolley) BuyUpgrade() int {
	m.count++
	return 50000
}

func (m *Pickaxe) GetName() string { return m.name }
func (m *Pickaxe) GetCount() int   { return m.count }
func (m *Pickaxe) GetPrice() int   { return m.price }

func (m *Ventilation) GetName() string { return m.name }
func (m *Ventilation) GetCount() int   { return m.count }
func (m *Ventilation) GetPrice() int   { return m.price }

func (m *Trolley) GetName() string { return m.name }
func (m *Trolley) GetCount() int   { return m.count }
func (m *Trolley) GetPrice() int   { return m.price }

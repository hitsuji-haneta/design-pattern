package manager

type Product interface {
	Use(string) string
}

type Manager struct {
	showcase map[string]Product
}

func New() *Manager {
	return &Manager{map[string]Product{}}
}

func (m *Manager) Register(key string, p Product) {
	m.showcase[key] = p
}

func (m *Manager) Create(key string) Product {
	copy := m.showcase[key]
	return copy
}

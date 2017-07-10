package manager

import (
	"github.com/minha-cidade/transparencia-crawler/transparencia"
	"github.com/minha-cidade/transparencia-crawler/db"
	"log"
)

type Manager struct {
	Transparencias []transparencia.Transparencia
}

func New() *Manager {
	return &Manager{
		Transparencias: make([]transparencia.Transparencia, 0, 10),
	}
}

func (m *Manager) Add(t transparencia.Transparencia) {
	m.Transparencias = append(m.Transparencias, t)
}

func (m *Manager) Run() {
	for _, t := range m.Transparencias {
		log.Printf("Rodando crawler para %s (%s)\n", t.Cidade(), t.Estado())
		db.EnviarDespesaAnual(t.Handle())
	}
}

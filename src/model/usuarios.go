package model

import "time"

type Usuario struct {
	ID       uint64    `json:"id, omitempty"`
	Nome     string    `json:"nome, omitempty"`
	Email    string    `json:"email, omitempty"`
	Senha    string    `json:"senha, omitempty"`
	criadoEm time.Time `json:"criadoEm, omitempty"`
}

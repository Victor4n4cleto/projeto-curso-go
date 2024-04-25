package modelos

import "time"


// estrutura que molda o json que vai vir para struc e vise e versa
type Usuario struct {
	ID        uint   `json: "id,omitempty"`
	Nome      string `json: "nome,omitempty"`
	Nick      string `json: "nick,omitempty"`
	Email     string `json: "email,omitempty"`
	Senha     string `json: "senha,omitempty"`
	CriandoEm time.Time `json: "CriadoEm,omitempty"`
}
package repositorios

import (
	"api/src/modelos"
	"database/sql"
)

type usuarios struct {
	db *sql.DB
}

// essa funçao cria um repositorios de usuarios
func NovoRepositioDeUsuarios(db *sql.DB) *usuarios {
	return &usuarios{db}
}


// cria e insere um usuario ao banco de dados
func (u usuarios) Criar(usuario modelos.Usuario) (uint64, error) {
	return 0, nil
}
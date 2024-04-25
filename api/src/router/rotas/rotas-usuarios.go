package rotas

import (
	"api/src/controllers"
	"net/http"
)

var rotasUsuarios = []Rotas{
	{
		URI:    "/usuarios", // essa cadastra os usuarios
		Metodo: http.MethodPost,
		Funcao: controllers.CriarUsuario,
		RequerAutenticacao: false,
	},

	{
		URI:    "/usuarios",
		Metodo: http.MethodGet, // essa pesquisa os usuarios
		Funcao: controllers.BuscarTodosUsuarios,
		RequerAutenticacao: false,
	},

	{
		URI:    "/usuarios/{usuarioId}", 
		Metodo: http.MethodGet, // pesquisa o usuario pelo ID 
		Funcao: controllers.BuscarUsuario,
		RequerAutenticacao: false,
	},

	{
		URI:    "/usuarios/{usuarioId}", // atualiza os dados de determinado user
		Metodo: http.MethodPut,
		Funcao: controllers.AtualizarUsuario,
		RequerAutenticacao: false,
	},

	{
		URI:    "/usuarios/{usuarioId}", // delete os dados do user
		Metodo: http.MethodDelete,
		Funcao: controllers.DeleterUsuario,
		RequerAutenticacao: false,
	},
}


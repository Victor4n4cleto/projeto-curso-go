package controllers

import "net/http"


// essa func cria usuarios
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Criar usuario!"))
}

// essa func busca todos os usuarios
func BuscarTodosUsuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando todos os Usuarios"))
}

// essa func busca um usuarios especifico
func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando Usuario"))
}

// essa func atualizar os dados do usuario
func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Atualziar usuario!"))
}

// essa func exclu os dados do usuario selecionado
func DeleterUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deletar usuario!"))
}
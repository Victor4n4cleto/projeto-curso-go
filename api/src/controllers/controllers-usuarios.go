package controllers

import (
	"api/src/banco"
	"api/src/modelos"
	"api/src/repositorios"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// essa func cria usuarios
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	corpoRequest, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		log.Fatal(erro)
	}

	var usuario modelos.Usuario
	if erro = json.Unmarshal(corpoRequest, &usuario); erro != nil {
		log.Fatal(erro)
	}

	db, erro := banco.Conectar()
	if erro != nil {
		log.Fatal(erro)
	}

	repositorio := repositorios.NovoRepositioDeUsuarios(db)
	repositorio.Criar(usuario)
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
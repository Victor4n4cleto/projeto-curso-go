package respostas

import (
	"encoding/json"
	"log"
	"net/http"
)

// as funçoes a seguir retornam uma resposta em JSON para a requisiçao
func JSON(w http.ResponseWriter, statusCode int, dados interface{}) {
	w.WriteHeader(statusCode)

	if erro := json.NewEncoder(w).Encode(dados); erro != nil {
		log.Fatal(erro)
	}
}

func Erro(w http.ResponseWriter, statusCode int, erro error) {
	JSON(w, statusCode, struct {
		Erro string `json: "erro"`
	}{
		Erro: erro.Error(),
	})
}
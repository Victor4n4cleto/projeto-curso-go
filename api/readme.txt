Explicações do sistema

1- todos os arquivo filhos subjacentes tem o nome da pasta mae sobrejacente exemplo:
src/router os arquivos localizados dentro desse path vai ter o nome assim router-NOMEDOARQUIVO

2 - rotas URI configuradas:
    localhost:8080/usuarios - GET
    localhost:8080/usuarios/id - GET
    localhost:8080/usuarios - POST
    localhost:8080/usuarios/id - PUT
    localhost:8080/usuarios - DELETE

problema - erro de syntax no SQL - RESOLVIDO

3 - Implementamos as variaveis de ambiente e utilizando godotenv package externo
    e agora uma função CARREGAR lê o .env e executa no main do codigo para abrir o DB
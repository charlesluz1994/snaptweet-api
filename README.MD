Rede Social tipo threads:
Possível criar publicacoes que contenham apenas textos



Domain Usuários:

CRUD
Seguir outro Usuários
Parar de seguir outro Usuário
Buscar todos usuarios que segue
Buscar todos usuarios que sao seguidos
Atualizar senha


2 tables:
users
followers


Domain Publicacoes:
CRUD
Buscar publicacoes de acordos com os usuarios que segue
Curtir publicacoes

Table:
publications


Packages:
Main: executa a app
Router: configura rotas
Controllers: Funcoes que sao executadas pelas rotas.
Modelos: entidades ou seja, structs e seus metodos.
Repositorios: interacao com o banco de dados

Packages Auxiliares:
Config: todas configs
Banco: abrir conexao com o db
Autenticacao_. criacao de login e tokens
Middleware: conjunto com Autenticacao, camada entre req e resposta, verificar se user ja esta authenticado.
Seguranca: validar senhas entre req e banco, etc..
respostas: respostas ao front.


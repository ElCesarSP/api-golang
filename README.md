# 🚀 Go API

API REST desenvolvida em Go utilizando Gin, Prisma ORM, PostgreSQL, JWT Authentication e Docker.

## 📚 Tecnologias

- Go
- Gin Gonic
- Prisma Client Go
- PostgreSQL
- JWT Authentication
- Docker
- bcrypt

---

# 📁 Estrutura do Projeto

```bash
.
├── cmd/
│   └── server/
├── internal/
│   ├── controller/
│   ├── middleware/
│   ├── model/
│   ├── repository/
│   ├── routes/
│   ├── service/
│   └── utils/
├── prisma/
├── Dockerfile
├── docker-compose.yml
├── .env
└── README.md

===

⚙️ Funcionalidades

✅ Cadastro de usuários
✅ Login com JWT
✅ Hash de senha com bcrypt
✅ Rotas protegidas
✅ Middleware JWT
✅ Busca de usuário autenticado
✅ Integração com PostgreSQL
✅ Docker support

---

🔐 Autenticação

A API utiliza JWT Bearer Token.

Header necessário
Authorization: Bearer SEU_TOKEN

---

📦 Instalação
Clone o projeto
git clone https://github.com/ElCesarSP/api-golang.git
Entre na pasta
cd api-golang
---

🐳 Rodando com Docker
Subir containers
docker compose up --build
--- 
▶️ Rodando sem Docker
Instalar dependências
go mod tidy
Gerar Prisma Client
npx prisma generate
Rodar migrations
npx prisma migrate dev
Iniciar API
go run cmd/server/main.go
---

Variáveis de Ambiente

Crie um arquivo .env

DATABASE_URL="postgresql://postgres:postgres@localhost:5432/go_api?schema=public"

JWT_SECRET="super_secret_key"

PORT=8080
📌 Rotas
Health Check
GET /ping
{
  "message": "pong"
}
---
Criar usuário
POST /users
{
  "name": "César",
  "email": "cesar@email.com",
  "password": "123456"
}
Login
POST /login
{
  "email": "cesar@email.com",
  "password": "123456"
}
Response
{
  "token": "jwt_token"
}

---

🗄️ Banco de Dados

A aplicação utiliza PostgreSQL com Prisma ORM.

Model User
model User {
  id        String   @id @default(uuid())
  name      String
  email     String   @unique
  password  String
  createdAt DateTime @default(now())
  updatedAt DateTime @updatedAt
}
🔒 Segurança
Senhas criptografadas com bcrypt
JWT Authentication
Middleware de autenticação
Rotas protegidas
===
👨‍💻 Autor

César Santos

GitHub: https://github.com/ElCesarSP
📄 Licença

Este projeto está sob a licença MIT.

---

# Depois faça

```bash id="5’winiq"
git add README.md
git commit -m "docs: add professional README"
git push

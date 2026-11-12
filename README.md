# Colégio Único: primeiro teste da API no Render

Esta é uma **API de teste de conectividade**, não a versão multiusuário final. Não possui rotas de login, alunos ou notas. Não migra dados nem permite acesso de usuários ao banco.

## Publicar no GitHub

Envie `main.go`, `go.mod`, `go.sum` (se existir), `.gitignore` e `README.md` para a raiz do repositório `Sync`. **Não envie `.env` nem credenciais**.

## Criar o Web Service no Render

- Repository: `arthurmagalhaesagsm-source/Sync`
- Branch: `main`
- Language: `Go`
- Build Command: `go mod download && go build -o app .`
- Start Command: `./app`
- Health Check Path: `/healthz`
- Environment Variable: `MONGODB_URI`, com o endereço completo do MongoDB Atlas inserido SOMENTE no painel do Render.

## Teste

Abra `https://SEU-SERVICO.onrender.com/healthz`. Deve exibir `running`.
Abra `https://SEU-SERVICO.onrender.com/readyz`. Deve exibir `database_connected` quando Render acessar o MongoDB.

Se `/readyz` exibir `database_unavailable`, verifique a URI e o Network Access no Atlas. Não cole a URI completa em conversas, issues ou logs.

## Segurança

Crie usuário de banco com função `readWrite` apenas no banco `colegio_unico`. Não use o Atlas Admin. Restrinja a lista de IPs de acesso do Atlas aos IPs de saída do Render aplicáveis ao seu serviço. **Nunca exponha o MongoDB diretamente no EXE**.

Este pacote é apenas um teste de conexão. O próximo desenvolvimento inclui autenticação, permissões, auditoria, sincronização, migração e backups.

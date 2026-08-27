# Pera API

Base independente para o e-commerce da Pera Tech, iniciada com Go e PostgreSQL para atender primeiro a loja de crochê.

## Estado

A base contém configuração por ambiente, pool PostgreSQL, migrations versionadas, health checks e desligamento gracioso. As regras de catálogo serão construídas no [Loop 001](docs/loops/001-fundacao-catalogo.md).

## Desenvolvimento local

1. Copie `.env.example` para `.env` e escolha uma senha local forte para `POSTGRES_PASSWORD`; use a mesma senha em `DATABASE_URL`.
2. Carregue as variáveis do `.env` no seu shell.
3. Inicie o banco: `make postgres-up`.
4. Aplique a estrutura: `make migrate-up`.
5. Inicie a API: `make run`.

Endpoints disponíveis:

- `GET /health/live`: processo em execução;
- `GET /health/ready`: processo e PostgreSQL disponíveis.

## Comandos

- `make test`
- `make migrate-up`
- `make run`

## Contexto de produto

Comece por [docs/CONTEXT.md](docs/CONTEXT.md). Decisões, regras e loops são versionados junto ao código.

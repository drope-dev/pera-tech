# Contexto ativo — Pera API / Pera Tech

> Leia este arquivo no início de cada loop. Ele aponta para as fontes de verdade e evita que decisões sejam inferidas do histórico de conversas.

## Produto

- **Produto:** plataforma de e-commerce da Pera Tech.
- **Primeira cliente:** loja de crochê artesanal.
- **Tese:** começar resolvendo a operação real de uma loja e evoluir para uma plataforma reutilizável, já com programa de afiliados.
- **Estágio:** descoberta e fundação técnica inicial.

## Estado atual

- A base independente usa Go, PostgreSQL, migrations versionadas e health checks; ela ainda não contém regras de catálogo.
- A jornada manual e a arquitetura alvo inicial estão em [discovery-jornada.md](discovery-jornada.md).
- Ainda não há decisão registrada para interface web, autenticação, provedor de pagamento, frete ou hospedagem.

## Ordem de leitura por tipo de trabalho

| Se for trabalhar em... | Leia antes |
| --- | --- |
| qualquer coisa | este arquivo e [visao-produto.md](product/visao-produto.md) |
| uma regra de negócio | [regras-de-dominio.md](product/regras-de-dominio.md) |
| uma decisão técnica | [decisoes.md](architecture/decisoes.md) |
| checkout, pedidos, estoque ou afiliados | [regras-de-dominio.md](product/regras-de-dominio.md) e [baseline.md](security/baseline.md) |
| um loop novo | [README.md](loops/README.md), [indice.md](loops/indice.md) e [template.md](loops/template.md) |

## Regras de contexto

1. Documento marcado como **Confirmado** tem precedência sobre suposições e conversas antigas.
2. Informação ainda não confirmada deve ser escrita como **Hipótese** ou **Decisão pendente**.
3. Toda decisão que afete mais de um loop entra em `docs/architecture/decisoes.md` antes da implementação.
4. Cada loop atualiza seu próprio registro e o índice; não reescreve a visão do produto para registrar progresso.
5. Se houver conflito, pare e registre a divergência como bloqueio — não escolha silenciosamente.

## Próximo loop recomendado

**Loop 001 — fundação do catálogo.** Definir contratos e critérios de aceite para produto, categoria, imagem e disponibilidade; depois evoluir a persistência e a vitrine em uma fatia vertical.

# Loop 001 — fundação do catálogo

## Estado

- **Estado:** proposto
- **Responsável:** Pera Tech
- **Início:** —
- **Última atualização:** 2026-08-26

## Problema e hipótese

**Problema:** a disponibilidade e as informações das peças estão dispersas, dificultando mostrar produtos e evitando vendas duplicadas.

**Hipótese:** se a administradora conseguir cadastrar e publicar um produto com estoque e imagem, o cliente poderá consultar um catálogo confiável sem iniciar uma conversa manual.

## Fatia entregue

- **Usuário:** administradora e cliente final.
- **Fluxo:** cadastrar produto → definir disponibilidade → publicar → cliente lista e abre o produto.
- **Fora de escopo:** carrinho, checkout, variações completas, upload em produção e afiliados.

## Critérios de aceite

- [ ] Administradora cadastra produto com os campos definidos na regra de domínio.
- [ ] Produto indisponível não aparece como comprável na vitrine.
- [ ] Cliente vê lista e detalhe de produto publicado.
- [ ] Alterações de estoque e publicação são rastreáveis.
- [ ] Regras de catálogo possuem testes automatizados.

## Decisões e dependências

- **Decisão tomada:** a persistência será PostgreSQL com migrations versionadas.
- **Decisão pendente:** stack da interface.
- **Dependência:** definir estratégia segura de mídia antes de upload real.
- **Referências:** [visão](../product/visao-produto.md), [regras](../product/regras-de-dominio.md), [segurança](../security/baseline.md).

## Segurança e qualidade

- Validar todos os campos no handler e aplicar limites de tamanho/valores permitidos.
- Restringir cadastro, alteração e publicação a administradora autorizada.
- Não aceitar URL arbitrária para imagem nem confiar no nome enviado em upload.
- Não expor dados administrativos na vitrine pública.

## Plano mínimo

1. Confirmar campos e regras de produto com a dona da loja.
2. Registrar decisões de persistência e interface.
3. Implementar domínio e persistência com testes.
4. Criar o fluxo administrativo mínimo e a vitrine correspondente.
5. Validar o cadastro de produtos reais com a loja.

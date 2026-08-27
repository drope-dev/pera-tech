# Índice de loops

| Loop | Estado | Resultado esperado | Bloqueio atual |
| --- | --- | --- | --- |
| 001 — fundação do catálogo | proposto | produto persistido e exibível para o cliente | definir interface e regras de produto |
| 002 — vitrine e carrinho | futuro | cliente monta carrinho com itens disponíveis | depende do 001 |
| 003 — pedido e checkout | futuro | pedido criado e pagamento confirmado com segurança | escolher gateway e frete |
| 004 — operação de pedidos | futuro | administradora acompanha preparação, envio e estoque | depende do 003 |
| 005 — afiliados essenciais | futuro | venda atribuída gera comissão auditável | regras pendentes de comissão |

## Como atualizar

- `proposto`: ainda não começou; hipótese e bloqueios podem estar incompletos.
- `em andamento`: há uma única pessoa/linha de trabalho responsável; o registro do loop está atualizado.
- `em validação`: solução pronta para teste com a loja.
- `concluído`: resultado e próximos aprendizados registrados.
- `cancelado`: não será feito; registrar o motivo no arquivo do loop.

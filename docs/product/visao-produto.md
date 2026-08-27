# Visão do produto

## Problema

A loja de crochê opera vendas, estoque, pagamento e pós-venda de forma majoritariamente manual. Isso limita o catálogo, aumenta o risco de vender uma peça única duas vezes e torna a operação dependente de conversas dispersas.

## Proposta

A Pera API é a base do primeiro e-commerce da Pera Tech: uma loja online própria para produtos artesanais, com operação administrativa simples e programa de afiliados incorporado desde a primeira versão comercial.

O objetivo de longo prazo é transformar os componentes validados — catálogo, pedido, estoque, entrega, checkout e afiliados — em uma plataforma que possa atender outras lojas. Multi-loja não é requisito do primeiro lançamento.

## Pessoas e resultados

| Pessoa | Resultado desejado |
| --- | --- |
| Cliente final | encontrar uma peça, entender disponibilidade e concluir uma compra com confiança |
| Dona da loja | cadastrar produtos, acompanhar pedido, entrega e estoque sem planilhas ou memória do WhatsApp |
| Afiliado | divulgar uma loja/produto, saber quais vendas foram atribuídas e acompanhar comissão |
| Pera Tech | aprender com uso real e transformar regras validadas em capacidades reutilizáveis |

## Escopo da primeira versão comercial

- catálogo público de produtos e páginas de produto;
- carrinho e checkout;
- pedido, pagamento confirmado e acompanhamento de entrega;
- painel administrativo de produtos, pedidos, estoque e vendas;
- afiliado aprovado, link ou cupom rastreável, comissão pendente/confirmada/revertida e consulta de resultado.

## Fora de escopo inicial

- marketplace aberto de afiliados;
- múltiplas lojas/tenants;
- aplicativo móvel nativo;
- múltiplos armazéns;
- automações complexas de CRM ou atendimento.

## Métricas iniciais

- pedidos concluídos sem intervenção manual;
- divergências de estoque;
- tempo entre pedido pago e despacho;
- percentual de pedidos atribuídos corretamente a afiliados;
- quantidade de correções manuais de comissão.

## Princípios de produto

1. **Operação real vence sofisticação.** Uma tela que a dona da loja usa vale mais que um módulo genérico não validado.
2. **Automatizar depois de compreender.** Um passo manual pode existir no começo se estiver explícito e auditável.
3. **Dinheiro e estoque são fonte de verdade no servidor.** A interface nunca decide preço, disponibilidade, status ou comissão.
4. **Afiliado é parte do fluxo comercial, não um adendo.** Atribuição, cancelamento e comissão fazem parte do pedido.

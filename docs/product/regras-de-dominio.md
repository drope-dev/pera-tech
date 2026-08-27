# Regras de domínio

Legenda: **Confirmado** representa uma regra decidida; **Hipótese** precisa de validação com a operação; **Pendente** bloqueia uma implementação definitiva.

## Produto e estoque

| Regra | Estado |
| --- | --- |
| Um produto possui nome, descrição, preço, imagens, status de publicação e disponibilidade. | Confirmado |
| Peças artesanais podem ser únicas; estoque não pode ficar negativo. | Confirmado |
| Variações (cor, tamanho e similares) existem, mas o modelo inicial ainda será definido. | Pendente |
| Produto sob encomenda terá prazo e estoque próprios. | Hipótese |

## Pedido e pagamento

| Regra | Estado |
| --- | --- |
| O preço, frete, desconto e total são calculados no servidor a partir do carrinho validado. | Confirmado |
| Um pedido só é considerado pago após confirmação confiável do provedor de pagamento. | Confirmado |
| Um mesmo evento de pagamento não pode confirmar o pedido duas vezes. | Confirmado |
| Cancelamento ou reembolso altera o pedido e pode afetar estoque e comissão. | Confirmado |
| Gateway de pagamento e métodos aceitos. | Pendente |

## Entrega

| Regra | Estado |
| --- | --- |
| O pedido possui status operacional de preparação, postagem, envio e entrega. | Confirmado |
| Frete será cotado e rastreado por integração; o provedor ainda não foi escolhido. | Pendente |
| Atendimento manual por WhatsApp segue como apoio no MVP. | Confirmado |

## Afiliados

| Regra | Estado |
| --- | --- |
| Um afiliado precisa ser aprovado antes de receber comissão. | Confirmado |
| Atribuição acontece por link e/ou cupom; a regra de precedência será documentada antes de codificar. | Confirmado |
| Comissão nasce como pendente e só fica disponível após o prazo de segurança definido para o pedido. | Confirmado |
| Cancelamento/reembolso reverte a comissão associada. | Confirmado |
| Percentual, janela de atribuição, prazo de liberação e forma de pagamento. | Pendente |

## Estados que devem ter transição controlada

`carrinho → pedido_criado → aguardando_pagamento → pago → em_preparacao → enviado → entregue`

Terminais ou de exceção: `cancelado`, `reembolsado`.

Toda transição será validada no servidor, registrada com data/origem e testada. Não é permitido que cliente, administrador ou afiliado envie diretamente um estado arbitrário.


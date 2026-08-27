# Baseline de segurança

Esta é uma lista de requisitos de produto e engenharia. Ela orienta os loops; não substitui revisão de código, testes ou obrigações legais.

## Regras não negociáveis

- Segredos de pagamento, banco e integrações ficam fora do repositório e são injetados por ambiente.
- O backend não recebe nem persiste dados brutos de cartão; o pagamento usa tokenização do provedor escolhido.
- Preço, desconto, frete, estoque, transição de pedido e comissão são recalculados e validados no servidor.
- Papéis de cliente, administrador e afiliado são autorizados no servidor. Um identificador recebido na URL ou corpo não prova posse de pedido, comissão ou recurso administrativo.
- Webhooks de pagamento e logística precisam validar autenticidade, suportar reenvio e ser idempotentes.
- Integrações externas usam destino permitido, HTTPS, timeout explícito e sem seguir redirecionamentos automaticamente.
- Logs não registram senha, token, dados financeiros, endereço completo, e-mail ou telefone sem necessidade e proteção apropriada.
- Upload de imagem valida tamanho, tipo real do arquivo, extensão e nome gerado no servidor.

## Ameaças prioritárias por domínio

| Domínio | Risco a prevenir |
| --- | --- |
| Pedido | acesso a pedidos de outro cliente; alteração de total ou status |
| Admin | alteração de preço/estoque por pessoa não autorizada |
| Afiliado | autoatribuição, fraude de cupom, comissão duplicada ou saque indevido |
| Pagamento | webhook falso, confirmação duplicada e exposição de dados de pagamento |
| Integrações | SSRF, timeout infinito e dados externos sem validação |

## Checklist obrigatório de um loop sensível

Para qualquer loop que toque autenticação, pedido, pagamento, dados pessoais, upload, integração ou afiliados, registrar no próprio loop:

1. quais entradas externas existem e como são validadas;
2. quem pode executar cada ação e como a autorização é verificada;
3. quais transições de estado são permitidas;
4. como reenvios/falhas serão tratados;
5. testes de abuso ou acesso indevido incluídos.


# Fluxo atual da loja

> Levantamento inicial da operação da Studio Pera, antes da construção do MVP.
> Afiliadas não fazem parte deste mapa.

## Visão geral

```text
Cliente acessa a loja
  -> navega pelos produtos
  -> adiciona itens ao carrinho
  -> informa endereço, escolhe frete e realiza o pagamento
  -> pagamento é aprovado
  -> venda e cliente ficam disponíveis no painel administrativo
  -> e-mail de confirmação é enviado
  -> Admin prepara o pacote e gera a etiqueta manualmente
  -> Admin paga a etiqueta via Pix
  -> etiqueta é liberada, impressa e colada no pacote
  -> Admin leva o pacote a um ponto de despacho
  -> cliente recebe rastreio, quando disponível
```

## Etapas e responsabilidades

| Etapa | Responsável | Como ocorre hoje |
|---|---|---|
| Catálogo | Admin | Cadastra, edita e remove produtos pelo painel da loja. |
| Compra | Cliente | Escolhe produtos, usa o carrinho e conclui checkout com endereço, frete e pagamento. |
| Pagamento | Provedor atual / sistema | A venda só avança após a aprovação do pagamento. O provedor ainda precisa ser confirmado. |
| Registro | Sistema | Cliente e venda ficam registrados no painel administrativo. O cliente é cadastrado com nome, e-mail e telefone. |
| Confirmação | Sistema | E-mail de confirmação da venda é enviado. |
| Expedição | Admin | Consulta vendas e prepara manualmente os pedidos para despacho. |
| Etiqueta | Admin + provedor de frete | Após o pagamento aprovado, o Admin gera a etiqueta, paga-a via Pix, imprime e cola no pacote. |
| Postagem | Admin | Leva o pacote a um ponto de despacho. Hoje são usadas Jadlog e Loggi. |
| Pós-despacho | Admin / sistema | Envia e-mail com URL de rastreio, quando a informação está disponível. |

## Ferramentas e fatos já conhecidos

- Loja pública: `https://www.studiopera.com.br/`.
- Frete: cotado por API do provedor atualmente usado, com Jadlog e Loggi como transportadoras.
- Etiqueta: comprada manualmente após a aprovação da compra; pagamento atual via Pix.
- Painel administrativo: concentra produtos, vendas e clientes; é um ponto de dor de operação e usabilidade.
- Cupons: são criados e editados pelo Admin.

## Pontos a confirmar no levantamento

1. Qual é o provedor de pagamento atual e quais meios de pagamento estão habilitados.
2. Qual serviço/agregador de frete fornece a API atual.
3. Onde são guardados os dados de estoque e como a disponibilidade é atualizada.
4. Em que condição o rastreio é enviado ao cliente e por qual canal.
5. Como são realizados cancelamentos, trocas, devoluções e reembolsos.
6. Qual é o ponto de despacho mais prático para a residência/operação e seu horário de corte.

## Dores observadas

- Painel administrativo difícil de operar.
- Geração, compra e impressão de etiqueta dependem de ações manuais.
- O despacho depende da proximidade e disponibilidade dos pontos de postagem.

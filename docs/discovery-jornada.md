# 🗺️ Discovery — jornada de venda e operação

> Task: [Discovery: mapear a jornada atual de venda e operação](https://app.notion.com/p/3bcb921d081181fa88b5cdbfccc59685)
> Referência: ADR-001 — produto-escola com migração progressiva.

Dois mapas: o **as-is** (operação manual de hoje) e o **to-be** (arquitetura técnica alvo da `pera-api`). A migração é progressiva: cada fase substitui um pedaço do manual.

---

## 1. As-is — jornada manual atual

> ⚠️ **A validar com quem opera a loja.** Montado por inferência (loja de crochê artesanal); confirmar cada passo antes de tratar como verdade.

```mermaid
flowchart TB
    subgraph disc["Descoberta"]
        insta["Instagram / feed / stories"]
    end
    subgraph atend["Atendimento manual (dona)"]
        dm["DM / WhatsApp"]
        estoque["Consulta estoque<br/>na cabeça / caderno"]
        freteM["Cotação de frete manual<br/>site Correios"]
    end
    subgraph pag["Pagamento manual"]
        pix["PIX chave manual"]
        comp["Cliente envia comprovante"]
        confere["Dona confere comprovante"]
    end
    subgraph log["Logística manual"]
        embala["Embala"]
        posta["Posta nos Correios/transportadora"]
        rastreio["Envia rastreio por WhatsApp"]
    end
    posvenda["Pós-venda: troca/devolução por WhatsApp"]

    insta --> dm --> estoque --> freteM --> pix --> comp --> confere
    confere --> embala --> posta --> rastreio --> posvenda

    classDef manual fill:#ffe9c7,stroke:#c90;
    class estoque,freteM,pix,comp,confere,embala,posta,rastreio manual
```

**Rotinas manuais / dores (evidência da task):**
- Estoque não persistido → risco de vender peça única duas vezes.
- Conciliação de pagamento manual (olhar comprovante).
- Sem catálogo → toda dúvida de preço/disponibilidade vira conversa.
- Sem histórico de pedido → pós-venda depende da memória do chat.

---

## 2. To-be — arquitetura técnica alvo

```mermaid
flowchart TB
    subgraph client["Cliente (browser)"]
        vitrine["Vitrine / Página de produto<br/>lista + destaques"]
        cart["Carrinho"]
        checkoutUI["Checkout"]
        admin["Admin · CRUD produtos"]
    end

    subgraph app["pera-api · Go (monólito)"]
        router["HTTP router + middleware<br/>logs · metrics · tracing · recover"]
        catalogSvc["Catalog svc<br/>listar / destaque / detalhe"]
        cartSvc["Cart svc"]
        shippingSvc["Shipping svc<br/>cotação de frete por CEP"]
        orderSvc["Order/Checkout svc<br/>cadastro cliente + pedido"]
        paymentSvc["Payment svc"]
    end

    subgraph data["Persistência"]
        db[("PostgreSQL<br/>produtos · estoque · pedidos · clientes")]
        blob[["Object storage<br/>imagens de produto"]]
    end

    subgraph ext["Integrações externas · pontos de falha SRE"]
        cep["CEP/Correios · ViaCEP"]
        carriers["Transportadoras<br/>Loggi · JadLog"]
        psp["Gateway pagamento<br/>Cartão + PIX"]
    end

    vitrine & cart & checkoutUI & admin --> router
    router --> catalogSvc & cartSvc & shippingSvc & orderSvc
    catalogSvc --> db & blob
    cartSvc --> db
    shippingSvc --> cep & carriers
    orderSvc --> db
    orderSvc --> paymentSvc --> psp
    orderSvc -. gera envio .-> carriers

    classDef fail fill:#fde,stroke:#c33,stroke-width:2px;
    class cep,carriers,psp fail
```

### Integrações e dados trocados

| Integração | Direção | Dados trocados | Falha impacta |
|---|---|---|---|
| ViaCEP/Correios | out | CEP → endereço/UF | cotação de frete |
| Loggi / JadLog | out (+webhook) | origem, destino, peso/dim → prazo+preço; etiqueta/rastreio | checkout e pós-venda |
| Gateway (Cartão/PIX) | out (+webhook) | valor, cartão/PIX → status, `payment_id` | conclusão do pedido |
| Object storage | out | imagens de produto | vitrine |

### Exceções (fora do happy path)
- **Estoque**: peça normalmente unidade única → reserva no carrinho/checkout p/ evitar venda dupla.
- **Cancelamento**: pago e não enviado → estorno no gateway.
- **Devolução**: pós-entrega → logística reversa + reembolso.
- **Atendimento**: hoje manual (WhatsApp) — candidato a integrar depois.

### Notas SRE / segurança
- Caixas vermelhas = falhas externas → timeout, retry c/ backoff, circuit breaker, **idempotência no webhook de pagamento**.
- Chaves do gateway em secret manager, nunca no código.
- Não trafegar dados de cartão pelo backend (tokenização no gateway) → fora de escopo PCI.

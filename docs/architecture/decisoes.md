# Decisões de arquitetura

Registre aqui somente decisões que atravessam mais de um loop. Uma entrada curta e clara é preferível a uma explicação longa e desatualizada.

## ADR-001 — produto-escola com migração progressiva

- **Status:** aceito
- **Contexto:** a loja já vende de modo manual e será a primeira cliente da Pera Tech.
- **Decisão:** evoluir a operação em fatias verticais, substituindo primeiro os pontos de maior dor e validando cada capacidade em uso real.
- **Consequência:** não construir multi-loja nem marketplace de afiliados antes de validar catálogo, pedido, operação e atribuição para esta loja.

## ADR-002 — monólito modular no início

- **Status:** aceito
- **Contexto:** a equipe e o produto estão no começo; os domínios ainda estão sendo descobertos.
- **Decisão:** manter catálogo, pedidos, estoque e afiliados no mesmo serviço, com módulos explícitos e fronteiras de domínio.
- **Consequência:** simplifica desenvolvimento e depuração. Extrações futuras exigem uma decisão nova baseada em necessidade real.

## ADR-003 — base independente em Go e PostgreSQL

- **Status:** aceito
- **Contexto:** o produto será desenvolvido fora do ecossistema Mercado Livre e deve permanecer simples de executar localmente.
- **Decisão:** usar Go padrão, `net/http`, PostgreSQL e dependências públicas. A configuração sensível é fornecida somente por variáveis de ambiente.
- **Consequência:** não usar Fury, SDKs internos ou credenciais versionadas. A escolha de framework web permanece aberta enquanto o `net/http` atende ao início do produto.

## Decisões pendentes

| Decisão | Por que importa | Quando decidir |
| --- | --- | --- |
| Interface e stack web | define a vitrine e o painel administrativo | antes do loop de vitrine |
| Autenticação/autorização | separa cliente, administrador e afiliado | antes de qualquer área privada |
| Gateway de pagamento | PIX/cartão, confirmação e reembolso | antes do checkout real |
| Frete/transportadora | preço, prazo, etiqueta e rastreio | antes do checkout real |
| Mídia de produto | upload, armazenamento e entrega de imagens | antes de imagens em produção |

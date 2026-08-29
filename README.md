# Pera API

Pera API is the independent backend foundation for **Pera Tech**, an e-commerce platform being built from the day-to-day needs of a small handmade crochet shop. The goal is to make selling, managing orders, and preparing shipments straightforward for a small business first, then evolve the platform without losing that operational clarity.

The project starts with Go and PostgreSQL so the core business data and workflows remain portable, self-hostable, and independent from a specific commerce platform.

> **Current scope:** this repository provides the backend foundation and operational documentation. The product catalog, storefront, checkout, payments, shipping integrations, and admin experience are being delivered incrementally; they are not all available yet.

## The problem

Small businesses often operate their store through an admin panel that is difficult to use and a set of disconnected manual steps. After a customer pays, the seller still needs to find the order, buy a shipping label, pay for it by Pix, print it, drop the parcel off, and share tracking information.

Pera Tech will bring these steps into a small, focused platform:

- a simple public storefront where customers can discover and buy products;
- a clear admin area to manage products, customers, orders, and coupons;
- dependable order state and customer records;
- shipping quotes and label workflows that match the seller's real dispatch process;
- integrations that can be introduced or replaced without coupling the business to one provider.

The first customer is the crochet shop, but the platform is deliberately being designed as a reusable base for similar small businesses.

## How the operation works today

The current workflow is the reference point for the MVP. Affiliate features are intentionally outside this flow.

```text
Customer visits the store
  -> browses products and adds items to the cart
  -> enters an address, chooses shipping, and pays
  -> payment is approved
  -> the system records the customer and sale, then sends a confirmation email
  -> the seller prepares the parcel
  -> the seller manually creates and pays for the shipping label by Pix
  -> the label is released, printed, and attached to the parcel
  -> the seller dispatches the parcel at a carrier drop-off point
  -> the customer receives tracking information when it is available
```

Today, Jadlog and Loggi are used for shipping through the current provider's API. The label purchase and dispatch are manual after payment approval. This is intentional context for the first versions of the platform: automation should improve the work without hiding the operational steps that still need a human decision.

For the detailed as-is map, including responsibilities and open questions, see [the current store flow](docs/fluxo-atual-loja.md).

## Initial architecture

```text
Storefront and Admin UI (planned)
              |
              v
          Pera API
              |
              v
        PostgreSQL database

Future integrations: payment provider, shipping provider, email, and tracking
```

The API currently includes:

- environment-based configuration;
- a PostgreSQL connection pool;
- versioned database migrations;
- liveness and readiness health checks;
- graceful shutdown.

The first business capability is the catalog foundation, specified in [Loop 001](docs/loops/001-fundacao-catalogo.md).

## Run locally

### Prerequisites

- Go 1.25 or later;
- Docker and Docker Compose;
- Make.

### Setup

1. Copy the environment template:

   ```sh
   cp .env.example .env
   ```

2. In `.env`, set a strong local `POSTGRES_PASSWORD` and use that same value in `DATABASE_URL`.
3. Load the environment variables into your shell. For example, with `zsh` or `bash`:

   ```sh
   set -a && source .env && set +a
   ```

4. Start PostgreSQL:

   ```sh
   make postgres-up
   ```

5. Apply the database migrations:

   ```sh
   make migrate-up
   ```

6. Start the API:

   ```sh
   make run
   ```

The API listens on `HTTP_ADDR` (by default, `:8080`).

### Health endpoints

| Endpoint | Meaning |
| --- | --- |
| `GET /health/live` | The API process is running. |
| `GET /health/ready` | The API process is running and PostgreSQL is reachable. |

### Useful commands

```sh
make test
make migrate-up
make run
make postgres-down
```

## Product and engineering context

Read [docs/CONTEXT.md](docs/CONTEXT.md) before starting a new product or implementation loop. It links to the source of truth for product decisions, domain rules, security baseline, and planned work.

The repository keeps product decisions and engineering work close to the code:

- [Product vision](docs/product/visao-produto.md)
- [Domain rules](docs/product/regras-de-dominio.md)
- [Architecture decisions](docs/architecture/decisoes.md)
- [Development loops](docs/loops/README.md)
- [Security baseline](docs/security/baseline.md)

## Future features

The following are planned directions, not promises or implemented functionality:

- product, category, image, and availability management;
- a fast storefront and a simple admin interface;
- cart, checkout, and order lifecycle;
- payment-provider integration with secure payment confirmation;
- shipping quotes, label purchase, dispatch support, and tracking updates;
- customer notifications and order history;
- stock controls, coupons, cancellations, returns, and refunds;
- affiliate capabilities only after the store's core operation is stable;
- deployment guidance for self-hosting and future managed infrastructure.

New features should be proposed and refined through a documented loop, keeping business decisions explicit and preserving the small-business workflow this project is meant to improve.

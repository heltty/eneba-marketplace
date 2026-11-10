# Decisions

## 001 – Scope
- Decision: Build catalog/search, product page with multi-seller offers, auth, cart + simulated checkout, and 3 homepage blocks.
- Reason: Assignment rewards a few features done in depth; these form one complete purchase flow.
- Skipped: Real payments, seller dashboards, reviews, wishlist, email verification.

## 002 – Backend router
- Decision: Go standard library `net/http` with method-based route patterns.
- Reason: Fewer dependencies, easier to explain line by line.
- Skipped: chi/gin (can adopt later if routing gets complex).

## 003 – Price representation
- Decision: Store and transmit prices as integer cents (`priceCents`); format as currency only in the React UI.
- Reason: Avoids floating-point rounding errors in totals and checkout.
- Skipped: Decimal/float fields, a money library.

## 004 – Database
- Decision: Postgres (hosted free on Neon), plain SQL, no ORM.
- Reason: Survives free-tier redeploys; plain SQL is easier to explain than ORM magic.
- Skipped: SQLite (data loss risk on deploy), ORMs like GORM.

## 005 – Platform vs region
- Decision: Each platform version of a game is its own `products` row; region is stored on `offers`.
- Reason: Platform defines the product; region varies per seller's key. Both stay filterable.
- Skipped: Separate `platforms`/`regions` tables (extra joins for little benefit at this scale).

## 006 – Offer uniqueness
- Decision: One offer per (seller, product, region); price > 0 and stock >= 0 enforced by CHECK constraints.
- Reason: Sellers can offer several regional keys of one game; the database blocks duplicates and invalid values.
- Skipped: Per-key inventory table (stock is a counter, keys are generated at purchase time).

## 007 – DB access
- Decision: `database/sql` with the pgx driver; connection string read from `DATABASE_URL` env var.
- Reason: Standard Go interface; secrets stay out of the repository.
- Skipped: ORM (GORM), pgxpool-specific API.

## 008 – Repository
- Decision: Single GitHub monorepo (`backend/`, `frontend/`); secrets only via env vars, `.env` git-ignored.
- Reason: One link for submission; no credentials in history.
- Skipped: Separate frontend/backend repos.

## 009 – Handler dependencies
- Decision: Handlers are functions that take `*sql.DB` and return `http.HandlerFunc` (closures).
- Reason: Explicit dependencies, no global state, easy to read.
- Skipped: Global `db` variable, handler structs with methods (can refactor if the app grows).
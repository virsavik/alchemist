CREATE TABLE IF NOT EXISTS "products" (
  "id"  BIGINT PRIMARY KEY,
  "sku"         TEXT        NOT NULL,
  "name"        TEXT        NOT NULL CONSTRAINT "name_check" CHECK ( "name" <> ''::TEXT ),
  "created_at"  TIMESTAMPTZ NOT NULL DEFAULT now(),
  "updated_at"  TIMESTAMPTZ NOT NULL DEFAULT now(),
  "deleted_at"  TIMESTAMPTZ
);
CREATE UNIQUE INDEX IF NOT EXISTS "sku_on_products" ON "products" ( "sku" );

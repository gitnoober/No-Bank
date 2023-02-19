-- Remove the Unique constraint
ALTER TABLE IF EXISTS "accounts" DROP CONSTRAINT IF EXISTS "owner_currency_key";

-- Remove the FK constraint
ALTER TABLE IF EXISTS "accounts" DROP CONSTRAINT IF EXISTS "accounts_owner_fkey";

-- Drop the users table
DROP TABLE IF EXISTS "users";
CREATE TABLE IF NOT EXISTS "data" ("id" integer primary key autoincrement,"created_at" datetime,"updated_at" datetime,"deleted_at" datetime );
CREATE INDEX idx_data_deleted_at ON "data"(deleted_at) ;

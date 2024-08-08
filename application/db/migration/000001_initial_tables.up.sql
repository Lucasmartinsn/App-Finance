-- Criação da extensão para UUID, se não existir
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Tabela de Users
CREATE TABLE "users" (
    "id" UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    "username" VARCHAR(100) NOT NULL,
    "password" VARCHAR(100) NOT NULL,
    "email" VARCHAR(100) UNIQUE NOT NULL,
    "create_at" TIMESTAMPTZ DEFAULT now() NOT NULL
);

-- Tabela de Categorias
CREATE TABLE "categories" (
    "id" UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    "user_id" UUID REFERENCES "users"(id) ON DELETE CASCADE,
    "title" VARCHAR(100) NOT NULL,
    "type" VARCHAR(100) NOT NULL,
    "description" VARCHAR(255) UNIQUE NOT NULL,
    "create_at" TIMESTAMPTZ DEFAULT now() NOT NULL
);

-- Tabela de Contas
CREATE TABLE "accounts" (
    "id" UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    "user_id" UUID REFERENCES "users"(id) ON DELETE CASCADE,
    "categories_id" UUID REFERENCES "categories"(id) ON DELETE CASCADE,
    "title" VARCHAR(100) NOT NULL,
    "type" VARCHAR(100) NOT NULL,
    "description" VARCHAR(255) UNIQUE NOT NULL,
    "value" INTEGER NOT NULL,
    "date" DATE NOT NULL,
    "create_at" TIMESTAMPTZ DEFAULT now() NOT NULL
);

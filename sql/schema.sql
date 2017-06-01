CREATE DATABASE IF NOT EXISTS teleport CHARACTER SET utf8 COLLATE utf8_general_ci;

USE teleport;

DROP TABLE IF EXISTS groups;
CREATE TABLE groups (
  id        VARCHAR(36) NOT NULL,
  parent_id VARCHAR(36),
  name      VARCHAR(3000),
  KEY record(id, parent_id)
);

DROP TABLE IF EXISTS properties;
CREATE TABLE properties (
  id   VARCHAR(36) NOT NULL,
  name VARCHAR(3000),
  type VARCHAR(3000),
  KEY record(id)
);

DROP TABLE IF EXISTS products;
CREATE TABLE products (
  id          VARCHAR(36) NOT NULL,
  name        VARCHAR(3000),
  description TEXT,
  barcode     VARCHAR(250),
  article     VARCHAR(250),
  full_name   VARCHAR(3000),
  country     VARCHAR(250),
  brand       VARCHAR(250),
  KEY record(id)
);

DROP TABLE IF EXISTS products_properties;
CREATE TABLE products_properties (
  parent_id VARCHAR(36) NOT NULL,
  id        VARCHAR(36) NOT NULL,
  value     VARCHAR(250),
  KEY record(parent_id, id)
);

DROP TABLE IF EXISTS products_taxes;
CREATE TABLE products_taxes (
  parent_id VARCHAR(36) NOT NULL,
  name      VARCHAR(36),
  rate      VARCHAR(20),
  KEY record(parent_id)
);

DROP TABLE IF EXISTS products_requisites;
CREATE TABLE products_requisites (
  parent_id VARCHAR(36) NOT NULL,
  name      VARCHAR(250),
  value     VARCHAR(250),
  KEY record(parent_id)
);

DROP TABLE IF EXISTS products_excises;
CREATE TABLE products_excises (
  parent_id VARCHAR(36) NOT NULL,
  name      VARCHAR(250),
  sum       FLOAT,
  currency  CHAR(3),
  KEY record(parent_id)
);

DROP TABLE IF EXISTS products_images;
CREATE TABLE products_images (
  parent_id VARCHAR(36) NOT NULL,
  url       VARCHAR(250),
  KEY record(parent_id)
);

DROP TABLE IF EXISTS products_groups;
CREATE TABLE products_groups (
  parent_id VARCHAR(36) NOT NULL,
  id        VARCHAR(36) NOT NULL,
  KEY record(parent_id)
);

DROP TABLE IF EXISTS products_contractor;
CREATE TABLE products_contractor (
  parent_id VARCHAR(36) NOT NULL,
  id        VARCHAR(36) NOT NULL,
  name      VARCHAR(250),
  title     VARCHAR(3000),
  full_name VARCHAR(3000),
  KEY record(parent_id, id)
);

DROP TABLE IF EXISTS products_component;
CREATE TABLE products_component (
  parent_id     VARCHAR(36) NOT NULL,
  catalog_id    VARCHAR(36) NOT NULL,
  classifier_id VARCHAR(36) NOT NULL,
  quantity      SMALLINT,
  KEY record(parent_id)
);

DROP TABLE IF EXISTS bundling;
CREATE TABLE bundling (
  id            VARCHAR(36) NOT NULL,
  name          VARCHAR(3000),
  catalog_id    VARCHAR(36),
  classifier_id VARCHAR(36),
  KEY record(id)
);

DROP TABLE IF EXISTS bundling_offers;
CREATE TABLE bundling_offers (
  parent_id        VARCHAR(36) NOT NULL,
  id               VARCHAR(36) NOT NULL,
  name             VARCHAR(3000),
  base_unit        VARCHAR(200),
  base_unit_name   VARCHAR(200),
  base_unit_code   VARCHAR(200),
  base_unit_global VARCHAR(200),
  quantity         FLOAT,
  KEY record(parent_id, id)
);

DROP TABLE IF EXISTS bundling_prices_types;
CREATE TABLE bundling_prices_types (
  parent_id VARCHAR(36) NOT NULL,
  KEY record(parent_id)
);

DROP TABLE IF EXISTS bundling_offers_prices;
CREATE TABLE bundling_offers_prices (
  parent_id     VARCHAR(36) NOT NULL,
  display       VARCHAR(200),
  price_type_id VARCHAR(36) NOT NULL,
  unit_price    VARCHAR(50),
  currency      CHAR(3),
  unit          VARCHAR(50),
  coefficient   SMALLINT,
  KEY record(parent_id)
);

CREATE DATABASE IF NOT EXISTS teleport CHARACTER SET utf8 COLLATE utf8_general_ci;

USE teleport;

DROP TABLE IF EXISTS groups;
CREATE TABLE groups (
  client_id VARCHAR(36) NOT NULL,
  id        VARCHAR(36) NOT NULL,
  parent_id VARCHAR(36),
  name      VARCHAR(3000),
  KEY record(client_id, id, parent_id)
);

DROP TABLE IF EXISTS properties;
CREATE TABLE properties (
  client_id VARCHAR(36) NOT NULL,
  id        VARCHAR(36) NOT NULL,
  name      VARCHAR(3000),
  type      VARCHAR(3000),
  KEY record(client_id, id)
);

DROP TABLE IF EXISTS products;
CREATE TABLE products (
  client_id    VARCHAR(36) NOT NULL,
  id           VARCHAR(36) NOT NULL,
  name         VARCHAR(3000),
  description  TEXT,
  barcode      VARCHAR(250),
  article      VARCHAR(250),
  full_name    VARCHAR(3000),
  groups       TEXT,
  image        VARCHAR(250),
  properties   TEXT,
  taxes        TEXT,
  requisites   TEXT,
  country      VARCHAR(250),
  brand        VARCHAR(250),
  owner_brand  VARCHAR(250),
  manufacturer VARCHAR(36),
  excises      TEXT,
  KEY record(client_id, id)
);

DROP TABLE IF EXISTS products_properties;
CREATE TABLE products_properties (
  client_id VARCHAR(36) NOT NULL,
  parent_id VARCHAR(36) NOT NULL,
  id        VARCHAR(36) NOT NULL,
  value     VARCHAR(250),
  KEY record(client_id, parent_id, id)
);

DROP TABLE IF EXISTS products_taxes;
CREATE TABLE products_taxes (
  client_id VARCHAR(36) NOT NULL,
  parent_id VARCHAR(36) NOT NULL,
  id        VARCHAR(36) NOT NULL,
  name      VARCHAR(36),
  rate      FLOAT,
  KEY record(client_id, parent_id, id)
);

DROP TABLE IF EXISTS products_requisites;
CREATE TABLE products_requisites (
  client_id VARCHAR(36) NOT NULL,
  parent_id VARCHAR(36) NOT NULL,
  id        VARCHAR(36) NOT NULL,
  name      VARCHAR(250),
  value     VARCHAR(250),
  KEY record(client_id, parent_id, id)
);

DROP TABLE IF EXISTS products_excises;
CREATE TABLE products_excises (
  client_id VARCHAR(36) NOT NULL,
  parent_id VARCHAR(36) NOT NULL,
  id        VARCHAR(36) NOT NULL,
  name      VARCHAR(250),
  KEY record(client_id, parent_id, id)
);

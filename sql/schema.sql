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
  id           VARCHAR(36) NOT NULL,
  name         VARCHAR(3000),
  description  TEXT,
  barcode      VARCHAR(250),
  article      VARCHAR(250),
  full_name    VARCHAR(3000),
  groups       TEXT,
  properties   TEXT,
  taxes        TEXT,
  requisites   TEXT,
  country      VARCHAR(250),
  brand        VARCHAR(250),
  owner_brand  VARCHAR(250),
  manufacturer VARCHAR(36),
  excises      TEXT,
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
  id        VARCHAR(36) NOT NULL,
  name      VARCHAR(36),
  rate      FLOAT,
  KEY record(parent_id, id)
);

DROP TABLE IF EXISTS products_requisites;
CREATE TABLE products_requisites (
  parent_id VARCHAR(36) NOT NULL,
  id        VARCHAR(36) NOT NULL,
  name      VARCHAR(250),
  value     VARCHAR(250),
  KEY record(parent_id, id)
);

DROP TABLE IF EXISTS products_excises;
CREATE TABLE products_excises (
  parent_id VARCHAR(36) NOT NULL,
  id        VARCHAR(36) NOT NULL,
  name      VARCHAR(250),
  KEY record(parent_id, id)
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

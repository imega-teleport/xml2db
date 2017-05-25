create database if not exists teleport character set utf8 collate utf8_general_ci;

use teleport;

DROP TABLE IF EXISTS groups;

CREATE TABLE groups (
  client_id VARCHAR(36) NOT NULL,
  id        VARCHAR(36) NOT NULL,
  parent_id VARCHAR(36),
  name      VARCHAR(3000),
  KEY record(client_id, id, parent_id)
);

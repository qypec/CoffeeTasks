drop database shop;
create database shop;
use shop;

CREATE TABLE IF NOT EXISTS category (
  `id` INT NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(128) NOT NULL,
  `discount` TINYINT NOT NULL,
  `alias_name` VARCHAR(128),
  PRIMARY KEY(`id`));

CREATE TABLE IF NOT EXISTS brand (
  `id` INT NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(128) NOT NULL,
  PRIMARY KEY(`id`));

CREATE TABLE IF NOT EXISTS product_type (
  `id` INT NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(128) NOT NULL,
  PRIMARY KEY(`id`));

CREATE TABLE IF NOT EXISTS product (
  `id` INT NOT NULL AUTO_INCREMENT,
  `brand_id` INT NOT NULL,
  `product_type_id` INT NOT NULL,
  `category_id` INT NOT NULL,
  `price` DECIMAL(10, 2) NOT NULL,  
  FOREIGN KEY(brand_id) REFERENCES brand(id) ON DELETE CASCADE ON UPDATE NO ACTION,
  FOREIGN KEY(product_type_id) REFERENCES product_type(id) ON DELETE CASCADE ON UPDATE NO ACTION,
  FOREIGN KEY(category_id) REFERENCES category(id) ON DELETE CASCADE ON UPDATE NO ACTION,
  PRIMARY KEY(`id`));

CREATE TABLE IF NOT EXISTS `order` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `user_name` VARCHAR(128) NOT NULL,
  `phone` VARCHAR(32) NOT NULL,
  `datetime` DATETIME NOT NULL,
  PRIMARY KEY(`id`));

CREATE TABLE IF NOT EXISTS order_products (
  `order_id` INT NOT NULL,
  `product_id` INT NOT NULL,
  `count` INT NOT NULL,
  FOREIGN KEY(order_id) REFERENCES `order`(id) ON DELETE CASCADE ON UPDATE NO ACTION,
  FOREIGN KEY(product_id) REFERENCES product(id) ON DELETE CASCADE ON UPDATE NO ACTION,
  PRIMARY KEY(`order_id`, `product_id`));

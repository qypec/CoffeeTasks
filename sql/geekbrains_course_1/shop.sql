-- shop (before 10 lesson)

CREATE TABLE IF NOT EXISTS product (
  `id` INT NOT NULL AUTO_INCREMENT,
  `brand_id` INT NOT NULL,
  `product_type_id` INT NOT NULL,
  `category_id` INT NOT NULL,
  `price` DECIMAL(10, 2) NOT NULL,
  PRIMARY KEY(`id`));
  
INSERT INTO product(`brand_id`, `product_type_id`, `category_id`, price) VALUES(1, 1, 1, 8999);

CREATE TABLE IF NOT EXISTS category (
  `id` INT NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(128) NOT NULL,
  `discount` TINYINT NOT NULL,
  `alias_name` VARCHAR(128),
  PRIMARY KEY(`id`));
  
INSERT INTO category(`name`, `discount`) VALUES('Женская одежда', '5');
INSERT INTO category(`name`, `discount`) VALUES('Мужская одежда', '0');
INSERT INTO category(`name`, `discount`) VALUES('Женская обувь', '10');
INSERT INTO category(`name`, `discount`) VALUES('Мужская обувь', '15');
INSERT INTO category(`name`, `discount`) VALUES('Шляпы', '0');
  
CREATE TABLE IF NOT EXISTS brand (
  `id` INT NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(128) NOT NULL,
  PRIMARY KEY(`id`));
  
INSERT INTO brand(`name`) VALUES("Marc O'Polo");
INSERT INTO brand(`name`) VALUES("ALCOTT");
INSERT INTO brand(`name`) VALUES("GUESS");
  
CREATE TABLE IF NOT EXISTS product_type (
  `id` INT NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(128) NOT NULL,
  PRIMARY KEY(`id`));

INSERT INTO product_type(`name`) VALUES("Платье");
INSERT INTO product_type(`name`) VALUES("Футболка");



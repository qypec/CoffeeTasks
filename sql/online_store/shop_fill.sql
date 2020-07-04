
INSERT INTO category(`name`, `discount`) VALUES('Женская одежда', '5');
INSERT INTO category(`name`, `discount`) VALUES('Мужская одежда', '0');
INSERT INTO category(`name`, `discount`) VALUES('Женская обувь', '10');
INSERT INTO category(`name`, `discount`) VALUES('Мужская обувь', '15');
INSERT INTO category(`name`, `discount`) VALUES('Шляпы', '0');

INSERT INTO brand(`name`) VALUES("Marc O'Polo");
INSERT INTO brand(`name`) VALUES("ALCOTT");
INSERT INTO brand(`name`) VALUES("GUESS");

INSERT INTO product_type(`name`) VALUES("Платье");
INSERT INTO product_type(`name`) VALUES("Футболка");

INSERT INTO product(`brand_id`, `product_type_id`, `category_id`, price) VALUES(1, 1, 1, 8999);
INSERT INTO product(`brand_id`, `product_type_id`, `category_id`, price) VALUES(1, 2, 1, 7999);
INSERT INTO product(`brand_id`, `product_type_id`, `category_id`, price) VALUES(2, 1, 3, 6999);

INSERT INTO `order`(user_name, phone, datetime) VALUES('Василий', '555-55-55', '2016-05-09 14:20');

INSERT INTO order_products(order_id, product_id, count) VALUES(1, 1, 1);
INSERT INTO order_products(order_id, product_id, count) VALUES(1, 2, 3);
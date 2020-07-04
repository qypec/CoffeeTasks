use shop;

SELECT `order`.user_name as customer, sum(product.price * order_products.`count`) as total_price FROM product
    INNER JOIN order_products ON product.id = order_products.product_id
    INNER JOIN `order` ON order_products.order_id = `order`.id
    GROUP BY `order`.user_name;

use shop;

SELECT product.id, category.name as category, product_type.name as type, brand.name as brand, product.price FROM product
    LEFT JOIN order_products ON product.id = order_products.product_id
    INNER JOIN brand ON product.brand_id = brand.id
    INNER JOIN category ON product.category_id = category.id
    INNER JOIN product_type ON product.product_type_id = product_type.id
    WHERE order_products.order_id IS NULL;

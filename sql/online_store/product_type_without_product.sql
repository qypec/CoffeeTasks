use shop;

SELECT product_type.* FROM product_type
    LEFT JOIN product on product_type.id = product.product_type_id
    WHERE product.id IS NULL;

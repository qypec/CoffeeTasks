SELECT product.id, brand.name, product_type.name, category.name, product.price FROM product
    inner join brand on product.brand_id = brand.id
    inner join product_type on product.product_type_id = product_type.id
    inner join category on product.category_id = category.id
    WHERE product_type.name = 'Футболка';
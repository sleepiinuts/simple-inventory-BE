-- name: GetAll
SELECT * FROM product

-- name: GetAllWithPaging
SELECT id,sku,name,image_url,dep,price FROM (
	SELECT id,sku,name,image_url,dep,price,
	ROW_NUMBER ()OVER (ORDER BY id) rn
	FROM product
)WHERE rn >= :1 AND rn <= :2
ORDER BY id

-- name: CountAll
SELECT count(id) as total FROM product
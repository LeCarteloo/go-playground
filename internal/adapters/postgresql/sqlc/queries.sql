-- name: ListProducts :many
SELECT
    *
FROM
    products;

-- name: GetProductById :one
SELECT
    *
FROM
    products
WHERE
    id = $1;

-- name: ListOrders :many
SELECT
    *
FROM
    orders;

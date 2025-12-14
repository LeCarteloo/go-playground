-- name: ListPosts :many
SELECT
    *
FROM
    posts;

-- name: GetPostById :one
SELECT
    *
FROM
    posts
WHERE
    id = $1;

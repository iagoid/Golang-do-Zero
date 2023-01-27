-- name: CreateCategory :exec
INSERT INTO categories (id, name) VALUES(?,?);

-- name: GetCategory :one
SELECT * FROM categories WHERE id = ?;

-- name: GetCategories :many
SELECT * FROM categories;

-- name: CreateCourse :exec
INSERT INTO courses (id, name, category_id) VALUES(?,?,?);

-- name: GetCourse :one
SELECT * FROM courses WHERE id = ?;

-- name: GetCourses :many
SELECT * FROM courses;

-- name: ListCourseCategories :many
SELECT * FROM courses;

-- name: InsertTemporaryTable :many
SELECT
    ca.id as category_id,
    ca.name as category_name,
    co.id,
    co.name
FROM courses co
INNER JOIN categories ca ON (ca.id = co.category_id)

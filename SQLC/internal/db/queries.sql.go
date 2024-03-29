// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: queries.sql

package db

import (
	"context"
)

const createCategory = `-- name: CreateCategory :exec
INSERT INTO categories (id, name) VALUES(?,?)
`

type CreateCategoryParams struct {
	ID   int32
	Name string
}

func (q *Queries) CreateCategory(ctx context.Context, arg CreateCategoryParams) error {
	_, err := q.db.ExecContext(ctx, createCategory, arg.ID, arg.Name)
	return err
}

const createCourse = `-- name: CreateCourse :exec
INSERT INTO courses (id, name, category_id) VALUES(?,?,?)
`

type CreateCourseParams struct {
	ID         int32
	Name       string
	CategoryID int32
}

func (q *Queries) CreateCourse(ctx context.Context, arg CreateCourseParams) error {
	_, err := q.db.ExecContext(ctx, createCourse, arg.ID, arg.Name, arg.CategoryID)
	return err
}

const getCategories = `-- name: GetCategories :many
SELECT id, name FROM categories
`

func (q *Queries) GetCategories(ctx context.Context) ([]Category, error) {
	rows, err := q.db.QueryContext(ctx, getCategories)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Category
	for rows.Next() {
		var i Category
		if err := rows.Scan(&i.ID, &i.Name); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getCategory = `-- name: GetCategory :one
SELECT id, name FROM categories WHERE id = ?
`

func (q *Queries) GetCategory(ctx context.Context, id int32) (Category, error) {
	row := q.db.QueryRowContext(ctx, getCategory, id)
	var i Category
	err := row.Scan(&i.ID, &i.Name)
	return i, err
}

const getCourse = `-- name: GetCourse :one
SELECT id, name, category_id FROM courses WHERE id = ?
`

func (q *Queries) GetCourse(ctx context.Context, id int32) (Course, error) {
	row := q.db.QueryRowContext(ctx, getCourse, id)
	var i Course
	err := row.Scan(&i.ID, &i.Name, &i.CategoryID)
	return i, err
}

const getCourses = `-- name: GetCourses :many
SELECT id, name, category_id FROM courses
`

func (q *Queries) GetCourses(ctx context.Context) ([]Course, error) {
	rows, err := q.db.QueryContext(ctx, getCourses)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Course
	for rows.Next() {
		var i Course
		if err := rows.Scan(&i.ID, &i.Name, &i.CategoryID); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const insertTemporaryTable = `-- name: InsertTemporaryTable :many
SELECT
    ca.id as category_id,
    ca.name as category_name,
    co.id,
    co.name
FROM courses co
INNER JOIN categories ca ON (ca.id = co.category_id)
`

type InsertTemporaryTableRow struct {
	CategoryID   int32
	CategoryName string
	ID           int32
	Name         string
}

func (q *Queries) InsertTemporaryTable(ctx context.Context) ([]InsertTemporaryTableRow, error) {
	rows, err := q.db.QueryContext(ctx, insertTemporaryTable)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []InsertTemporaryTableRow
	for rows.Next() {
		var i InsertTemporaryTableRow
		if err := rows.Scan(
			&i.CategoryID,
			&i.CategoryName,
			&i.ID,
			&i.Name,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listCourseCategories = `-- name: ListCourseCategories :many
SELECT id, name, category_id FROM courses
`

func (q *Queries) ListCourseCategories(ctx context.Context) ([]Course, error) {
	rows, err := q.db.QueryContext(ctx, listCourseCategories)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Course
	for rows.Next() {
		var i Course
		if err := rows.Scan(&i.ID, &i.Name, &i.CategoryID); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

package courses

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{db: db}
}

//
// ------------------------------------------------------------
// COURSES
// ------------------------------------------------------------
//

func (r *Repository) CreateCourse(ctx context.Context, c *Course) error {
	return r.db.QueryRow(ctx, `
        INSERT INTO courses (title, description, price_cents, teacher_id, status, age_range)
        VALUES ($1, $2, $3, $4, $5, $6)
        RETURNING id, created_at, updated_at
    `, c.Title, c.Description, c.PriceCents, c.TeacherID, c.Status, c.AgeRange).Scan(
		&c.ID, &c.CreatedAt, &c.UpdatedAt,
	)
}

func (r *Repository) GetCourse(ctx context.Context, id int) (*Course, error) {
	var c Course
	err := r.db.QueryRow(ctx, `
        SELECT id, title, description, price_cents, teacher_id,
               thumbnail_file, status, age_range,
               created_at, updated_at
        FROM courses
        WHERE id=$1
    `, id).Scan(
		&c.ID, &c.Title, &c.Description, &c.PriceCents,
		&c.TeacherID, &c.ThumbnailFile, &c.Status, &c.AgeRange,
		&c.CreatedAt, &c.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func (r *Repository) ListCourses(ctx context.Context) ([]Course, error) {
	rows, err := r.db.Query(ctx, `
        SELECT id, title, description, price_cents, teacher_id,
               thumbnail_file, status, age_range,
               created_at, updated_at
        FROM courses
        ORDER BY created_at DESC
    `)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []Course
	for rows.Next() {
		var c Course
		rows.Scan(
			&c.ID, &c.Title, &c.Description, &c.PriceCents,
			&c.TeacherID, &c.ThumbnailFile, &c.Status, &c.AgeRange,
			&c.CreatedAt, &c.UpdatedAt,
		)
		list = append(list, c)
	}
	return list, nil
}

//
// ------------------------------------------------------------
// LESSONS
// ------------------------------------------------------------
//

func (r *Repository) AddLesson(ctx context.Context, l *Lesson) error {
	return r.db.QueryRow(ctx, `
        INSERT INTO lessons (course_id, title, file_id, order_index, type)
        VALUES ($1, $2, $3, $4, $5)
        RETURNING id, created_at, updated_at
    `, l.CourseID, l.Title, l.FileID, l.OrderIndex, l.Type).Scan(
		&l.ID, &l.CreatedAt, &l.UpdatedAt,
	)
}

func (r *Repository) ListLessons(ctx context.Context, courseID int) ([]Lesson, error) {
	rows, err := r.db.Query(ctx, `
        SELECT id, course_id, title, file_id, order_index, type,
               created_at, updated_at
        FROM lessons
        WHERE course_id=$1
        ORDER BY order_index ASC
    `, courseID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []Lesson
	for rows.Next() {
		var l Lesson
		rows.Scan(
			&l.ID, &l.CourseID, &l.Title,
			&l.FileID, &l.OrderIndex, &l.Type,
			&l.CreatedAt, &l.UpdatedAt,
		)
		list = append(list, l)
	}
	return list, nil
}

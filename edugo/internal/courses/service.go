package courses

import (
	"context"
	"fmt"
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreateCourse(ctx context.Context, teacherID int, title, description string, priceCents int) (*Course, error) {
	if title == "" {
		return nil, fmt.Errorf("title required")
	}

	c := &Course{
		Title:       title,
		Description: description,
		PriceCents:  priceCents, // FIXED
		TeacherID:   teacherID,
		Status:      "draft", // optional default
		AgeRange:    "",      // optional default
	}

	if err := s.repo.CreateCourse(ctx, c); err != nil {
		return nil, err
	}

	return c, nil
}

func (s *Service) GetCourse(ctx context.Context, id int) (*Course, error) {
	return s.repo.GetCourse(ctx, id)
}

func (s *Service) ListCourses(ctx context.Context) ([]Course, error) {
	return s.repo.ListCourses(ctx)
}

func (s *Service) AddLesson(ctx context.Context, courseID int, title string, fileID int, orderIndex int, lessonType string) (*Lesson, error) {
	if title == "" {
		return nil, fmt.Errorf("title required")
	}

	l := &Lesson{
		CourseID:   courseID,
		Title:      title,
		FileID:     fileID,
		OrderIndex: orderIndex, // FIXED
		Type:       lessonType, // FIXED
	}

	if err := s.repo.AddLesson(ctx, l); err != nil {
		return nil, err
	}

	return l, nil
}

func (s *Service) ListLessons(ctx context.Context, courseID int) ([]Lesson, error) {
	return s.repo.ListLessons(ctx, courseID)
}

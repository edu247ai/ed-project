package courses

import "time"

// Course represents a full course created by a teacher.
type Course struct {
	ID            int       `json:"id"`
	Title         string    `json:"title"`
	Description   string    `json:"description"`
	PriceCents    int       `json:"price_cents"`
	TeacherID     int       `json:"teacher_id"`
	ThumbnailFile int       `json:"thumbnail_file,omitempty"`
	Status        string    `json:"status"`    // draft, published, archived
	AgeRange      string    `json:"age_range"` // "7-10", "10-14", etc.
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

// Lesson represents a single lesson inside a course.
type Lesson struct {
	ID         int       `json:"id"`
	CourseID   int       `json:"course_id"`
	Title      string    `json:"title"`
	FileID     int       `json:"file_id"`     // PDF, MP3, EPUB, video, etc.
	OrderIndex int       `json:"order_index"` // lesson ordering
	Type       string    `json:"type"`        // pdf, video, audio, quiz
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

// CourseWithLessons is used for full course detail views.
type CourseWithLessons struct {
	Course  Course   `json:"course"`
	Lessons []Lesson `json:"lessons"`
}

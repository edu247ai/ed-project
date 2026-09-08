-- Users
CREATE INDEX idx_users_email ON users(email);

-- Courses
CREATE INDEX idx_courses_teacher ON courses(teacher_id);

-- Lessons
CREATE INDEX idx_lessons_course ON lessons(course_id);

-- Student progress
CREATE INDEX idx_progress_student ON student_progress(student_user_id);

-- Marketplace
CREATE INDEX idx_products_type ON products(type);
CREATE INDEX idx_orders_user ON orders(user_id);

-- Chat
CREATE INDEX idx_chat_messages_room ON chat_messages(room_id);

-- ML events
CREATE INDEX idx_events_user ON events(user_id);

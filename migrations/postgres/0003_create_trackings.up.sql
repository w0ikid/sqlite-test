CREATE TABLE habit_trackings (
    id SERIAL PRIMARY KEY,
    habit_id INT REFERENCES habits(id) ON DELETE CASCADE,
    tracked_date DATE NOT NULL,
    status BOOLEAN DEFAULT true,
    UNIQUE (habit_id, tracked_date)
);
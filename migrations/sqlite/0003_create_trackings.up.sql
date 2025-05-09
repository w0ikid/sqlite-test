CREATE TABLE habit_trackings (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    habit_id INTEGER,
    tracked_date DATE NOT NULL,
    status BOOLEAN DEFAULT 1,
    UNIQUE (habit_id, tracked_date),
    FOREIGN KEY (habit_id) REFERENCES habits(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS ideas (
  id integer not null primary key autoincrement,
  idea text not null,
  created_at DATETIME not null default DATETIME('now'),
  updated_at DATETIME not null default DATETIME('now'),
);

CREATE INDEX IF NOT EXISTS ideas_id_idx ON ideas(id);


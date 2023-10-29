CREATE TABLE IF NOT EXISTS tests (
  id serial PRIMARY KEY,
  title VARCHAR(100) NOT NULL,
  description VARCHAR(511) NOT NULL,
  published_date timestamp NOT NULL
);

CREATE TABLE IF NOT EXISTS questions (
  id serial PRIMARY KEY,
  test_id int NOT NULL,
  question_text VARCHAR(255),
  options JSON,
  FOREIGN KEY (test_id) REFERENCES tests(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS usersResults(
    sessionId UUID PRIMARY KEY,
    results JSON
);


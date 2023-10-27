
-- Inserting real test data
INSERT INTO tests (title, description, published_date)
VALUES
  ('History Quiz', 'Test your knowledge of historical events.', NOW()),
  ('Science Quiz', 'Explore the wonders of science.', NOW()),
  ('Math Quiz', 'Solve challenging math problems.', NOW());

-- Inserting real question data for the History Quiz
INSERT INTO questions (test_id, question_text, options)
VALUES
  (1, 'Who was the first President of the United States?', '[{"option_text": "George Washington", "is_correct": true}, {"option_text": "Thomas Jefferson", "is_correct": false}, {"option_text": "Abraham Lincoln", "is_correct": false}]'),
  (1, 'In which year did World War II end?', '[{"option_text": "1945", "is_correct": true}, {"option_text": "1918", "is_correct": false}, {"option_text": "1955", "is_correct": false}]');

-- Inserting real question data for the Science Quiz
INSERT INTO questions (test_id, question_text, options)
VALUES
  (2, 'What is the chemical symbol for water?', '[{"option_text": "H2O", "is_correct": true}, {"option_text": "CO2", "is_correct": false}, {"option_text": "O2", "is_correct": false}]'),
  (2, 'Which planet is known as the Red Planet?', '[{"option_text": "Mars", "is_correct": true}, {"option_text": "Venus", "is_correct": false}, {"option_text": "Jupiter", "is_correct": false}]');

-- Inserting real question data for the Math Quiz
INSERT INTO questions (test_id, question_text, options)
VALUES
  (3, 'What is the value of pi (π)?', '[{"option_text": "3.14159", "is_correct": true}, {"option_text": "2.71828", "is_correct": false}, {"option_text": "1.61803", "is_correct": false}]'),
  (3, 'What is the result of 7 multiplied by 8?', '[{"option_text": "56", "is_correct": true}, {"option_text": "64", "is_correct": false}, {"option_text": "42", "is_correct": false}]');
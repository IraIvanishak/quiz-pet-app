# Quiz Pet App

## Main Application

The main application serves as a REST API to interact with the quiz data.

### API Endpoints

- **GET /**: Retrieve previews of all quizzes.
- **GET /test?id=<quiz_id>**: Retrieve a specific quiz by its ID.

## Getting Started

Follow these steps to get the project up and running on your local machine.

### Prerequisites

- Go programming language installed.
- PostgreSQL database set up.
  
### Dependencies

Run the following commands to fetch and install the project dependencies:
```bash
   cd back
```
```bash
   go mod download
```
### Database Configuration

1. Create a PostgreSQL database for the application.

2. Configure the `.env` file with the appropriate database credentials:
```
DB_HOST=<database_host>
DB_USER=<database_user>
DB_PASS=<database_password>
DB_NAME=<database_name>
```
Replace `<database_host>`, `<database_user>`, `<database_password>`, and `<database_name>` with your database details.

### Database Schema

To set up the database for this application, execute the following SQL statements in your PostgreSQL database:

#### `tests` Table

```sql
CREATE TABLE tests (
  id serial PRIMARY KEY,
  title VARCHAR(100) NOT NULL,
  description VARCHAR(511) NOT NULL,
  published_date timestamp NOT NULL
);
```
#### `questions` Table

```sql
CREATE TABLE questions (
  id serial PRIMARY KEY,
  test_id int NOT NULL,
  question_text VARCHAR(255),
  options JSON,
  FOREIGN KEY (test_id) REFERENCES tests(id) ON DELETE CASCADE
);
```
#### Mock Database Setup 

To facilitate testing and development in isolated environments, you can create a mock database using SQL statements.
```sql
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
```
### Installation

1. Clone the repository to your local machine:

```bash
git clone https://github.com/IraIvanishak/quiz-pet-app.git
```
2. Navigate to the project directory:
```bash
cd quiz-pet-app
```
3. Build and run the application:
```bash
go build ./quiz-pet-app
```


## Contributing
If you'd like to contribute to this project, please follow these guidelines:

Fork the repository.
Create a new branch for your feature or bug fix.
Make your changes and submit a pull request.

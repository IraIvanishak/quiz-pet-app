# Quiz Pet App

## Main Application

The main application serves as a REST API to interact with the quiz data.

### API Endpoints

- **GET /**: Retrieve previews of all quizzes.
- **GET /test?id=<quiz_id>**: Retrieve a specific quiz by its ID.

## Getting Started

1. Configure the `.env` file with the appropriate database credentials:
```
DB_HOST=<database_host>
DB_USER=<database_user>
DB_PASS=<database_password>
DB_NAME=<database_name>
```
Replace `<database_host>`, `<database_user>`, `<database_password>`, and `<database_name>` with your database details.


2. Run the application:
```bash
docker-compose up --build
```


## Contributing
If you'd like to contribute to this project, please follow these guidelines:

Fork the repository.
Create a new branch for your feature or bug fix.
Make your changes and submit a pull request.

-- name: QuizesAll :many
SELECT 
    id, 
    title, 
    description 
FROM 
    tests;

-- name: QuizeByID :one
SELECT 
    title, 
    description 
FROM 
    tests 
WHERE 
    id = $1;

-- name: QuestionsByTestID :many
SELECT 
    question_text, 
    options 
FROM 
    questions 
WHERE 
    test_id = $1;

-- name: ResultsByID :one
SELECT 
    results 
FROM 
    usersresults 
WHERE 
    sessionid = $1;

-- name: InsertResult :exec
INSERT INTO 
    usersresults (sessionid, results)
VALUES 
    ($1, $2)
ON CONFLICT 
    (sessionid)
DO UPDATE SET 
    results = $2;

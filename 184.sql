# Write your MySQL query statement below


WITH max_salary_department AS (
    SELECT D.id AS id, MAX(salary) AS Salary
    FROM Employee E INNER JOIN Department D
    ON E.departmentId = D.id
    GROUP BY D.name
)

SELECT D.name AS Department, E.name AS Employee,
    salary AS SALARY
FROM Employee E INNER JOIN Department D
ON E.departmentId = D.id
WHERE salary = (SELECT Salary 
               FROM max_salary_department
               WHERE id = D.id)

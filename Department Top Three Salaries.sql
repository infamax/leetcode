with salary_by_ranks as (select departmentId, name, salary,
dense_rank() over (partition by departmentId order by salary desc) as r
from Employee)

select d.name as Department, s.name as Employee, salary as Salary
from salary_by_ranks s inner join Department d
on s.departmentId = d.id
where r <= 3

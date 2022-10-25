# Write your MySQL query statement below
# SELECT
#     d.Name AS 'Department', e1.Name AS 'Employee', e1.Salary
# FROM
#     Employee e1
#         JOIN
#     Department d ON e1.DepartmentId = d.Id
# WHERE
#      (SELECT
#             COUNT(DISTINCT e2.Salary)
#         FROM
#             Employee e2
#         WHERE
#             e2.Salary > e1.Salary
#                 AND e1.DepartmentId = e2.DepartmentId
#         ) < 3
# ;

select B.name AS Department
    ,A.Employee
    ,A.Salary 
from (
    select DepartmentId 
        ,name as employee
        ,salary
        ,dense_rank() over (partition by departmentid order by salary desc) as rk
    from employee
) A left join department B 
on A.departmentid = B.id
where A.rk <= 3
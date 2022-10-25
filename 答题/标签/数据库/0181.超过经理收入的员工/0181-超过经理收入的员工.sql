# Write your MySQL query statement below
select e.name as Employee
from Employee e
where e.managerId is not null and 
e.salary > 
(
    select salary
    from Employee e2
    where e2.id = e.managerId
    )
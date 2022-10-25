# Write your MySQL query statement below
# 1025 ms
# select  d.name as 'Department', e.name as 'Employee', e.Salary as Salary
# from Employee e join Department d
# on e.departmentId = d.id 
# where
# 0 = (
#     select count(distinct e2.Salary)
#     from Employee e2
#     where e2.departmentId = e.departmentId and e2.Salary > e.Salary
# )

# 673ms
# select Department, Employee, Salary 
# from (
#     select d.name Department, e.Name  Employee, e.Salary Salary ,
#             rank() over(partition by d.Id  order by Salary desc)  rk
#     from
#         Employee  e join  Department  d on e.DepartmentId  =d.Id 
# ) t1 

# where rk=1

# 811ms
SELECT
    Department.name AS 'Department',
    Employee.name AS 'Employee',
    Salary
FROM
    Employee
        JOIN
    Department ON Employee.DepartmentId = Department.Id
WHERE
    (Employee.DepartmentId , Salary) IN
    (   SELECT
            DepartmentId, MAX(Salary)
        FROM
            Employee
        GROUP BY DepartmentId
	)
;
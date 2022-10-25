# Write your MySQL query statement below
select  dept_name as dept_name, COUNT(student_id) as student_number
from Department left join Student
on Department.dept_id = Student.dept_id
group by Department.dept_name
order by student_number desc, department.dept_name
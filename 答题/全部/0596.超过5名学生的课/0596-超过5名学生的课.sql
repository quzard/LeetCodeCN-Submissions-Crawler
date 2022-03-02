# Write your MySQL query statement below
select temp.class
from (
    select count(Courses.student) as cnt, Courses.class
    from Courses
    group by Courses.class
) as temp
where temp.cnt >= 5
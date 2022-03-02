# Write your MySQL query statement below
select distinct a.Num as ConsecutiveNums
from Logs a, Logs b, Logs c
where a.Num = b.Num and
    b.Num = c.Num and
    a.id = b.id-1 and
    b.id = c.id-1
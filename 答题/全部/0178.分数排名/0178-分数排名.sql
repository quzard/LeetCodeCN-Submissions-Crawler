# Write your MySQL query statement below
select a.Score as Score,
(select count(distinct b.Score) from Scores b where b.Score >= a.Score) as 'Rank'
from Scores a
order by a.Score DESC

# select
#     score,
#     (dense_rank() over (order by Score desc)) AS "rank" 
# from
#     Scores 

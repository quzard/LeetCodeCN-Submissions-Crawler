# Write your MySQL query statement below
select player_id, min(event_date) as first_login
from Activity
group by player_id

# select player_id, event_date as first_login
# from (
# 	select
# 		player_id,
# 		event_date,
# 		dense_rank()
# 		 over(partition by player_id order by event_date) as 排名
# 	from activity
# ) as temp
# where 排名 = 1;

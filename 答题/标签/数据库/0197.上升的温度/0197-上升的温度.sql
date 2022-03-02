# Write your MySQL query statement below
# https://www.w3schools.com/sql/func_mysql_adddate.asp
select w1.id
from Weather w1, Weather w2
where w1.Temperature > w2.Temperature
and w1.recordDate = ADDDATE(w2.recordDate, INTERVAL 1 DAY)

# SELECT
#     weather.id AS 'Id'
# FROM
#     weather
#         JOIN
#     weather w ON DATEDIFF(weather.date, w.date) = 1
#         AND weather.Temperature > w.Temperature
# ;

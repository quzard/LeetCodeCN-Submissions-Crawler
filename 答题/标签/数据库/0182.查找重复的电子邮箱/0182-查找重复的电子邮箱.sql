# Write your MySQL query statement below
# select distinct p1.Email
# from person p1, person p2
# where p1.Email = p2.Email and p1.id != p2.id

# select Email
# from Person
# group by Email
# having count(Email) > 1;

select Email from
(
  select Email, count(Email) as num
  from Person
  group by Email
) as statistic
where num > 1

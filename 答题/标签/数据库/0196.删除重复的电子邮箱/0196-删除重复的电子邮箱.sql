# Write your MySQL query statement below
DELETE p1 
FROM 
    Person p1,
    (
        select MIN(Id) as id, email
        from Person 
        group by Email
        having count(email) > 1
    ) as p2
where p1.email = p2.email and p1.id != p2.id

# DELETE p1 

# FROM Person p1,
#     Person p2
# WHERE
#     p1.Email = p2.Email AND p1.Id > p2.Id


# delete person 
# from 
#     person,
#     (
#         select min(id) as id, email 
#         from person 
#         group by email 
#         having count(email) > 1
#     )   as p2
# where
#     person.email = p2.email and person.id != p2.id;


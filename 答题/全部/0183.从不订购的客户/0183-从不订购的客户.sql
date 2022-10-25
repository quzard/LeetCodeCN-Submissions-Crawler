# Write your MySQL query statement below

# select  cus.Name as 'Customers'
# from Customers cus left join orders
# on cus.id = orders.CustomerId
# where orders.CustomerId is null

select customers.name as 'Customers'
from customers
where customers.id not in
(
    select customerid from orders
);
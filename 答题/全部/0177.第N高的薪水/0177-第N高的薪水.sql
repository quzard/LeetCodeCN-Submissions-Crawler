CREATE FUNCTION getNthHighestSalary(N INT) RETURNS INT
BEGIN
set N=N-1;
  RETURN (
      # Write your MySQL query statement below.
      select distinct salary
      from Employee
      order by Employee.salary desc
      limit N, 1
  );
END
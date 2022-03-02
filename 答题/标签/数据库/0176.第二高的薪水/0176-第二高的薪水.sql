# Write your MySQL query statement below
# select ifnull(
#     (
#         select distinct salary 
#         from Employee 
#         order by  salary desc   
          # limit m,n : 表示从第m+1条开始，取n条数据；  跳过m条，读取n条
          # 等价于 offset m limit n
#         limit 1, 1
#     ),null
# ) as SecondHighestSalary 

select ifnull(
    # distinct 去重
    (select distinct salary
    from employee
    # desc 降序
    order by salary desc
    # limit n offset m 分句表示查询结果跳过 m 条数据，读取前 n 条数据
    limit 1 offset 1),
    null
) as SecondHighestSalary

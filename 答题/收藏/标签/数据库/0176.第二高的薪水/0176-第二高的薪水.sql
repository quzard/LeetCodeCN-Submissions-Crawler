select ifnull(
    # distinct ȥ��
    (
        select distinct salary
        from employee
        # desc ����
        order by salary desc
        # limit m,n : ��ʾ�ӵ�m+1����ʼ��ȡn�����ݣ�  ����m������ȡn��
        # limit n offset m �־��ʾ��ѯ������� m �����ݣ���ȡǰ n ������
        limit 1 offset 1
        ),
    null
) as SecondHighestSalary

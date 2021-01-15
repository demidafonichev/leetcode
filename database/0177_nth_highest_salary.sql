CREATE FUNCTION getNthHighestSalary(N INT) RETURNS INT
BEGIN
    declare M int;
    set M = N-1;
    return (
        select distinct Salary from Employee
        order by -Salary
        limit 1 offset M
    );
END

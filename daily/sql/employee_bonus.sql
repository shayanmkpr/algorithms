-- do a left outer join on employee with bonus table. then return those values that are less than 100 or null.
SELECT e.name, b.bonus
FROM Employee e
LEFT JOIN Bonus b ON e.empId = b.empId  -- Replace with actual FK column
WHERE b.bonus < 1000 OR b.bonus IS NULL;

select c.name
from customer c
left join orders o on c.id = o.customer_id
where o.id is null;

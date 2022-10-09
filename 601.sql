# Write your MySQL query statement below

with people_count_greater_than_100 as (select id, visit_date, people,
lag(id) over (order by id) as prev,
lead(id) over (order by id) as next
from Stadium
where people >= 100), 
id_list_not_full as (
    select id, visit_date, people
    from people_count_greater_than_100
    where next - prev = 2
), id_list_people_greater_than_100 as (
    select id
    from Stadium
    where people >= 100
),  id_list_full as (
    select distinct id - 1 as id_
    from id_list_not_full
    where id - 1 in (select id from id_list_people_greater_than_100)
    union 
    select distinct id + 1 as id_
    from id_list_not_full
    where id + 1 in (select id from id_list_people_greater_than_100)
    union
    select id as id_
    from id_list_not_full
)

select id, visit_date, people
from people_count_greater_than_100
where id in (select id_ from id_list_full)

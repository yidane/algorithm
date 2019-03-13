drop table salary

Create table salary(Id int,employee_id int,amount int,pay_date date);

insert into salary values(1,1,9000,'2017-03-31');
insert into salary values(2,2,6000,'2017-03-31');
insert into salary values(3,3,10000,'2017-03-31');
insert into salary values(4,1,7000,'2017-02-28');
insert into salary values(5,2,6000,'2017-02-28');
insert into salary values(6,3,8000,'2017-02-28');

drop table employee
Create table employee(Id int,department_id int);

insert into employee values(1,1);
insert into employee values(2,2);
insert into employee values(3,2);
 

Given two tables as above, write a query to display the comparison result (higher/lower/same) of the average salary of employees in a department to the company's average salary.

 

So for the sample data above, the result is:

| pay_month | department_id | comparison  |
|-----------|---------------|-------------|
| 2017-03   | 1             | higher      |
| 2017-03   | 2             | lower       |
| 2017-02   | 1             | same        |
| 2017-02   | 2             | same        |
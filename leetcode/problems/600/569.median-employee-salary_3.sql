drop table employees

Create table employees(Id int,Company varchar(2),salary int,);

insert into employees values(1,'A',2341);
insert into employees values(2,'A',341);
insert into employees values(3,'A',15);
insert into employees values(4,'A',15314);
insert into employees values(5,'A',451);
insert into employees values(6,'A',513);
insert into employees values(7,'B',15);
insert into employees values(8,'B',13);
insert into employees values(9,'B',1154);
insert into employees values(10,'B',1345);
insert into employees values(11,'B',1221);
insert into employees values(12,'B',234);
insert into employees values(13,'C',2345);
insert into employees values(14,'C',2645);
insert into employees values(15,'C',2645);
insert into employees values(16,'C',2652);
insert into employees values(17,'C',65);
select * from employees

 

The Employee table holds all employees. The employee table has three columns: Employee Id, Company Name, and Salary.

 

+-----+------------+--------+
|Id   | Company    | Salary |
+-----+------------+--------+
|1    | A          | 2341   |
|2    | A          | 341    |
|3    | A          | 15     |
|4    | A          | 15314  |
|5    | A          | 451    |
|6    | A          | 513    |
|7    | B          | 15     |
|8    | B          | 13     |
|9    | B          | 1154   |
|10   | B          | 1345   |
|11   | B          | 1221   |
|12   | B          | 234    |
|13   | C          | 2345   |
|14   | C          | 2645   |
|15   | C          | 2645   |
|16   | C          | 2652   |
|17   | C          | 65     |
+-----+------------+--------+
Write a SQL query to find the median salary of each company. Bonus points if you can solve it without using any built-in SQL functions.

+-----+------------+--------+
|Id   | Company    | Salary |
+-----+------------+--------+
|5    | A          | 451    |
|6    | A          | 513    |
|12   | B          | 234    |
|9    | B          | 1154   |
|14   | C          | 2645   |
+-----+------------+--------+
Solution:

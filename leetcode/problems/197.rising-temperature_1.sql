--
-- @lc app=leetcode id=197 lang=mysql
--
-- [197] Rising Temperature
--
-- https://leetcode.com/problems/rising-temperature/description/
--
-- database
-- Easy (33.57%)
-- Total Accepted:    67.5K
-- Total Submissions: 201.2K
-- Testcase Example:  '{"headers": {"Weather": ["Id", "RecordDate", "Temperature"]}, "rows": {"Weather": [[1, "2015-01-01", 10], [2, "2015-01-02", 25], [3, "2015-01-03", 20], [4, "2015-01-04", 30]]}}'
--
-- Given a Weather table, write a SQL query to find all dates' Ids with higher
-- temperature compared to its previous (yesterday's) dates.
-- 
-- 
-- +---------+------------------+------------------+
-- | Id(INT) | RecordDate(DATE) | Temperature(INT) |
-- +---------+------------------+------------------+
-- |       1 |       2015-01-01 |               10 |
-- |       2 |       2015-01-02 |               25 |
-- |       3 |       2015-01-03 |               20 |
-- |       4 |       2015-01-04 |               30 |
-- +---------+------------------+------------------+
-- 
-- 
-- For example, return the following Ids for the above Weather table:
-- 
-- 
-- +----+
-- | Id |
-- +----+
-- |  2 |
-- |  4 |
-- +----+
-- 
-- 
--
# Write your MySQL query statement below


#1 worst
select w2.ID from Weather w1
inner join Weather w2
on DATE_ADD(w1.RecordDate,INTERVAL 1 DAY) = w2.RecordDate
and w1.Temperature < w2.Temperature;

#2 TODO 需要再考虑考虑

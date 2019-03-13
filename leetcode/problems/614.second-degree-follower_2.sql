In facebook, there is a follow table with two columns: followee, follower.

Please write a sql query to get the amount of each follower’s follower if he/she has one.

For example:

1
2
3
4
5
6
7
8
+-------------+------------+
| followee    | follower   |
+-------------+------------+
|     A       |     B      |
|     B       |     C      |
|     B       |     D      |
|     D       |     E      |
+-------------+------------+
should output:

1
2
3
4
5
6
+-------------+------------+
| follower    | num        |
+-------------+------------+
|     B       |  2         |
|     D       |  1         |
+-------------+------------+
Explaination:
Both B and D exist in the follower list, when as a followee, B’s follower is C and D, and D’s follower is E. A does not exist in follower list.
Note:
Followee would not follow himself/herself in all cases.
Please display the result in follower’s alphabet order.

help: https://nifannn.github.io/2017/10/27/SQL-Notes-Leetcode-614-Second-Degree-Follower/
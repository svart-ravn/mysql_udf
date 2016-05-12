# mysql_udf
Sample mysql user defined function written in go

build_so.sh used to build so file

to install stored procedure do
````
mysql -e "create function goup returns string soname 'goup.so'"
````

Sample calls:
````
mysql -BN -e "select goup('golang', 1)'"
ERROR 1123 (HY000) at line 1: Can't initialize function 'goup'; functions requires 1 argument
````
````
mysql -BN -e "select goup('golang')'"
GOLANG
````

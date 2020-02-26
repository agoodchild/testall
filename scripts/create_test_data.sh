HOST="127.0.0.1:9093"
TABLENAME="test"
URL=http://$HOST/${TABLENAME}/schema
curl --location --request PUT $URL \
--header 'Content-Type: text/xml' \
--header 'Accept: text/xml' \
-d '<TableSchema name="tableschema"><ColumnSchema name="id"/><ColumnSchema name="info"/><ColumnSchema name="data"/></TableSchema>'
echo '创建表'$TABLENAME
DATA='{"Row":[{"key":"MTExMTExMTExMTEx","Cell":[{"column":"aWQ6aWQ=","$":"MTExMTExMTEx"},{"column":"aW5mbzpzdGFydHRpbWU=","$":"MTU4MjYzMzYyMg=="},{"column":"aW5mbzplbmR0aW1l","$":"MTU4MjcyMDAyMg=="},{"column":"aW5mbzpqc29uZGF0YQ==","$":"MjAyMC8yLzI1IDIwOjI3OjItMjAyMC8yLzI2IDIwOjI3OjI="}]},{"key":"MjIyMjIyMjIyMjIyMg==","Cell":[{"column":"aWQ6aWQ=","$":"MjIyMjIyMjIyMjIyMjI="},{"column":"aW5mbzpzdGFydHRpbWU=","$":"MTU4MjU0NzIyMg=="},{"column":"aW5mbzplbmR0aW1l","$":"MTU4MjYzMzYyMg=="},{"column":"aW5mbzpqc29uZGF0YQ==","$":"MjAyMC8yLzI0IDIwOjI3OjItMjAyMC8yLzI1IDIwOjI3OjI="}]},{"key":"MzMzMzMzMzMzMzMzMzM=","Cell":[{"column":"aWQ6aWQ=","$":"MzMzMzMzMzMzMzMz"},{"column":"aW5mbzpzdGFydHRpbWU=","$":"MTU4MjQ0NzIyMg=="},{"column":"aW5mbzplbmR0aW1l","$":"MTU4MjUzMzYyMg=="},{"column":"aW5mbzpqc29uZGF0YQ==","$":"MjAyMC8yLzIzIDIwOjI3OjItMjAyMC8yLzI0IDIwOjI3OjI="}]},{"key":"NDQ0NDQ0NDQ0NDQ0NA==","Cell":[{"column":"aWQ6aWQ=","$":"NDQ0NDQ0NDQ0NDQ0"},{"column":"aW5mbzpzdGFydHRpbWU=","$":"MTU4MjM0NzIyMg=="},{"column":"aW5mbzplbmR0aW1l","$":"MTU4MjQzMzYyMg=="},{"column":"aW5mbzpqc29uZGF0YQ==","$":"MjAyMC8yLzIyIDIwOjI3OjItMjAyMC8yLzIzIDIwOjI3OjI="}]},{"key":"NTU1NTU1NTU1NTU=","Cell":[{"column":"aWQ6aWQ=","$":"NTU1NTU1NTU1NTU1NTU="},{"column":"aW5mbzpzdGFydHRpbWU=","$":"MTU4MjI0NzIyMg=="},{"column":"aW5mbzplbmR0aW1l","$":"MTU4MjMzMzYyMg=="},{"column":"aW5mbzpqc29uZGF0YQ==","$":"MjAyMC8yLzIxIDIwOjI3OjItMjAyMC8yLzIyIDIwOjI3OjI="}]}]}'
URL=http://$HOST/${TABLENAME}/row_key
curl --location --request PUT $URL \
--header 'Content-Type: application/json' \
--header 'Accept: application/json' \
-d $DATA
echo '添加5条数据'

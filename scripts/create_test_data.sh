HOST="127.0.0.1:9093"
TABLENAME="test"
URL=http://$HOST/${TABLENAME}/schema
echo $URL
curl --location --request PUT $URL \
--header 'Content-Type: text/xml' \
--header 'Accept: text/xml' \
-d '<TableSchema name="tableschema"><ColumnSchema name="id"/><ColumnSchema name="info"/><ColumnSchema name="data"/></TableSchema>'

echo '创建表'$TABLENAME
DATA='{"Row":[{"key":"MTEx","Cell":[{"column":"aWQ6aWQ=","$":"MTEx"},{"column":"aW5mbzpzdGFydHRpbWU=","$":"MTU4MjYzMzYyMg=="},{"column":"aW5mbzplbmR0aW1l","$":"MTU4MjcyMDAyMg=="},{"column":"aW5mbzpqc29uZGF0YQ==","$":"ZGF0YTExMQ=="}]},{"key":"MjIy","Cell":[{"column":"aWQ6aWQ=","$":"MjIy"},{"column":"aW5mbzpzdGFydHRpbWU=","$":"MTU4MjU0NzIyMg=="},{"column":"aW5mbzplbmR0aW1l","$":"MTU4MjYzMzYyMg=="},{"column":"aW5mbzpqc29uZGF0YQ==","$":"ZGF0YTIyMg=="}]},{"key":"MzMz","Cell":[{"column":"aWQ6aWQ=","$":"MzMz"},{"column":"aW5mbzpzdGFydHRpbWU=","$":"MTU4MjQ0NzIyMg=="},{"column":"aW5mbzplbmR0aW1l","$":"MTU4MjUzMzYyMg=="},{"column":"aW5mbzpqc29uZGF0YQ==","$":"ZGF0YTMzMw=="}]},{"key":"NDQ0","Cell":[{"column":"aWQ6aWQ=","$":"NDQ0"},{"column":"aW5mbzpzdGFydHRpbWU=","$":"MTU4MjM0NzIyMg=="},{"column":"aW5mbzplbmR0aW1l","$":"MTU4MjQzMzYyMg=="},{"column":"aW5mbzpqc29uZGF0YQ==","$":"ZGF0YTQ0NA=="}]},{"key":"NTU1","Cell":[{"column":"aWQ6aWQ=","$":"NTU1"},{"column":"aW5mbzpzdGFydHRpbWU=","$":"MTU4MjI0NzIyMg=="},{"column":"aW5mbzplbmR0aW1l","$":"MTU4MjMzMzYyMg=="},{"column":"aW5mbzpqc29uZGF0YQ==","$":"ZGF0YTU1NQ=="}]}]}'
URL=http://$HOST/${TABLENAME}/row_key
curl --location --request PUT $URL \
--header 'Content-Type: application/json' \
--header 'Accept: application/json' \
-d $DATA

echo '添加5条数据'

HOST="127.0.0.1:9093/"

#创建表test1
TABLENAME="test"
XML="<TableSchema name=\"tableschema\"><ColumnSchema name=\"id\"/><ColumnSchema name=\"info\"/><ColumnSchema name=\"data\"/></TableSchema>"
DATA=$(curl -x PUT "http://$HOST$TABLENAME/schema" -H "Content-Type: application/json" -d ${XML})
echo $DATA

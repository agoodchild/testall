HOST="127.0.0.1:9093/"
HTTPHOST=http://$HOST
DATA=$(curl -s "${HTTPHOST}")

#用空格分割所有表名到arr数组
OLD_IFS="$IFS"
IFS=" "
arr=($DATA)
IFS="$OLD_IFS"

#循环表名,删除表
for s in ${arr[@]}
do
  URL=$HTTPHOST${s}/schema
  DATA=$(curl -s -X DELETE "${URL}")
  echo `printf "删除表："$s`
done
echo `printf "删除所有表"`
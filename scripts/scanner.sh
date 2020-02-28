D=$(curl -s  -w %{Location}  -X PUT  -H "Accept: text/xml"  -H "Content-Type:text/xml"  -d @filter.xml  "http://127.0.0.1:9093/test/scanner/")
echo $D

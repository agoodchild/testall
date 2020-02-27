curl -vi -X PUT \
  -H "Accept: text/xml" \
  -H "Content-Type:application/json" \
  -d @filterjson.txt \
  "http://127.0.0.1:9093/test/scanner/"
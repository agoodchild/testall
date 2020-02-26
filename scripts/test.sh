create 'member','id','address','info'
put 'member', 'debugo','id','11'
put 'member', 'debugo','info:age','27'
put 'member', 'debugo','info:birthday','1987-04-04'
put 'member', 'debugo','info:industry', 'it'
put 'member', 'debugo','address:city','beijing'
put 'member', 'debugo','address:country','china'
put 'member', 'Sariel', 'id', '21'
put 'member', 'Sariel','info:age', '26'
put 'member', 'Sariel','info:birthday', '1988-05-09 '
put 'member', 'Sariel','info:industry', 'it'
put 'member', 'Sariel','address:city', 'beijing'
put 'member', 'Sariel','address:country', 'china'
put 'member', 'Elvis', 'id', '22'
put 'member', 'Elvis','info:age', '26'
put 'member', 'Elvis','info:birthday', '1988-09-14 '
put 'member', 'Elvis','info:industry', 'it'
put 'member', 'Elvis','address:city', 'beijing'
put 'member', 'Elvis','address:country', 'china'


count 'member'
get 'member', ‘Sariel'
get 'member', 'Sariel', 'info'
scan 'member'
scan 'member', {COLUMN=>'info'}
scan 'member', {COLUMNS=> 'info:birthday'}
scan 'member', { STARTROW => 'Sariel', LIMIT=>1, VERSIONS=>1}
scan 'member', FILTER=>"ValueFilter(=,'binary:26’)"
scan 'member', FILTER=>"ValueFilter(=,'substring:6')"
scan 'member', FILTER=>"ColumnPrefixFilter('birth')"
scan 'member', FILTER=>"ColumnPrefixFilter('birth') AND ValueFilter ValueFilter(=,'substring:1988')"
scan 'member', FILTER=>"PrefixFilter('E')"
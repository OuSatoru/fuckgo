#! /usr/bin/env sh

db2 connect to bcas user bcas using bcas
db2 "export to /fr/data/xms/bcas_kmh_$1.del of del
select * from bcas.$1 where ETLDT = to_char(current_date - 1 day, 'yyyy-mm-dd')"
DATE=/fr/data/xms/kmhexp/$(date -d yesterday +%Y%m%d).$1
touch $DATE
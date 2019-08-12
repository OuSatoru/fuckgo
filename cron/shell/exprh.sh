#! /usr/bin/env sh

db2 connect to jsbods user ods using ods@98
db2 "export to /fr/data/xms/ODS_RHRBTJB.del of del
SELECT zl, sjrq, ZBSJ FROM REPORT.ODS_RHRBTJB WHERE SJRQ = to_char(CURRENT_DATE - 1 day, 'yyyymmdd')"
DATE=/fr/data/xms/$(date -d yesterday +%Y%m%d).rh
touch $DATE
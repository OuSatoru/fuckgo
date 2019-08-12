#! /usr/bin/env sh

db2 connect to jsbods user ods using ods@98
db2 "export to /fr/data/xms/FDM_DKWJFLQDMXB.del of del
SELECT current_date, a.* FROM REPORT.FDM_DKWJFLQDMXB a"
db2 "export to /fr/data/xms/CMIS_USERINFO.del of del
SELECT USERID, USERNAME, BELONGORG FROM CMIS.USER_INFO"
DATE=/fr/data/xms/$(date -d yesterday +%Y%m%d).yq
touch $DATE
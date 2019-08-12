#! /usr/bin/env sh

db2 connect to jsbods user ods using ods@98
db2 "export to /fr/data/xms/HQCKZJD_pg_cz.csv of del modified by codepage=1208
select a.* from REPORT.ODS_HQCKZJD a left join REPORT.ODS_HQCKZD b on a.ZH = b.ZH where substr(b.CPDM, 6 ,1) = '1'"
db2 "export to /fr/data/xms/HQCKZD_pg_cz.csv of del modified by codepage=1208
select * from REPORT.ODS_HQCKZD where substr(CPDM,6,1) = '1'"
db2 "export to /fr/data/xms/HQCKMXZ_pg_cz.csv of del modified by codepage=1208
select * from REPORT.ODS_HQCKMXZ_NEW where JYRQ1 >= to_char(current_date - 3 days, 'yyyymmdd') and substr(CPDM, 6, 1) = '1'"
DATE=/fr/data/xms/$(date -d yesterday +%Y%m%d).pgcz
touch $DATE

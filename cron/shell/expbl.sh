#! /usr/bin/env sh

db2 connect to jsbods user ods using ods@98
db2 "export to /fr/data/xms/gg_all_data_bl.del of del
SELECT
  CURRENT_DATE sjrq,
  jgm,
  sum(CASE WHEN substr(dqrq, 1, 4) = year(CURRENT_DATE - 1 DAY)
    THEN dkye
      ELSE 0 END) / 10000 bnyqye,
  sum(CASE WHEN substr(dqrq, 1, 4) != year(CURRENT_DATE - 1 DAY)
    THEN dkye
      ELSE 0 END) / 10000 cqyqye,
  CURRENT_DATE - 1 DAY    scrq,
  sum(dkye) / 10000       yqye
FROM REPORT.FDM_DKWJFLQDMXB a
WHERE substr(kmh, 1, 4) IN ('1301', '1302', '1303', '1304') AND dkye > 0 AND dqrq < to_char(CURRENT_DATE, 'yyyymmdd')
GROUP BY jgm"
db2 "export to /fr/data/xms/gg_all_data_bw.del of del
SELECT
  CURRENT_DATE sjrq,
  JGM,
  sum(jfye) / 10000 bwye
FROM REPORT.ODS_HXZZYEB a
WHERE a.SJRQ = to_char(CURRENT_DATE - 1 DAY, 'yyyymmdd') AND a.KMH = '91310000'
GROUP BY jgm"

DATE=/fr/data/xms/$(date -d yesterday +%Y%m%d).bl
touch $DATE
#! /usr/bin/env sh

db2 connect to jsbods user ods using ods@98
db2 "export to /fr/data/xms/gg_all_data_save_top.del of del
SELECT
  CURRENT_DATE sjrq,
  row_number()
  OVER ()      serial,
  a.*,
  1 AS         updated
FROM (
       SELECT
         a.*,
         (CASE JGJC
          WHEN '��̨������֧��'
            THEN '����'
          ELSE trim(substr(replace(JGJC, '֧��', ''), 5, length(replace(JGJC, '֧��', '')) / 2 * 3 - 4)) END) JGJC
       FROM (
              SELECT
                JGM,
                SUM(CASE WHEN SJRQ = to_char(CURRENT_DATE - 1 DAY, 'yyyymmdd')
                  THEN ZBSJ1 END) / 10000 ck,
                SUM(CASE WHEN SJRQ = to_char(CURRENT_DATE - 1 DAY, 'yyyymmdd')
                  THEN ZBSJ1 END) / 10000 - SUM(CASE WHEN SJRQ = to_char(CURRENT_DATE - 2 DAY, 'yyyymmdd')
                  THEN ZBSJ1 END) / 10000 ckz,
                SUM(CASE WHEN SJRQ = to_char(CURRENT_DATE - 1 DAY, 'yyyymmdd')
                  THEN ZBSJ1 END) / 10000 -
                SUM(CASE WHEN SJRQ = to_char(last_day(add_months(CURRENT_DATE - 1 day, -1)), 'yyyymmdd')
                  THEN ZBSJ1 END) / 10000 ckm,
                SUM(CASE WHEN SJRQ = to_char(CURRENT_DATE - 1 DAY, 'yyyymmdd')
                  THEN ZBSJ1 END) / 10000 - SUM(CASE WHEN SJRQ = to_char(
                    CURRENT_DATE - (MONTH(CURRENT_DATE) - 1) MONTH - day(CURRENT_DATE) DAY, 'yyyymmdd')
                  THEN ZBSJ1 END) / 10000 cky
              FROM REPORT.ODS_ZBTJB
              WHERE ZBDH = '0001'
              GROUP BY JGM
            ) a LEFT JOIN REPORT.ODS_HXXTJGM b ON a.JGM = b.JGM WHERE a.JGM NOT IN ('320919900', '320919031', '320919931')) a"
db2 "export to /fr/data/xms/gg_all_data_loan_top.del of del
SELECT
  CURRENT_DATE sjrq,
  row_number()
  OVER ()      serial,
  a.*,
  1 AS         updated
FROM (
       SELECT
         a.*,
         (CASE JGJC
          WHEN '��̨������֧��'
            THEN '����'
          ELSE trim(substr(replace(JGJC, '֧��', ''), 5, length(replace(JGJC, '֧��', '')) / 2 * 3 - 4)) END) JGJC
       FROM (
              SELECT
                JGM,
                SUM(CASE WHEN SJRQ = to_char(CURRENT_DATE - 1 DAY, 'yyyymmdd')
                  THEN ZBSJ1 END) / 10000 ck,
                SUM(CASE WHEN SJRQ = to_char(CURRENT_DATE - 1 DAY, 'yyyymmdd')
                  THEN ZBSJ1 END) / 10000 - SUM(CASE WHEN SJRQ = to_char(CURRENT_DATE - 2 DAY, 'yyyymmdd')
                  THEN ZBSJ1 END) / 10000 ckz,
                SUM(CASE WHEN SJRQ = to_char(CURRENT_DATE - 1 DAY, 'yyyymmdd')
                  THEN ZBSJ1 END) / 10000 -
                SUM(CASE WHEN SJRQ = to_char(last_day(add_months(CURRENT_DATE - 1 day, -1)), 'yyyymmdd')
                  THEN ZBSJ1 END) / 10000 ckm,
                SUM(CASE WHEN SJRQ = to_char(CURRENT_DATE - 1 DAY, 'yyyymmdd')
                  THEN ZBSJ1 END) / 10000 - SUM(CASE WHEN SJRQ = to_char(
                    CURRENT_DATE - (MONTH(CURRENT_DATE) - 1) MONTH - day(CURRENT_DATE) DAY, 'yyyymmdd')
                  THEN ZBSJ1 END) / 10000 cky
              FROM REPORT.ODS_ZBTJB
              WHERE ZBDH = '0002'
              GROUP BY JGM
            ) a LEFT JOIN REPORT.ODS_HXXTJGM b ON a.JGM = b.JGM WHERE a.JGM NOT IN ('320919900', '320919031', '320919931')) a"
db2 "export to /fr/data/xms/gg_all_data_zbsj.del of del
SELECT
  CURRENT_DATE data_date,
  sum(CASE WHEN ZBDH = '0001' AND SJRQ = to_char(CURRENT_DATE - 1 DAY, 'yyyymmdd')
    THEN ZBSJ1 END) / 10000 ����,
  sum(CASE WHEN ZBDH = '0002' AND SJRQ = to_char(CURRENT_DATE - 1 DAY, 'yyyymmdd')
    THEN ZBSJ1 END) / 10000 ����,
  sum(CASE WHEN ZBDH = '0009' AND SJRQ = to_char(CURRENT_DATE - 1 DAY, 'yyyymmdd')
    THEN ZBSJ1 END) / 10000 ����,
  sum(CASE WHEN ZBDH = '0010' AND SJRQ = to_char(CURRENT_DATE - 1 DAY, 'yyyymmdd')
    THEN ZBSJ1 END) / 10000 ����,
  sum(CASE WHEN ZBDH = '0011' AND SJRQ = to_char(CURRENT_DATE - 1 DAY, 'yyyymmdd')
    THEN ZBSJ1 END) / 10000 ����,
  sum(CASE WHEN ZBDH = '0004' AND SJRQ = to_char(CURRENT_DATE - 1 DAY, 'yyyymmdd')
    THEN ZBSJ1 END) / 10000 ����,
  sum(CASE WHEN ZBDH = '0005' AND SJRQ = to_char(CURRENT_DATE - 1 DAY, 'yyyymmdd')
    THEN ZBSJ1 END) / 10000 ������,
  sum(CASE WHEN ZBDH = '0006' AND SJRQ = to_char(CURRENT_DATE - 1 DAY, 'yyyymmdd')
    THEN ZBSJ1 END) / 10000 ũ����,
  sum(CASE WHEN ZBDH = '0008' AND SJRQ = to_char(CURRENT_DATE - 1 DAY, 'yyyymmdd')
    THEN ZBSJ1 END) / 10000 ������,
  ----
  sum(CASE WHEN ZBDH = '0001' AND SJRQ = to_char(CURRENT_DATE - 1 DAY, 'yyyymmdd')
    THEN ZBSJ1 END) / 10000 - sum(CASE WHEN ZBDH = '0001' AND SJRQ = to_char(
      CURRENT_DATE - (MONTH(CURRENT_DATE) - 1) MONTH - day(CURRENT_DATE) DAY, 'yyyymmdd')
    THEN ZBSJ1 END) / 10000 ����,
  sum(CASE WHEN ZBDH = '0002' AND SJRQ = to_char(CURRENT_DATE - 1 DAY, 'yyyymmdd')
    THEN ZBSJ1 END) / 10000 - sum(CASE WHEN ZBDH = '0002' AND SJRQ = to_char(
      CURRENT_DATE - (MONTH(CURRENT_DATE) - 1) MONTH - day(CURRENT_DATE) DAY, 'yyyymmdd')
    THEN ZBSJ1 END) / 10000 ����,
  sum(CASE WHEN ZBDH = '0009' AND SJRQ = to_char(CURRENT_DATE - 1 DAY, 'yyyymmdd')
    THEN ZBSJ1 END) / 10000 - sum(CASE WHEN ZBDH = '0009' AND SJRQ = to_char(
      CURRENT_DATE - (MONTH(CURRENT_DATE) - 1) MONTH - day(CURRENT_DATE) DAY, 'yyyymmdd')
    THEN ZBSJ1 END) / 10000 ����,
  sum(CASE WHEN ZBDH = '0010' AND SJRQ = to_char(CURRENT_DATE - 1 DAY, 'yyyymmdd')
    THEN ZBSJ1 END) / 10000 - sum(CASE WHEN ZBDH = '0010' AND SJRQ = to_char(
      CURRENT_DATE - (MONTH(CURRENT_DATE) - 1) MONTH - day(CURRENT_DATE) DAY, 'yyyymmdd')
    THEN ZBSJ1 END) / 10000 ����,
  sum(CASE WHEN ZBDH = '0011' AND SJRQ = to_char(CURRENT_DATE - 1 DAY, 'yyyymmdd')
    THEN ZBSJ1 END) / 10000 - sum(CASE WHEN ZBDH = '0011' AND SJRQ = to_char(
      CURRENT_DATE - (MONTH(CURRENT_DATE) - 1) MONTH - day(CURRENT_DATE) DAY, 'yyyymmdd')
    THEN ZBSJ1 END) / 10000 ����,
  sum(CASE WHEN ZBDH = '0004' AND SJRQ = to_char(CURRENT_DATE - 1 DAY, 'yyyymmdd')
    THEN ZBSJ1 END) / 10000 - sum(CASE WHEN ZBDH = '0004' AND SJRQ = to_char(
      CURRENT_DATE - (MONTH(CURRENT_DATE) - 1) MONTH - day(CURRENT_DATE) DAY, 'yyyymmdd')
    THEN ZBSJ1 END) / 10000 ����,
  sum(CASE WHEN ZBDH = '0005' AND SJRQ = to_char(CURRENT_DATE - 1 DAY, 'yyyymmdd')
    THEN ZBSJ1 END) / 10000 - sum(CASE WHEN ZBDH = '0005' AND SJRQ = to_char(
      CURRENT_DATE - (MONTH(CURRENT_DATE) - 1) MONTH - day(CURRENT_DATE) DAY, 'yyyymmdd')
    THEN ZBSJ1 END) / 10000 ������,
  sum(CASE WHEN ZBDH = '0006' AND SJRQ = to_char(CURRENT_DATE - 1 DAY, 'yyyymmdd')
    THEN ZBSJ1 END) / 10000 - sum(CASE WHEN ZBDH = '0006' AND SJRQ = to_char(
      CURRENT_DATE - (MONTH(CURRENT_DATE) - 1) MONTH - day(CURRENT_DATE) DAY, 'yyyymmdd')
    THEN ZBSJ1 END) / 10000 ũ����,
  sum(CASE WHEN ZBDH = '0008' AND SJRQ = to_char(CURRENT_DATE - 1 DAY, 'yyyymmdd')
    THEN ZBSJ1 END) / 10000 - sum(CASE WHEN ZBDH = '0008' AND SJRQ = to_char(
      CURRENT_DATE - (MONTH(CURRENT_DATE) - 1) MONTH - day(CURRENT_DATE) DAY, 'yyyymmdd')
    THEN ZBSJ1 END) / 10000 ������,
  -----
  sum(CASE WHEN ZBDH = '0001' AND SJRQ = to_char(CURRENT_DATE - 1 DAY, 'yyyymmdd')
    THEN ZBSJ1 END) / 10000 -
  sum(CASE WHEN ZBDH = '0001' AND SJRQ = to_char(last_day(add_months(CURRENT_DATE - 1 day, -1)), 'yyyymmdd')
    THEN ZBSJ1 END) / 10000 ����,
  sum(CASE WHEN ZBDH = '0002' AND SJRQ = to_char(CURRENT_DATE - 1 DAY, 'yyyymmdd')
    THEN ZBSJ1 END) / 10000 -
  sum(CASE WHEN ZBDH = '0002' AND SJRQ = to_char(last_day(add_months(CURRENT_DATE - 1 day, -1)), 'yyyymmdd')
    THEN ZBSJ1 END) / 10000 ����,
  sum(CASE WHEN ZBDH = '0009' AND SJRQ = to_char(CURRENT_DATE - 1 DAY, 'yyyymmdd')
    THEN ZBSJ1 END) / 10000 -
  sum(CASE WHEN ZBDH = '0009' AND SJRQ = to_char(last_day(add_months(CURRENT_DATE - 1 day, -1)), 'yyyymmdd')
    THEN ZBSJ1 END) / 10000 ����,
  sum(CASE WHEN ZBDH = '0010' AND SJRQ = to_char(CURRENT_DATE - 1 DAY, 'yyyymmdd')
    THEN ZBSJ1 END) / 10000 -
  sum(CASE WHEN ZBDH = '0010' AND SJRQ = to_char(last_day(add_months(CURRENT_DATE - 1 day, -1)), 'yyyymmdd')
    THEN ZBSJ1 END) / 10000 ����,
  sum(CASE WHEN ZBDH = '0011' AND SJRQ = to_char(CURRENT_DATE - 1 DAY, 'yyyymmdd')
    THEN ZBSJ1 END) / 10000 -
  sum(CASE WHEN ZBDH = '0011' AND SJRQ = to_char(last_day(add_months(CURRENT_DATE - 1 day, -1)), 'yyyymmdd')
    THEN ZBSJ1 END) / 10000 ����,
  sum(CASE WHEN ZBDH = '0004' AND SJRQ = to_char(CURRENT_DATE - 1 DAY, 'yyyymmdd')
    THEN ZBSJ1 END) / 10000 -
  sum(CASE WHEN ZBDH = '0004' AND SJRQ = to_char(last_day(add_months(CURRENT_DATE - 1 day, -1)), 'yyyymmdd')
    THEN ZBSJ1 END) / 10000 ����,
  sum(CASE WHEN ZBDH = '0005' AND SJRQ = to_char(CURRENT_DATE - 1 DAY, 'yyyymmdd')
    THEN ZBSJ1 END) / 10000 -
  sum(CASE WHEN ZBDH = '0005' AND SJRQ = to_char(last_day(add_months(CURRENT_DATE - 1 day, -1)), 'yyyymmdd')
    THEN ZBSJ1 END) / 10000 ������,
  sum(CASE WHEN ZBDH = '0006' AND SJRQ = to_char(CURRENT_DATE - 1 DAY, 'yyyymmdd')
    THEN ZBSJ1 END) / 10000 -
  sum(CASE WHEN ZBDH = '0006' AND SJRQ = to_char(last_day(add_months(CURRENT_DATE - 1 day, -1)), 'yyyymmdd')
    THEN ZBSJ1 END) / 10000 ũ����,
  sum(CASE WHEN ZBDH = '0008' AND SJRQ = to_char(CURRENT_DATE - 1 DAY, 'yyyymmdd')
    THEN ZBSJ1 END) / 10000 -
  sum(CASE WHEN ZBDH = '0008' AND SJRQ = to_char(last_day(add_months(CURRENT_DATE - 1 day, -1)), 'yyyymmdd')
    THEN ZBSJ1 END) / 10000 ������,
  -----
  sum(CASE WHEN ZBDH = '0001' AND SJRQ = to_char(CURRENT_DATE - 1 DAY, 'yyyymmdd')
    THEN ZBSJ1 END) / 10000 - sum(CASE WHEN ZBDH = '0001' AND SJRQ = to_char(CURRENT_DATE - 2 DAY, 'yyyymmdd')
    THEN ZBSJ1 END) / 10000 ����,
  sum(CASE WHEN ZBDH = '0002' AND SJRQ = to_char(CURRENT_DATE - 1 DAY, 'yyyymmdd')
    THEN ZBSJ1 END) / 10000 - sum(CASE WHEN ZBDH = '0002' AND SJRQ = to_char(CURRENT_DATE - 2 DAY, 'yyyymmdd')
    THEN ZBSJ1 END) / 10000 ����,
  sum(CASE WHEN ZBDH = '0009' AND SJRQ = to_char(CURRENT_DATE - 1 DAY, 'yyyymmdd')
    THEN ZBSJ1 END) / 10000 - sum(CASE WHEN ZBDH = '0009' AND SJRQ = to_char(CURRENT_DATE - 2 DAY, 'yyyymmdd')
    THEN ZBSJ1 END) / 10000 ����,
  sum(CASE WHEN ZBDH = '0010' AND SJRQ = to_char(CURRENT_DATE - 1 DAY, 'yyyymmdd')
    THEN ZBSJ1 END) / 10000 - sum(CASE WHEN ZBDH = '0010' AND SJRQ = to_char(CURRENT_DATE - 2 DAY, 'yyyymmdd')
    THEN ZBSJ1 END) / 10000 ����,
  sum(CASE WHEN ZBDH = '0011' AND SJRQ = to_char(CURRENT_DATE - 1 DAY, 'yyyymmdd')
    THEN ZBSJ1 END) / 10000 - sum(CASE WHEN ZBDH = '0011' AND SJRQ = to_char(CURRENT_DATE - 2 DAY, 'yyyymmdd')
    THEN ZBSJ1 END) / 10000 ����,
  sum(CASE WHEN ZBDH = '0004' AND SJRQ = to_char(CURRENT_DATE - 1 DAY, 'yyyymmdd')
    THEN ZBSJ1 END) / 10000 - sum(CASE WHEN ZBDH = '0004' AND SJRQ = to_char(CURRENT_DATE - 2 DAY, 'yyyymmdd')
    THEN ZBSJ1 END) / 10000 ����,
  sum(CASE WHEN ZBDH = '0005' AND SJRQ = to_char(CURRENT_DATE - 1 DAY, 'yyyymmdd')
    THEN ZBSJ1 END) / 10000 - sum(CASE WHEN ZBDH = '0005' AND SJRQ = to_char(CURRENT_DATE - 2 DAY, 'yyyymmdd')
    THEN ZBSJ1 END) / 10000 ������,
  sum(CASE WHEN ZBDH = '0006' AND SJRQ = to_char(CURRENT_DATE - 1 DAY, 'yyyymmdd')
    THEN ZBSJ1 END) / 10000 - sum(CASE WHEN ZBDH = '0006' AND SJRQ = to_char(CURRENT_DATE - 2 DAY, 'yyyymmdd')
    THEN ZBSJ1 END) / 10000 ũ����,
  sum(CASE WHEN ZBDH = '0008' AND SJRQ = to_char(CURRENT_DATE - 1 DAY, 'yyyymmdd')
    THEN ZBSJ1 END) / 10000 - sum(CASE WHEN ZBDH = '0008' AND SJRQ = to_char(CURRENT_DATE - 2 DAY, 'yyyymmdd')
    THEN ZBSJ1 END) / 10000 ������,
  1 as updated
FROM
  (SELECT *
   FROM REPORT.ODS_ZBTJB
   WHERE SJRQ IN (to_char(CURRENT_DATE - 1 DAY, 'yyyymmdd'), to_char(
       CURRENT_DATE - (MONTH(CURRENT_DATE) - 1) MONTH - day(CURRENT_DATE) DAY, 'yyyymmdd'),
                  to_char(last_day(add_months(CURRENT_DATE - 1 day, -1)), 'yyyymmdd'), to_char(CURRENT_DATE - 2 DAY,
                                                                                       'yyyymmdd')))"
DATE=/fr/data/xms/$(date -d yesterday +%Y%m%d).zb
touch $DATE
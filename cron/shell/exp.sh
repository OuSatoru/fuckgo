#! /usr/bin/env sh

db2 connect to jsbods user ods using ods@98
db2 "export to /fr/data/xms/bank.del of del SELECT * FROM REPORT.ODS_HXXTJGM"
db2 "export to /fr/data/xms/gg_all_data_buliang.del of del
SELECT
  CURRENT_DATE data_date,
  sum(CASE WHEN DQRQ < to_char(CURRENT_DATE - 1 DAY, 'yyyymmdd') AND DKYE > 0 AND
                                DZHXBZ = 'N'
  THEN DKYE END) / 10000 bl,
  sum(CASE WHEN DQRQ < to_char(CURRENT_DATE - 1 DAY, 'yyyymmdd') AND DKYE > 0 AND
                                DZHXBZ = 'N'
  THEN DKYE END) / 10000 - sum(CASE WHEN DQRQ < to_char(last_day(add_months(CURRENT_DATE - 1 day, -1)), 'yyyymmdd') AND DKYE > 0 AND
                                DZHXBZ = 'N'
  THEN DKYE END) / 10000 bl_m,
  sum(CASE WHEN DQRQ < to_char(CURRENT_DATE - 1 DAY, 'yyyymmdd') AND DKYE > 0 AND
                                DZHXBZ = 'N'
  THEN DKYE END) / 10000 - sum(CASE WHEN DQRQ < to_char(CURRENT_DATE - 2 DAY, 'yyyymmdd') AND DKYE > 0 AND
                                DZHXBZ = 'N'
  THEN DKYE END) / 10000 bl_d,
  sum(CASE WHEN DQRQ < to_char(CURRENT_DATE - 1 DAY, 'yyyymmdd') AND DKYE > 0 AND
                                DZHXBZ = 'N'
  THEN DKYE END) / 10000 - sum(CASE WHEN DQRQ < to_char(
      CURRENT_DATE - (MONTH(CURRENT_DATE) - 1) MONTH - day(CURRENT_DATE) DAY, 'yyyymmdd') AND DKYE > 0 AND
                                DZHXBZ = 'N'
  THEN DKYE END) / 10000 bl_y,
  1 as updated

           FROM REPORT.ODS_DKFHZ a
             LEFT JOIN REPORT.ODS_XDYWJJXXB b ON a.ZH = b.LSH"
db2 "export to /fr/data/xms/gg_all_data_save_top_10.del of del
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
          WHEN '东台弓京港支行'
            THEN '弶港'
          ELSE trim(substr(replace(JGJC, '支行', ''), 5, length(replace(JGJC, '支行', '')) / 2 * 3 - 4)) END) JGJC
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
            ) a LEFT JOIN REPORT.ODS_HXXTJGM b ON a.JGM = b.JGM WHERE a.JGM <> '320919900'
       ORDER BY ckz DESC
       FETCH FIRST 10 ROWS ONLY) a"
db2 "export to /fr/data/xms/gg_all_data_save_last_10.del of del
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
          WHEN '东台弓京港支行'
            THEN '弶港'
          ELSE trim(substr(replace(JGJC, '支行', ''), 5, length(replace(JGJC, '支行', '')) / 2 * 3 - 4)) END) JGJC
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
            ) a LEFT JOIN REPORT.ODS_HXXTJGM b ON a.JGM = b.JGM WHERE a.JGM <> '320919900'
       ORDER BY ckz
       FETCH FIRST 10 ROWS ONLY) a"

db2 "export to /fr/data/xms/gg_all_data_loan_top_10.del of del
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
          WHEN '东台弓京港支行'
            THEN '弶港'
          ELSE trim(substr(replace(JGJC, '支行', ''), 5, length(replace(JGJC, '支行', '')) / 2 * 3 - 4)) END) JGJC
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
            ) a LEFT JOIN REPORT.ODS_HXXTJGM b ON a.JGM = b.JGM WHERE a.JGM <> '320919900'
       ORDER BY ckz DESC
       FETCH FIRST 10 ROWS ONLY) a"
db2 "export to /fr/data/xms/gg_all_data_loan_last_10.del of del
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
          WHEN '东台弓京港支行'
            THEN '弶港'
          ELSE trim(substr(replace(JGJC, '支行', ''), 5, length(replace(JGJC, '支行', '')) / 2 * 3 - 4)) END) JGJC
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
            ) a LEFT JOIN REPORT.ODS_HXXTJGM b ON a.JGM = b.JGM WHERE a.JGM <> '320919900'
       ORDER BY ckz
       FETCH FIRST 10 ROWS ONLY) a"

db2 "export to /fr/data/xms/gg_all_data_map.del of del
SELECT
  CURRENT_DATE data_date,
  z.*,
  1            updated
FROM (
       SELECT
         '000'              JGM,
         sum(ZBSJ1) / 10000 ck,
         '台城片'              JGJC
       FROM REPORT.ODS_ZBTJB
       WHERE SJRQ = to_char(CURRENT_DATE - 1 DAY, 'yyyymmdd') AND ZBDH = '0001'
             AND JGM IN
                 ('320919028', '320919029', '320919030', '320919032', '320919037', '320919038', '320919045', '320919046', '320919047', '320919048')
       UNION ALL
       SELECT
         '001'              JGM,
         sum(ZBSJ1) / 10000 ck,
         '清算中心及卡部'          JGJC
       FROM REPORT.ODS_ZBTJB
       WHERE SJRQ = to_char(CURRENT_DATE - 1 DAY, 'yyyymmdd') AND ZBDH = '0001'
             AND JGM IN
                 ('320919900', '320919991')
       UNION ALL
       SELECT
         '002'              JGM,
         sum(ZBSJ1) / 10000 ck,
         '东蹲片'              JGJC
       FROM REPORT.ODS_ZBTJB
       WHERE SJRQ = to_char(CURRENT_DATE - 1 DAY, 'yyyymmdd') AND ZBDH = '0001'
             AND JGM IN
                 ('320919009', '320919021', '320919022', '320919023', '320919024', '320919025', '320919026', '320919043')
       UNION ALL
       SELECT
         '004'              JGM,
         sum(ZBSJ1) / 10000 ck,
         '东南片'              JGJC
       FROM REPORT.ODS_ZBTJB
       WHERE SJRQ = to_char(CURRENT_DATE - 1 DAY, 'yyyymmdd') AND ZBDH = '0001'
             AND JGM IN
                 ('320919015', '320919016', '320919017', '320919019', '320919020', '320919027', '320919041', '320919042')
       UNION ALL
       SELECT
         '005'              JGM,
         sum(ZBSJ1) / 10000 ck,
         '堤西片'              JGJC
       FROM REPORT.ODS_ZBTJB
       WHERE SJRQ = to_char(CURRENT_DATE - 1 DAY, 'yyyymmdd') AND ZBDH = '0001'
             AND JGM IN
                 ('320919001', '320919002', '320919003', '320919004', '320919005', '320919006', '320919007', '320919039')
       UNION ALL
       SELECT
         '006'              JGM,
         sum(ZBSJ1) / 10000 ck,
         '通榆片'              JGJC
       FROM REPORT.ODS_ZBTJB
       WHERE SJRQ = to_char(CURRENT_DATE - 1 DAY, 'yyyymmdd') AND ZBDH = '0001'
             AND JGM IN
                 ('320919008', '320919010', '320919011', '320919012', '320919013', '320919014', '320919018', '320919034', '320919040')
     ) z"

db2 "export to /fr/data/xms/ggzhcx_buliang.del of del
SELECT
  CURRENT_DATE             sjrq,
  JGJC,
  sum(CASE WHEN DQRQ < to_char(CURRENT_DATE - 1 DAY, 'yyyymmdd') AND DKYE > 0 AND
                DZHXBZ = 'N'
    THEN DKYE END) / 10000 不良,
  sum(CASE WHEN DQRQ < to_char(CURRENT_DATE - 1 DAY, 'yyyymmdd') AND DKYE > 0 AND
                DZHXBZ = 'N'
    THEN DKYE END) / 10000 -
  sum(CASE WHEN DQRQ < to_char(last_day(add_months(CURRENT_DATE - 1 day, -1)), 'yyyymmdd') AND DKYE > 0 AND
                DZHXBZ = 'N'
    THEN DKYE END) / 10000 不良m,
  sum(CASE WHEN DQRQ < to_char(CURRENT_DATE - 1 DAY, 'yyyymmdd') AND DKYE > 0 AND
                DZHXBZ = 'N'
    THEN DKYE END) / 10000 - sum(CASE WHEN DQRQ < to_char(CURRENT_DATE - 2 DAY, 'yyyymmdd') AND DKYE > 0 AND
                                           DZHXBZ = 'N'
    THEN DKYE END) / 10000 不良d,
  sum(CASE WHEN DQRQ < to_char(CURRENT_DATE - 1 DAY, 'yyyymmdd') AND DKYE > 0 AND
                DZHXBZ = 'N'
    THEN DKYE END) / 10000 - sum(CASE WHEN DQRQ < to_char(
      CURRENT_DATE - (MONTH(CURRENT_DATE) - 1) MONTH - day(CURRENT_DATE) DAY, 'yyyymmdd') AND DKYE > 0 AND
                                           DZHXBZ = 'N'
    THEN DKYE END) / 10000 不良y,
  c.JGM,
  JGQC,
  1 as updated
FROM REPORT.ODS_DKFHZ a
  LEFT JOIN REPORT.ODS_XDYWJJXXB b ON a.ZH = b.LSH
  LEFT JOIN REPORT.ODS_HXXTJGM c ON a.JGM = c.JGM
GROUP BY JGJC, c.JGM, JGQC"
db2 "export to /fr/data/xms/ggzhcx_zbsj.del of del
SELECT
  CURRENT_DATE              sjrq,
  JGJC,
  sum(CASE WHEN ZBDH = '0001' AND SJRQ = to_char(CURRENT_DATE - 1 DAY, 'yyyymmdd')
    THEN ZBSJ1 END) / 10000 存现,
  sum(CASE WHEN ZBDH = '0002' AND SJRQ = to_char(CURRENT_DATE - 1 DAY, 'yyyymmdd')
    THEN ZBSJ1 END) / 10000 贷现,
  sum(CASE WHEN ZBDH = '0010' AND SJRQ = to_char(CURRENT_DATE - 1 DAY, 'yyyymmdd')
    THEN ZBSJ1 END) / 10000 定现,
  sum(CASE WHEN ZBDH = '0011' AND SJRQ = to_char(CURRENT_DATE - 1 DAY, 'yyyymmdd')
    THEN ZBSJ1 END) / 10000 公现,
  sum(CASE WHEN ZBDH = '0006' AND SJRQ = to_char(CURRENT_DATE - 1 DAY, 'yyyymmdd')
    THEN ZBSJ1 END) / 10000 农贷现,
  ----
  sum(CASE WHEN ZBDH = '0001' AND SJRQ = to_char(CURRENT_DATE - 1 DAY, 'yyyymmdd')
    THEN ZBSJ1 END) / 10000 - sum(CASE WHEN ZBDH = '0001' AND SJRQ = to_char(
      CURRENT_DATE - (MONTH(CURRENT_DATE) - 1) MONTH - day(CURRENT_DATE) DAY, 'yyyymmdd')
    THEN ZBSJ1 END) / 10000 存年,
  sum(CASE WHEN ZBDH = '0002' AND SJRQ = to_char(CURRENT_DATE - 1 DAY, 'yyyymmdd')
    THEN ZBSJ1 END) / 10000 - sum(CASE WHEN ZBDH = '0002' AND SJRQ = to_char(
      CURRENT_DATE - (MONTH(CURRENT_DATE) - 1) MONTH - day(CURRENT_DATE) DAY, 'yyyymmdd')
    THEN ZBSJ1 END) / 10000 贷年,
  sum(CASE WHEN ZBDH = '0010' AND SJRQ = to_char(CURRENT_DATE - 1 DAY, 'yyyymmdd')
    THEN ZBSJ1 END) / 10000 - sum(CASE WHEN ZBDH = '0010' AND SJRQ = to_char(
      CURRENT_DATE - (MONTH(CURRENT_DATE) - 1) MONTH - day(CURRENT_DATE) DAY, 'yyyymmdd')
    THEN ZBSJ1 END) / 10000 定年,
  sum(CASE WHEN ZBDH = '0011' AND SJRQ = to_char(CURRENT_DATE - 1 DAY, 'yyyymmdd')
    THEN ZBSJ1 END) / 10000 - sum(CASE WHEN ZBDH = '0011' AND SJRQ = to_char(
      CURRENT_DATE - (MONTH(CURRENT_DATE) - 1) MONTH - day(CURRENT_DATE) DAY, 'yyyymmdd')
    THEN ZBSJ1 END) / 10000 公年,
  sum(CASE WHEN ZBDH = '0006' AND SJRQ = to_char(CURRENT_DATE - 1 DAY, 'yyyymmdd')
    THEN ZBSJ1 END) / 10000 - sum(CASE WHEN ZBDH = '0006' AND SJRQ = to_char(
      CURRENT_DATE - (MONTH(CURRENT_DATE) - 1) MONTH - day(CURRENT_DATE) DAY, 'yyyymmdd')
    THEN ZBSJ1 END) / 10000 农贷年,
  -----
  sum(CASE WHEN ZBDH = '0001' AND SJRQ = to_char(CURRENT_DATE - 1 DAY, 'yyyymmdd')
    THEN ZBSJ1 END) / 10000 -
  sum(CASE WHEN ZBDH = '0001' AND SJRQ = to_char(last_day(add_months(CURRENT_DATE - 1 day, -1)), 'yyyymmdd')
    THEN ZBSJ1 END) / 10000 存月,
  sum(CASE WHEN ZBDH = '0002' AND SJRQ = to_char(CURRENT_DATE - 1 DAY, 'yyyymmdd')
    THEN ZBSJ1 END) / 10000 -
  sum(CASE WHEN ZBDH = '0002' AND SJRQ = to_char(last_day(add_months(CURRENT_DATE - 1 day, -1)), 'yyyymmdd')
    THEN ZBSJ1 END) / 10000 贷月,
  sum(CASE WHEN ZBDH = '0010' AND SJRQ = to_char(CURRENT_DATE - 1 DAY, 'yyyymmdd')
    THEN ZBSJ1 END) / 10000 -
  sum(CASE WHEN ZBDH = '0010' AND SJRQ = to_char(last_day(add_months(CURRENT_DATE - 1 day, -1)), 'yyyymmdd')
    THEN ZBSJ1 END) / 10000 定月,
  sum(CASE WHEN ZBDH = '0011' AND SJRQ = to_char(CURRENT_DATE - 1 DAY, 'yyyymmdd')
    THEN ZBSJ1 END) / 10000 -
  sum(CASE WHEN ZBDH = '0011' AND SJRQ = to_char(last_day(add_months(CURRENT_DATE - 1 day, -1)), 'yyyymmdd')
    THEN ZBSJ1 END) / 10000 公月,
  sum(CASE WHEN ZBDH = '0006' AND SJRQ = to_char(CURRENT_DATE - 1 DAY, 'yyyymmdd')
    THEN ZBSJ1 END) / 10000 -
  sum(CASE WHEN ZBDH = '0006' AND SJRQ = to_char(last_day(add_months(CURRENT_DATE - 1 day, -1)), 'yyyymmdd')
    THEN ZBSJ1 END) / 10000 农贷月,
  JGM,
  JGQC,
  1 as updated,
  -----
  sum(CASE WHEN ZBDH = '0001' AND SJRQ = to_char(CURRENT_DATE - 1 DAY, 'yyyymmdd')
    THEN ZBSJ1 END) / 10000 -
  sum(CASE WHEN ZBDH = '0001' AND SJRQ = to_char(CURRENT_DATE - 2 DAY, 'yyyymmdd')
    THEN ZBSJ1 END) / 10000 存日,
  sum(CASE WHEN ZBDH = '0002' AND SJRQ = to_char(CURRENT_DATE - 1 DAY, 'yyyymmdd')
    THEN ZBSJ1 END) / 10000 -
  sum(CASE WHEN ZBDH = '0002' AND SJRQ = to_char(CURRENT_DATE - 2 DAY, 'yyyymmdd')
    THEN ZBSJ1 END) / 10000 贷日,
  sum(CASE WHEN ZBDH = '0010' AND SJRQ = to_char(CURRENT_DATE - 1 DAY, 'yyyymmdd')
    THEN ZBSJ1 END) / 10000 -
  sum(CASE WHEN ZBDH = '0010' AND SJRQ = to_char(CURRENT_DATE - 2 DAY, 'yyyymmdd')
    THEN ZBSJ1 END) / 10000 定日,
  sum(CASE WHEN ZBDH = '0011' AND SJRQ = to_char(CURRENT_DATE - 1 DAY, 'yyyymmdd')
    THEN ZBSJ1 END) / 10000 -
  sum(CASE WHEN ZBDH = '0011' AND SJRQ = to_char(CURRENT_DATE - 2 DAY, 'yyyymmdd')
    THEN ZBSJ1 END) / 10000 公日,
  sum(CASE WHEN ZBDH = '0006' AND SJRQ = to_char(CURRENT_DATE - 1 DAY, 'yyyymmdd')
    THEN ZBSJ1 END) / 10000 -
  sum(CASE WHEN ZBDH = '0006' AND SJRQ = to_char(CURRENT_DATE - 2 DAY, 'yyyymmdd')
    THEN ZBSJ1 END) / 10000 农贷日
FROM
  (SELECT
     a.*,
     b.JGJC,
     b.JGQC
   FROM REPORT.ODS_ZBTJB a LEFT JOIN REPORT.ODS_HXXTJGM b ON a.JGM = b.JGM
   WHERE SJRQ IN (to_char(CURRENT_DATE - 1 DAY, 'yyyymmdd'), to_char(
       CURRENT_DATE - (MONTH(CURRENT_DATE) - 1) MONTH - day(CURRENT_DATE) DAY, 'yyyymmdd'),
                  to_char(last_day(add_months(CURRENT_DATE - 1 day, -1)), 'yyyymmdd'), to_char(CURRENT_DATE - 2 DAY,
                                                                                       'yyyymmdd')))
GROUP BY JGJC, JGM, JGQC"
db2 "export to /fr/data/xms/trans_f10rec_all_daikuan.del of del
SELECT a.*,
  1 as updated
FROM (SELECT
        CURRENT_DATE                sjrq,
        tmp1.LN_LN_ACCT_NO,
        tmp2.LN_CUST_NAME,
        (tmp1.LN_TX_AMT / 10000)    LN_TX_AMT,
        (tmp1.LN_ATX_BAL / 10000)   LN_ATX_BAL,
        trim(tmp3.JGJC)             SNAME,
        tmp1.LN_BELONG_INSTN_COD,
        rank()
        OVER (
          PARTITION BY tmp1.LN_BELONG_INSTN_COD
          ORDER BY LN_TX_AMT DESC ) rn
      FROM CBOD.LNLNSJRN0 tmp1
        LEFT JOIN CBOD.LNLNSLNS tmp2 ON tmp1.LN_LN_ACCT_NO = tmp2.LN_LN_ACCT_NO
        LEFT JOIN REPORT.ODS_HXXTJGM tmp3 ON tmp1.LN_BELONG_INSTN_COD = tmp3.JGM
      WHERE
        tmp1.LN_ENTR_DT_N = to_char(current date - 1 DAY, 'YYYYMMDD')
        AND tmp1.LN_DR_CR_COD = 'DR') a
WHERE rn <= 10"
db2 "export to /fr/data/xms/trans_f10rec_all_dingcun.del of del
SELECT a.*,
  1 as updated
FROM (SELECT
        CURRENT_DATE                sjrq,
        tmp1.FK_TDACN_KEY,
        tmp2.TD_CUST_NAME,
        (tmp1.TD_TX_AMT / 10000)    TD_TX_AMT,
        (tmp1.TD_ACCT_BAL / 10000)  TD_ACCT_BAL,
        trim(tmp3.JGJC)             SNAME,
        tmp1.TD_BELONG_INSTN_COD,
        rank()
        OVER (
          PARTITION BY tmp1.TD_BELONG_INSTN_COD
          ORDER BY TD_TX_AMT DESC ) rn
      FROM CBOD.TDACNINT tmp1
        LEFT JOIN CBOD.TDACNACN tmp2 ON tmp1.FK_TDACN_KEY = tmp2.TD_TD_ACCT_NO
        LEFT JOIN REPORT.ODS_HXXTJGM tmp3 ON tmp1.TD_BELONG_INSTN_COD = tmp3.JGM
      WHERE
        tmp1.SRC_DT = to_char(current date - 1 DAY, 'YYYYMMDD')
        AND tmp1.TD_DR_CR_COD = 'CR') a
WHERE rn <= 10"
db2 "export to /fr/data/xms/trans_f10rec_all_dingqu.del of del
SELECT a.*,
  1 as updated
FROM (
  SELECT
    CURRENT_DATE                sjrq,
    tmp1.FK_TDACN_KEY,
    tmp2.TD_CUST_NAME,
    (tmp1.TD_TX_AMT / 10000)    TD_TX_AMT,
    (tmp1.TD_ACCT_BAL / 10000)  TD_ACCT_BAL,
    trim(tmp3.JGJC)             SNAME,
    tmp1.TD_BELONG_INSTN_COD,
    rank()
    OVER (
      PARTITION BY tmp1.TD_BELONG_INSTN_COD
      ORDER BY TD_TX_AMT DESC ) rn
  FROM CBOD.TDACNINT tmp1
    LEFT JOIN CBOD.TDACNACN tmp2 ON tmp1.FK_TDACN_KEY = tmp2.TD_TD_ACCT_NO
    LEFT JOIN REPORT.ODS_HXXTJGM tmp3 ON tmp1.TD_BELONG_INSTN_COD = tmp3.JGM
  WHERE
    tmp1.SRC_DT = to_char(current date - 1 DAY, 'YYYYMMDD')
    AND tmp1.TD_DR_CR_COD = 'DR') a
WHERE rn <= 10"
db2 "export to /fr/data/xms/trans_f10rec_all_huankuan.del of del
SELECT a.*,
  1 as updated
FROM (
  SELECT
    CURRENT_DATE                sjrq,
    tmp1.LN_LN_ACCT_NO,
    tmp2.LN_CUST_NAME,
    (tmp1.LN_TX_AMT / 10000)    LN_TX_AMT,
    (tmp1.LN_ATX_BAL / 10000)   LN_ATX_BAL,
    trim(tmp3.JGJC)             SNAME,
    tmp1.LN_BELONG_INSTN_COD,
    rank()
    OVER (
      PARTITION BY tmp1.LN_BELONG_INSTN_COD
      ORDER BY LN_TX_AMT DESC ) rn
  FROM CBOD.LNLNSJRN0 tmp1
    LEFT JOIN CBOD.LNLNSLNS tmp2 ON tmp1.LN_LN_ACCT_NO = tmp2.LN_LN_ACCT_NO
    LEFT JOIN REPORT.ODS_HXXTJGM tmp3 ON tmp1.LN_BELONG_INSTN_COD = tmp3.JGM
  WHERE
    tmp1.LN_ENTR_DT_N = to_char(current date - 1 DAY, 'YYYYMMDD')
    AND tmp1.LN_DR_CR_COD = 'CR') a
WHERE rn <= 10"
db2 "export to /fr/data/xms/trans_f10rec_all_huocun.del of del
SELECT a.*,
  1 as updated
FROM (
  SELECT
    CURRENT_DATE                   sjrq,
    tmp1.FK_SAACN_KEY,
    tmp2.SA_CUST_NAME,
    (tmp1.SA_TX_AMT / 10000)       SA_TX_AMT,
    (tmp1.SA_DDP_ACCT_BAL / 10000) SA_DDP_ACCT_BAL,
    trim(tmp3.JGJC)                SNAME,
    tmp1.SA_BELONG_INSTN_COD,
    rank()
    OVER (
      PARTITION BY tmp1.SA_BELONG_INSTN_COD
      ORDER BY SA_TX_AMT DESC )    rn
  FROM CBOD.SAACNTXN tmp1
    LEFT JOIN CBOD.SAACNACN tmp2 ON tmp1.FK_SAACN_KEY = tmp2.SA_ACCT_NO
    LEFT JOIN REPORT.ODS_HXXTJGM tmp3 ON tmp1.SA_BELONG_INSTN_COD = tmp3.JGM
  WHERE
    tmp1.SRC_DT = to_char(current date - 1 DAY, 'YYYYMMDD')
    AND tmp1.SA_EC_DET_NO_N = '0'
    AND tmp1.SA_DR_AMT = '0') a
WHERE rn <= 10"
db2 "export to /fr/data/xms/trans_f10rec_all_huoqu.del of del
SELECT a.*,
  1 as updated
FROM (
  SELECT
    CURRENT_DATE                   sjrq,
    tmp1.FK_SAACN_KEY,
    tmp2.SA_CUST_NAME,
    (tmp1.SA_TX_AMT / 10000)       SA_TX_AMT,
    (tmp1.SA_DDP_ACCT_BAL / 10000) SA_DDP_ACCT_BAL,
    trim(tmp3.JGJC)                SNAME,
    tmp1.SA_BELONG_INSTN_COD,
    rank()
    OVER (
      PARTITION BY tmp1.SA_BELONG_INSTN_COD
      ORDER BY SA_TX_AMT DESC )    rn
  FROM CBOD.SAACNTXN tmp1
    LEFT JOIN CBOD.SAACNACN tmp2 ON tmp1.FK_SAACN_KEY = tmp2.SA_ACCT_NO
    LEFT JOIN REPORT.ODS_HXXTJGM tmp3 ON tmp1.SA_BELONG_INSTN_COD = tmp3.JGM
  WHERE
    tmp1.SRC_DT = to_char(current date - 1 DAY, 'YYYYMMDD')
    AND tmp1.SA_EC_DET_NO_N = '0'
    AND tmp1.SA_CR_AMT = '0'
) a
WHERE rn <= 10"
DATE=/fr/data/xms/$(date -d yesterday +%Y%m%d).over
touch $DATE

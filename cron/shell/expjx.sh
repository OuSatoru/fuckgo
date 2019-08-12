#! /usr/bin/env sh

db2 connect to bcas user bcas using bcas
db2 "export to /fr/data/xms/bcas_zzzj.del of del
SELECT
  CURRENT_DATE SJRQ,
  B.DEPT_ID             JGM,
  B.DEPT_NAME           机构名,
  B.PARENT_ID, b.FULLNAME,
  SUM(NVL(T.ZCKYE, 0))  总存款,
  SUM(NVL(T.CXCK, 0))   储蓄存款,
  SUM(NVL(T.HQCK, 0))   活期存款,
  SUM(NVL(T.DQ1YE, 0))  定期一年及以下,
  SUM(NVL(T.DQ35YE, 0)) 定期二三年,
  SUM(NVL(T.DQ5YE, 0))  定期五年,
  SUM(NVL(T.DGCK, 0))   对公存款,
  SUM(NVL(T.BZJCK, 0))  保证金存款,
  SUM(NVL(T.KNHQCK, 0)) 卡内活期,
  SUM(NVL(T.KNDQCK, 0)) 卡内定期,
  B.SORT
FROM (SELECT
        USER_ID,
        SUM(CASE WHEN FORMULA_CODE = '21000000_TC'
          THEN BALANCE
            ELSE 0 END) ZCKYE,
        SUM(CASE WHEN FORMULA_CODE IN
                      ('21020101_TC', '21020102_TC', '21020201_TC', '21020202_TC', '21020203_TC', '21020204_TC', '21020205_TC', '21020211_TC', '21020212_TC', '21020213_TC', '21020207_TC', '21020208_TC', '21020209_TC', '21020210_TC')
          THEN BALANCE
            ELSE 0 END) CXCK,
        SUM(CASE WHEN FORMULA_CODE IN ('21020101_TC', '21020102_TC')
          THEN BALANCE
            ELSE 0 END) HQCK,
        SUM(CASE WHEN FORMULA_CODE IN
                      ('21020201_TC', '21020202_TC', '21020203_TC', '21020204_TC', '21020205_TC', '21020211_TC', '21020212_TC', '21020213_TC')
          THEN BALANCE
            ELSE 0 END) DQ1YE,
        SUM(CASE WHEN FORMULA_CODE IN ('21020207_TC', '21020208_TC')
          THEN BALANCE
            ELSE 0 END) DQ35YE,
        SUM(CASE WHEN FORMULA_CODE IN ('21020209_TC', '21020210_TC')
          THEN BALANCE
            ELSE 0 END) DQ5YE,
        SUM(CASE WHEN FORMULA_CODE IN ('21010000_TC', '21030000_TC', '21040000_TC', '21050000_TC', '21060000_TC', '21070000_TC')
          THEN BALANCE
            ELSE 0 END) DGCK,
        SUM(CASE WHEN FORMULA_CODE = '21030000_TC'
          THEN BALANCE
            ELSE 0 END) BZJCK,
        SUM(CASE WHEN FORMULA_CODE IN ('21010102_TC', '21020101_TC')
          THEN BALANCE
            ELSE 0 END) KNHQCK,
        SUM(CASE WHEN FORMULA_CODE = '21100000_TC'
          THEN BALANCE
            ELSE 0 END) KNDQCK
      FROM BCAS.D_PER_ACHV_2019
      WHERE ETLDT = CURRENT_DATE - 1 DAY AND LEFT(FORMULA_CODE, 1) LIKE '2%' AND
            LEFT(FORMULA_CODE, 4) NOT IN ('2109', '2509', '2609', '2709')
      GROUP BY USER_ID) T INNER JOIN (SELECT
                                        su.id,
                                        sd.DEPT_ID,
                                        sd.DEPT_NAME,
  sd.PARENT_ID, sd.FULLNAME,
                                        sd.SORT
                                      FROM (SELECT *
                                            FROM BCAS.SS_USER_BAK
                                            WHERE SSSQ = CURRENT_DATE - 1 DAY AND USERJB != '99') su LEFT JOIN bcas.ss_dept sd
                                          ON su.USERDEPT = sd.DEPT_ID) B ON T.USER_ID = B.ID
GROUP BY B.DEPT_ID, B.DEPT_NAME, b.PARENT_ID, b.FULLNAME, B.SORT
ORDER BY B.DEPT_ID"
db2 "export to /fr/data/xms/bcas_gr.del of del
SELECT
  CURRENT_DATE SJRQ,
  USER_ID,
  b.CHSNAME,
  b.DEPT_ID,
  b.DEPT_NAME,
  b.PARENT_ID,
  b.FULLNAME,
  sum(BALANCE)   total,
  sum(AVG_MONTH) avg_month,
  sum(AVG_QTR)   avg_qtr,
  sum(AVG_YEAR)  avg_year
FROM BCAS.D_PER_ACHV_2019 a
  LEFT JOIN (SELECT
               su.id,
               su.CHSNAME,
               sd.DEPT_ID,
               sd.DEPT_NAME,
               sd.PARENT_ID,
               sd.FULLNAME,
               sd.SORT
             FROM (SELECT *
                   FROM BCAS.SS_USER_BAK
                   WHERE SSSQ = CURRENT_DATE - 1 DAY AND USERJB != '99') su LEFT JOIN bcas.ss_dept sd
                 ON su.USERDEPT = sd.DEPT_ID) b ON a.USER_ID = b.ID
WHERE ETLDT = CURRENT_DATE - 1 DAY AND FORMULA_CODE = '21000000_TC'
GROUP BY USER_ID, b.CHSNAME, b.DEPT_ID, b.DEPT_NAME, b.PARENT_ID, b.FULLNAME"
DATE=/fr/data/xms/$(date -d yesterday +%Y%m%d).jx
touch $DATE
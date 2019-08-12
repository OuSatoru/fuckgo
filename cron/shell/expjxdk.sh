#! /usr/bin/env sh

db2 connect to bcas user bcas using bcas
db2 "export to /fr/data/xms/bcas_jgdk.del of del
select
  CURRENT_DATE SJRQ,
  x8,
  x1,
  PARENT_ID, FULLNAME,
  x2,
  x3,
  x4,
  x5,
  x6,
  x7
from (SELECT
        B.DEPT_NAME            X1,
        SUM(NVL(T.GXDK, 0.00)) X2,
        SUM(NVL(T.TXZC, 0.00)) X3,
        SUM(NVL(T.GRDK, 0.00)) X4,
        SUM(NVL(T.FNGR, 0.00)) X5,
        SUM(NVL(T.GSDK, 0.00)) X6,
        SUM(NVL(T.FNGS, 0.00)) X7,
        B.DEPT_ID              x8,
        b.PARENT_ID,
  b.FULLNAME
      FROM (SELECT
              user_dept,
              SUM(CASE WHEN FORMULA_CODE = '11000000'
                THEN balance
                  ELSE 0.00 END) GXDK,
              SUM(CASE WHEN FORMULA_CODE = '11030001'
                THEN balance
                  ELSE 0.00 END) TXZC,
              SUM(CASE WHEN FORMULA_CODE = '11020000'
                THEN balance
                  ELSE 0.00 END) GRDK,
              SUM(CASE WHEN FORMULA_CODE = '11020400'
                THEN balance
                  ELSE 0.00 END) FNGR,
              SUM(CASE WHEN FORMULA_CODE = '11010000'
                THEN balance
                  ELSE 0.00 END) GSDK,
              SUM(CASE WHEN FORMULA_CODE = '11010300'
                THEN balance
                  ELSE 0.00 END) FNGS
            FROM D_PER_ACHV_2019
            WHERE ETLDT = CURRENT_DATE - 1 DAY AND LEFT(FORMULA_CODE, 1) LIKE '1%'
            GROUP BY user_dept) T INNER JOIN (SELECT *
                                              FROM SS_DEPT A
                                              WHERE STATUS = '0') B ON (T.user_dept = B.DEPT_ID)
      GROUP BY B.DEPT_ID, B.DEPT_NAME, b.PARENT_ID, b.FULLNAME, B.SORT)
where 1 = 1
ORDER BY x8"
DATE=/fr/data/xms/$(date -d yesterday +%Y%m%d).jxdk
touch $DATE
#! /usr/bin/env sh

db2 connect to bcas user bcas using bcas
db2 "export to /fr/data/xms/bcas_sjyh.del of del
select current_date sjrq, t.*
from (SELECT
        b.dept_id            ,
        b.DEPT_NAME, b.PARENT_ID, b.FULLNAME,
        INT(SUM(CASE WHEN FORMULA_CODE = '57000001'
          THEN a.BALANCE
                ELSE 0 END)) zhs,
        INT(SUM(CASE WHEN FORMULA_CODE = '57000002'
          THEN a.BALANCE
                ELSE 0 END)) dyxkh,
        INT(SUM(CASE WHEN FORMULA_CODE = '57000003'
          THEN a.BALANCE
                ELSE 0 END)) yxh,
        INT(SUM(CASE WHEN FORMULA_CODE = '57000004'
          THEN a.BALANCE
                ELSE 0 END)) dyxzyxh,
        INT(SUM(CASE WHEN FORMULA_CODE = '57000005'
          THEN a.BALANCE
                ELSE 0 END)) X5,
        INT(SUM(CASE WHEN FORMULA_CODE = '57000006'
          THEN a.BALANCE
                ELSE 0 END)) X6,
        INT(SUM(CASE WHEN FORMULA_CODE = '57000007'
          THEN a.BALANCE
                ELSE 0 END)) bdh,
        INT(SUM(CASE WHEN FORMULA_CODE = '57000013'
          THEN a.BALANCE
                ELSE 0 END)) dnxkh,
        SUM(CASE WHEN FORMULA_CODE = '57000008'
          THEN a.BALANCE
            ELSE 0 END)      dyjyje,
        INT(SUM(CASE WHEN FORMULA_CODE = '57000009'
          THEN a.BALANCE
                ELSE 0 END)) dyjybs,
        SUM(CASE WHEN FORMULA_CODE = '57000010'
          THEN a.BALANCE
            ELSE 0 END)      qnjyje,
        INT(SUM(CASE WHEN FORMULA_CODE = '57000011'
          THEN a.BALANCE
                ELSE 0 END)) qnjybs
      from BCAS.D_PER_ACHV_2019 a inner join (select
                                           dept_id,
                                           dept_name, PARENT_ID, FULLNAME
                                         from BCAS.ss_dept) b on a.chk_instn_cod = b.dept_id
      where 1 = 1 AND a.FORMULA_CODE LIKE '57%' and a.etldt = CURRENT_DATE - 1 DAY
      group by b.dept_id, b.DEPT_NAME, b.PARENT_ID, b.FULLNAME) t
order by t.DEPT_ID"
DATE=/fr/data/xms/$(date -d yesterday +%Y%m%d).jxsjyh
touch $DATE
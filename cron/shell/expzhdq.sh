#! /usr/bin/env sh

db2 connect to jsbods user ods using ods@98
db2 "export to /fr/data/xms/zh_ck_dq_mx.del of del
SELECT
  CURRENT_DATE   sjrq,
  TD_BELONG_INSTN_COD,
  TD_ACCT_STS,
  '存款到期'      AS FLAG,
  TD_TD_ACCT_NO,
  TD_CUST_NAME,
  TD_DUE_DT,
  TD_ACTU_AMT,
  CASE WHEN tmp1.TD_DEP_PRD_N IN ('ST02', 'ST03')
    THEN '通知存款'
  ELSE '' END AS REMARK,
  1 as updated,
  TD_OPAC_DT
FROM CBOD.TDACNACN tmp1 LEFT JOIN CBOD.CMEMPEMP tmp2 ON tmp1.TD_CONNTR_NO = tmp2.CM_EMP_NO
WHERE
  tmp1.TD_DUE_DT >= to_char(current date, 'YYYYMMDD') AND tmp1.TD_DUE_DT < to_char(current date + 30 DAYS, 'YYYYMMDD')
  AND TD_ACCT_STS = '01' AND TD_ACTU_AMT >= '50000'
ORDER BY tmp1.TD_DUE_DT"
db2 "export to /fr/data/xms/zh_dk_dq_mx.del of del
SELECT
  CURRENT_DATE sjrq,
  LN_BELONG_INSTN_COD,
  LN_LN_ACCT_NO,
  LN_CUST_NAME,
  LN_DUE_DT_N,
  LN_LN_BAL,
  tmp2.USERNAME,
  '贷款到期' AS TYPE,
  1 updated,
  LN_CRNT_PRD_PAYRBL_INT
FROM CBOD.LNLNSLNS tmp1
  LEFT JOIN CMIS.USER_INFO tmp2 ON tmp1.LN_ZHIBIAO_NO = tmp2.USERID
WHERE LN_ACCT_STS = '01' AND LN_DUE_DT_N >= to_char(current date, 'YYYYMMDD') AND
      LN_DUE_DT_N <= to_char(current date + 30 DAYS, 'YYYYMMDD')
ORDER BY LN_DUE_DT_N ASC, LN_LN_BAL DESC"
DATE=/fr/data/xms/$(date -d yesterday +%Y%m%d).zhdq
touch $DATE
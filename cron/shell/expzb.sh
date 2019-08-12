#! /usr/bin/env sh

db2 connect to jsbods user ods using ods@98
db2 "export to /fr/data/xms/ODS_HXZZYEB.del of del
SELECT
  a.*,
  b.GSBZ
FROM (
       SELECT
         GL_OPUN_COD                                                            jgm,
         GL_FIRST_LEVEL_LG_COD || GL_SECOND_LEVEL_LG_CD || GL_THIRD_LEVEL_LG_CD kmh,
         GL_DR_AMT                                                              jffse,
         GL_CR_AMT                                                              dffse,
         GL_DR_BAL                                                              jfye,
         GL_CR_BAL                                                              dfye,
         CASE WHEN GL_SECOND_LEVEL_LG_CD = '00' AND GL_THIRD_LEVEL_LG_CD = '00'
           THEN '1'
         WHEN GL_SECOND_LEVEL_LG_CD <> '00' AND GL_THIRD_LEVEL_LG_CD = '00'
           THEN '2'
         ELSE '3'
         END                                                                    kmjb,
         SRC_DT                                                                 sjrq
       FROM CBOD.GLGLGGLG
       WHERE SRC_DT = to_char(CURRENT_DATE - 1 DAY, 'yyyymmdd')) a LEFT JOIN REPORT.ODS_CDKKMDGDSSXB b ON a.kmh = b.KMH"
DATE=/fr/data/xms/$(date -d yesterday +%Y%m%d).zb
touch $DATE
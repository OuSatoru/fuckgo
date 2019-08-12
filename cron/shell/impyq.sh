#! /usr/bin/env sh

db2 connect to jsbods user ods using ods@98

#db2 delete from fr.FDM_DKWJFLQDMXB
db2 import from /fr/data/xms/FDM_DKWJFLQDMXB.del of del insert into fr.FDM_DKWJFLQDMXB


db2 "UPDATE fr.CANUPDATE SET yq = CURRENT_DATE"

DATE=/fr/data/xms/odsimp/$(date -d yesterday +%Y%m%d).endyq
touch $DATE
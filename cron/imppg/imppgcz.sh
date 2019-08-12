#!/bin/bash

psql -U postgres -d postgres -c "delete from report.ods_hqckzd"
psql -U postgres -d postgres -c "COPY report.ods_hqckzd from '/fr/data/xms/HQCKZD_pg_cz.csv' with csv"
psql -U postgres -d postgres -c "delete from report.ods_hqckzjd"
psql -U postgres -d postgres -c "COPY report.ods_hqckzjd from '/fr/data/xms/HQCKZJD_pg_cz.csv' with csv"
psql -U postgres -d postgres -c "delete from report.ods_hqckmxz_new_lsb"
psql -U postgres -d postgres -c "COPY report.ods_hqckmxz_new_lsb from '/fr/data/xms/HQCKMXZ_pg_cz.csv' with csv"
psql -U postgres -d postgres -c "insert into report.ods_hqckmxz_new
select * from report.ods_hqckmxz_new_lsb
on conflict do nothing"
psql -U postgres -d postgres -c "delete from report.ods_hqckmxz_new_lsb"
DATE=/fr/data/xms/$(date -d yesterday +%Y%m%d).endpgcz
touch $DATE
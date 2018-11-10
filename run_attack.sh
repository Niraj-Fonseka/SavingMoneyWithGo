echo "---- Starting attack -----"

echo "---- Attacking Datasource Server ----"
vegeta attack -targets=targets_normal.txt -duration=10s -rate=10 | tee results_normal.bin | vegeta report


echo "\n\n"

echo "---- Attacking Datasource with singleflight Server ----"
vegeta attack -targets=targets_singleflight.txt -duration=10s -rate=10 | tee results_singleflight.bin | vegeta report


echo "---- Ending  attack -----"

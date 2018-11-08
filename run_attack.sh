#!/bin/bash

echo "---- Starting attack -----"

vegeta attack -targets=targets.txt -duration=5s -rate=10 | tee results.bin | vegeta report


echo "---- Ending  attack -----"

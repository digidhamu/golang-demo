#!/bin/bash

export PATH="$PATH:/Users/dhamukrish/Documents/digidhamu/k8s.do.digidhamu.com/tools/sonar-scanner-4.2.0.1873-macosx/bin"

sonar-scanner

echo "Checking if analysis is finished.."

SONAR_PROJECT_KEY=golang-daas-demo
SONAR_STATUS_URL=$(cat .scannerwork/report-task.txt | grep ceTaskUrl | sed -e 's/ceTaskUrl=//')
SONAR_STATUS=$(curl -skg "${SONAR_STATUS_URL}" | sed -e 's/.*status":"//' | sed -e 's/",.*//')

while ! [ "${SONAR_STATUS}" = "SUCCESS" ] || [ "${SONAR_STATUS}" = "CANCELED" ] || [ "${SONAR_STATUS}" = "FAILED" ];
do                                    
    echo "Sonar analysis is: ${SONAR_STATUS}. Taking a nap while we wait..."
    sleep 5
    SONAR_STATUS=$(curl -skg ${SONAR_STATUS_URL} | sed -e 's/.*status":"//' | sed -e 's/",.*//')                    
done

echo "Sonar Task returned: ${SONAR_STATUS}"

if [ "${SONAR_STATUS}" = "FAILED" ]; then                
    exit -1
fi

exit 0

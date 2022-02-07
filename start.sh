#!/bin/bash
#PROJECT_NAME=$1
PROJECT_NAME=go-admin
LOG_DIR=tmp/logs

echo "${PROJECT_NAME} shell start!"
source /etc/profile
go build
if [ ! -d ${LOG_DIR} ]; then
  mkdir -p ${LOG_DIR}
fi
pid=$(ps -ef | grep ${PROJECT_NAME} | grep -v grep | awk '{print $2}')
if [ "$pid" ]; then
  kill -15 $pid
  echo "killed service pid $pid !"
fi
./${PROJECT_NAME} >>${LOG_DIR}/$(date +%Y-%m-%d).log &
pid=$(ps -ef | grep ${PROJECT_NAME} | grep -v grep | awk '{print $2}')
echo "Service started with pid $pid !"
echo "shell end!"

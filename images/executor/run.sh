#!/bin/bash

set -a

AWS_ACCESS_KEY_ID=minio
AWS_SECRET_ACCESS_KEY=minio123
MLFLOW_TRACKING_URI=http://mlflow-service:5000/
MLFLOW_S3_ENDPOINT_URL=http://minio-service:9000/

while getopts b:s: flag
do
    case "${flag}" in
        b) backend=${OPTARG};;
        s) source=${OPTARG};;
    esac
done
basename=$(basename $source)
folder=${basename%.*}

echo "Preparing environment...."
git clone $source

sleep 2

echo "Ready to execute pipeline..."
if [ $backend == 'kfp' ];
then
    
    sleep 2

    sleep 2

fi

if [ $backend == 'mlflow' ];
then
    sleep 1
    echo 'export $AWS_ACCESS_KEY_ID'
    echo 'export $AWS_SECRET_ACCESS_KEY'
    echo 'export $MLFLOW_TRACKING_URI'
    echo 'export $MLFLOW_S3_ENDPOINT_URL'
    sleep 1
    mlflow run $folder
fi

if [ $backend == 'airflow' ];
then
    echo $backend
fi
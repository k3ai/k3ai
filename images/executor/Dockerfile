# Defining base image
FROM continuumio/miniconda3:latest AS build
WORKDIR /root
# COPY k3ai-run
RUN  apt-get update && apt-get install git -y && \
pip install --upgrade pip && \
pip install Cython mlflow[extras] psycopg2-binary boto3 kfp apache-airflow --upgrade --no-cache-dir && \
echo 'export AWS_ACCESS_KEY_ID=minio\n \
export AWS_SECRET_ACCESS_KEY=minio123\n \
export MLFLOW_TRACKING_URI=http://mlflow-service.default:5000\n \
export MLFLOW_S3_ENDPOINT_URL=http://minio-service.default:9000\n' \
>> /etc/environment

FROM build

COPY k3ai-run  /opt
RUN chmod +x /opt/k3ai-run
CMD /bin/bash -c "source /etc/environment"
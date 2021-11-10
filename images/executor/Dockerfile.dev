# Defining base image
FROM continuumio/miniconda3:latest AS build
WORKDIR /root
RUN  apt-get update && apt-get install git -y && \
pip install --upgrade pip && \
pip install Cython mlflow[extras] psycopg2-binary boto3 kfp apache-airflow --upgrade --no-cache-dir && \
echo 'export AWS_ACCESS_KEY_ID=minio\n \
export AWS_SECRET_ACCESS_KEY=minio123\n \
export MLFLOW_TRACKING_URI=http://mlflow-service.default.svc.cluster.local:5000\n \
export MLFLOW_S3_ENDPOINT_URL=http://minio-service.default.svc.cluster.local:9000\n' \
>> /etc/environment

FROM build
ENV AWS_ACCESS_KEY_ID=minio
ENV AWS_SECRET_ACCESS_KEY=minio123
ENV MLFLOW_TRACKING_URI=http://mlflow-service.default:5000
ENV MLFLOW_S3_ENDPOINT_URL=http://minio-service.default:9000

COPY k3ai-executor  /opt
RUN chmod +x /opt/k3ai-executor && conda clean --all
CMD /bin/bash -c "source /etc/environment"
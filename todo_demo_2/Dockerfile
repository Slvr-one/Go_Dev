FROM python:3.8-slim-buster
# FROM python:3-slim

WORKDIR /app

# RUN apt-get update && apt-get install -y wget gnupg2
# RUN wget -qO - https://artifacts.elastic

COPY requirements.txt requirements.txt

RUN pip3 install -r requirements.txt
# RUN pip install flask prometheus_client


COPY app.py /app.py

ENTRYPOINT [ "python", "app.py" ]



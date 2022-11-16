FROM python:3.8.0-alpine

# ソースを置くディレクトリを変数として格納                                                  
ARG project_dir=/projects/

# 必要なファイルをローカルからコンテナにコピー
ADD requirements.txt $project_dir

# requirements.txtに記載されたパッケージをインストール                         
WORKDIR $project_dir

RUN set -x \
    && apk update \
    && apk add linux-headers \
    && apk add build-base \
    && apk add python3 py3-pip python3-dev \
    #&& pip3 install --no-cache --upgrade pip setuptools \
    && pip3 install --no-cache pip setuptools==49.6.0  \
    && pip3 install wheel

RUN pip install -r requirements.txt
RUN pip install --upgrade google-api-python-client


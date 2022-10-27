FROM python:3.8.0-alpine

# ソースを置くディレクトリを変数として格納                                                  
ARG project_dir=/projects/

# 必要なファイルをローカルからコンテナにコピー
ADD requirements.txt $project_dir

# requirements.txtに記載されたパッケージをインストール                         
WORKDIR $project_dir
RUN pip install -r requirements.txt
RUN pip install --upgrade google-api-python-client

# # （コンテナ内で作業する場合）必要なパッケージをインストール
# RUN apk update                  
# RUN apk add zsh vim tmux git tig
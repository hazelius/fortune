# nginxdocker

docker buildコマンドを実行。
下記はコンテナーイメージの名前を nginxdocker/nginx、タグを1.0に指定する場合

    cd dockertest/nginxdocker
    docker build -t nginxdocker/nginx:1.0 .

8080ポートで起動

    docker run -d -p 8080:80 --name ng nginxdocker/nginx:1.0

ローカルでWebページが開けることを確認  
http://localhost:8080/

起動した名前を指定して停止

    docker stop ng

コンテナーの削除

    docker rm ng

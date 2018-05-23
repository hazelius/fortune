# php

Amazon Linuxでphpサーバーの起動

    cd dockertest/php
    docker build -t dockertest/php .

8080ポートで起動

    docker run -d -p 8080:80 --name az dockertest/php

ローカルでWebページが開けることを確認  
http://localhost:8080/

起動した名前を指定して停止

    docker stop az

コンテナーの削除

    docker rm az

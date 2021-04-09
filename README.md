# DockerでGoとNginxの環境を作る

Nginxが80番でリッスンしており、/api 以下にリクエストが来たら8080番で動いているfcgi(Go)にリクエストを投げる

## 使い方

サービスをビルドする（初回のみ）

```docker-compose build```

コンテナを起動

```docker-compose up```

コンテナを削除する

```docker-compose down -v```

## 動作確認

ブラウザからも確認できる
```
curl http://localhost:8000 
curl http://localhost:8000/api/v1/health
```
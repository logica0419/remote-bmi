# Remote-BMI

[![Client CI Pipeline](https://github.com/logica0419/remote-bmi/actions/workflows/client.yml/badge.svg)](https://github.com/logica0419/remote-bmi/actions/workflows/client.yml)
[![Server CI Pipeline](https://github.com/logica0419/remote-bmi/actions/workflows/server.yml/badge.svg)](https://github.com/logica0419/remote-bmi/actions/workflows/server.yml)
[![CodeQL](https://github.com/logica0419/remote-bmi/actions/workflows/codeql-analysis.yml/badge.svg)](https://github.com/logica0419/remote-bmi/actions/workflows/codeql-analysis.yml)

(BMI = Bench Marker Interface)  
ISUCONの練習時に使える、簡易的なベンチマーカーWebUI  
プライベートネットワークで競技用サーバーと繋がったベンチマークサーバーで動かすことを想定しています。

## 設定

リポジトリ直下に`config.json`というファイルを作って下さい。  
同様のスキーマであればyamlでも大丈夫です。

```json
{
  "address": ":3000",             // Remote-BMIがリッスンするアドレス
  "mysql": {                      // MySQLの設定
    "hostname": "localhost",        // ホスト
    "port": 3306,                   // ポート
    "username": "isucon",           // ユーザー名
    "password": "isucon",           // パスワード
    "database": "remote_bmi"        // データベース名
  },
  "version": "isucon11-qualify",  // 実行するベンチマークコマンドのバージョン
  // server/benchmark/command.go に全てのコマンドをmapでまとめ、そのキーを"version"に入れます。
  "bench_ip": "localhost"         // ベンチマークサーバーのプライベートIPアドレス
}
```

環境変数でも設定できます。

```sh
export ADDRESS=":3000"
export MYSQL_HOSTNAME="localhost"
export MYSQL_PORT="3306"
export MYSQL_USERNAME="isucon"
export MYSQL_PASSWORD="isucon"
export MYSQL_DATABASE="remote_bmi"
export VERSION="isucon11-qualify"
export BENCH_IP="localhost"
```

## ビルド

Node.js環境とGo環境が必要です。
リポジトリをクローンした上で、リポジトリのルートディレクトリでコマンドを実行してください。

```sh
make init
make build
```

リポジトリのルートディレクトリに`remote-bmi`という名前の実行用バイナリが生成されます。

## 実行

ビルドを実行した上で、バイナリファイルと設定ファイルをお好きな場所に設置し、下記のコマンドを入力するとサーバーが立ち上がります。

```sh
{バイナリまでのパス}/remote-bmi serve -c {設定ファイルのパス}
```

ベンチマークサーバー内でのsystemdによる管理をオススメします。

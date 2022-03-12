# Remote-BMI

(BMI = Bench Marker Interface)  
ISUCONの練習時に使える、簡易的なベンチマーカーWebUI  
プライベートネットワークで競技用サーバーと繋がったベンチマークサーバーで動かすことを想定しています。

## 設定

リポジトリ直下に`config.json`というファイルを作って下さい。

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

## ビルド

Node.js環境とGo環境が必要です。
リポジトリをクローンした上で、リポジトリのルートディレクトリでコマンドを実行してください。

```sh
make init
make build
```

- サーバーサイド - リポジトリのルートディレクトリに`remote-bmi`という名前の実行用バイナリ
- クライアントサイド - `client/dist`にコンパイル & バンドルされた静的ファイル

がそれぞれ生成されます。

## 実行

ビルドを実行した上で、下記のコマンドを入力するとサーバーが立ち上がります。

```sh
./remote-bmi serve
```

ベンチマークサーバー内でのsystemdによる管理をオススメします。

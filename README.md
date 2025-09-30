# about this project

会員登録とログイン後の処理を実装してみる練習です。完成していないので、あしからず。

[[_TOC_]]

# 環境構築

## sql の設定
作業ディレクトリ：プロジェクトのルート

1. sql が動作するか確認 (値は例)

```shell
$ podman run --rm -p 3307:3306 -e MYSQL_ROOT_PASSWORD=dev -e MYSQL_USER=dev -e MYSQL_PASSWORD=dev -e MYSQL_DATABASE=dev -d mysql:latest
```
2. 下記設定ファイルの作成（値は例）

```shell
$ cat .env
DB_DRIVER=mysql
DB_USER=dev
DB_PASS=dev
DB_NAME=dev
DB_HOST=localhost
DB_PORT=3307
```

3. `go run main.go` 実行

4. ターミナル別タブなどで、`Users` テーブルが作成されたことを確認
```
$ mysql -h 127.0.0.1 -P 3307 -u dev -pdev dev -e "show tables;"
mysql: [Warning] Using a password on the command line interface can be insecure.
+---------------+
| Tables_in_dev |
+---------------+
| users         |
+---------------+
```


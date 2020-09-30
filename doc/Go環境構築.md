# 起動前のGo環境構築

## linux

### Goダウンロードとパスを通す

[公式の説明](https://golang.org/doc/install)

1. ダウンロード後/usr/localに入れる

    例:

    ```bash
    curl -OL  https://golang.org/dl/go1.15.2.linux-amd64.tar.gz
    sudo tar -xf go1.15.2.linux-amd64.tar.gz  -C /usr/local/
    ```

2. PATH環境変数に追加

    vi ~/.bashrc

    ```bash
    # golang
    export PATH=$PATH:/usr/local/go/bin
    ```

3. PATHの更新
    ```source ~/.bashrc```

4. バージョン確認
    `go version`
    出力例: go version go1.15.2 linux/amd64

### VSCODEの設定

1. vscodeいれる

2. Goの拡張機能をインストール
    発行元はgolang.go

3. 適当な開発フォルダを作り移動する

    例:

    ```bash
    mkdir dev
    cd dev
    ```

4. vscodeでフォルダを開く

## GOMODULEを使えるようにする

[GOMODULE--Goのパッケージ管理](https://qiita.com/Syoitu/items/f221b52231703cebe8ff)
Go1.11バージョン以降追加された機能

1. goの環境変数を確認
    `go env`

2. go modulesをonにする
    `go env -w GO111MODULE=on`

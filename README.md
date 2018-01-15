# go-backup

ファイルシステムのバックアップ。

「Go言語によるWebアプリケーション開発」の写経。

## コマンドラインツール

パスの追加:  
`backup -db=/データベースへのパス add {パス} [パス...]`

パスの削除:  
`backup -db=/データベースへのパス remove {パス} [パス...]`

すべてのパスの一覧表示:  
`backup -db=/データベースへのパス list`

定期的にバックアップ
`backupd -db=./backupdata/ -archive=./archive -interval=5`

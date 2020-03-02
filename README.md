# GRPCのCRUDデモ

## 環境設定
### Protobufのインストール
```
brew install protobuf 
```

### ライブラリのインストール
```
go get -u google.golang.org/grpc
go get -u github.com/golang/protobuf/protoc-gen-go
```

## 生成
```
protoc server/proto/user.proto --go_out=plugins=grpc:.
```

## サーバーの起動
```
go run server/server.go 
```

## クライアントの実行
```
go run client/client.go
```

## ディレクトリ構成
```
├── client
│   └── client.go
├── proto
│   ├── user.pb.go
│   └── user.proto
├── server
│   ├── config
│   │   └── config.go
│   ├── config.ini
│   ├── logs
│   │   └── server.log
│   ├── model
│   │   └── userItem.go
│   ├── repository
│   │   └── userRepository.go
│   ├── server.go
│   └── utils
│       └── logging.go
```

# gRPCのCRUDサンプル

## 環境構築
### Homebrew経由でProtocをインストール
```
brew install protobuf 
```

### ライブラリのインストール
```
go get google.golang.org/grpc
go get github.com/golang/protobuf/protoc-gen-go
go get go.mongodb.org/mongo-driver/mongo
```

### .protoファイルからコードを生成
```
protoc server/proto/user.proto --go_out=plugins=grpc:.
```

### サーバーの起動
```
go run server/server.go 
```

### クライアントの実行
```
go run client/client.go
```

### ディレクトリ構成
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

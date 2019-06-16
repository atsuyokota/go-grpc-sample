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
protoc unary/greetpb/greetpb.proto --go_out=plugins=grpc:.
```

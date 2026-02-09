# gRPC セットアップガイド

このプロジェクトはRESTful APIとgRPCの両方をサポートしています。

## 必要なツール

### 1. Protocol Buffers コンパイラ (protoc) のインストール

#### Windows
```powershell
# Chocolateyを使用する場合
choco install protoc

# または手動でダウンロード
# https://github.com/protocolbuffers/protobuf/releases
# protoc-xx.x-win64.zip をダウンロードして解凍
# binフォルダをPATHに追加
```

#### macOS
```bash
brew install protobuf
```

#### Linux
```bash
sudo apt install -y protobuf-compiler
```

### 2. Go用のprotocプラグインのインストール

```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

PATHに追加:
```bash
export PATH="$PATH:$(go env GOPATH)/bin"
```

## protoファイルからコード生成

### Windows (PowerShell)
```powershell
cd backend
.\generate.ps1
```

### Linux/macOS
```bash
cd backend
chmod +x generate.sh
./generate.sh
```

## gRPCサーバーの実装

`backend/proto/kakeibo.proto` にサービス定義があります。

生成されたコードは `backend/proto/pb/` に配置されます。

### サーバー実装例

```go
// backend/grpc/server.go を作成
package grpc

import (
    "context"
    pb "kakeibo-backend/proto/pb"
    "kakeibo-backend/models"
    "gorm.io/gorm"
)

type UserServiceServer struct {
    pb.UnimplementedUserServiceServer
    DB *gorm.DB
}

func (s *UserServiceServer) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.UserResponse, error) {
    user := &models.User{
        Name:        req.Name,
        RealName:    req.RealName,
        Email:       req.Email,
        Password:    req.Password,
        Icon:        req.Icon,
        ProfileMemo: req.ProfileMemo,
    }
    
    if err := s.DB.Create(user).Error; err != nil {
        return nil, err
    }
    
    return &pb.UserResponse{
        User: &pb.User{
            Id:          uint32(user.ID),
            Name:        user.Name,
            RealName:    user.RealName,
            Email:       user.Email,
            Icon:        user.Icon,
            ProfileMemo: user.ProfileMemo,
        },
    }, nil
}
```

### main.goでgRPCサーバーを起動

```go
import (
    "google.golang.org/grpc"
    pb "kakeibo-backend/proto/pb"
    grpcserver "kakeibo-backend/grpc"
)

// gRPCサーバーを別ポートで起動
go func() {
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    
    grpcServer := grpc.NewServer()
    pb.RegisterUserServiceServer(grpcServer, &grpcserver.UserServiceServer{DB: db})
    pb.RegisterCategoryServiceServer(grpcServer, &grpcserver.CategoryServiceServer{DB: db})
    pb.RegisterExpenseServiceServer(grpcServer, &grpcserver.ExpenseServiceServer{DB: db})
    pb.RegisterSubscriptionServiceServer(grpcServer, &grpcserver.SubscriptionServiceServer{DB: db})
    
    log.Println("gRPC server listening on :50051")
    if err := grpcServer.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}()
```

## フロントエンドでgRPCを使用

### gRPC-Webの設定

```bash
cd frontend
npm install @grpc/grpc-js @grpc/proto-loader
```

### Envoyプロキシの設定 (gRPC-Web用)

gRPC-Webを使用する場合、Envoyプロキシが必要です。

```yaml
# docker-compose.yml に追加
envoy:
  image: envoyproxy/envoy:v1.28-latest
  ports:
    - "8081:8081"
  volumes:
    - ./envoy.yaml:/etc/envoy/envoy.yaml
```

## 現在の状態

- ✅ protoファイル作成済み (`backend/proto/kakeibo.proto`)
- ✅ RESTful API実装済み
- ⏳ protocのインストールが必要
- ⏳ gRPCサーバー実装が必要
- ⏳ フロントエンドのgRPCクライアント実装が必要

## 次のステップ

1. `protoc` をインストール
2. `generate.ps1` を実行してGoコードを生成
3. gRPCサーバーを実装
4. フロントエンドでgRPC-Webクライアントを実装

現時点では、RESTful APIが完全に動作しているため、そちらを使用できます。

# Go grpc moon Repo

### Setup

To set up this project run these commands sequentially

```shell
  proto install go --pin
```

```shell
  proto plugin search moon
```

```shell
  proto plugin add <name> <url>
```

```shell
  proto plugin add moon https://raw.githubusercontent.com/moonrepo/moon/master/proto-plugin.toml
  proto plugin add protoc  https://raw.githubusercontent.com/b4nst/proto-plugins/refs/heads/main/toml/protoc.toml 
```

```shell
  moon init    
```

```shell
  proto run protoc -- --go-grpc_out=apps/gateway/pb/gen --go_out=apps/gateway/pb/gen  apps/gateway/pb/chat.proto
```

```shell
  go mod init https://github.com/rmmbdev/demo-go-grpc-moon.git 
```

```shell
  go mod tidy
```
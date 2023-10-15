# 生成protobuf 生成到上层目录 micro_out, go_out 两个都得生成 package service
#protoc --proto_path=. --micro_out=../ --go_out=../ *.proto
# 这边遇到的问题生成的micro.go 需要加上/v2
protoc --proto_path=. --micro_out=./ --go_out=./ *.proto
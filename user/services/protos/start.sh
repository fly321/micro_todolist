# 生成protobuf 生成到上层目录 micro_out, go_out 两个都得生成
protoc --proto_path=. --micro_out=../ --go_out=../ *.proto
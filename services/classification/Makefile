.PHONY: proto

proto:
	poetry run python3 -m grpc_tools.protoc -I=./proto --python_out=./proto --pyi_out=./proto --grpc_python_out=./proto ./proto/classification.proto
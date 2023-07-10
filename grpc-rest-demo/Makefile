build_go:
	protoc --go_out="./compiled-go" --go-grpc_out="./compiled-go" ./services/*.proto

build_ts:
	grpc_tools_node_protoc \
		--js_out=import_style=commonjs,binary:./compiled-ts \
		--grpc_out=grpc_js:./compiled-ts \
		--plugin=protoc-gen-grpc=`which grpc_tools_node_protoc_plugin` \
		-I ./services \
		./services/*.proto
	protoc \
		--plugin=protoc-gen-ts=./node_modules/.bin/protoc-gen-ts \
		--ts_out=grpc_js:./compiled-ts \
		-I ./services \
		./services/*.proto

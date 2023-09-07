gen_proto_ruby:
	mkdir -p lib/grpc
	gem install grpc grpc-tools
	grpc_tools_ruby_protoc -I ./../../proto --ruby_out=lib/grpc --grpc_out=lib/grpc ./../../proto/*.proto
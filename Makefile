gen:
	@protoc \
		--proto_path=/Users/mrrishi/Documents/Goo/coffeeshop \
		coffeeshop.proto \
		--go_out=proto_files --go_opt=paths=source_relative \
		--go-grpc_out=proto_files --go-grpc_opt=paths=source_relative

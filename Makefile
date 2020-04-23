proto:
	cd $(shell pwd)/api/orders/models && protoc --go_out=:../ ./orders.proto	
	cd $(shell pwd)/api/products/models && protoc --go_out=:../ ./products.proto
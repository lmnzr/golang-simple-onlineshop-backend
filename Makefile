# Build And Development
clean:
	@ rm -f *.out || rm -f *.html || rm -f *.exe 
test:
	@ go test simpleshop/test/... 
test-cover:
	@ go test simpleshop/test/... -coverpkg=./... -coverprofile=cover.txt
	@ go tool cover -html=cover.txt -o cover.html   
run:
	@ go build && ./simpleshop	                                                                                 
.PHONY: run test test-cover clean
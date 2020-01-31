packagename = simpleshop
# Build And Development
init:
	@ go mod init $(packagename)
	@ go mod vendor 
clean:
	@ sudo rm -f $(packagename).bin $(packagename).exe cover.txt cover.html cover.out
test:
	@ go test $(packagename)/test/... 
test-cover:
	@ go test $(packagename)/test/... -coverpkg=./... -coverprofile=cover.out
	@ go tool cover -html=cover.out -o cover.html   
run:
	@ go build -o $(packagename).bin && ./$(packagename).bin	                                                                                 
.PHONY: init run test test-cover clean
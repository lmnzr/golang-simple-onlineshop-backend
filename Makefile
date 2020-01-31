packagename = simpleshop
# Build And Development
clean:
	@ sudo rm -f $(packagename).bin $(packagename).exe cover.txt cover.html cover.out
test:
	@ go test $(packagename)/test/... 
test-cover:
	@ go test $(packagename)/test/... -coverpkg=./... -coverprofile=cover.txt
	@ go tool cover -html=cover.txt -o cover.html   
run:
	@ go build -o $(packagename).bin && ./$(packagename).bin	                                                                                 
.PHONY: run test test-cover clean
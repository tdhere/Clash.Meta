CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o bin/linux/clash ./
 CGO_ENABLED=0 GOOS=linux GOARCH=arm go build -ldflags "-s -w" -o bin/arm/clash ./
#CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags "-s -w" -o bin/win/clash.exe ./
#CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -ldflags "-s -w" -o bin/mac/clash ./
upx bin/*/*
ls -hla bin/*/*
cp -rf bin/* /d/mycloud/dl
build :
	go mod -C ./wasm/sisyphus  tidy
	GOOS=js GOARCH=wasm go build -C ./wasm/perlinwalking -o ../../deployment/perlinwalking/app.wasm .
	go mod -C ./wasm/sisyphus tidy
	GOOS=js GOARCH=wasm go build -C ./wasm/sisyphus -o ../../deployment/sisyphus/app.wasm .
	go mod -C ./wasm/fussball tidy
	GOOS=js GOARCH=wasm go build -C ./wasm/fussball -o ../../deployment/fussball/app.wasm .

push :
	git add .
	git commit -a
	git push
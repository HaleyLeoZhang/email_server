all: debug

debug:
	@clear
	@echo "App debug is loading"
	@go run ./build/main.go -conf=./build/app.yaml

ini:
	@echo "Copy conf/app.ini To /usr/local/etc/mail.ops.hlzblog.top.ini ---DOING"
	@cp conf/app.ini /usr/local/etc/mail.ops.hlzblog.top.ini
	@echo "Copy ---DONE"

compile:
	@rm -rf email_server
	@echo "App is creating. Please wait ..."
	@go build -o email_server -v ./build/main.go
	@echo "App is created"

run:
	@./email_server -conf=./build/app.yaml

deploy:
	@make ini
	@make compile
	@make run

tool:
	@go vet ./...; true
	@gofmt -w .

json:
    # 以下会为 model/ 所有二级目录下所有结构体生成 easyjson 文件
	@echo "Creating easyjson file for model/*/*.go"
	@rm -rf model/*/*_easyjson.go
	@easyjson -all model/*/*.go # 格式化 json ，需要 https://github.com/mailru/easyjson
	@echo "Created model/*/*.go easyjson file Success"
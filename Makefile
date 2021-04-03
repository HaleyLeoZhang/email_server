all: debugapi

# 调试
debugapi:
	@clear
	@echo "API debug is loading"
	@go run ./api/build/main.go -conf=./api/build/app.yaml

debugjob:
	@clear
	@echo "JOB debug is loading"
	@go run ./job/build/main.go -conf=./job/build/app.yaml

# 初始化配置文件
iniapi:
	@echo "Copy api/build/app.example.yaml To api/build/app.yaml ---DOING"
	@cp api/build/app.example.yaml api/build/app.yaml
	@echo "Copy ---DONE"

inijob:
	@echo "Copy job/build/app.example.yaml To job/build/app.yaml ---DOING"
	@cp job/build/app.example.yaml job/build/app.yaml
	@echo "Copy ---DONE"

# 编译项目
capi:
	@rm -rf app_api
	@echo "API is creating. Please wait ..."
	@go build -o app_api -v ./api/build/main.go
	@echo "API is created"
	
cjob:
	@rm -rf app_job
	@echo "JOB is creating. Please wait ..."
	@go build -o app_job -v ./job/build/main.go
	@echo "JOB is created"

# 运行项目
rapi:
	@./app_api -conf=./api/build/app.yaml

rjob:
	@./app_job -conf=./job/build/app.yaml

# 编译并运行项目
deployapi:
	@clear
	@make -s capi
	@make -s rapi

deployjob:
	@clear
	@make -s cjob
	@make -s rjob

# 一键格式化代码
tool:
	@clear
	@echo "Using gofmt..."
	@go vet ./...; true
	@gofmt -w .

# 一键生成 easyjson
json:
	@echo "Creating easyjson file for common/model/*/*.go"
	@rm -rf common/model/*/*_easyjson.go
	@easyjson -all common/model/*/*.go # 格式化 json ，需要 https://github.com/mailru/easyjson
	@echo "Created common/model/*/*.go easyjson file Success"
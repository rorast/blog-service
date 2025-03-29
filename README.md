# blog-service
A Go http blog service than using gin grom  viper lumberjack 

## 1. 在 GitHub 建立新專案
-   進入 GitHub → 點選 New repository（新建倉庫）。
-   填寫專案名稱（如 blog-service），選擇 Public 或 Private。
-   初始化倉庫（建議勾選 Add a README file）。
-   點擊 Create repository，完成建立。

## 2. 在 GoLand 拉取 GitHub 專案
-   打開 GoLand，點選 File → New → Project from Version Control (在開始面板 Get for VCS)。
-   選擇 GitHub，然後輸入你的 GitHub 倉庫 URL（例如 https://github.com/rorast/blog-service.git）。
-   選擇本地存放路徑（如 ~/go/src/blog-service）。
-   點擊 Clone，拉取 GitHub 專案到本地。

## 3. 初始化 Go 專案
-    在 GoLand 中打開剛剛拉取的專案。
-    點選 File → New → Go Module，輸入 Module name（如 blog-service）。
-    點擊 OK，完成 Go 專案初始化。
-   go mod init github.com/rorast/blog-service
-   go mod tidy
## 4. 安裝 Gin
-   Gin 是一個用於構建 Web 應用程序和微服務的 Go 框架。在 GoLand 中安裝 Gin：
-   go get -u github.com/gin-gonic/gin@v1.6.3
-   go mod tidy (更新 go.mod 和 go.sum)
## 5. 提交並推送到 GitHub
-   初始化 Git（如果未自動設定）：
-   git init
-   git branch -M main
-   git remote add origin https://github.com/rorast/blog-service.git
-   提交變更：
-   git add .
-   git commit -m "Initial commit"
-   git push -u origin main
## 最適合的流程 ✅
### 1️⃣ 先在 GitHub 建立倉庫 → 2️⃣ 從 GoLand 拉取 → 3️⃣ 初始化 Go 專案 → 4️⃣ 提交並推送。這樣可以確保本地與遠端同步，並且不會有 Git Remote 設置錯誤的問題。🎯

## 6. Go 目錄結構
### blog-service/
- │── cmd/                   # 應用程式入口點
- │   ├── server/            # 主 Web 服務入口
- │   │   ├── main.go
- │   │   ├── config.yaml    # 設定檔案
- │   │   ├── .env           # 環境變數
- │   ├── migrate/           # 數據庫遷移工具入口
- │   │   ├── main.go
- │
- │── configs/               # 設定檔案 (如 database, JWT, Redis, etc.)
- │── docs/                  # API 文件與 Swagger 規範
- │── internal/              # 內部應用邏輯
- │   ├── dao/               # 資料庫訪問層 (Data Access Object)
- │   ├── middleware/        # HTTP 中間件 (如身份驗證、日誌)
- │   ├── model/             # 資料結構與 ORM 模型
- │   ├── routers/           # 路由與控制器
- │   ├── service/           # 核心業務邏輯
- │   ├── repository/        # 資料儲存層，封裝 DAO 的操作
- │   ├── handler/           # 用於處理 API 請求的控制器
- │
- │── pkg/                   # 可重用的公用函式庫
- │── storage/               # 預設存儲 (檔案上傳、臨時存放)
- │── scripts/               # 构建、安装、分析等脚本 (如 CI/CD)
- │── third_party/           # 第三方依賴 (Swagger UI, JWT, Logger)
- │── tests/                 # 測試檔案 (單元測試 & 整合測試)
- │── go.mod                 # 依賴管理
- │── go.sum                 # 依賴鎖定檔案
- │── README.md              # 項目說明文件
- │── Makefile               # 常用指令 (建置、測試、自動化流程)

## 7.安裝 viper
-   viper 是一個 Go 語言的設定管理庫，支持多種格式的設定文件，如 JSON、TOML、YAML、HCL、envfile 和 Java properties config files。
-   安裝 viper：
-   go get -u github.com/spf13/viper@v1.4.0
-   go mod tidy
   
## 8. 安裝 GORM (有 v1 跟 v2 版本)
-   GORM 是一個 Go 語言的 ORM 框架，支持 MySQL、PostgreSQL、SQLite 和 SQL Server 等多種數據庫。
-   安裝 GORM v1：
-   go get -u github.com/jinzhu/gorm@v1.9.12
-   go get -u github.com/jinzhu/gorm/dialects/mysql@v1.9.16
-   go mod tidy
-   安裝 GORM v2：
-   go get -u gorm.io/gorm@v1.21.12
-   go get -u gorm.io/driver/mysql@v1.5.7
-   go mod tidy
   
## 9. 安裝 otelgorm (先用舊版為otgorm)
-   otgorm 是一個 OpenTelemetry 的 GORM 插件，用於收集數據庫操作的性能指標。
-   安裝 otgorm：
-   go get -u github.com/eddycjy/opentracing-gorm
-   go mod tidy

-   otelgorm 是一個 OpenTelemetry 的 新版 GORM 插件，用於收集數據庫操作的性能指標。
-   安裝 otelgorm：
-   go get -u go.opentelemetry.io/contrib/instrumentation/gorm.io/gorm/otelgorm@v1.2.0  (指定較舊的穩定版本)
-   go mod tidy (要等安裝完成，才能重整 mod 進行 import)
   
## 10. 安裝 Lumberjack
-   Lumberjack 是一個 Go 語言的日誌庫，支持日誌切割、壓縮和清理。
-   安裝 Lumberjack：
-   go get -u gopkg.in/natefinch/lumberjack.v2
-   go mod tidy
   
## 11. 安裝 Swagger
-   Swagger 是一個用於設計、構建、文件化和消費 RESTful API 的工具。
-   安裝 Swagger：
-   go get -u github.com/swaggo/swag/cmd/swag@v1.6.5
-   go get -u github.com/swaggo/gin-swagger@v1.2.0
-   go get -u github.com/swaggo/files
-   go get -u github.com/alecthomas/template
-   go mod tidy
   
## 12. 安裝 Validator
-   Validator 是一個 Go 語言的驗證庫，支持結構體驗證、自定義驗證規則和多語言錯誤信息。
-   安裝 Validator：
-   go get -u github.com/go-playground/validator/v10
-   go mod tidy
   
## 13. 測試檔案上傳方式
-   透過 Postman 測試檔案上傳方式：
-   命令列測試 : curl -X POST http://127.0.0.1:8000/upload/file -F file=@{file_path} -F type=1
-   開啟 Postman → 選擇 POST 請求方式 → 輸入 URL（如 http://localhost:8000/upload/file）。
-   點擊 Body → 選擇 form-data → 輸入 Key 為 file  → 選擇 File（如 test.jpg），並點選上傳到 Postman。
-   點擊 Body → 選擇 form-data → 輸入 Key 為 type → 設值為 1。
-   點擊 Send，測試檔案上傳功能。

## 14. API 訪問控制  安裝 JWT
-   JWT 是一種用於簽署 JSON 格式的 Web Token 的標準。
-   安裝 JWT：
-   go get -u github.com/dgrijalva/jwt-go@v3.2.0
-   go mod tidy 

## 15. 安裝 Gomail (寄送郵件)
-   Gomail 是一個 Go 語言的郵件庫，支持發送郵件。
-   安裝 Gomail：
-   go get -u gopkg.in/gomail.v2
-   go mod tidy

## 16. 安裝 限流器 ratelimit
-   ratelimit 是一個 Go 語言的限流器庫，支持基於內存和 Redis 的限流。
-   安裝 ratelimit：
-   go get -u github.com/juju/ratelimit@v1.0.1
-   go mod tidy

## 17. 安裝 鏈路追蹤 opentracing jaeger-client-go
-   opentracing 是一個 Go 語言的鏈路追蹤庫，支持 Jaeger、Zipkin 和 LightStep 等追蹤系統。
-  安裝 opentracing：
- go get -u github.com/opentracing/opentracing-go@v1.2.0
- go get -u github.com/uber/jaeger-client-go@v2.26.0
- go mod tidy

## 18. 實現 SQL 追蹤
   實現 SQL 追蹤，需要安裝 opentracing 和 gorm。
-  安裝 opentracing-gorm:
-  go get -u github.com/eddycjy/opentracing-gorm
-  go mod tidy

## 19. 安裝 go-bindata (資源打包) 可將數據文件轉換為 Go 代碼 (正常go編譯是不會將數據文件打包進去的)
-   go-bindata 是一個 Go 語言的資源打包工具，支持將數據文件轉換為 Go 代碼。
-  安裝 go-bindata：
- go get -u github.com/go-bindata/go-bindata/...
- go mod tidy
- 執行 go-bindata -o config/config.go -pkg=configs configs/config.yaml (將 configs/ 目錄下的所有文件打包到 data.go 中)
- 執行以上程式後通過代碼 b, _ := Asset("config/config.yaml") 可以獲取到文件內容。

## 20. 安裝 fsnotify (文件變更監控) 跨平台文件系統監聽事件庫
-   fsnotify 是一個 Go 語言的文件變更監控庫，支持跨平台文件系統監聽事件。
- 安裝 fsnotify：
- go get -u golang.org/x/sys/...
- go get -u github.com/fsnotify/fsnotify
- go mod tidy

## 21. 安裝 UPX (可執行文件壓縮) 
-   UPX 是一個可執行文件壓縮工具，支持多種操作系統。
- 安裝 UPX：https://github.com/upx/upx
- 下載 UPX：https://upx.github.io/
- 使用 : upx -9 -o blog-service blog-service (壓縮 blog-service 可執行文件)
  - 簡單使用 : upx blog-service (已編譯出執行檔)

## 22. 編譯附帶版本信息
-   透過 go build 編譯時附帶版本信息：
-  go build -ldflags "-X main.Version=1.0.0 -X 'main.BuildTime=`date`' -X 'main.GoVersion=`go version`'"
-  使用 : go build -ldflags "-X main.appName=Go Blog API" -o blog-service main.go
-  程式中需加入變數 : 
   - var (
   -   buildTime    string
   -   buildVersion string
   -   gitCommitID  string
   - )

   - if isVersion {
   -   fmt.Printf("build time: %s\n", buildTime)
   -   fmt.Printf("build version: %s\n", buildVersion)
   -   fmt.Printf("build git commit: %s\n", gitCommitID)
   -   os.Exit(0)
   - }
- windows PowerShell 下:
  -  $buildTime = Get-Date -Format "yyyy-MM-dd HH:mm:ss"
  -  $gitCommitID = git rev-parse HEAD
  -  go build -ldflags "-X 'main.buildTime=$buildTime' -X 'main.buildVersion=1.0.0' -X 'main.gitCommitID=$gitCommitID'"
- 執行 : ./blog-service --version (查看版本信息)
  - build time: 2025-03-01 13:25:12
  - build version: 1.0.0
  - build git commit: 5914afc2ea2a8bfaa7fdd3113fd07f1ba8060470
  
## 23. 安裝 git-cz
- 安裝 git-cz：npm install -g git-cz
- 使用 : 
git add . => git commit -m 'first' => git push
改成
git add . => git cz => git push
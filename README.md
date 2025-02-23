# blog-service
A Go http blog service than using gin grom  viper lumberjack 

## 1. 在 GitHub 建立新專案
   進入 GitHub → 點選 New repository（新建倉庫）。
   填寫專案名稱（如 blog-service），選擇 Public 或 Private。
   初始化倉庫（建議勾選 Add a README file）。
   點擊 Create repository，完成建立。

## 2. 在 GoLand 拉取 GitHub 專案
   打開 GoLand，點選 File → New → Project from Version Control (在開始面板 Get for VCS)。
   選擇 GitHub，然後輸入你的 GitHub 倉庫 URL（例如 https://github.com/rorast/blog-service.git）。
   選擇本地存放路徑（如 ~/go/src/blog-service）。
   點擊 Clone，拉取 GitHub 專案到本地。

## 3. 初始化 Go 專案
    在 GoLand 中打開剛剛拉取的專案。
    點選 File → New → Go Module，輸入 Module name（如 blog-service）。
    點擊 OK，完成 Go 專案初始化。
   go mod init github.com/rorast/blog-service
   go mod tidy
## 4. 安裝 Gin
   Gin 是一個用於構建 Web 應用程序和微服務的 Go 框架。在 GoLand 中安裝 Gin：
   go get -u github.com/gin-gonic/gin@v1.6.3
   go mod tidy (更新 go.mod 和 go.sum)
## 5. 提交並推送到 GitHub
   初始化 Git（如果未自動設定）：
   git init
   git branch -M main
   git remote add origin https://github.com/rorast/blog-service.git
   提交變更：
   git add .
   git commit -m "Initial commit"
   git push -u origin main
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


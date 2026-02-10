# サンプルデータ作成スクリプト
Write-Host "サンプルデータを作成します..." -ForegroundColor Green

# 環境変数を設定（Dockerコンテナ内のDBに接続）
$env:DB_HOST = "localhost"
$env:DB_USER = "postgres"
$env:DB_PASSWORD = "postgres"
$env:DB_NAME = "kakeibo"
$env:DB_PORT = "5433"

# シードスクリプトを実行
go run cmd/seed/main.go

if ($LASTEXITCODE -eq 0) {
    Write-Host "`nサンプルデータの作成が完了しました！" -ForegroundColor Green
    Write-Host "以下のコマンドでユーザー一覧を確認できます:" -ForegroundColor Cyan
    Write-Host "  curl http://localhost:8080/api/users" -ForegroundColor Yellow
} else {
    Write-Host "`nエラーが発生しました。" -ForegroundColor Red
}

package todo

// 定義錯誤碼常量
const (
	Error6000 = "error6000" // 無法綁定 JSON 請求體
	Error6001 = "error6001" // 任務內容不能為空
	Error6002 = "error6002" // 創建 Todo 項目失敗
	Error6003 = "error6003" // 獲取 Todo 項目失敗
	Error6004 = "error6004" // 更新 Todo 項目失敗
	Error6005 = "error6005" // Todo 項目未找到
	Error6006 = "error6006" // 刪除 Todo 項目失敗
	// 根據需求添加更多錯誤碼
)

// 定義錯誤碼對應的錯誤訊息
var ErrorMessages = map[string]string{
	Error6000: "無法綁定 JSON 請求體",
	Error6001: "任務內容不能為空",
	Error6002: "創建 Todo 項目失敗",
	Error6003: "獲取 Todo 項目失敗",
	Error6004: "更新 Todo 項目失敗",
	Error6005: "Todo 項目未找到",
	Error6006: "刪除 Todo 項目失敗",
}

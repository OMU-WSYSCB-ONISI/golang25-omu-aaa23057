package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	// /info にアクセスが来たときに infoHandler を呼び出す
	http.HandleFunc("/info", infoHandler)

	// ポート8080でサーバを起動
	fmt.Println("Server is running on http://localhost:8080/info")
	http.ListenAndServe(":8080", nil)
}

// /info にアクセスがあったときに実行される関数
func infoHandler(w http.ResponseWriter, r *http.Request) {
	// JST(日本時間)を取得
	jst, _ := time.LoadLocation("Asia/Tokyo")
	now := time.Now().In(jst).Format("2006年01月02日 15:04:05")

	// ブラウザ情報(User-Agent)を取得
	ua := r.Header.Get("User-Agent")

	// クライアントにメッセージを返す
	fmt.Fprintf(w, "今の時刻は%sで，利用しているブラウザは「%s」，ですね。", now, ua)
}

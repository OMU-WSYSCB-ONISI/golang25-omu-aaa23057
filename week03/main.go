package main

import (
    "fmt"
    "math/rand"
    "net/http"
    "time"
)

func main() {
    // 乱数のタネを現在時刻で初期化
    rand.Seed(time.Now().UnixNano())

    // /webfortune にアクセスが来たら fortuneHandler を呼ぶ
    http.HandleFunc("/webfortune", fortuneHandler)

    // ポート8080でWebサーバを起動
    http.ListenAndServe(":8080", nil)
}

// おみくじ用ハンドラ
func fortuneHandler(w http.ResponseWriter, r *http.Request) {
    fortunes := []string{"大吉", "中吉", "吉", "凶"}

    // 0〜len(fortunes)-1 の乱数
    n := rand.Intn(len(fortunes))

    // ブラウザに表示する文言
    fmt.Fprintf(w, "今の運勢は%sです。\n", fortunes[n])
}

package main

import (
	"encoding/json"
	"fmt"
	"html"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

const logFile = "public/logs.json"

type Log struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Body  string `json:"body"`
	Image string `json:"image"` // 追加: 画像URL（例: /uploads/xxx.png）
	CTime int64  `json:"ctime"`
}

func main() {
	fmt.Printf("Go version: %s\n", runtime.Version())

	// public/ 以下を配信（uploads もここで見える）
	http.Handle("/", http.FileServer(http.Dir("public/")))

	http.HandleFunc("/hello", hellohandler)
	http.HandleFunc("/bbs", showHandler)
	http.HandleFunc("/write", writeHandler)

	fmt.Println("Launch server...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Failed to launch server: %v", err)
	}
}

func hellohandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "こんにちは from Codespace !")
}

// 表示（新着順 + 改行保持 + 検索）
func showHandler(w http.ResponseWriter, r *http.Request) {
	logs := loadLogs()

	// 検索（q=）
	q := strings.TrimSpace(r.URL.Query().Get("q"))
	qLower := strings.ToLower(q)

	style := "<style>" +
		"body { font-family: sans-serif; margin: 24px; }" +
		"p { border: 1px solid silver; padding: 1em; border-radius: 8px; }" +
		"span { background-color: #eef; padding: 2px 6px; border-radius: 6px; }" +
		".body { white-space: pre-wrap; line-height: 1.6; margin-top: 8px; }" + // 改行保持
		".meta { color: #666; font-size: 12px; margin-top: 8px; }" +
		".img { margin-top: 8px; }" +
		".img img { max-width: 360px; height: auto; border: 1px solid #ddd; border-radius: 8px; }" +
		".box { max-width: 900px; margin: 0 auto; }" +
		"</style>"

	htmlLog := ""

	// 新着順（後ろから）
	for idx := len(logs) - 1; idx >= 0; idx-- {
		i := logs[idx]

		// 検索フィルタ
		if q != "" {
			target := strings.ToLower(i.Name + " " + i.Body)
			if !strings.Contains(target, qLower) {
				continue
			}
		}

		imgHTML := ""
		if i.Image != "" {
			imgHTML = fmt.Sprintf(
				"<div class='img'><img src='%s' alt='posted image'></div>",
				html.EscapeString(i.Image),
			)
		}

		htmlLog += fmt.Sprintf(
			"<p>(%d) <span>%s</span>"+
				"<div class='body'>%s</div>%s"+
				"<div class='meta'>%s</div></p>",
			i.ID,
			html.EscapeString(i.Name),
			html.EscapeString(i.Body),
			imgHTML,
			time.Unix(i.CTime, 0).Format("2006/01/02 15:04"),
		)
	}

	htmlBody := "<html><head><meta charset='utf-8'>" + style +
		"</head><body><div class='box'><h1>BBS</h1>" +
		getSearchForm(q) +
		getForm() +
		"<hr>" + htmlLog +
		"</div></body></html>"

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write([]byte(htmlBody))
}

// 書き込み（画像対応）
func writeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/bbs", 302)
		return
	}

	// 最大 5MB（課題なら十分）
	r.Body = http.MaxBytesReader(w, r.Body, 5<<20)
	if err := r.ParseMultipartForm(5 << 20); err != nil {
		http.Error(w, "フォームが不正か、ファイルが大きすぎます（最大5MB）", 400)
		return
	}

	name := strings.TrimSpace(r.FormValue("name"))
	body := strings.TrimSpace(r.FormValue("body"))
	if name == "" {
		name = "名無し"
	}

	// 画像保存（任意）
	var imgPath string
	file, header, err := r.FormFile("image")
	if err == nil {
		defer file.Close()

		ext := strings.ToLower(filepath.Ext(header.Filename))
		if ext != ".png" && ext != ".jpg" && ext != ".jpeg" && ext != ".gif" && ext != ".webp" {
			http.Error(w, "画像は png/jpg/jpeg/gif/webp のみ対応です", 400)
			return
		}

		_ = os.MkdirAll("public/uploads", 0755)

		saveName := fmt.Sprintf("%d%s", time.Now().UnixNano(), ext)
		saveFile := filepath.Join("public/uploads", saveName)

		out, err := os.Create(saveFile)
		if err != nil {
			http.Error(w, "画像の保存に失敗しました", 500)
			return
		}
		defer out.Close()

		if _, err := io.Copy(out, file); err != nil {
			http.Error(w, "画像の保存に失敗しました", 500)
			return
		}

		imgPath = "/uploads/" + saveName
	}

	// ログ保存
	logs := loadLogs()
	log := Log{
		ID:    len(logs) + 1,
		Name:  name,
		Body:  body,
		Image: imgPath,
		CTime: time.Now().Unix(),
	}
	logs = append(logs, log)
	saveLogs(logs)

	http.Redirect(w, r, "/bbs", 302)
}

// 検索フォーム（q=）
func getSearchForm(q string) string {
	return "<div><form action='/bbs' method='get'>" +
		"検索: <input type='text' name='q' value='" + html.EscapeString(q) + "' placeholder='名前/本文のキーワード' style='width:20em;'>" +
		"<input type='submit' value='検索'>" +
		" <a href='/bbs'>クリア</a>" +
		"</form></div><hr>"
}

// 書き込みフォーム（POST + multipart）
func getForm() string {
	return "<div><form action='/write' method='post' enctype='multipart/form-data'>" +
		"名前: <input type='text' name='name'><br>" +
		"本文:<br><textarea name='body' style='width:30em; height:6em;'></textarea><br>" +
		"画像: <input type='file' name='image' accept='image/*'><br>" +
		"<input type='submit' value='書込'>" +
		"</form></div>"
}

// ログ読み込み
func loadLogs() []Log {
	text, err := os.ReadFile(logFile)
	if err != nil {
		return make([]Log, 0)
	}
	var logs []Log
	_ = json.Unmarshal(text, &logs)
	return logs
}

// ログ保存
func saveLogs(logs []Log) {
	bytes, _ := json.Marshal(logs)
	_ = os.WriteFile(logFile, bytes, 0644)
}

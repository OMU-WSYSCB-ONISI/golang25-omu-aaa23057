package main
import (
	"fmt"
	"net/http"
	"runtime"
  "strconv"
  "strings"
)

func main() {
	fmt.Printf("Go version: %s\n", runtime.Version())

	http.Handle("/", http.FileServer(http.Dir("public/")))
	http.HandleFunc("/hello", hellohandler)
    http.HandleFunc("/enq", enqhandler)
	http.HandleFunc("/fdump", fdump)
    http.HandleFunc("/cal00", cal00handler)
    http.HandleFunc("/cal01", calpmhandler)
    http.HandleFunc("/sum", sumhandler)

//追加ハンドラ
    http.HandleFunc("/cal02", calpmtdhandler)
	http.HandleFunc("/bmi", bmihandler)
    http.HandleFunc("/avg", avghandler)

	fmt.Println("Launch server...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Failed to launch server: %v", err)
	}
}

func hellohandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "こんにちは from Codespace !")
}

func fdump(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Println("errorだよ")
	}
	// フォームはマップとして利用でき以下で内容を確認できる．
	for k, v := range r.Form {
		fmt.Printf("%v : %v\n", k, v)
	}
}

func enqhandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Println("errorだよ")
	}
	// r.FormValue("name")として，フォーム中name欄の値を得る
	fmt.Fprintln(w, r.FormValue("name")+"さん，ご協力ありがとうございます.\n年齢は"+r.FormValue("age")+"で，性別は"+r.FormValue("gend")+"で，出身地は"+r.FormValue("birthplace")+"ですね")
}

func cal00handler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Println("errorだよ")
	}
	price, _ := strconv.Atoi(r.FormValue("price"))
	num, _ := strconv.Atoi(r.FormValue("num"))
	fmt.Fprint(w, "合計金額は ")
	fmt.Fprintln(w, price*num)
}

func calpmhandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Println("errorだよ")
	}
	x, _ := strconv.Atoi(r.FormValue("x"))
	y, _ := strconv.Atoi(r.FormValue("y"))
	switch r.FormValue("cal0") {
	case "+":
		fmt.Fprintln(w, x+y)
	case "-":
		fmt.Fprintln(w, x-y)
	}
}

func sumhandler(w http.ResponseWriter, r *http.Request) {
	var sum, tt int
	if err := r.ParseForm(); err != nil {
		fmt.Println("errorだよ")
	}
	tokuten := strings.Split(r.FormValue("dd"), ",")
	fmt.Println(tokuten)
	for i := range tokuten {
		tt, _ = strconv.Atoi(tokuten[i])
		sum += tt
	}
	fmt.Fprintln(w, sum)
	fmt.Println(sum)
}

// BMI ハンドラ
func bmihandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Println("errorだよ")
	}

	hStr := r.FormValue("height")
	wStr := r.FormValue("weight")

	heightCm, err1 := strconv.Atoi(hStr)
	weight, err2 := strconv.ParseFloat(wStr, 64)

	if err1 != nil || err2 != nil || heightCm <= 0 {
		fmt.Fprintln(w, "身長・体重を正しく入力してください")
		return
	}

	heightM := float64(heightCm) / 100.0
	bmi := weight / (heightM * heightM)

	fmt.Fprintf(w, "身長: %d cm\n体重: %.1f kg\nBMI: %.2f\n", heightCm, weight, bmi)
}

// 計算（加減乗除） cal02
func calpmtdhandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Println("errorだよ")
	}
	x, _ := strconv.Atoi(r.FormValue("x"))
	y, _ := strconv.Atoi(r.FormValue("y"))

	switch r.FormValue("cal0") {
	case "+":
		fmt.Fprintln(w, x+y)
	case "-":
		fmt.Fprintln(w, x-y)
	case "*":
		fmt.Fprintln(w, x*y)
	case "/":
		if y == 0 {
			fmt.Fprintln(w, "0で割ることはできません")
			return
		}
		result := float64(x) / float64(y)
		fmt.Fprintf(w, "%.2f\n", result)
	default:
		fmt.Fprintln(w, "不明な計算です")
	}
}

// 平均と分布のハンドラ
func avghandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Println("errorだよ")
	}

	raw := r.FormValue("scores")
	if raw == "" {
		fmt.Fprintln(w, "得点を入力してください（例: 80,90,65,100）")
		return
	}

	items := strings.Split(raw, ",")
	var sum int
	var count int

	// 0-9, 10-19, ..., 90-100 の 11 区分
	dist := make([]int, 11)

	for _, s := range items {
		s = strings.TrimSpace(s)
		if s == "" {
			continue
		}
		v, err := strconv.Atoi(s)
		if err != nil {
			continue
		}
		if v < 0 || v > 100 {
			continue
		}

		sum += v
		count++

		idx := v / 10
		if v == 100 {
			idx = 10
		}
		dist[idx]++
	}

	if count == 0 {
		fmt.Fprintln(w, "有効な得点がありません")
		return
	}

	avg := float64(sum) / float64(count)
	fmt.Fprintf(w, "人数: %d人\n合計: %d点\n平均: %.2f点\n\n", count, sum, avg)

	labels := []string{
		"0〜9点",
		"10〜19点",
		"20〜29点",
		"30〜39点",
		"40〜49点",
		"50〜59点",
		"60〜69点",
		"70〜79点",
		"80〜89点",
		"90〜99点",
		"100点",
	}

	for i, label := range labels {
		fmt.Fprintf(w, "%s: %d人\n", label, dist[i])
	}
}

# 🚀 Go言語プログラミング - 学生向けクイックスタートガイド

<!-- TOC -->
## 📑 目次

### 🎯 [初心者向け - 必読セクション](#初心者向け---必読セクション)
- [重要な注意事項](#重要な注意事項)
- [初回セットアップ（第1週のみ）](#初回セットアップ第1週のみ約2分)
- [毎週の課題の進め方](#毎週の課題の進め方第2週以降も同じ)
- [よくある質問（FAQ）](#よくある質問faq)

### 📚 [経験者向け - 詳細情報](#経験者向け---詳細情報)
- [便利な機能とショートカット](#便利な機能とショートカット)
- [トラブルシューティング](#トラブルシューティング)
- [高度な使用方法](#高度な使用方法)

### 🔗 [その他の重要なガイド](#その他の重要なガイド)
- [セキュリティガイド（別文書）](#セキュリティガイド別文書)
- [エラーメッセージ集（別文書）](#エラーメッセージ集別文書)

### 📋 [チェックリストと成功のコツ](#チェックリストと成功のコツ)

### 📞 [サポート情報](#サポート情報)

---

## 初心者向け - 必読セクション

> 💡 **プログラミング初心者の方へ**: このセクションを順番に読めば、課題に取り組めます。

---

## 重要な注意事項

### 📋 このリポジトリの使い方

**1つのリポジトリで全週の課題に取り組みます。**

- ✅ 初回セットアップは第1週のみ
- ✅ 第2週以降は同じリポジトリで作業継続
- ✅ Git操作は不要（簡単な提出方法を使用）
- ⚠️ **作業終了後は必ずCodespaceを停止**（最重要）

### ⚠️ 必ず守ること

1. **Codespaceを停止する** - 停止しないと課金され続けます
2. **セキュリティを意識する** - 詳細は **[SECURITY_GUIDE.md](SECURITY_GUIDE.md)** を参照
   - 個人情報を含めない
   - API Keyをコードに書かない
   - ポート公開時は内容を確認
3. **ブランチ名をLMSで報告** - 報告しないと教員が確認できません

---

## 初回セットアップ（第1週のみ）

### ステップ1: リポジトリの作成

1. 教員から配布された課題リンクにアクセス
   ```
   https://classroom.github.com/a/xxxxxxxx
   ```

<img src="images/empty-browser-window-assignment.png" alt="課題リンク" width="600">

2. GitHubにサインインします

<img src="images/Sign-in-to-GitHub-to-continue-to-GitHub-Classroom.png" alt="GitHubサインイン" width="500">

<img src="images/Sign-in-to-GitHub-to-continue-to-GitHub-Classroom-with-id-pass.png" alt="IDとパスワード入力" width="500">

3. 二要素認証が求められた場合は、認証アプリのコードを入力します

<img src="images/Two-factor-authentication-request.png" alt="二要素認証リクエスト" width="500">

<img src="images/Two-factor-authentication-option.png" alt="二要素認証オプション" width="500">

<img src="images/Two-factor-authentication-Enter-the-code-from-your-two-factor-authentication-app.png" alt="認証コード入力" width="500">

4. 「Accept this assignment」をクリック

<img src="images/Accept-the-assignment-websys-golang-xx.png" alt="課題の受け入れ" width="700">

<img src="images/Accept-the-assignment-websys-golang-xx-cont.png" alt="課題受け入れ処理中" width="700">

5. しばらく待つと「Your assignment has been created」と表示される

<img src="images/You-are-ready-to-go-You-accepted-the-assignment.png" alt="課題受け入れ完了" width="700">

6. 表示されたURLをクリック

<img src="images/created-repo.png" alt="作成されたリポジトリ" width="800">

7. **重要**: 表示されたURL を**ブックマーク**してください
   ```
   例: https://github.com/classroom-org/repo-name-yourname
   ```
   この個人専用URLは学期中ずっと使用します！

### ステップ2: Codespacesの起動

1. リポジトリページで緑色の「**Code**」ボタンをクリック

2. 「**Codespaces**」タブを選択

3. 「**Create codespace on main**」をクリック

<img src="images/created-repo-code-create.png" alt="Codespace作成" width="600">

4. しばらく待つ（Codespace構築中）

<img src="images/codesapce-building-0.png" alt="構築開始" width="700">

<img src="images/codesapce-building-1.png" alt="構築中1" width="700">

<img src="images/codesapce-building-2.png" alt="構築中2" width="700">

<img src="images/codesapce-building-3.png" alt="構築中3" width="700">

<img src="images/codesapce-building-4.png" alt="構築中4" width="700">

<img src="images/codesapce-building-5.png" alt="構築中5" width="700">

5. 「Codespace ready!」と表示のあるVS Codeのような画面が表示されたら準備完了 ✅

<img src="images/codesapce-ready.png" alt="Codespace準備完了" width="800">

6. 「Copilot」関連の領域は不要なので閉じること

---

## 毎週の課題の進め方（第2週以降も同じ）

### 【ステップ1】LMSで課題内容を確認

1. LMSにログイン

2. コースを開く

3. 今週の課題ページで以下を確認(例)：
   - ✅ 詳細な課題内容・要件
   - ✅ 入出力例
   - ✅ 締切日時

### 【ステップ2】Codespacesで作業

#### 環境の起動

1. ブックマークしたリポジトリURLにアクセス
<img src="images/created-repo.png" alt="作成されたリポジトリ" width="800">

2. 「Code」→「Codespaces」

3. 既存のCodespaceを「Open in browser」で選択
<img src="images/codespaces-Open-in-Browser2.png"  width="800">
もしくは以下から
<img src="images/codespaces-Open-in-Browser.png"  width="800">


#### プログラムの作成と実行（GUI操作のみ）

1. 左側のエクスプローラーで未更新のフォルダを開く
 
<img src="images/codesapce-open-maingo.png" alt="main.goを開く" width="800">

2. `main.go`ファイルを開いて編集

3. **プログラムの実行（ボタンをクリックするだけ）**
   
   **方法A: Run Codeボタン（推奨）**
   - 画面右上の「▶ Run Code」ボタンをクリック
   - 画面下部のターミナルに結果が表示される

<img src="images/codesapce-run-maingo-0.png" alt="Run Codeボタン" width="800">
<img src="images/codesapce-run-maingo-5.png" alt="プレビュー表示" width="800">

   **方法B: 右クリックメニュー**
   - ファイル上で右クリック
   - 「Run Code」を選択

4A. **プレビューを見る（Webサーバー課題のみ）**

   **注意**: LMSの課題内容で「Webサーバー」「HTTP」「ブラウザで確認」などの指示がある場合のみ実施してください。
   
   a. 実行中であることが確認できれば，以下，画面下部の「ポート」を選択
   - ポート: 実行中のプログラムが待機しているポート
   - 転送されたアドレス
       - アドレス
       - ボタン[コピー，ブラウザーで開く，エディターでのプレビュー]  
   - 実行中のプロセス
   - 表示範囲
   - 配信元
<img src="images/cs-w2-port.png" width="800">

   b. 「エディターでのプレビュー」を選ぶ
<img src="images/cs-w2-sbrw-connect.png" width="800">
<img src="images/ cs-w1-sample.png" width="800">
   
   c. 右クリックメニューからも可能
   - ブラウザーで開く
   - エディターでのプレビュー
   - ポートの表示範囲
       - Private
       - Private to Org.
       - **Public**
<img src="images/cs-w2-port-pub.png" width="800">

4B. **プレビューを見る（Webサーバー課題のみ）**
   
   **注意**: LMSの課題内容で「Webサーバー」「HTTP」「ブラウザで確認」などの指示がある場合のみ実施してください。
   
   a. `Ctrl+Shift+P`（Mac: `Cmd+Shift+P`）を押す
   
   b. 「**simple browser**」と入力
   
   c. 「**Simple Browser: Show**」を選択
  
   d. `http://localhost:8080` と入力してEnter
   
   e. 右側にブラウザ画面が表示される

   
   ```
   ┌──────────────┬──────────────┐
   │ main.go      │ プレビュー    │
   │ (編集)        │ (表示)        │
   └──────────────┴──────────────┘
   ```

5. **コードを変更した後**
   
   a. ファイルを保存（`Ctrl+S` / `Cmd+S`）
   
   b. 右上の「▶ Run Code」ボタンをクリック
   
   c. ブラウザ画面で右クリック → 「Reload」を選択

### 【ステップ3】提出（Git操作不要）

#### 簡単3ステップで提出完了

```
1. 画面左下の「≡」メニューをクリック
2. 「My Codespaces」で表示される画面から，作業中のCodespace のメニューを展開．「Export changes to a branch」を選択
3. 「Create branch」をクリック
```

<img src="images/mycodespaces-from-codesapce.png" alt="メニューを開く" width="700">

<img src="images/page-https--github.com-codespaces-stopped-codespace-Export-changes-to-a-branch.png" alt="Export changes to a branch" width="700">

<img src="images/page-https--github.com-codespaces-stopped-codespace-new-branch.png" alt="ブランチ作成" width="700">

#### ブランチ名の報告

1. 表示されたブランチ名をコピー
   ```
   例: codespace-symmetrical-guacamole-q74vxx47qp92x7wv
   ```

<img src="images/page-https--github.com-codespaces-stopped-codespace-new-branch-cont.png" alt="ブランチ名表示" width="700">

<img src="images/page-https--github.com-codespaces-stopped-codespace-new-branch-finished.png" alt="ブランチ作成完了" width="700">

2. LMSの提出ボックスにブランチ名を貼り付け

3. LMSで「提出」をクリック

4. 提出完了 ✅

### 【ステップ4】Codespaceの停止（超重要）⚠️

**必ず実行してください！停止しないと課金され続けます。**

#### 停止方法（推奨）

1. ブラウザで自分のリポジトリページを開く

2. 「**Code**」→「**Codespaces**」タブをクリック

<img src="images/page-https--github.com-codespaces.png" alt="Codespacesタブ" width="700">

3. 実行中のCodespace（緑色の●マーク）の右側「**...**」をクリック

<img src="images/page-https--github.com-codespaces-actions.png" alt="アクションメニュー" width="700">

4. 「**Stop codespace**」を選択

<img src="images/page-https--github.com-codespaces-stop-codespace.png" alt="Stop codespace" width="700">

5. ステータスが「**Stopped**」になれば完了 ✅

<img src="images/page-https--github.com-codespaces-stopped-codespace.png" alt="停止完了" width="700">

もしくは，以下から「Stop Codespaces」を選択
<img src="images/codespaces-menu-stop-codespaces.png" alt="停止完了" width="700">

#### 停止のタイミング

以下の場合は**必ず**停止してください:

- ✅ 課題提出後（すぐに停止）
- ✅ 30分以上離席する時
- ✅ その日の作業を終える時
- ✅ 他の授業に移る時
- ✅ PCをシャットダウンする前

#### なぜ停止が重要？

```
無料枠: 月180コア時間（2コアマシンで90時間/月）
※GitHub Student Developer Pack取得済みの場合

停止を忘れると:
例）24時間放置 = 48コア時間消費
→ 2日で無料枠の半分以上を使用
→ 無料枠超過で追加料金が発生

停止を習慣化すれば:
週3-4時間の使用 = 月12-16時間
→ 無料枠内で十分足ります
```

---

## よくある質問（FAQ）

### Q1: 「Export changes to a branch」が見つかりません

**A**: 以下を確認してください

```
1. 画面左下の「≡」（3本線）メニューをクリック
2. メニューから「Codespaces」を探す
3. その中に「Export changes to a branch」があります

別の場所:
- 画面上部の「...」（3点メニュー）からもアクセス可能
```

<img src="images/page-https---github.com-codespaces-menu.png" alt="Codespacesメニュー" width="700">

### Q2: ブランチ名を忘れてしまいました

**A**: GitHubで確認できます

```
1. 自分のリポジトリページを開く
2. 「X branches」をクリック
3. 最新のブランチ名を確認
4. ブランチ名をコピーしてLMSで報告

または:
- 再度「Export changes to a branch」を実行
- 新しいブランチ名をLMSで報告
```

### Q3: どのフォルダで作業すればいいですか？

**A**: LMSの課題ページに記載されています

```
例: 「week03/main.go を編集してください」
→ week03 フォルダで作業
```

### Q4: Codespacesが起動しません

**A**: 以下を試してください

1. ブラウザのキャッシュをクリア
2. 別のブラウザで試す（Chrome、Edge推奨）
3. 既存のCodespaceを削除して新規作成
4. それでもダメなら教員に連絡

### Q5: 過去の週のコードを見たいです

**A**: 同じリポジトリ内にすべて保存されています

```
- 左側のエクスプローラーで過去のフォルダを開く
- 例: week01, week02 など
- いつでも参照・コピー可能
```

### Q6: Git操作を覚える必要がありますか？

**A**: いいえ、不要です

```
「Export changes to a branch」機能を使えば、
Git操作の知識は一切不要です。

この授業では:
- commit
- push
- pull
などの操作は使いません。
```

### Q7: 「▶ Run Code」ボタンが表示されません

**A**: 以下を確認してください

```
1. main.goファイルを開いているか
2. ファイルの種類が「Go」と認識されているか（右下に表示）

代替方法:
- ファイル上で右クリック → 「Run Code」
- ターミナルで `go run main.go` を実行
```

### Q8: プレビューが更新されません

**A**: 以下を確認してください

```
1. ファイルを保存したか（Ctrl+S）
2. プログラムを再実行したか（▶ Run Code）
3. ブラウザをリロードしたか（右クリック → Reload）

それでも更新されない場合:
- ターミナルで Ctrl+C（プログラム停止）
- 再度「▶ Run Code」で実行
```

### Q9: エラーメッセージが表示されました

**A**: エラーメッセージ集を確認してください

**詳細は [ERROR_MESSAGES_GUIDE.md](ERROR_MESSAGES_GUIDE.md) を参照してください。**

主なエラー:
- `permission denied` - 権限エラー
- `cannot find package` - パッケージが見つからない
- `syntax error` - 構文エラー
- `port already in use` - ポートが使用中
- その他19種類のエラーと解決方法を掲載

### Q10: 間違ったファイルを編集してしまいました

**A**: 正しいファイルを編集し直してください

```
1. 正しいファイル（該当週のmain.go）を編集
2. 再度「Export changes to a branch」を実行
3. 新しいブランチ名をLMSで報告
4. Codespaceを停止

教師は最新のブランチを確認するので問題ありません。
```

### Q11: セキュリティについて詳しく知りたい

**A**: セキュリティガイドを参照してください

**詳細は [SECURITY_GUIDE.md](SECURITY_GUIDE.md) を参照してください。**

内容:
- 個人情報の取り扱い
- API Keyと機密情報の管理
- 環境変数の使用方法
- ポート公開の注意事項
- 安全なコーディング例
- 提出前チェックリスト

### Q12: Codespaceを停止し忘れました

**A**: すぐに停止してください

```
1. リポジトリページで「Code」→「Codespaces」
2. 実行中のCodespaceの「...」→「Stop codespace」
3. 次回から作業終了後すぐに停止する習慣を

課金時間の確認:
- GitHub設定 → Billing and plans
- 「Codespaces」セクションで使用時間を確認
```

### Q13: 無料枠を超えたらどうなりますか？

**A**: Codespacesが使用できなくなります

```
無料枠超過後:
- 新しいCodespaceを起動できない
- 既存のCodespaceも起動できない
- 翌月1日まで使用不可

対策:
- 停止を徹底する
- 使用時間を定期的に確認
- 計画的に使用する

停止を習慣化すれば無料枠内で十分です。
```

### Q14: 複数のCodespaceが実行中です

**A**: すべて停止してください

```
同時に複数実行していると、
使用時間が倍速で消費されます。

1. リポジトリの「Code」→「Codespaces」
2. すべての実行中Codespaceを停止
3. 使っていないCodespaceは削除
```

<img src="images/page-https---github.com-codespaces-0.png" alt="複数のCodespace" width="700">

### Q15: 停止すると作業内容が消えますか？

**A**: いいえ、消えません

```
- ファイルはすべて保存されます
- 次回再開時に同じ状態から始められます
- 安心して停止してください
```

---

## 経験者向け - 詳細情報

> 💡 **プログラミング経験者の方へ**: 基本操作をスキップして、詳細機能を確認できます。

---

## 便利な機能とショートカット

### コードのフォーマット

```bash
# 自動フォーマット
go fmt main.go
```

または

- ファイルを保存すると自動フォーマット（設定済み）

### 文法チェック

```bash
# リント実行
golangci-lint run
```

### ターミナルの使い方

```bash
# ディレクトリ移動
cd week03

# ファイル一覧
ls

# プログラム実行
go run main.go

# テスト実行（テストファイルがある場合）
go test

# ビルド
go build -o myapp main.go
```

### デバッグ

VS Codeのデバッガがインストールされています：

1. ブレークポイントを設定（行番号の左をクリック）
2. F5キーを押してデバッグ開始
3. ステップ実行、変数の確認が可能

---

## トラブルシューティング

### パフォーマンス最適化

```bash
# Go module cacheのクリーンアップ
go clean -modcache

# ビルドキャッシュのクリーンアップ
go clean -cache
```

### 環境のリセット

```bash
# Codespace内で環境変数を確認
go env

# Go toolsの再インストール（必要な場合）
go install golang.org/x/tools/gopls@latest
```

---

## 高度な使用方法

### 複数ファイルのプロジェクト

```go
// main.go
package main

import "./mypackage"

func main() {
    mypackage.Hello()
}
```

```go
// mypackage/hello.go
package mypackage

import "fmt"

func Hello() {
    fmt.Println("Hello from package!")
}
```

### 外部パッケージの使用

```bash
# go.modファイルの初期化
go mod init example.com/myproject

# パッケージのインストール
go get github.com/some/package

# 依存関係の整理
go mod tidy
```

---

## その他の重要なガイド

### セキュリティガイド（別文書）

**[SECURITY_GUIDE.md](SECURITY_GUIDE.md) を必ずお読みください。**

このガイドには以下の重要な情報が含まれています：

#### 📋 主な内容
- ✅ **個人情報の取り扱い** - 何が個人情報か、コードに含めてはいけない情報
- ✅ **API Keyと機密情報** - 環境変数の正しい使い方
- ✅ **ポート公開の注意事項** - Webサーバーを公開する際の注意点
- ✅ **安全なコーディング例** - ユーザー登録、DB接続などの実装例
- ✅ **提出前チェックリスト** - 具体的な確認項目

#### ⚠️ 特に重要なポイント

**絶対にコードに含めてはいけない情報：**
- 実名、住所、電話番号、メールアドレス
- API Key、パスワード、アクセストークン
- データベース接続情報

**詳細な例とチェック方法は [SECURITY_GUIDE.md](SECURITY_GUIDE.md) を参照してください。**

---

### エラーメッセージ集（別文書）

**[ERROR_MESSAGES_GUIDE.md](ERROR_MESSAGES_GUIDE.md) を困ったときに参照してください。**

このガイドには以下のエラー解決方法が含まれています：

#### 📋 収録内容（全19種類のエラー）

**環境関連のエラー (3件)**
- `permission denied` - 権限エラーの解決方法
- `go: command not found` - Goが認識されない場合
- `GOPATH not set` - 環境変数の問題

**Go言語の文法エラー (5件)**
- `cannot find package` - パッケージが見つからない
- `undefined: xxxxx` - 関数や変数が未定義
- `syntax error` - 構文エラー
- `declared and not used` - 未使用変数
- `cannot use xxx as type yyy` - 型の不一致

**実行時エラー (3件)**
- `index out of range` - 配列の範囲外アクセス
- `nil pointer dereference` - nilポインタの参照
- `integer divide by zero` - ゼロ除算

**ネットワーク関連のエラー (3件)**
- `port is already allocated` - ポートが使用中
- `no such host` - ホストが見つからない
- `context deadline exceeded` - タイムアウト

**Codespaces関連のエラー (2件)**
- Codespaceが起動しない
- ファイルが保存できない

**提出関連のエラー (3件)**
- 「Export changes to a branch」が押せない
- ブランチ名が見つからない
- LMSで提出できない

#### 🔍 各エラーの情報

各エラーには以下の情報が含まれています：
- ✅ エラーメッセージの例
- ✅ 原因の詳細説明
- ✅ 具体的な解決方法（コード例付き）
- ✅ 悪い例と良い例の比較

**詳細は [ERROR_MESSAGES_GUIDE.md](ERROR_MESSAGES_GUIDE.md) を参照してください。**

---

## 提出確認方法

### GitHubでの確認

1. ブラウザで自分のリポジトリページを開く

2. 「**X branches**」をクリック

3. 自分が作成したブランチが表示されているか確認

4. ブランチ名をクリック→該当週のフォルダを開く

5. 自分のコードが表示されていればOK ✅

### LMSでの確認

1. LMSの課題ページで「提出済み」と表示されているか確認

2. ブランチ名が正しく記録されているか確認

---

## 修正して再提出する場合

1. Codespacesで該当ファイルを修正

2. 再度「Export changes to a branch」を実行
   - 新しいブランチが自動作成されます

3. 新しいブランチ名をLMSで報告
   - LMSで「編集」または「再提出」
   - 新しいブランチ名に更新

4. **Codespaceを停止** ⚠️

**重要**: 最新のブランチ名を報告してください

---

## チェックリストと成功のコツ

### 初回セットアップ（第1週）

- [ ] 課題リンクから「Accept this assignment」
- [ ] リポジトリURLをブックマーク
- [ ] Codespacesを起動
- [ ] week01のmain.goを編集
- [ ] 「Export changes to a branch」で提出
- [ ] ブランチ名をLMSで報告
- [ ] **Codespaceを停止** ⚠️

### 毎週の作業（第2週以降）

- [ ] LMSで課題内容を確認
- [ ] ブックマークからリポジトリにアクセス
- [ ] Codespacesを起動（約3秒）
- [ ] 該当週のフォルダで作業
- [ ] プログラムを実行してテスト
- [ ] **セキュリティ確認** - [SECURITY_GUIDE.md](SECURITY_GUIDE.md) のチェックリスト参照
- [ ] 「Export changes to a branch」で提出
- [ ] ブランチ名をLMSで報告
- [ ] **Codespaceを停止** ⚠️

### 停止確認

- [ ] リポジトリの「Codespaces」タブで確認
- [ ] ステータスが「Stopped」になっている
- [ ] 緑色の●マークが消えている

---

## 成功のコツ

1. **リポジトリURLを必ずブックマーク**
   - 学期中ずっと使います

2. **LMSで課題内容を毎週確認**
   - 詳細な要件はLMSに記載

3. **セキュリティを意識**
   - [SECURITY_GUIDE.md](SECURITY_GUIDE.md) を参照
   - 個人情報を含めない
   - API Keyはコードに書かない

4. **提出後は必ずブランチ名を報告**
   - 報告しないと教師が確認できません

5. **作業終了後は必ずCodespaceを停止** ⚠️
   - 停止し忘れると課金され続けます
   - 習慣化が重要です
   - 課題提出後すぐに停止

6. **過去のコードを活用**
   - 同じリポジトリ内に全週のコードが残ります

7. **困ったらすぐに質問**
   - [ERROR_MESSAGES_GUIDE.md](ERROR_MESSAGES_GUIDE.md) を確認
   - LMSの質問フォーラムで質問
   - 早めに解決することが大切

---

## 最重要ポイント

### 4つの必須習慣

```
1. 編集
2. 実行
3. 提出
4. 停止 ← 必ず実行！
```

### 参照すべき重要ドキュメント

```
1. SECURITY_GUIDE.md - セキュリティの3原則
2. ERROR_MESSAGES_GUIDE.md - エラー解決方法
3. CODESPACE_STOP_GUIDE.md - 停止の詳細手順
```

### 停止を忘れると

- 無料枠を使い切る
- 追加料金が発生する可能性
- 他の課題で使えなくなる

### 停止の確認方法

```
リポジトリページ → Code → Codespaces
→ ステータスが「Stopped」になっているか確認
```

---

## ブラウザでの確認方法

### 方法1: エディタ内プレビュー（推奨）
サーバー起動後、自動的にエディタ内に表示されます

### 方法2: 外部ブラウザ
1. VS Code下部の「ポート」タブを開く
2. `8080` の行の地球儀アイコンをクリック

### 方法3: Simple Browser（シンプル表示）
1. `Ctrl+Shift+P` でコマンドパレットを開く
2. "Simple Browser: Show" を選択
3. `http://localhost:8080` を入力

---

## サポート情報

困ったときは：

1. **このガイド**: よくある質問を再確認
2. **[ERROR_MESSAGES_GUIDE.md](ERROR_MESSAGES_GUIDE.md)**: エラー解決方法（19種類のエラーを収録）
3. **[SECURITY_GUIDE.md](SECURITY_GUIDE.md)**: セキュリティ情報とチェックリスト
4. **LMS**: 質問フォーラムで質問
5. **授業**: 質疑応答時間に質問
6. **オフィスアワー**: [LMSで確認]

---

**セキュリティを意識し、停止を習慣化して、素晴らしいプログラミング学習を！** 🚀

---

## 📚 関連ドキュメント一覧

### 学生向けドキュメント
- **[SECURITY_GUIDE.md](SECURITY_GUIDE.md)** - セキュリティとプライバシーの完全ガイド
- **[ERROR_MESSAGES_GUIDE.md](ERROR_MESSAGES_GUIDE.md)** - 19種類のエラーと解決方法
- **[CODESPACE_STOP_GUIDE.md](CODESPACE_STOP_GUIDE.md)** - Codespace停止の詳細ガイド

### 教員・TA向けドキュメント
- **[TEACHER_REVIEW_GUIDE.md](TEACHER_REVIEW_GUIDE.md)** - 課題確認・評価ガイド
- **[CONTAINER_SETUP_MANUAL.md](CONTAINER_SETUP_MANUAL.md)** - コンテナイメージ構築ガイド（Podman版）












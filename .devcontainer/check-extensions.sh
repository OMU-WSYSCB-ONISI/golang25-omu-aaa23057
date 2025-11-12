#!/bin/bash
# 拡張機能監視スクリプト
# 用途: 許可されていない拡張機能をチェック

# 許可された拡張機能リスト
ALLOWED_EXTENSIONS=(
    "golang.go"
    "oderwat.indent-rainbow"
    "mosapride.zenkaku"
    "shardulm94.trailing-spaces"
    "usernamehw.errorlens"
    "formulahendry.code-runner"
    "aaron-bond.better-comments"
    "wayou.vscode-todo-highlight"
    "mhutchie.git-graph"
    "pkief.material-icon-theme"
    "ms-ceintl.vscode-language-pack-ja"
    "yzhang.markdown-all-in-one"
    "bierner.markdown-preview-github-styling"
)

# AI補完系の禁止拡張機能
FORBIDDEN_EXTENSIONS=(
    "github.copilot"
    "github.copilot-chat"
    "github.copilot-labs"
    "github.copilot-nightly"
    "visualstudioexptteam.vscodeintellicode"
    "visualstudioexptteam.intellicode-api-usage-examples"
    "tabnine.tabnine-vscode"
    "codeium.codeium"
    "amazonwebservices.aws-toolkit-vscode"
)

echo "🔍 拡張機能チェック中..."
echo ""

# インストール済み拡張機能を取得
INSTALLED=$(code --list-extensions 2>/dev/null)

if [ -z "$INSTALLED" ]; then
    echo "⚠️  拡張機能リストを取得できませんでした"
    exit 0
fi

# 禁止拡張機能のチェック
FOUND_FORBIDDEN=0
echo "❌ 禁止されている拡張機能:"
for ext in "${FORBIDDEN_EXTENSIONS[@]}"; do
    if echo "$INSTALLED" | grep -qi "^$ext\$"; then
        echo "  - $ext"
        FOUND_FORBIDDEN=1
    fi
done

if [ $FOUND_FORBIDDEN -eq 0 ]; then
    echo "  なし ✅"
fi
echo ""

# 許可されていない拡張機能のチェック
echo "⚠️  許可リストにない拡張機能:"
FOUND_UNAUTHORIZED=0
while IFS= read -r ext; do
    ext_lower=$(echo "$ext" | tr '[:upper:]' '[:lower:]')
    
    # 許可リストにあるかチェック
    IS_ALLOWED=0
    for allowed in "${ALLOWED_EXTENSIONS[@]}"; do
        if [ "$ext_lower" = "$allowed" ]; then
            IS_ALLOWED=1
            break
        fi
    done
    
    # 禁止リストにあるかチェック（既に報告済み）
    IS_FORBIDDEN=0
    for forbidden in "${FORBIDDEN_EXTENSIONS[@]}"; do
        if [ "$ext_lower" = "$forbidden" ]; then
            IS_FORBIDDEN=1
            break
        fi
    done
    
    if [ $IS_ALLOWED -eq 0 ] && [ $IS_FORBIDDEN -eq 0 ]; then
        echo "  - $ext"
        FOUND_UNAUTHORIZED=1
    fi
done <<< "$INSTALLED"

if [ $FOUND_UNAUTHORIZED -eq 0 ]; then
    echo "  なし ✅"
fi
echo ""

# 結果サマリー
if [ $FOUND_FORBIDDEN -eq 1 ]; then
    echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
    echo "🚨 禁止されている拡張機能が検出されました"
    echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
    echo ""
    echo "以下のコマンドで削除してください:"
    echo ""
    for ext in "${FORBIDDEN_EXTENSIONS[@]}"; do
        if echo "$INSTALLED" | grep -qi "^$ext\$"; then
            echo "  code --uninstall-extension $ext"
        fi
    done
    echo ""
    echo "または、Codespaceを再起動すると自動削除されます。"
    echo ""
    exit 1
elif [ $FOUND_UNAUTHORIZED -eq 1 ]; then
    echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
    echo "⚠️  許可リストにない拡張機能があります"
    echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
    echo ""
    echo "必要な場合は教員に確認してください。"
    echo ""
    exit 0
else
    echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
    echo "✅ すべての拡張機能が適切です"
    echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
    exit 0
fi
#!/bin/bash
# 拡張機能監視スクリプト
# 用途: 学生が許可されていない拡張機能をインストールしていないかチェック

# 許可された拡張機能リスト
ALLOWED_EXTENSIONS=(
    "golang.go"
    "oderwat.indent-rainbow"
    "mosapride.zenkaku"
    "shardulm94.trailing-spaces"
    "usernamehw.errorlens"
    "formulahendry.code-runner"
    "aaron-bond.better-comments"
    "wayou.vscode-todo-highlight"
    "mhutchie.git-graph"
    "pkief.material-icon-theme"
    "ms-ceintl.vscode-language-pack-ja"
    "yzhang.markdown-all-in-one"
    "bierner.markdown-preview-github-styling"
)

# AI補完系の禁止拡張機能（検出したら警告）
FORBIDDEN_EXTENSIONS=(
    "github.copilot"
    "github.copilot-chat"
    "visualstudioexptteam.vscodeintellicode"
    "tabnine.tabnine-vscode"
    "codeium.codeium"
    "amazonwebservices.aws-toolkit-vscode"  # CodeWhisperer含む
)

echo "🔍 拡張機能チェック中..."
echo ""

# インストール済み拡張機能を取得
INSTALLED=$(code --list-extensions 2>/dev/null)

if [ -z "$INSTALLED" ]; then
    echo "⚠️  拡張機能リストを取得できませんでした"
    exit 0
fi

# 禁止拡張機能のチェック
FOUND_FORBIDDEN=0
echo "❌ 禁止されている拡張機能:"
for ext in "${FORBIDDEN_EXTENSIONS[@]}"; do
    if echo "$INSTALLED" | grep -qi "$ext"; then
        echo "  - $ext （AI補完機能 - 削除してください）"
        FOUND_FORBIDDEN=1
    fi
done

if [ $FOUND_FORBIDDEN -eq 0 ]; then
    echo "  なし ✅"
fi
echo ""

# 許可されていない拡張機能のチェック
echo "⚠️  許可リストにない拡張機能:"
FOUND_UNAUTHORIZED=0
while IFS= read -r ext; do
    ext_lower=$(echo "$ext" | tr '[:upper:]' '[:lower:]')
    
    # 許可リストにあるかチェック
    IS_ALLOWED=0
    for allowed in "${ALLOWED_EXTENSIONS[@]}"; do
        if [ "$ext_lower" = "$allowed" ]; then
            IS_ALLOWED=1
            break
        fi
    done
    
    if [ $IS_ALLOWED -eq 0 ]; then
        echo "  - $ext"
        FOUND_UNAUTHORIZED=1
    fi
done <<< "$INSTALLED"

if [ $FOUND_UNAUTHORIZED -eq 0 ]; then
    echo "  なし ✅"
fi
echo ""

# 結果サマリー
if [ $FOUND_FORBIDDEN -eq 1 ]; then
    echo "🚨 禁止されている拡張機能が検出されました！"
    echo "   課題評価に影響する可能性があります。"
    echo "   拡張機能タブから削除してください。"
    exit 1
elif [ $FOUND_UNAUTHORIZED -eq 1 ]; then
    echo "⚠️  許可リストにない拡張機能が検出されました。"
    echo "   必要な場合は教員に確認してください。"
    exit 0
else
    echo "✅ すべての拡張機能が適切です！"
    exit 0
fi

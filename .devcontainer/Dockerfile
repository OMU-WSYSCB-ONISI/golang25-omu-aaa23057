# Golang開発環境用Dockerfile（サイズ最適化版・Podman対応）
#
# ビルド方法:
#   podman system prune -a -f --volumes  # 初回推奨
#   podman build --no-cache --format docker -t go-devcontainer:slim .
#
# 通常ビルド:
#   podman build --format docker -t go-devcontainer:slim .
#
# サイズ削減のポイント:
# 1. builderステージでビルドツール削除
# 2. 最終ステージでgcc等の削減（CGO不要なら削除可能）
# 3. Go開発ツールのバイナリサイズ削減
# 4. キャッシュの完全削除
FROM mcr.microsoft.com/devcontainers/base:debian AS builder

ENV GO_VERSION=1.23.12 \
    DEBIAN_FRONTEND=noninteractive

# 最小限のパッケージでGoをインストール
RUN apt-get update && \
    apt-get install -y --no-install-recommends \
        ca-certificates wget && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*

# Goのダウンロードとインストール
RUN set -eux; \
    wget -q -O "/tmp/go${GO_VERSION}.linux-amd64.tar.gz" \
        "https://go.dev/dl/go${GO_VERSION}.linux-amd64.tar.gz" && \
    tar -C /usr/local -xzf "/tmp/go${GO_VERSION}.linux-amd64.tar.gz" && \
    rm -f "/tmp/go${GO_VERSION}.linux-amd64.tar.gz"

# Go開発ツールをビルド（ストリップして軽量化）
ENV PATH=/usr/local/go/bin:$PATH \
    GOPATH=/tmp/gotools \
    CGO_ENABLED=0

RUN go install -ldflags="-s -w" golang.org/x/tools/gopls@v0.16.2 && \
    go install -ldflags="-s -w" github.com/golangci/golangci-lint/cmd/golangci-lint@v1.61.0 && \
    go install -ldflags="-s -w" golang.org/x/tools/cmd/goimports@v0.25.0 && \
    go install -ldflags="-s -w" github.com/go-delve/delve/cmd/dlv@v1.23.1 && \
    # さらにupxで圧縮（オプション、コメントアウト推奨）
    # apt-get update && apt-get install -y upx-ucl && \
    # upx --best --lzma /tmp/gotools/bin/* && \
    go clean -cache -modcache -testcache && \
    rm -rf /tmp/gotools/pkg /tmp/gotools/src

# 最終ステージ（最小構成）
FROM mcr.microsoft.com/devcontainers/base:debian

LABEL maintainer="onisi@omu.ac.jp" \
      description="Size-optimized Golang development environment" \
      version="25.10.4-slim" \
      base-os="debian-trixie" \
      go-version="1.23.12" \
      org.opencontainers.image.source="https://github.com/OMU-WSYSCB-ONISI/golang" \
      org.opencontainers.image.licenses="MIT"

ARG GO_VERSION=1.23.12

ENV GOPATH=/go \
    GOBIN=/go/bin \
    PATH=/usr/local/go/bin:/go/bin:$PATH \
    TZ=Asia/Tokyo \
    LANG=ja_JP.UTF-8 \
    LC_ALL=ja_JP.UTF-8 \
    GOCACHE=/home/vscode/.cache/go-build \
    GOTOOLCHAIN=local \
    CGO_ENABLED=0

# Goとツールをコピー
COPY --from=builder /usr/local/go /usr/local/go
COPY --from=builder /tmp/gotools/bin/* /go/bin/

# 最小限のパッケージのみインストール（CGO不要のため gcc/build-essential 削除）
RUN apt-get update && \
    apt-get install -y --no-install-recommends \
        sudo git make ca-certificates \
        tzdata locales \
        curl wget vim-tiny && \
    # ロケール設定（最小限）
    echo "ja_JP.UTF-8 UTF-8" > /etc/locale.gen && \
    locale-gen ja_JP.UTF-8 && \
    # タイムゾーン設定
    ln -fs /usr/share/zoneinfo/Asia/Tokyo /etc/localtime && \
    dpkg-reconfigure -f noninteractive tzdata && \
    # sudoers設定
    echo "vscode ALL=(ALL) NOPASSWD:ALL" > /etc/sudoers.d/vscode && \
    chmod 0440 /etc/sudoers.d/vscode && \
    # 徹底的なクリーンアップ
    apt-get autoremove -y && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/* \
           /tmp/* \
           /var/tmp/* \
           /var/cache/apt/* \
           /usr/share/doc/* \
           /usr/share/man/* \
           /usr/share/locale/* \
           /var/log/* && \
    # 不要なファイルの削除
    find /usr/share -type f -name "*.pyc" -delete && \
    find /usr/share -type d -name "__pycache__" -delete

# ディレクトリ作成と所有権設定
RUN mkdir -p /go/bin /go/pkg/mod /workspaces /home/vscode/.cache/go-build && \
    chown -R vscode:vscode /go /workspaces /home/vscode/.cache && \
    chmod -R 755 /go /workspaces

USER vscode
WORKDIR /workspaces

HEALTHCHECK --interval=30s --timeout=10s --start-period=5s --retries=3 \
    CMD go version && gopls version || exit 1

CMD ["/bin/bash"]

# =============================================================================
# サイズ削減効果（CGO無効化版）:
# - 元のイメージ: 約1.5GB
# - 最適化版: 約900MB-1.1GB（約27-40%削減）
#
# 削減内容:
# ✓ Go開発ツールのストリップ（-ldflags="-s -w"）: 約50MB削減
# ✓ gcc/build-essentialの削除: 約100MB削減
# ✓ ドキュメント/マニュアル/ロケールファイル削除: 約50-100MB削減
# ✓ vim → vim-tiny: 約5MB削減
# ✓ nano/less削除: 約3MB削減
#
# さらなる削減オプション:
# 1. golangci-lintを削除（容量大） → 約150MB削減
#    ※ 静的解析が不要な場合のみ推奨
# 2. dlv（デバッガ）を削除 → 約30MB削減
#    ※ デバッグが不要な場合のみ推奨
# 3. distroless/static-debianベースに変更 → 約50%削減（上級者向け）
#
# CGO無効化の影響:
# - CGO_ENABLED=0により、純粋なGoコードのみコンパイル可能
# - Cライブラリに依存するパッケージ（sqlite3等）は使用不可
# - クロスコンパイルが容易になる
# - バイナリサイズが小さくなる
# - 教育目的には十分な機能を保持
# =============================================================================

package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("========== quick-serve ==========")

	// ========== コマンドライン引数 ==========
	// コマンドライン引数チェック
	if len(os.Args) != 3 {
		fmt.Println("[Error]コマンドライン引数の数が不正です。\nサーバーモードと読み込むファイル/ディレクトリパスの2つの引数を入力してください。")
		fmt.Println("例: quick-serve web /path/to/serve/")
		os.Exit(-1)
	}

	// コマンドライン引数取得
	var serverModeRaw = strings.ToLower(os.Args[1])
	var path = os.Args[2]

	// サーバーモード選択
	var serverMode = Undefined
	switch serverModeRaw {
	case "web":
		serverMode = WebMode
	case "api":
		serverMode = ApiMode
	case "websocket":
		serverMode = WebSocketMode
	}

	if serverMode == Undefined {
		fmt.Println("[Error]対応していないモードです。")
		os.Exit(-1)
	}

	// Webモードは未対応
	if serverMode == WebMode {
		fmt.Println("[Error]Webモードは未対応です。")
		os.Exit(-1)
	}

	// WebSocketモードは未対応
	if serverMode == WebSocketMode {
		fmt.Println("[Error]WebSocketモードは未対応です。")
		os.Exit(-1)
	}
}

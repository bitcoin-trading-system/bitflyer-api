package main

import (
	"os/exec"
	"testing"
	"time"
)

func TestMainFunction(t *testing.T) {
	// main関数をゴルーチンで実行
	go main()

	type args struct {
		curl []string
	}

	tests := []struct {
		name string
		args args
	}{
		{
			name: "health check",
			args: args{
				curl: []string{"curl", "-w", "%{http_code}", "-o", "/dev/null", "-s", "http://localhost:8001/healthcheck"},
			},
		},
		{
			name: "board",
			args: args{
				curl: []string{"curl", "-w", "%{http_code}", "-o", "/dev/null", "-s", "http://localhost:8001/board/"},
			},
		},
		{
			name: "ticker",
			args: args{
				curl: []string{"curl", "-w", "%{http_code}", "-o", "/dev/null", "-s", "http://localhost:8001/ticker/"},
			},
		},
		{
			name: "executions",
			args: args{
				curl: []string{"curl", "-w", "%{http_code}", "-o", "/dev/null", "-s", "http://localhost:8001/executions/"},
			},
		},
		{
			name: "board_state",
			args: args{
				curl: []string{"curl", "-w", "%{http_code}", "-o", "/dev/null", "-s", "http://localhost:8001/board_state/"},
			},
		},
		{
			name: "health",
			args: args{
				curl: []string{"curl", "-w", "%{http_code}", "-o", "/dev/null", "-s", "http://localhost:8001/health/"},
			},
		},
		{
			name: "getbalance",
			args: args{
				curl: []string{"curl", "-w", "%{http_code}", "-o", "/dev/null", "-s", "http://localhost:8001/getbalance/"},
			},
		},
		{
			name: "getcollateral",
			args: args{
				curl: []string{"curl", "-w", "%{http_code}", "-o", "/dev/null", "-s", "http://localhost:8001/getcollateral/"},
			},
		},
		{
			name: "sendchildorder_dryrun",
			args: args{
				curl: []string{"curl", "-w", "%{http_code}", "-o", "/dev/null", "-s", "-X", "POST", "-H", "Content-Type: application/json", "-d", `{"product_code": "BTC_JPY", "child_order_type": "LIMIT", "side": "BUY", "price": 1, "size": 0.0001, "minute_to_expire": 1, "time_in_force": "GTC"}`, "http://localhost:8001/sendchildorder/?dry_run=1"},
			},
		},
		{
			name: "cancelchildorder_dryrun",
			args: args{
				curl: []string{"curl", "-w", "%{http_code}", "-o", "/dev/null", "-s", "-X", "POST", "-H", "Content-Type: application/json", "-d", `{"product_code": "BTC_JPY", "child_order_id": "hogehoge-id"}`, "http://localhost:8001/cancelchildorder/?dry_run=1"},
			},
		},
		{
			name: "getchildorders",
			args: args{
				curl: []string{"curl", "-w", "%{http_code}", "-o", "/dev/null", "-s", "http://localhost:8001/getchildorders/"},
			},
		},
	}

	// サーバーが起動するのを待つ
	time.Sleep(1 * time.Second)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// curlコマンドを実行
			cmd := exec.Command("curl", tt.args.curl...)
			output, err := cmd.CombinedOutput()
			if err != nil {
				t.Fatalf("Failed to execute curl command: %v", err)
			}

			// ステータスコードを確認
			statusCode := string(output[len(output)-3:]) // 最後の3文字がステータスコード
			if statusCode != "200" {
				t.Fatalf("Server returned non-200 status code: %s", statusCode)
			}
		})
	}
}

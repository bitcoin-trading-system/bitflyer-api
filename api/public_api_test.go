package api

import (
	"reflect"
	"testing"

	"github.com/bitcoin-trading-system/bitflyer-api/api/models"
	"github.com/bitcoin-trading-system/bitflyer-api/config"
)

func TestNewPublicAPI(t *testing.T) {
	type args struct {
		cfg config.Config
	}
	tests := []struct {
		name string
		args args
		want PublicAPI
	}{{
		name: "TestNewPublicAPI",
		args: args{
			cfg: APITestConfig,
		},
		want: API{
			BaseUrl:   BaseUrl(APITestConfig.BitflyerConfig.BaseEndPoint),
			ApiKey:    APITestConfig.BitflyerConfig.ApiKey,
			ApiSecret: APITestConfig.BitflyerConfig.ApiSecret,
		},
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewPublicAPI(tt.args.cfg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPublicAPI() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAPI_GetBoard(t *testing.T) {
	type args struct {
		productCode string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "TestAPI_GetBoard",
			args: args{
				productCode: "BTC_JPY",
			},
			wantErr: false,
		},
		{
			name: "TestAPI_GetBoard_No_Product_Code",
			args: args{
				productCode: "",
			},
			wantErr: false,
		},
		{
			name: "TestAPI_GetBoard_Invalid_Product_Code",
			args: args{
				productCode: "invalid_product_code",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			api := NewPublicAPI(APITestConfig)
			got, err := api.GetBoard(tt.args.productCode)
			if err != nil {
				if !tt.wantErr {
					t.Errorf("API.GetBoard() error = %v, wantErr %v", err, tt.wantErr)
				}
				return
			}

			if got.Asks == nil {
				t.Errorf("API.GetBoard() = %v, want %v", got.Asks, "not nil")
			}
			if got.Bids == nil {
				t.Errorf("API.GetBoard() = %v, want %v", got.Bids, "not nil")
			}
			if got.MidPrice == 0 {
				t.Errorf("API.GetBoard() = %v, want %v", got.MidPrice, "not 0")
			}
		})
	}
}

func TestAPI_GetTicker(t *testing.T) {
	type args struct {
		productCode string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "TestAPI_GetTicker",
			args: args{
				productCode: "BTC_JPY",
			},
			wantErr: false,
		},
		{
			name: "TestAPI_GetTicker_No_Product_Code",
			args: args{
				productCode: "",
			},
			wantErr: false,
		},
		{
			name: "TestAPI_GetTicker_Invalid_Product_Code",
			args: args{
				productCode: "invalid_product_code",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			api := NewPublicAPI(APITestConfig)
			got, err := api.GetTicker(tt.args.productCode)
			if err != nil {
				if !tt.wantErr {
					t.Errorf("API.GetTicker() error = %v, wantErr %v", err, tt.wantErr)
				}
				return
			}

			if got.ProductCode == "" {
				t.Errorf("API.GetTicker() = %v, want %v", got.ProductCode, "not empty")
			}
			if got.Timestamp == "" {
				t.Errorf("API.GetTicker() = %v, want %v", got.Timestamp, "not empty")
			}
			if got.BestAsk == 0 {
				t.Errorf("API.GetTicker() = %v, want %v", got.BestAsk, "not 0")
			}
			if got.BestBid == 0 {
				t.Errorf("API.GetTicker() = %v, want %v", got.BestBid, "not 0")
			}
			if got.BestAskSize == 0 {
				t.Errorf("API.GetTicker() = %v, want %v", got.BestAskSize, "not 0")
			}
			if got.BestBidSize == 0 {
				t.Errorf("API.GetTicker() = %v, want %v", got.BestBidSize, "not 0")
			}
			if got.Ltp == 0 {
				t.Errorf("API.GetTicker() = %v, want %v", got.Ltp, "not 0")
			}
			if got.Volume == 0 {
				t.Errorf("API.GetTicker() = %v, want %v", got.Volume, "not 0")
			}
			if got.VolumeByProduct == 0 {
				t.Errorf("API.GetTicker() = %v, want %v", got.VolumeByProduct, "not 0")
			}
			if got.State == "" {
				t.Errorf("API.GetTicker() = %v, want %v", got.State, "not empty")
			}
			if got.TickID == 0 {
				t.Errorf("API.GetTicker() = %v, want %v", got.TickID, "not 0")
			}
			if got.TotalAskDepth == 0 {
				t.Errorf("API.GetTicker() = %v, want %v", got.TotalAskDepth, "not 0")
			}
			if got.TotalBidDepth == 0 {
				t.Errorf("API.GetTicker() = %v, want %v", got.TotalBidDepth, "not 0")
			}
		})
	}
}

func TestAPI_GetExecutions(t *testing.T) {
	type args struct {
		productCode string
		count       string
		before      string
		after       string
	}
	tests := []struct {
		name    string
		args    args
		lenGot  int
		wantErr bool
	}{
		{
			name: "TestAPI_GetExecutions",
			args: args{
				productCode: "BTC_JPY",
				count:       "110",
				before:      "",
				after:       "",
			},
			lenGot:  110,
			wantErr: false,
		},
		{
			name: "TestAPI_GetExecutions_No_Product_Code",
			args: args{
				productCode: "",
				count:       "110",
				before:      "",
				after:       "",
			},
			lenGot:  110,
			wantErr: false,
		},
		{
			name: "TestAPI_GetExecutions_Invalid_Product_Code",
			args: args{
				productCode: "invalid_product_code",
				count:       "110",
				before:      "",
				after:       "",
			},
			lenGot:  0,
			wantErr: true,
		},
		{
			name: "TestAPI_GetExecutions_No_Count",
			args: args{
				productCode: "BTC_JPY",
				count:       "",
				before:      "",
				after:       "",
			},
			lenGot:  100, // default count
			wantErr: false,
		},
		{
			name: "TestAPI_GetExecutions_Invalid_Count",
			args: args{
				productCode: "BTC_JPY",
				count:       "invalid_count",
				before:      "",
				after:       "",
			},
			lenGot:  100, // default count
			wantErr: false,
		},
		{
			name: "TestAPI_GetExecutions_Before_AND_After",
			args: args{
				productCode: "BTC_JPY",
				count:       "110",
				before:      "3",
				after:       "1",
			},
			lenGot:  0,
			wantErr: true, // データの有効期限が31日。対象のレコードが存在しないためエラーになる。beforeとafterの値は反映されることは確認できる。
		},
		{
			name: "TestAPI_GetExecutions_Before",
			args: args{
				productCode: "BTC_JPY",
				count:       "110",
				before:      "",
				after:       "2000",
			},
			lenGot:  110,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			api := NewPublicAPI(APITestConfig)
			got, err := api.GetExecutions(tt.args.productCode, tt.args.count, tt.args.before, tt.args.after)
			if err != nil {
				if !tt.wantErr {
					t.Errorf("API.GetExecutions() error = %v, wantErr %v", err, tt.wantErr)
				}
				return
			}

			if len(got) != tt.lenGot {
				t.Errorf("API.GetExecutions() = %v, want %v", len(got), tt.lenGot)
			}
		})
	}
}

func TestAPI_GetBoardState(t *testing.T) {
	type args struct {
		productCode string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "TestAPI_GetBoardState",
			args: args{
				productCode: "BTC_JPY",
			},
			wantErr: false,
		},
		{
			name: "TestAPI_GetBoardState_No_Product_Code",
			args: args{
				productCode: "",
			},
			wantErr: false,
		},
		{
			name: "TestAPI_GetBoardState_Invalid_Product_Code",
			args: args{
				productCode: "invalid_product_code",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			api := NewPublicAPI(APITestConfig)
			got, err := api.GetBoardState(tt.args.productCode)
			if err != nil {
				if !tt.wantErr {
					t.Errorf("API.GetBoardState() error = %v, wantErr %v", err, tt.wantErr)
				}
				return
			}

			if got.Health == "" {
				t.Errorf("API.GetBoardState() = %v, want %v", got.Health, "not empty")
			}
			if got.State == "" {
				t.Errorf("API.GetBoardState() = %v, want %v", got.State, "not empty")
			}
		})
	}
}

func TestAPI_GetHealth(t *testing.T) {
	type args struct {
		productCode string
	}
	tests := []struct {
		name    string
		args    args
		want    models.Health
		wantErr bool
	}{
		{
			name: "TestAPI_GetHealth",
			args: args{
				productCode: "BTC_JPY",
			},
			want: models.Health{
				Status: "NORMAL",
			},
			wantErr: false,
		},
		{
			name: "TestAPI_GetHealth_No_Product_Code",
			args: args{
				productCode: "",
			},
			want: models.Health{
				Status: "NORMAL",
			},
			wantErr: false,
		},
		{
			name: "TestAPI_GetHealth_Invalid_Product_Code",
			args: args{
				productCode: "invalid_product_code",
			},
			want: models.Health{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			api := NewPublicAPI(APITestConfig)
			got, err := api.GetHealth(tt.args.productCode)
			if (err != nil) != tt.wantErr {
				t.Errorf("API.GetHealth() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("API.GetHealth() = %v, want %v", got, tt.want)
			}
		})
	}
}

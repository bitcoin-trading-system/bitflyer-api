package api

import (
	"reflect"
	"testing"

	"github.com/bitcoin-trading-system/bitflyer-api/api/models"
	"github.com/bitcoin-trading-system/bitflyer-api/config"
)

func TestNewPrivateAPI(t *testing.T) {
	type args struct {
		cfg config.Config
	}
	tests := []struct {
		name string
		args args
		want PrivateAPI
	}{{
		name: "TestNewPrivateAPI",
		args: args{
			cfg: APITestConfig,
		},
		want: &API{
			BaseUrl:   BaseUrl(APITestConfig.BitflyerConfig.BaseEndPoint),
			ApiKey:    APITestConfig.BitflyerConfig.ApiKey,
			ApiSecret: APITestConfig.BitflyerConfig.ApiSecret,
		},
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewPrivateAPI(tt.args.cfg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPrivateAPI() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAPI_GetBalance(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "TestAPI_GetBalance",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			api := NewPrivateAPI(APITestConfig)
			got, err := api.GetBalance()
			if err != nil {
				if !tt.wantErr {
					t.Errorf("API.GetBalance() error = %v, wantErr %v", err, tt.wantErr)
				}
				return
			}

			if len(got) == 0 {
				t.Errorf("API.GetBalance() = %v, want not empty", got)
			}
		})
	}
}

func TestAPI_GetCollateral(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "TestAPI_GetCollateral",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			api := NewPrivateAPI(APITestConfig)
			_, err := api.GetCollateral()
			if err != nil {
				if !tt.wantErr {
					t.Errorf("API.GetCollateral() error = %v, wantErr %v", err, tt.wantErr)
				}
				return
			}
		})
	}
}

func TestAPI_PostSendChildOrder(t *testing.T) {
	type args struct {
		req   models.SendChildOrderRequest
		isDry bool
	}
	tests := []struct {
		name    string
		args    args
		want    models.ChildOrder
		wantErr bool
	}{
		{
			name: "TestAPI_PostSendChildOrder",
			args: args{
				req: models.SendChildOrderRequest{
					ProductCode: "BTC_JPY",
					ChildOrderType: "MARKET",
					Side: "BUY",
					Size: 0.001,
					MinuteToExpire: 10000,
					TimeInForce: "GTC",
				},
				isDry: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			api := NewPrivateAPI(APITestConfig)
			got, err := api.PostSendChildOrder(tt.args.req, tt.args.isDry)
			if err != nil {
				if !tt.wantErr {
					t.Errorf("API.PostSendChildOrder() error = %v, wantErr %v", err, tt.wantErr)
				}
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("API.PostSendChildOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAPI_PostCancelChildOrder(t *testing.T) {
	type args struct {
		req   models.CancelChildOrderRequest
		isDry bool
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			api := NewPrivateAPI(APITestConfig)
			if err := api.PostCancelChildOrder(tt.args.req, tt.args.isDry); (err != nil) != tt.wantErr {
				t.Errorf("API.PostCancelChildOrder() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAPI_GetChildOrders(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			api := NewPrivateAPI(APITestConfig)
			got, err := api.GetChildOrders()
			if err != nil {
				if !tt.wantErr {
					t.Errorf("API.GetChildOrders() error = %v, wantErr %v", err, tt.wantErr)
				}
				return
			}

			if len(got) == 0 {
				t.Errorf("API.GetChildOrders() = %v, want not empty", got)
			}
		})
	}
}

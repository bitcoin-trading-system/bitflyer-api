package usecase

import (
	"testing"

	"github.com/bitcoin-trading-system/bitflyer-api/api/models"
	"github.com/bitcoin-trading-system/bitflyer-api/config"
)

var UseCaseTestConfig config.Config

func init() {
	UseCaseTestConfig = config.NewConfig("../conf/local.toml", "../env/.env.local")
}

func TestBitflyerUseCase_GetBoard(t *testing.T) {
	type args struct {
		productCode string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "TestBitflyerUseCase_GetBoard",
			args: args{
				productCode: ProductCodeBTCJPY,
			},
			wantErr: false,
		},
		{
			name: "TestBitflyerUseCase_GetBoard_No_Product_Code",
			args: args{
				productCode: "",
			},
			wantErr: false,
		},
		{
			name: "TestBitflyerUseCase_GetBoard_Invalid_Product_Code",
			args: args{
				productCode: "INVALID",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bu := NewBitflyerUseCase(UseCaseTestConfig)
			_, err := bu.GetBoard(tt.args.productCode)
			if (err != nil) != tt.wantErr {
				t.Errorf("BitflyerUseCase.GetBoard() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestBitflyerUseCase_GetTicker(t *testing.T) {
	type args struct {
		productCode string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "TestBitflyerUseCase_GetTicker",
			args: args{
				productCode: ProductCodeBTCJPY,
			},
			wantErr: false,
		},
		{
			name: "TestBitflyerUseCase_GetTicker_No_Product_Code",
			args: args{
				productCode: "",
			},
			wantErr: false,
		},
		{
			name: "TestBitflyerUseCase_GetTicker_Invalid_Product_Code",
			args: args{
				productCode: "INVALID",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bu := NewBitflyerUseCase(UseCaseTestConfig)
			_, err := bu.GetTicker(tt.args.productCode)
			if (err != nil) != tt.wantErr {
				t.Errorf("BitflyerUseCase.GetTicker() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestBitflyerUseCase_GetExecutions(t *testing.T) {
	type args struct {
		productCode string
		count       string
		before      string
		after       string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		lenGot  int
	}{
		{
			name: "TestBitflyerUseCase_GetExecutions",
			args: args{
				productCode: ProductCodeBTCJPY,
				count:       "10",
				before:      "",
				after:       "",
			},
			wantErr: false,
			lenGot:  10,
		},
		{
			name: "TestBitflyerUseCase_GetExecutions_No_Product_Code",
			args: args{
				productCode: "",
				count:       "10",
				before:      "",
				after:       "",
			},
			wantErr: false,
			lenGot:  10,
		},
		{
			name: "TestBitflyerUseCase_GetExecutions_Invalid_Product_Code",
			args: args{
				productCode: "INVALID",
				count:       "10",
				before:      "",
				after:       "",
			},
			wantErr: true,
			lenGot:  0,
		},
		{
			name: "TestBitflyerUseCase_GetExecutions_No_Count",
			args: args{
				productCode: ProductCodeBTCJPY,
				count:       "",
				before:      "",
				after:       "",
			},
			wantErr: false,
			lenGot:  100,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bu := NewBitflyerUseCase(UseCaseTestConfig)
			got, err := bu.GetExecutions(tt.args.productCode, tt.args.count, tt.args.before, tt.args.after)
			if (err != nil) != tt.wantErr {
				t.Errorf("BitflyerUseCase.GetExecutions() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(got) != tt.lenGot {
				t.Errorf("BitflyerUseCase.GetExecutions() = %v, want %v", len(got), tt.lenGot)
			}
		})
	}
}

func TestBitflyerUseCase_GetBoardState(t *testing.T) {
	type args struct {
		productCode string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "TestBitflyerUseCase_GetBoardState",
			args: args{
				productCode: ProductCodeBTCJPY,
			},
			wantErr: false,
		},
		{
			name: "TestBitflyerUseCase_GetBoardState_No_Product_Code",
			args: args{
				productCode: "",
			},
			wantErr: false,
		},
		{
			name: "TestBitflyerUseCase_GetBoardState_Invalid_Product_Code",
			args: args{
				productCode: "INVALID",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bu := NewBitflyerUseCase(UseCaseTestConfig)
			_, err := bu.GetBoardState(tt.args.productCode)
			if (err != nil) != tt.wantErr {
				t.Errorf("BitflyerUseCase.GetBoardState() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestBitflyerUseCase_GetHealth(t *testing.T) {
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
			name: "TestBitflyerUseCase_GetHealth",
			args: args{
				productCode: ProductCodeBTCJPY,
			},
			wantErr: false,
		},
		{
			name: "TestBitflyerUseCase_GetHealth_No_Product_Code",
			args: args{
				productCode: "",
			},
			wantErr: false,
		},
		{
			name: "TestBitflyerUseCase_GetHealth_Invalid_Product_Code",
			args: args{
				productCode: "INVALID",
			},
			wantErr: true,
		},
		{
			name: "TestBitflyerUseCase_GetHealth_Other_Product_Code",
			args: args{
				productCode: "ETH_JPY",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bu := NewBitflyerUseCase(UseCaseTestConfig)
			_, err := bu.GetHealth(tt.args.productCode)
			if (err != nil) != tt.wantErr {
				t.Errorf("BitflyerUseCase.GetHealth() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestBitflyerUseCase_GetBalance(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "TestBitflyerUseCase_GetBalance",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bu := NewBitflyerUseCase(UseCaseTestConfig)
			got, err := bu.GetBalance()
			if err != nil {
				if !tt.wantErr {
					t.Errorf("BitflyerUseCase.GetBalance() error = %v, wantErr %v", err, tt.wantErr)
				}
				return
			}
			if len(got) == 0 {
				t.Errorf("BitflyerUseCase.GetBalance() = %v, want not empty", got)
			}
		})
	}
}

func TestBitflyerUseCase_GetCollateral(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "TestBitflyerUseCase_GetCollateral",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bu := NewBitflyerUseCase(UseCaseTestConfig)
			_, err := bu.GetCollateral()
			if err != nil {
				if !tt.wantErr {
					t.Errorf("BitflyerUseCase.GetCollateral() error = %v, wantErr %v", err, tt.wantErr)
				}
				return
			}
		})
	}
}

func TestBitflyerUseCase_PostSendChildOrder(t *testing.T) {
	type args struct {
		productCode    string
		ChildOrderType string
		side           string
		price          int
		size           float64
		MinuteToExpire int
		TimeInForce    string
		isDry          bool
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "TestBitflyerUseCase_PostSendChildOrder",
			args: args{
				productCode:    ProductCodeBTCJPY,
				ChildOrderType: "LIMIT",
				side:           "BUY",
				price:          1000000,
				size:           0.01,
				MinuteToExpire: 1,
				TimeInForce:    "GTC",
				isDry:          true,
			},
			wantErr: false,
		},
		{
			name: "TestBitflyerUseCase_PostSendChildOrder_No_Product_Code",
			args: args{
				productCode:    "",
				ChildOrderType: "LIMIT",
				side:           "BUY",
				price:          1000000,
				size:           0.01,
				MinuteToExpire: 1,
				TimeInForce:    "GTC",
				isDry:          true,
			},
			wantErr: false,
		},
		{
			name: "TestBitflyerUseCase_PostSendChildOrder_Invalid_Product_Code",
			args: args{
				productCode:    "INVALID",
				ChildOrderType: "LIMIT",
				side:           "BUY",
				price:          1000000,
				size:           0.01,
				MinuteToExpire: 1,
				TimeInForce:    "GTC",
				isDry:          true,
			},
			wantErr: true,
		},
		{
			name: "TestBitflyerUseCase_PostSendChildOrder_No_ChildOrderType",
			args: args{
				productCode:    ProductCodeBTCJPY,
				ChildOrderType: "",
				side:           "BUY",
				price:          1000000,
				size:           0.01,
				MinuteToExpire: 1,
				TimeInForce:    "GTC",
				isDry:          true,
			},
			wantErr: true,
		},
		{
			name: "TestBitflyerUseCase_PostSendChildOrder_No_Side",
			args: args{
				productCode:    ProductCodeBTCJPY,
				ChildOrderType: "LIMIT",
				side:           "",
				price:          1000000,
				size:           0.01,
				MinuteToExpire: 1,
				TimeInForce:    "GTC",
				isDry:          true,
			},
			wantErr: true,
		},
		{
			name: "TestBitflyerUseCase_PostSendChildOrder_No_TimeInForce",
			args: args{
				productCode:    ProductCodeBTCJPY,
				ChildOrderType: "LIMIT",
				side:           "BUY",
				price:          1000000,
				size:           0.01,
				MinuteToExpire: 1,
				TimeInForce:    "",
				isDry:          true,
			},
			wantErr: true,
		},
		{
			name: "TestBitflyerUseCase_PostSendChildOrder_Invalid_TimeInForce",
			args: args{
				productCode:    ProductCodeBTCJPY,
				ChildOrderType: "LIMIT",
				side:           "BUY",
				price:          1000000,
				size:           0.01,
				MinuteToExpire: 1,
				TimeInForce:    "INVALID",
				isDry:          true,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bu := NewBitflyerUseCase(UseCaseTestConfig)
			_, err := bu.PostSendChildOrder(tt.args.productCode, tt.args.ChildOrderType, tt.args.side, tt.args.price, tt.args.size, tt.args.MinuteToExpire, tt.args.TimeInForce, tt.args.isDry)
			if (err != nil) != tt.wantErr {
				t.Errorf("BitflyerUseCase.PostSendChildOrder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestBitflyerUseCase_PostCancelChildOrder(t *testing.T) {
	type args struct {
		productCode  string
		ChildOrderID string
		isDry        bool
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "TestBitflyerUseCase_PostCancelChildOrder",
			args: args{
				productCode:  ProductCodeBTCJPY,
				ChildOrderID: "30", // テキトウな値
				isDry:        true,
			},
			wantErr: false,
		},
		{
			name: "TestBitflyerUseCase_PostCancelChildOrder_No_Product_Code",
			args: args{
				productCode:  "",
				ChildOrderID: "30", // テキトウな値
				isDry:        true,
			},
			wantErr: false,
		},
		{
			name: "TestBitflyerUseCase_PostCancelChildOrder_Invalid_Product_Code",
			args: args{
				productCode:  "INVALID",
				ChildOrderID: "30", // テキトウな値
				isDry:        true,
			},
			wantErr: true,
		},
		{
			name: "TestBitflyerUseCase_PostCancelChildOrder_No_ChildOrderID",
			args: args{
				productCode:  ProductCodeBTCJPY,
				ChildOrderID: "",
				isDry:        true,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bu := NewBitflyerUseCase(UseCaseTestConfig)
			if err := bu.PostCancelChildOrder(tt.args.productCode, tt.args.ChildOrderID, tt.args.isDry); (err != nil) != tt.wantErr {
				t.Errorf("BitflyerUseCase.PostCancelChildOrder() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestBitflyerUseCase_GetChildOrders(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "TestBitflyerUseCase_GetChildOrders",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bu := NewBitflyerUseCase(UseCaseTestConfig)
			got, err := bu.GetChildOrders()
			if err != nil {
				if !tt.wantErr {
					t.Errorf("BitflyerUseCase.GetChildOrders() error = %v, wantErr %v", err, tt.wantErr)
				}
				return
			}
			if len(got) == 0 {
				t.Errorf("BitflyerUseCase.GetChildOrders() = %v, want not empty", got)
			}
		})
	}
}

func Test_validateProductCode(t *testing.T) {
	type args struct {
		productCode string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test_validateProductCode",
			args: args{
				productCode: ProductCodeBTCJPY,
			},
			want: true,
		},
		{
			name: "Test_validateProductCode_Invalid",
			args: args{
				productCode: "INVALID",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validateProductCode(tt.args.productCode); got != tt.want {
				t.Errorf("validateProductCode() = %v, want %v", got, tt.want)
			}
		})
	}
}

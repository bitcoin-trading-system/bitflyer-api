package usecase

// TODO modelsをmock化したい

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

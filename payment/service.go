package payment

import (
	"strconv"
	"yayasuryana/user"

	midtrans "github.com/veritrans/go-midtrans"
)

type Service interface{
	GetPaymentURL(transaksi Transaksi, user user.User) (string, error)
}

type service struct{}

func NewService() *service{
	return &service{}
}

func (s *service) GetPaymentURL(transaksi Transaksi, user user.User) (string, error){
	midclient := midtrans.NewClient()
	midclient.ServerKey = "SB-Mid-server-LECMIol5nxhaFTg-PiRJU_7R"
	midclient.ClientKey = "SB-Mid-client-vmAtuGvI4OF3v1XA"
	midclient.APIEnvType = midtrans.Sandbox

	snapGateway := midtrans.SnapGateway{
		Client: midclient,
	}

	snapReq := &midtrans.SnapReq{
		CustomerDetail: &midtrans.CustDetail{
			Email: user.Email,
			FName: user.Name,
		},
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  strconv.Itoa(transaksi.ID),
			GrossAmt: int64(transaksi.Amount),
		},
	}

	snapTokenResp, err := snapGateway.GetToken(snapReq)
	if err != nil {
		return "", err
	}

	return snapTokenResp.RedirectURL, nil
	
}
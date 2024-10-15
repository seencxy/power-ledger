package mains

import (
	"PowerLedgerGo/application"
	"PowerLedgerGo/infrastructure/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type M map[string]interface{}

type MainHandler struct {
	MainService *application.MainService `inject:""`
}

type PriceDetail struct {
	Id     int64 `json:"id"`
	Price  int64 `json:"price"`
	Amount int64 `json:"amount"`
}

func (m *MainHandler) SubmitBid(ctx *gin.Context) {
	var req PriceDetail

	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.CommonResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	tx, err := m.MainService.SubmitBid(req.Price, req.Amount, req.Id)
	if err != nil {
		utils.CommonResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	utils.Success(ctx, "提交购买成功", M{"tx": tx})
}

func (m *MainHandler) SubmitOffer(ctx *gin.Context) {
	var req PriceDetail

	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.CommonResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	tx, err := m.MainService.SubmitOffer(req.Price, req.Amount, req.Id)
	if err != nil {
		utils.CommonResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	utils.Success(ctx, "提交报价成功", M{"tx": tx})
}

func (m *MainHandler) QueryTrade(ctx *gin.Context) {
	userIdString := ctx.Query("id")
	modeString := ctx.Query("mode")
	// mode 1查询购买 2查询出售的
	userId, err := strconv.Atoi(userIdString)
	if err != nil {
		utils.CommonResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	mode, err := strconv.Atoi(modeString)
	if err != nil {
		utils.CommonResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if mode != 1 && mode != 2 {
		utils.CommonResponse(ctx, http.StatusBadRequest, "mode参数只能为1或者2")
		return
	}

	tradeInfos, err := m.MainService.QueryTradeBySelf(int64(userId), int64(mode))
	if err != nil {
		utils.CommonResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	utils.Success(ctx, "查询成功", M{"trades": tradeInfos})
}

func (m *MainHandler) TradePayment(ctx *gin.Context) {
	req := struct {
		Id      int64 `json:"id"`
		TradeId int64 `json:"tradeId"`
	}{}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.CommonResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if err := m.MainService.SettleTradePayments(req.Id, req.TradeId); err != nil {
		utils.CommonResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	utils.Success(ctx, "交易成功", nil)
}

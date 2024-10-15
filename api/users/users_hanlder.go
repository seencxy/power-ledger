package users

import (
	"PowerLedgerGo/application"
	"PowerLedgerGo/domain/entity"
	"PowerLedgerGo/infrastructure/utils"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/seencxy/lsr/common"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type UserHandler struct {
	UserService *application.UserService `inject:""`
}

type M map[string]interface{}

// UserRegister 用户注册
func (u *UserHandler) UserRegister(ctx *gin.Context) {
	var err error
	decoder := json.NewDecoder(ctx.Request.Body)
	var userInfo entity.UserInfo
	if err = decoder.Decode(&userInfo); err != nil {
		utils.CommonResponse(ctx, http.StatusNotFound, err.Error())
		return
	}
	// 检查用户名是否被注册
	if _, err := u.UserService.QueryUserByUsername(userInfo.UserName); err == nil || !errors.Is(err, gorm.ErrRecordNotFound) {
		utils.CommonResponse(ctx, http.StatusNotFound, "用户名已被注册")
		return
	}

	if userInfo.Prv, userInfo.Address, err = common.CreateOneAddress(); err != nil {
		utils.CommonResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if err = u.UserService.CreateUser(&userInfo); err != nil {
		utils.CommonResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	utils.Success(ctx, "用户注册成功", M{"id": userInfo.ID})
}

// GetBalance 查询用户账户余额
func (u *UserHandler) GetBalance(ctx *gin.Context) {
	userIdString := ctx.Query("id")
	userId, err := strconv.Atoi(userIdString)
	if err != nil {
		utils.CommonResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	balance, err := u.UserService.QueryBalance(int64(userId))
	if err != nil {
		utils.CommonResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	utils.Success(ctx, "查询用户余额成功", map[string]interface{}{"balance": balance})
}

// UserRecharge 用户充值
func (u *UserHandler) UserRecharge(ctx *gin.Context) {
	reqParam := struct {
		Id    int64 `json:"id"`
		Money int64 `json:"money"`
	}{}
	err := ctx.ShouldBindBodyWithJSON(&reqParam)
	if err != nil {
		utils.CommonResponse(ctx, http.StatusBadRequest, err.Error())
	}

	err = u.UserService.Recharge(reqParam.Id, reqParam.Money)
	if err != nil {
		utils.CommonResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	utils.CommonResponse(ctx, http.StatusOK, "用户充值成功")
}

// UserWithdraw 提取余额
func (u *UserHandler) UserWithdraw(ctx *gin.Context) {
	type requestBody struct {
		ID int64 `json:"id"`
	}
	var req requestBody
	if err := ctx.BindJSON(&req); err != nil {
		utils.CommonResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if err := u.UserService.Withdraw(req.ID); err != nil {
		utils.CommonResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	utils.CommonResponse(ctx, http.StatusOK, "提取余额成功")
}

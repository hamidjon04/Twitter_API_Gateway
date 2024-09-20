package handler

import (
	"api/api/token"
	pb "api/generated/users"
	"api/model"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetClaim(c *gin.Context, h *handlerImpl) *token.Claim {
	access := c.GetHeader("Authorization")
	claim, err := token.ExtractClaimToken(access)
	if err != nil {
		h.Logger.Error(fmt.Sprintf("Claim ma'lumotlarini olishda xatolik: %v", err))
		c.JSON(http.StatusBadRequest, model.Error{
			Message: "Claim ma'lumotlarini olishda xatolik",
		})
		return nil
	}
	return claim
}

// @Summary Foydalanuvchilar ro'yxatini olish
// @Description Limit va offset query parametrlari orqali foydalanuvchilar ro'yxatini olish
// @Tags Users
// @Accept json
// @Produce json
// @Param limit query int true "Limit"
// @Param offset query int true "Offset"
// @Success 200 {object} users.GetUserRes
// @Failure 400 {object} model.Error "Noto'g'ri limit yoki offset qiymatlari"
// @Failure 500 {object} model.Error "GetUsers so'rov xatosi"
// @Router /user/getUsers [get]
// @Security ApiKeyAuth
func (h *handlerImpl) GetUsers(c *gin.Context) {
	claim := GetClaim(c, h)

	var limit, offset int
	limit, err := strconv.Atoi(c.Param("limit"))
	if err != nil {
		h.Logger.Error(fmt.Sprintf("Limit uchun noto'gri ma'lumot kiritildi: %v", err))
		c.JSON(http.StatusBadRequest, model.Error{
			Message: "Limit uchhun noto'g'ri ma'lumot kiritildi",
		})
		return
	}
	offset, err = strconv.Atoi(c.Param("offset"))
	if err != nil {
		h.Logger.Error(fmt.Sprintf("Offset uchun noto'gri ma'lumot kiritildi: %v", err))
		c.JSON(http.StatusBadRequest, model.Error{
			Message: "Offset uchhun noto'g'ri ma'lumot kiritildi",
		})
		return
	}

	resp, err := h.UserService.GetUsers(c, &pb.GetUserReq{
		Id:    claim.Id,
		Limit: int32(limit),
		Page:  int32(offset),
	})
	if err != nil {
		h.Logger.Error(fmt.Sprintf("GetUsers request error: %v", err))
		c.JSON(500, model.Error{
			Message: "GetUsers request error: %v",
		})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// @Summary Foydalanuvchini o'chirish
// @Description Foydalanuvchini ID orqali o'chirish
// @Tags Users
// @Accept json
// @Produce json
// @Param id body string true "Foydalanuvchi ID"
// @Success 200 {object} users.Massage
// @Failure 400 {object} model.Error "Ma'lumotlarni olishda xatolik"
// @Failure 500 {object} model.Error "DeleteUsers so'rov xatosi"
// @Router /user/deleteUser [delete]
// @Security ApiKeyAuth
func (h *handlerImpl) DeleteUsers(c *gin.Context) {
	var id string
	err := c.ShouldBindJSON(&id)
	if err != nil {
		h.Logger.Error(fmt.Sprintf("Ma'lumotlarni olishda xatolik: %v", err))
		c.JSON(http.StatusBadRequest, model.Error{
			Message: "Ma'lumotlarni olishda xatolik",
		})
		return
	}
	resp, err := h.UserService.DeleteUsers(c, &pb.Id{Id: id})
	if err != nil {
		h.Logger.Error(fmt.Sprintf("DeleteUsers request error: %v", err))
		c.JSON(500, model.Error{
			Message: "DeleteUsers request error",
		})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// @Summary Foydalanuvchini ID orqali olish
// @Description Foydalanuvchi ma'lumotlarini ID orqali olish
// @Tags Users
// @Accept json
// @Produce json
// @Param id path string true "Foydalanuvchi ID"
// @Success 200 {object} users.User
// @Failure 400 {object} model.Error "Noto'g'ri ID"
// @Failure 500 {object} model.Error "GetByIdUsers so'rov xatosi"
// @Router /users/getUserById [get]
// @Security ApiKeyAuth
func (h *handlerImpl) GetByIdUsers(c *gin.Context) {
	var id string
	err := c.ShouldBindJSON(&id)
	if err != nil {
		h.Logger.Error(fmt.Sprintf("Ma'lumotlarni olishda xatolik: %v", err))
		c.JSON(http.StatusBadRequest, model.Error{
			Message: "Ma'lumotlarni olishda xatolik",
		})
		return
	}

	resp, err := h.UserService.GetByIdUsers(c, &pb.Id{Id: id})
	if err != nil {
		h.Logger.Error(fmt.Sprintf("GetByIdUsers request error: %v", err))
		c.JSON(500, model.Error{
			Message: "GetByIdUsers request error",
		})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// @Summary Foydalanuvchilar ro'yxatidagi kuzatuvchilarni olish
// @Description Limit va offset query parametrlari orqali foydalanuvchi kuzatuvchilarini olish
// @Tags Users
// @Accept json
// @Produce json
// @Param limit query int true "Limit"
// @Param offset query int true "Offset"
// @Success 200 {object} users.GetaFollowersRes
// @Failure 400 {object} model.Error "Noto'g'ri limit yoki offset qiymatlari"
// @Failure 500 {object} model.Error "GetFollowers so'rov xatosi"
// @Router /user/getFollowers [get]
// @Security ApiKeyAuth
func (h *handlerImpl) GetFollowers(c *gin.Context) {
	claim := GetClaim(c, h)

	var limit, offset int
	limit, err := strconv.Atoi(c.Query("limit"))
	if err != nil {
		h.Logger.Error(fmt.Sprintf("Limit uchun noto'gri ma'lumot kiritildi: %v", err))
		c.JSON(http.StatusBadRequest, model.Error{
			Message: "Limit uchhun noto'g'ri ma'lumot kiritildi",
		})
		return
	}
	offset, err = strconv.Atoi(c.Query("offset"))
	if err != nil {
		h.Logger.Error(fmt.Sprintf("Offset uchun noto'gri ma'lumot kiritildi: %v", err))
		c.JSON(http.StatusBadRequest, model.Error{
			Message: "Offset uchhun noto'g'ri ma'lumot kiritildi",
		})
		return
	}

	resp, err := h.UserService.GetFollowers(c, &pb.GetFollowersReq{
		Id:    claim.Id,
		Limit: int32(limit),
		Page:  int32(offset),
	})
	if err != nil {
		h.Logger.Error(fmt.Sprintf("GetFollowers request error: %v", err))
		c.JSON(500, model.Error{
			Message: "GetFollowers request error",
		})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// @Summary Foydalanuvchi kuzatuvchisini o'chirish
// @Description Foydalanuvchi kuzatuvchisini o'chirish
// @Tags Users
// @Accept json
// @Produce json
// @Param followerId path string true "Kuzatuvchi ID'si"
// @Success 200 {object} users.Massage
// @Failure 400 {object} model.Error "Noto'g'ri kuzatuvchi ID'si"
// @Failure 500 {object} model.Error "DeleteFollower request error"
// @Router /user/deleteFollower [delete]
// @Security ApiKeyAuth
func (h *handlerImpl) DeleteFollower(c *gin.Context) {
	claim := GetClaim(c, h)

	var followerId string
	err := c.ShouldBindJSON(&followerId)
	if err != nil {
		h.Logger.Error(fmt.Sprintf("Ma'lumotlarni olishda xatolik: %v", err))
		c.JSON(http.StatusBadRequest, model.Error{
			Message: "ma'lumotlarni olishda xatolik",
		})
		return
	}

	resp, err := h.UserService.DeleteFollower(c, &pb.DeleteFollowerReq{
		UserId:     claim.Id,
		FollowerId: followerId,
	})
	if err != nil {
		h.Logger.Error(fmt.Sprintf("DeleteFollower request error: %v", err))
		c.JSON(500, model.Error{
			Message: "DeleteFollower request error",
		})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// @Summary Kuzatuvchini ID bo'yicha olish
// @Description Kuzatuvchini ID bo'yicha olish
// @Tags Users
// @Accept json
// @Produce json
// @Param followerId path string true "Kuzatuvchi ID'si"
// @Success 200 {object} users.Follow
// @Failure 400 {object} model.Error "Noto'g'ri kuzatuvchi ID'si"
// @Failure 500 {object} model.Error "GetByIdFollower request error"
// @Router /user/getByIdFollower [get]
// @Security ApiKeyAuth
func (h *handlerImpl) GetByIdFollower(c *gin.Context) {
	claim := GetClaim(c, h)

	var followerId string
	err := c.ShouldBindJSON(&followerId)
	if err != nil {
		h.Logger.Error(fmt.Sprintf("Ma'lumotlarni olishda xatolik: %v", err))
		c.JSON(http.StatusBadRequest, model.Error{
			Message: "ma'lumotlarni olishda xatolik",
		})
		return
	}

	resp, err := h.UserService.GetByIdFollower(c, &pb.DeleteFollowerReq{
		UserId:     claim.Id,
		FollowerId: followerId,
	})
	if err != nil {
		h.Logger.Error(fmt.Sprintf("GetByIdFollower request error: %v", err))
		c.JSON(500, model.Error{
			Message: "GetByIdFollower request error",
		})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// @Summary Kuzatuvchi qo'shish (subscribe)
// @Description Foydalanuvchi boshqa bir foydalanuvchini kuzatadi
// @Tags Users
// @Accept json
// @Produce json
// @Param data body users.FollowingReq1 true "Following ID ma'lumotlari"
// @Success 200 {object} users.Massage
// @Failure 400 {object} model.Error "Noto'g'ri kuzatilayotgan foydalanuvchi ID'si"
// @Failure 500 {object} model.Error "AddFollowing request error"
// @Router /user/subscribe [post]
// @Security ApiKeyAuth
func (h *handlerImpl) Subscribe(c *gin.Context) {
	claim := GetClaim(c, h)

	var req pb.FollowingReq1
	err := c.ShouldBindJSON(&req)
	if err != nil {
		h.Logger.Error(fmt.Sprintf("Ma'lumotlarni olishda xatolik: %v", err))
		c.JSON(http.StatusBadRequest, model.Error{
			Message: "ma'lumotlarni olishda xatolik",
		})
		return
	}

	resp, err := h.UserService.AddFollowing(c, &pb.FollowingReq{
		UserId:      claim.Id,
		FollowingId: req.FollowingId,
	})
	if err != nil {
		h.Logger.Error(fmt.Sprintf("AddFollowing request error: %v", err))
		c.JSON(500, model.Error{
			Message: "AddFollowing request error",
		})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// @Summary Foydalanuvchi kuzatayotgan odamlarni olish
// @Description Foydalanuvchi kuzatayotgan odamlarning ro'yxatini oladi
// @Tags Users
// @Accept json
// @Produce json
// @Param limit query int true "Sahifada ko'rsatiladigan elementlar soni"
// @Param offset query int true "Qaysi sahifadan boshlab olish kerakligi"
// @Success 200 {object} users.GetaFollowingRes
// @Failure 400 {object} model.Error "Limit yoki offset noto'g'ri"
// @Failure 500 {object} model.Error "GetFollowing request error"
// @Router /user/getFollowing [get]
// @Security ApiKeyAuth
func (h *handlerImpl) GetFollowing(c *gin.Context) {
	claim := GetClaim(c, h)

	// limit query parametri
	limitStr := c.Query("limit")
	if limitStr == "" {
		h.Logger.Error("Limit parametri kiritilmadi")
		c.JSON(http.StatusBadRequest, model.Error{
			Message: "Limit parametri kiritilmadi",
		})
		return
	}
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		h.Logger.Error(fmt.Sprintf("Limit uchun noto'g'ri ma'lumot kiritildi: %v", err))
		c.JSON(http.StatusBadRequest, model.Error{
			Message: "Limit uchun noto'g'ri ma'lumot kiritildi",
		})
		return
	}

	// offset query parametri
	offsetStr := c.Query("offset")
	if offsetStr == "" {
		h.Logger.Error("Offset parametri kiritilmadi")
		c.JSON(http.StatusBadRequest, model.Error{
			Message: "Offset parametri kiritilmadi",
		})
		return
	}
	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		h.Logger.Error(fmt.Sprintf("Offset uchun noto'g'ri ma'lumot kiritildi: %v", err))
		c.JSON(http.StatusBadRequest, model.Error{
			Message: "Offset uchun noto'g'ri ma'lumot kiritildi",
		})
		return
	}

	// GetFollowing requestini yuborish
	resp, err := h.UserService.GetFollowing(c, &pb.GetFollowingReq{
		Id:    claim.Id,
		Limit: int32(limit),
		Page:  int32(offset),
	})
	if err != nil {
		h.Logger.Error(fmt.Sprintf("GetFollowing request error: %v", err))
		c.JSON(http.StatusInternalServerError, model.Error{
			Message: "GetFollowing request error",
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Summary Foydalanuvchini kuzatishni bekor qilish
// @Description Foydalanuvchi kuzatayotgan boshqa foydalanuvchini kuzatishni bekor qiladi
// @Tags Users
// @Accept json
// @Produce json
// @Param followingId body string true "Bekor qilinadigan following foydalanuvchi IDsi"
// @Success 200 {object} users.Massage
// @Failure 400 {object} model.Error "Ma'lumotlarni olishda xatolik"
// @Failure 500 {object} model.Error "DeleteFollowing request error"
// @Router /user/deleteFollowing [delete]
// @Security ApiKeyAuth
func (h *handlerImpl) DeleteFollowing(c *gin.Context) {
	claim := GetClaim(c, h)

	var followingId string
	err := c.ShouldBindJSON(&followingId)
	if err != nil {
		h.Logger.Error(fmt.Sprintf("Ma'lumotlarni olishda xatolik: %v", err))
		c.JSON(http.StatusBadRequest, model.Error{
			Message: "ma'lumotlarni olishda xatolik",
		})
		return
	}

	resp, err := h.UserService.DeleteFollowing(c, &pb.DeleteFollowerReq{
		UserId:     claim.Id,
		FollowerId: followingId,
	})
	if err != nil {
		h.Logger.Error(fmt.Sprintf("DeleteFollowing request error: %v", err))
		c.JSON(500, model.Error{
			Message: "DeleteFollowing request error",
		})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// @Summary Kuzatib boruvchi foydalanuvchi haqida ma'lumot olish
// @Description Foydalanuvchi tomonidan kuzatib borilayotgan ma'lum bir foydalanuvchi haqida ma'lumot olish
// @Tags Users
// @Accept json
// @Produce json
// @Param followingId body string true "Kuzatilayotgan foydalanuvchi IDsi"
// @Success 200 {object} users.Follow
// @Failure 400 {object} model.Error "Ma'lumotlarni olishda xatolik"
// @Failure 500 {object} model.Error "GetByIdFollowing request error"
// @Router /users/getByIdFollowing [get]
// @Security ApiKeyAuth
func (h *handlerImpl) GetByIdFollowing(c *gin.Context) {
	claim := GetClaim(c, h)

	var followingId string
	err := c.ShouldBindJSON(&followingId)
	if err != nil {
		h.Logger.Error(fmt.Sprintf("Ma'lumotlarni olishda xatolik: %v", err))
		c.JSON(http.StatusBadRequest, model.Error{
			Message: "ma'lumotlarni olishda xatolik",
		})
		return
	}

	resp, err := h.UserService.GetByIdFollowing(c, &pb.DeleteFollowerReq{
		UserId:     claim.Id,
		FollowerId: followingId,
	})
	if err != nil {
		h.Logger.Error(fmt.Sprintf("GetByIdFollowing request error: %v", err))
		c.JSON(500, model.Error{
			Message: "GetByIdFollowing request error",
		})
		return
	}
	c.JSON(http.StatusOK, resp)
}

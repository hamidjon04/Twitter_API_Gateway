package handler

import (
	"api/generated/twit"
	"api/model"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Yangi twit yaratish
// @Description Foydalanuvchi tomonidan yangi twit yaratadi
// @Tags Twits
// @Accept json
// @Produce json
// @Param CreateTwitReq body twit.CreateTwitReq1 true "Twit ma'lumotlari"
// @Success 200 {object} twit.CreateTwitResp
// @Failure 400 {object} model.Error "Ma'lumotlarni olishda xatolik"
// @Failure 500 {object} model.Error "CreateTwit request error"
// @Router /twits/createTwit [post]
// @Security ApiKeyAuth
func (h *handlerImpl) CreateTwit(c *gin.Context) {
	claim := GetClaim(c, h)
	req := twit.CreateTwitReq1{}

	err := c.ShouldBindJSON(&req)
	if err != nil {
		h.Logger.Error(fmt.Sprintf("Ma'lumotlarni olishda xatolik: %v", err))
		c.JSON(http.StatusBadRequest, model.Error{
			Message: "Ma'lumotlarni olishda xatolik",
		})
		return
	}

	resp, err := h.TwitService.CreateTwit(c, &twit.CreateTwitReq{
		UserId: claim.Id,
		Content: req.Content,
		Media: req.Media,
	})
	if err != nil{
		h.Logger.Error(fmt.Sprintf("Service bilan bog'lanib bo'lmadi, CreateTwit request error: %s", err))
		c.JSON(500, model.Error{
			Message: "Error",
		})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// @Summary Mavjud twitni yangilash
// @Description Foydalanuvchi tomonidan mavjud twitni yangilaydi
// @Tags Twits
// @Accept json
// @Produce json
// @Param UpdateTwitReq body twit.UpdateReq1 true "Yangilash uchun twit ma'lumotlari"
// @Success 200 {object} twit.UpdateTwitResp
// @Failure 400 {object} model.Error "Ma'lumotlarni olishda xatolik"
// @Failure 500 {object} model.Error "UpdateTwit request error"
// @Router /twits/updateTwit [put]
// @Security ApiKeyAuth
func(h *handlerImpl) UpdateTwit(c *gin.Context){
	claim := GetClaim(c, h)
	req := twit.UpdateReq1{}

	err := c.ShouldBindJSON(&req)
	if err != nil{
		h.Logger.Error(fmt.Sprintf("Ma'lumotlarni olishda xatolik: %v", err))
		c.JSON(http.StatusBadRequest, model.Error{
			Message: "Ma'lumotlarni olishda xatolik",
		})
		return
	}

	resp, err := h.TwitService.UpdateTwit(c, &twit.UpdateReq{
		UserId: claim.Id,
		Id: req.Id,
		Content: req.Content,
		Media: req.Media,
	})
	if err != nil{
		h.Logger.Error(fmt.Sprintf("Service bilan bog'lanib bo'lmadi, UpdateTwit request error: %s", err))
		c.JSON(500, model.Error{
			Message: "Error",
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Summary Twitni o'chirish
// @Description Foydalanuvchi tomonidan yaratilgan twitni o'chiradi
// @Tags Twits
// @Produce json
// @Param id path string true "Twit ID"
// @Success 200 {object} twit.Message
// @Failure 500 {object} model.Error "DeleteTwit request error"
// @Router /twits/deleteTwit/{id} [delete]
// @Security ApiKeyAuth
func(h *handlerImpl) DeleteTwit(c *gin.Context){
	claim := GetClaim(c, h)
	id := c.Param("id")

	resp, err := h.TwitService.DeleteTwit(c, &twit.DeleteTwitReq{
		Id: id,
		UserId: claim.Id,
	})
	if err != nil{
		h.Logger.Error(fmt.Sprintf("Service bilan bog'lanib bo'lmadi, DeleteTwit request error: %s", err))
		c.JSON(500, model.Error{
			Message: "Error",
		})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// @Summary Foydalanuvchi twitlarini olish
// @Description Foydalanuvchining barcha twitlarini qaytaradi
// @Tags Twits
// @Produce json
// @Success 200 {object} twit.GetTwitsResp
// @Failure 500 {object} model.Error "GetTwits request error"
// @Router /twits/getUserTwit [get]
// @Security ApiKeyAuth
func(h *handlerImpl) GetTwits(c *gin.Context){
	claim := GetClaim(c, h)

	resp, err := h.TwitService.GetTwits(c, &twit.GetTwitsReq{UserId: claim.Id})
	if err != nil{
		h.Logger.Error(fmt.Sprintf("Service bilan bog'lanib bo'lmadi, GetUserTwits request error: %s", err))
		c.JSON(500, model.Error{
			Message: "Error",
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Summary Kuzatuvchilar twitlarini olish
// @Description Foydalanuvchining kuzatuvchilari tomonidan yaratilgan barcha twitlarni qaytaradi
// @Tags Twits
// @Produce json
// @Success 200 {object} twit.GetTwitsResp
// @Failure 500 {object} model.Error "GetFollowerTwit request error"
// @Router /twits/getFollowerTwit [get]
// @Security ApiKeyAuth
func(h *handlerImpl) GetFollowerTwit(c *gin.Context){
	claim := GetClaim(c, h)
	
	resp, err := h.TwitService.GetFollowerTwit(c, &twit.GetTwitsReq{
		UserId: claim.Id,
	})
	if err != nil{
		h.Logger.Error(fmt.Sprintf("Service bilan bog'lanib bo'lmadi, GetFollowerTwit request error: %s", err))
		c.JSON(500, model.Error{
			Message: "Error",
		})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// @Summary Twitga like qo'shish
// @Description Foydalanuvchi tomonidan tanlangan twitga like qo'shadi
// @Tags Likes
// @Accept json
// @Produce json
// @Param data body twit.AddLikeReq1 true "Twit ID ma'lumotlari"
// @Success 200 {object} twit.AddLikeResp "Like qo'shish muvaffaqiyatli bo'ldi"
// @Failure 400 {object} model.Error "Ma'lumotlarni olishda xatolik"
// @Failure 500 {object} model.Error "AddLike request error"
// @Router /twits/addLike [post]
// @Security ApiKeyAuth
func(h *handlerImpl) AddLike(c *gin.Context){
	claim := GetClaim(c, h)

	var req twit.AddLikeReq1
	err := c.ShouldBindJSON(&req)
	if err != nil{
		h.Logger.Error(fmt.Sprintf("Ma'lumotlarni olishda xatolik: %v", err))
		c.JSON(http.StatusBadRequest, model.Error{
			Message: "Ma'lumotlarni olishda xatolik",
		})
		return
	}

	resp, err := h.TwitService.AddLike(c, &twit.AddLikeReq{
		ClickerId: claim.Id,
		TwitId: req.TwitId,
	})
	if err != nil{
		h.Logger.Error(fmt.Sprintf("Service bilan bog'lanib bo'lmadi, AddLike request error: %s", err))
		c.JSON(500, model.Error{
			Message: "Error",
		})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// @Summary Twitdan like olib tashlash
// @Description Foydalanuvchi tomonidan tanlangan twitdan like olib tashlanadi
// @Tags Likes
// @Accept json
// @Produce json
// @Param data body twit.DeleteTwitReq1 true "Twit ID ma'lumotlari"
// @Success 200 {object} twit.Message "Muvaffaqiyatli olib tashlandi"
// @Failure 400 {object} model.Error "Ma'lumotlarni olishda xatolik"
// @Failure 500 {object} model.Error "RemoveLike request error"
// @Router /twits/unlike [put]
// @Security ApiKeyAuth
func(h *handlerImpl) RemoveLike(c *gin.Context){
	claim := GetClaim(c, h)

	var req twit.DeleteTwitReq1
	err := c.ShouldBindJSON(&req)
	if err != nil{
		h.Logger.Error(fmt.Sprintf("Ma'lumotlarni olishda xatolik: %v", err))
		c.JSON(http.StatusBadRequest, model.Error{
			Message: "Ma'lumotlarni olishda xatolik",
		})
		return
	}

	resp, err := h.TwitService.RemoveLike(c, &twit.DeleleLikeReq{
		ClickerId: claim.Id,
		TwitId: req.TwitId,
	})
	if err != nil{
		h.Logger.Error(fmt.Sprintf("Service bilan bog'lanib bo'lmadi, AddLike request error: %s", err))
		c.JSON(500, model.Error{
			Message: "Error",
		})
		return
	}
	c.JSON(http.StatusOK, resp)
}
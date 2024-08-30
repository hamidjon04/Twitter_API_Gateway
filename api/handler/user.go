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

func GetClaim(c *gin.Context, h *handlerImpl)*token.Claim{
	access := c.GetHeader("Authorization")
	claim, err := token.ExtractClaimToken(access)
	if err != nil{
		h.Logger.Error(fmt.Sprintf("Claim ma'lumotlarini olishda xatolik: %v", err))
		c.JSON(http.StatusBadRequest, model.Error{
			Message: "Claim ma'lumotlarini olishda xatolik",
		})
		return nil 
	}
	return claim
}

func(h *handlerImpl) GetUsers(c *gin.Context){
	claim := GetClaim(c, h)	

	var limit, offset int
	limit, err := strconv.Atoi(c.Param("limit"))
	if err != nil{
		h.Logger.Error(fmt.Sprintf("Limit uchun noto'gri ma'lumot kiritildi: %v", err))
		c.JSON(http.StatusBadRequest, model.Error{
			Message: "Limit uchhun noto'g'ri ma'lumot kiritildi",
		})
		return 
	}
	offset, err = strconv.Atoi(c.Param("offset"))
	if err != nil{
		h.Logger.Error(fmt.Sprintf("Offset uchun noto'gri ma'lumot kiritildi: %v", err))
		c.JSON(http.StatusBadRequest, model.Error{
			Message: "Offset uchhun noto'g'ri ma'lumot kiritildi",
		})
		return 
	}

	resp, err := h.UserService.GetUsers(c, &pb.GetUserReq{
		Id: claim.Id, 
		Limit: int32(limit),
		Page: int32(offset),
	})
	if err != nil{
		h.Logger.Error(fmt.Sprintf("GetUsers request error: %v", err))
		c.JSON(500, model.Error{
			Message: "GetUsers request error: %v",
		})
		return 
	}
	c.JSON(http.StatusOK, resp)
}

func(h *handlerImpl) DeleteUsers(c *gin.Context){
	var id string
	err := c.ShouldBindJSON(&id)
	if err != nil{
		h.Logger.Error(fmt.Sprintf("Ma'lumotlarni olishda xatolik: %v", err))
		c.JSON(http.StatusBadRequest, model.Error{
			Message: "Ma'lumotlarni olishda xatolik",
		})
		return 
	}
	resp, err := h.UserService.DeleteUsers(c, &pb.Id{Id: id})
	if err != nil{
		h.Logger.Error(fmt.Sprintf("DeleteUsers request error: %v", err))
		c.JSON(500, model.Error{
			Message: "DeleteUsers request error",
		})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func(h *handlerImpl) GetByIdUsers(c *gin.Context){
	var id string 
	err := c.ShouldBindJSON(&id)
	if err != nil{
		h.Logger.Error(fmt.Sprintf("Ma'lumotlarni olishda xatolik: %v", err))
		c.JSON(http.StatusBadRequest, model.Error{
			Message: "Ma'lumotlarni olishda xatolik",
		})
		return 
	}

	resp, err := h.UserService.GetByIdUsers(c, &pb.Id{Id: id})
	if err != nil{
		h.Logger.Error(fmt.Sprintf("GetByIdUsers request error: %v", err))
		c.JSON(500, model.Error{
			Message: "GetByIdUsers request error",
		})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func(h *handlerImpl) GetFollowers(c *gin.Context){
	claim := GetClaim(c, h)

	var limit, offset int
	limit, err := strconv.Atoi(c.Param("limit"))
	if err != nil{
		h.Logger.Error(fmt.Sprintf("Limit uchun noto'gri ma'lumot kiritildi: %v", err))
		c.JSON(http.StatusBadRequest, model.Error{
			Message: "Limit uchhun noto'g'ri ma'lumot kiritildi",
		})
		return 
	}
	offset, err = strconv.Atoi(c.Param("offset"))
	if err != nil{
		h.Logger.Error(fmt.Sprintf("Offset uchun noto'gri ma'lumot kiritildi: %v", err))
		c.JSON(http.StatusBadRequest, model.Error{
			Message: "Offset uchhun noto'g'ri ma'lumot kiritildi",
		})
		return 
	}

	resp, err := h.UserService.GetFollowers(c, &pb.GetFollowersReq{
		Id: claim.Id,
		Limit: int32(limit),
		Page: int32(offset),
	})
	if err != nil{
		h.Logger.Error(fmt.Sprintf("GetFollowers request error: %v", err))
		c.JSON(500, model.Error{
			Message: "GetFollowers request error",
		})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func(h *handlerImpl) DeleteFollower(c *gin.Context){
	claim := GetClaim(c, h)

	var followerId string
	err := c.ShouldBindJSON(&followerId)
	if err != nil{
		h.Logger.Error(fmt.Sprintf("Ma'lumotlarni olishda xatolik: %v", err))
		c.JSON(http.StatusBadRequest, model.Error{
			Message: "ma'lumotlarni olishda xatolik",
		})
		return
	}
	
	resp, err := h.UserService.DeleteFollower(c, &pb.DeleteFollowerReq{
		UserId: claim.Id,
		FollowerId: followerId,
	})
	if err != nil{
		h.Logger.Error(fmt.Sprintf("DeleteFollower request error: %v", err))
		c.JSON(500, model.Error{
			Message: "DeleteFollower request error",
		})
		return 
	}
	c.JSON(http.StatusOK, resp)
}

func(h *handlerImpl) GetByIdFollower(c *gin.Context){
	claim := GetClaim(c, h)

	var followerId string
	err := c.ShouldBindJSON(&followerId)
	if err != nil{
		h.Logger.Error(fmt.Sprintf("Ma'lumotlarni olishda xatolik: %v", err))
		c.JSON(http.StatusBadRequest, model.Error{
			Message: "ma'lumotlarni olishda xatolik",
		})
		return
	}

	resp, err := h.UserService.GetByIdFollower(c, &pb.DeleteFollowerReq{
		UserId: claim.Id,
		FollowerId: followerId,
	})
	if err != nil{
		h.Logger.Error(fmt.Sprintf("GetByIdFollower request error: %v", err))
		c.JSON(500, model.Error{
			Message: "GetByIdFollower request error",
		})
		return 
	}
	c.JSON(http.StatusOK, resp)
}

func(h *handlerImpl) Subscribe(c *gin.Context){
	claim := GetClaim(c, h)

	var followingId string
	err := c.ShouldBindJSON(&followingId)
	if err != nil{
		h.Logger.Error(fmt.Sprintf("Ma'lumotlarni olishda xatolik: %v", err))
		c.JSON(http.StatusBadRequest, model.Error{
			Message: "ma'lumotlarni olishda xatolik",
		})
		return
	}

	resp, err := h.UserService.AddFollowing(c, &pb.FollowingReq{
		UserId: claim.Id,
		FollowingId: followingId,
	})
	if err != nil{
		h.Logger.Error(fmt.Sprintf("AddFollowing request error: %v", err))
		c.JSON(500, model.Error{
			Message: "AddFollowing request error",
		})
		return 
	}
	c.JSON(http.StatusOK, resp)
}

func(h *handlerImpl) GetFollowing(c *gin.Context){
	claim := GetClaim(c, h)

	var limit, offset int
	limit, err := strconv.Atoi(c.Param("limit"))
	if err != nil{
		h.Logger.Error(fmt.Sprintf("Limit uchun noto'gri ma'lumot kiritildi: %v", err))
		c.JSON(http.StatusBadRequest, model.Error{
			Message: "Limit uchhun noto'g'ri ma'lumot kiritildi",
		})
		return 
	}
	offset, err = strconv.Atoi(c.Param("offset"))
	if err != nil{
		h.Logger.Error(fmt.Sprintf("Offset uchun noto'gri ma'lumot kiritildi: %v", err))
		c.JSON(http.StatusBadRequest, model.Error{
			Message: "Offset uchhun noto'g'ri ma'lumot kiritildi",
		})
		return 
	}

	resp, err := h.UserService.GetFollowing(c, &pb.GetFollowingReq{
		Id: claim.Id,
		Limit: int32(limit),
		Page: int32(offset),
	})
	if err != nil{
		h.Logger.Error(fmt.Sprintf("GetFollowing request error: %v", err))
		c.JSON(500, model.Error{
			Message: "GetFollowing request error",
		})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func(h *handlerImpl) DeleteFollowing(c *gin.Context){
	claim := GetClaim(c, h)

	var followingId string
	err := c.ShouldBindJSON(&followingId)
	if err != nil{
		h.Logger.Error(fmt.Sprintf("Ma'lumotlarni olishda xatolik: %v", err))
		c.JSON(http.StatusBadRequest, model.Error{
			Message: "ma'lumotlarni olishda xatolik",
		})
		return
	}

	resp, err := h.UserService.DeleteFollowing(c, &pb.DeleteFollowerReq{
		UserId: claim.Id,
		FollowerId: followingId,
	})
	if err != nil{
		h.Logger.Error(fmt.Sprintf("DeleteFollowing request error: %v", err))
		c.JSON(500, model.Error{
			Message: "DeleteFollowing request error",
		})
		return 
	}
	c.JSON(http.StatusOK, resp)
}

func(h *handlerImpl) GetByIdFollowing(c *gin.Context){
	claim := GetClaim(c, h)

	var followingId string
	err := c.ShouldBindJSON(&followingId)
	if err != nil{
		h.Logger.Error(fmt.Sprintf("Ma'lumotlarni olishda xatolik: %v", err))
		c.JSON(http.StatusBadRequest, model.Error{
			Message: "ma'lumotlarni olishda xatolik",
		})
		return
	}

	resp , err := h.UserService.GetByIdFollowing(c, &pb.DeleteFollowerReq{
		UserId: claim.Id,
		FollowerId: followingId,
	})
	if err != nil{
		h.Logger.Error(fmt.Sprintf("GetByIdFollowing request error: %v", err))
		c.JSON(500, model.Error{
			Message: "GetByIdFollowing request error",
		})
		return
	}
	c.JSON(http.StatusOK, resp)
}
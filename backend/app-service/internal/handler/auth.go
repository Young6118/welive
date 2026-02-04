package handler

import (
	"net/http"
	"time"

	"ai-egg/app-service/internal/config"
	"ai-egg/app-service/internal/model"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type RegisterRequest struct {
	Username string `json:"username" binding:"required,min=3,max=20"`
	Password string `json:"password" binding:"required,min=6"`
	Email    string `json:"email"`
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// Response 统一响应格式
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

var jwtSecret = []byte("your-secret-key-change-in-production")

// GenerateToken 生成JWT token
func GenerateToken(userID uint) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24 * 7).Unix(), // 7天过期
	})
	return token.SignedString(jwtSecret)
}

func Register(c *gin.Context) {
	db := config.GetDB()

	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, Response{
			Code:    400,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	// 检查用户名是否已存在
	var existingUser model.User
	if result := db.Where("username = ?", req.Username).First(&existingUser); result.Error == nil {
		c.JSON(http.StatusOK, Response{
			Code:    400,
			Message: "用户名已存在",
			Data:    nil,
		})
		return
	}

	// 加密密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusOK, Response{
			Code:    500,
			Message: "密码加密失败",
			Data:    nil,
		})
		return
	}

	// 创建用户
	user := model.User{
		Username:     req.Username,
		PasswordHash: string(hashedPassword),
		Email:        req.Email,
		Avatar:       "",
		Bio:          "",
		Status:       1,
	}

	if result := db.Create(&user); result.Error != nil {
		c.JSON(http.StatusOK, Response{
			Code:    500,
			Message: "创建用户失败",
			Data:    nil,
		})
		return
	}

	// 生成token
	token, err := GenerateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusOK, Response{
			Code:    500,
			Message: "生成token失败",
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		Code:    200,
		Message: "注册成功",
		Data: gin.H{
			"userId":  user.ID,
			"token":   token,
			"username": user.Username,
		},
	})
}

func Login(c *gin.Context) {
	db := config.GetDB()

	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, Response{
			Code:    400,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	// 查找用户
	var user model.User
	if result := db.Where("username = ?", req.Username).First(&user); result.Error != nil {
		c.JSON(http.StatusOK, Response{
			Code:    400,
			Message: "用户名或密码错误",
			Data:    nil,
		})
		return
	}

	// 验证密码
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil {
		c.JSON(http.StatusOK, Response{
			Code:    400,
			Message: "用户名或密码错误",
			Data:    nil,
		})
		return
	}

	// 检查用户状态
	if user.Status != 1 {
		c.JSON(http.StatusOK, Response{
			Code:    403,
			Message: "账号已被禁用",
			Data:    nil,
		})
		return
	}

	// 生成token
	token, err := GenerateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusOK, Response{
			Code:    500,
			Message: "生成token失败",
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		Code:    200,
		Message: "登录成功",
		Data: gin.H{
			"token":    token,
			"userId":   user.ID,
			"username": user.Username,
			"avatar":   user.Avatar,
		},
	})
}

func Logout(c *gin.Context) {
	// JWT是无状态的，客户端只需删除token即可
	c.JSON(http.StatusOK, Response{
		Code:    200,
		Message: "退出成功",
		Data:    nil,
	})
}

func CheckLogin(c *gin.Context) {
	// 从上下文中获取用户信息（由Auth中间件设置）
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusOK, Response{
			Code:    200,
			Message: "",
			Data: gin.H{
				"isLoggedIn": false,
			},
		})
		return
	}

	db := config.GetDB()
	var user model.User
	if result := db.First(&user, userID); result.Error != nil {
		c.JSON(http.StatusOK, Response{
			Code:    200,
			Message: "",
			Data: gin.H{
				"isLoggedIn": false,
			},
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		Code:    200,
		Message: "",
		Data: gin.H{
			"isLoggedIn": true,
			"userId":     user.ID,
			"username":   user.Username,
			"avatar":     user.Avatar,
		},
	})
}

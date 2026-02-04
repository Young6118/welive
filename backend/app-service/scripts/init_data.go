package main

import (
	"ai-egg/app-service/internal/config"
	"ai-egg/app-service/internal/model"
	"fmt"
	"math/rand"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func main() {
	// 加载配置
	cfg := config.Load()

	// 初始化数据库
	config.InitDB(cfg)
	db := config.GetDB()

	// 清空现有数据
	cleanData(db)

	// 初始化数据
	initUsers(db)
	initQuestions(db)
	initNotes(db)
	initVillages(db)
	initComments(db)

	fmt.Println("数据初始化完成！")
}

func cleanData(db *gorm.DB) {
	db.Exec("SET FOREIGN_KEY_CHECKS = 0")
	db.Exec("TRUNCATE TABLE comments")
	db.Exec("TRUNCATE TABLE posts")
	db.Exec("TRUNCATE TABLE villages")
	db.Exec("TRUNCATE TABLE notes")
	db.Exec("TRUNCATE TABLE questions")
	db.Exec("TRUNCATE TABLE users")
	db.Exec("SET FOREIGN_KEY_CHECKS = 1")
	fmt.Println("已清空现有数据")
}

func initUsers(db *gorm.DB) {
	users := []model.User{
		{Username: "tech_guru", Email: "tech@guru.com", Bio: "科技爱好者，分享前沿技术资讯"},
		{Username: "life_artist", Email: "life@artist.com", Bio: "记录生活中的美好瞬间"},
		{Username: "finance_wiz", Email: "finance@wiz.com", Bio: "财经分析师，投资理财顾问"},
		{Username: "emotion_healer", Email: "emotion@healer.com", Bio: "心理咨询师，倾听你的故事"},
		{Username: "code_master", Email: "code@master.com", Bio: "全栈开发者，热爱开源"},
		{Username: "traveler", Email: "travel@er.com", Bio: "背包客，走遍世界各地"},
		{Username: "foodie", Email: "food@ie.com", Bio: "美食探店达人"},
		{Username: "fitness_coach", Email: "fitness@coach.com", Bio: "健身教练，科学运动指导"},
	}

	for i := range users {
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)
		users[i].PasswordHash = string(hashedPassword)
		users[i].Avatar = fmt.Sprintf("https://api.dicebear.com/7.x/avataaars/svg?seed=%s", users[i].Username)
	}

	db.Create(&users)
	fmt.Printf("已创建 %d 个用户\n", len(users))
}

func initQuestions(db *gorm.DB) {
	questions := []model.Question{
		{
			Title:    "2024年最值得学习的编程语言是什么？",
			Content:  "随着技术的不断发展，新的编程语言层出不穷。对于想要进入IT行业或者提升技能的人来说，2024年最值得投资的编程语言有哪些？请从就业前景、薪资水平、学习难度等方面分析。",
			AuthorID: 5,
			Tags:     "科技,编程,职业发展",
			Likes:    156,
		},
		{
			Title:    "如何培养良好的理财习惯？",
			Content:  "刚工作不久，想要开始理财但不知道从何入手。请问有哪些适合新手的理财方法？如何制定合理的预算和储蓄计划？",
			AuthorID: 3,
			Tags:     "财经,理财,生活",
			Likes:    234,
		},
		{
			Title:    "异地恋如何维持感情？",
			Content:  "和女朋友异地两年了，最近感觉沟通越来越少，感情有些变淡。有没有成功维持异地恋的朋友分享一下经验？",
			AuthorID: 4,
			Tags:     "情感,恋爱,生活",
			Likes:    189,
		},
		{
			Title:    "推荐几个适合周末短途旅行的地方",
			Content:  "坐标北京，周末想出去走走， preferably 2-3小时车程。喜欢自然风光和古镇，有什么推荐吗？",
			AuthorID: 6,
			Tags:     "生活,旅行,推荐",
			Likes:    312,
		},
		{
			Title:    "AI会取代程序员吗？",
			Content:  "最近ChatGPT等AI工具越来越强大，可以写代码、debug。作为程序员有些焦虑，想听听大家的看法，AI真的会取代我们吗？",
			AuthorID: 1,
			Tags:     "科技,AI,职业",
			Likes:    567,
		},
		{
			Title:    "如何做出完美的牛排？",
			Content:  "在家尝试做牛排总是失败，要么太老要么太生。求大神分享煎牛排的技巧，包括选材、腌制、火候控制等。",
			AuthorID: 7,
			Tags:     "生活,美食,烹饪",
			Likes:    145,
		},
		{
			Title:    "健身新手应该如何制定训练计划？",
			Content:  "办了健身卡三个月了，但去健身房的次数屈指可数。不知道如何制定适合自己的训练计划，求指导！",
			AuthorID: 8,
			Tags:     "生活,健身,健康",
			Likes:    278,
		},
		{
			Title:    "有哪些提升工作效率的工具推荐？",
			Content:  "感觉每天工作效率很低，经常被打断。大家有什么好用的时间管理或效率工具推荐吗？",
			AuthorID: 2,
			Tags:     "科技,效率,工具",
			Likes:    423,
		},
	}

	for i := range questions {
		questions[i].CreatedAt = time.Now().Add(-time.Duration(rand.Intn(30)) * 24 * time.Hour)
	}

	db.Create(&questions)
	fmt.Printf("已创建 %d 个问题\n", len(questions))
}

func initNotes(db *gorm.DB) {
	notes := []model.Note{
		{
			Title:    "Go语言并发编程笔记",
			Content:  "Goroutine是Go语言的轻量级线程，由Go运行时管理。使用go关键字即可启动一个新的goroutine。Channel用于goroutine之间的通信，遵循CSP模型...",
			Category: "学习",
			Tags:     "Go,并发,编程",
			AuthorID: 5,
		},
		{
			Title:    "我的2024年阅读清单",
			Content:  "1. 《人类简史》- 尤瓦尔·赫拉利\n2. 《原则》- 瑞·达利欧\n3. 《深度工作》- 卡尔·纽波特\n4. 《思考，快与慢》- 丹尼尔·卡尼曼\n5. 《Atomic Habits》- James Clear",
			Category: "生活",
			Tags:     "阅读,书单,成长",
			AuthorID: 2,
		},
		{
			Title:    "投资理财入门指南",
			Content:  "理财的第一步是建立应急基金，建议储备3-6个月的生活费。然后是保险配置，包括重疾险、医疗险、意外险等。接下来可以考虑定投指数基金...",
			Category: "财经",
			Tags:     "理财,投资,基金",
			AuthorID: 3,
		},
		{
			Title:    "日本关西7日游攻略",
			Content:  "Day1: 大阪城公园→心斋桥→道顿堀\nDay2: 环球影城\nDay3: 奈良公园→东大寺\nDay4: 京都清水寺→二年坂三年坂\nDay5: 伏见稻荷大社→岚山\nDay6: 神户牛肉→有马温泉\nDay7: 大阪购物返程",
			Category: "旅行",
			Tags:     "日本,旅行,攻略",
			AuthorID: 6,
		},
		{
			Title:    "ChatGPT使用技巧总结",
			Content:  "1. 角色扮演：让AI扮演特定角色回答问题\n2. 分步骤提问：复杂问题拆解成多个小问题\n3. 提供示例：给AI参考示例提高回答质量\n4. 迭代优化：根据回答不断调整提示词",
			Category: "科技",
			Tags:     "AI,ChatGPT,效率",
			AuthorID: 1,
		},
		{
			Title:    "30天减脂计划",
			Content:  "Week1: 适应期，每天30分钟有氧\nWeek2: 增加力量训练，每周3次\nWeek3: 提高强度，加入HIIT\nWeek4: 综合训练，注意饮食控制\n\n饮食原则：高蛋白、低碳水、适量脂肪",
			Category: "健康",
			Tags:     "健身,减脂,计划",
			AuthorID: 8,
		},
		{
			Title:    "摄影入门：构图技巧",
			Content:  "三分法则：将画面分成九宫格，主体放在交点处\n对称构图：适合建筑、倒影等\n引导线：利用道路、河流等引导视线\n框架构图：利用门窗等作为前景框架",
			Category: "兴趣",
			Tags:     "摄影,构图,技巧",
			AuthorID: 2,
		},
		{
			Title:    "情绪管理的5个方法",
			Content:  "1. 觉察情绪：识别自己当下的情绪状态\n2. 接纳情绪：允许自己有负面情绪，不抗拒\n3. 表达情绪：通过写日记、倾诉等方式释放\n4. 转移注意力：运动、听音乐等\n5. 认知重构：换个角度看问题",
			Category: "心理",
			Tags:     "情绪,心理,健康",
			AuthorID: 4,
		},
	}

	for i := range notes {
		notes[i].CreatedAt = time.Now().Add(-time.Duration(rand.Intn(60)) * 24 * time.Hour)
	}

	db.Create(&notes)
	fmt.Printf("已创建 %d 个笔记\n", len(notes))
}

func initVillages(db *gorm.DB) {
	villages := []model.Village{
		{
			Name:        "科技前沿",
			Description: "探讨最新科技趋势，分享技术干货",
			MemberCount: 12580,
		},
		{
			Name:        "理财投资",
			Description: "交流投资理财心得，实现财富增长",
			MemberCount: 8932,
		},
		{
			Name:        "情感树洞",
			Description: "倾诉情感困惑，分享恋爱经验",
			MemberCount: 15670,
		},
		{
			Name:        "旅行日记",
			Description: "分享旅行故事，推荐旅游目的地",
			MemberCount: 6789,
		},
		{
			Name:        "美食天地",
			Description: "分享美食制作，推荐餐厅美食",
			MemberCount: 23450,
		},
		{
			Name:        "健身打卡",
			Description: "记录健身历程，分享运动经验",
			MemberCount: 9876,
		},
	}

	db.Create(&villages)
	fmt.Printf("已创建 %d 个村落\n", len(villages))

	// 为每个村落创建一些帖子
	var posts []model.Post
	postContents := []string{
		"今天学到了很多新知识，分享给大家！",
		"有个问题想请教大家，希望能得到帮助。",
		"最近的一些心得体会，记录一下。",
		"推荐一个很好用的工具/方法。",
		"大家对这个话题怎么看？欢迎讨论。",
	}

	for _, village := range villages {
		for i := 0; i < 5; i++ {
			posts = append(posts, model.Post{
				Content:   postContents[i],
				AuthorID:  uint(rand.Intn(8) + 1),
				VillageID: village.ID,
				Likes:     rand.Intn(100),
			})
		}
	}

	db.Create(&posts)
	fmt.Printf("已创建 %d 个帖子\n", len(posts))
}

func initComments(db *gorm.DB) {
	comments := []model.Comment{
		{
			Content:    "非常有用的分享，感谢！",
			AuthorID:   2,
			TargetID:   1,
			TargetType: "question",
			Likes:      23,
		},
		{
			Content:    "我也遇到过类似的问题，后来通过...解决了",
			AuthorID:   3,
			TargetID:   1,
			TargetType: "question",
			Likes:      45,
		},
		{
			Content:    "说得很对，补充一点...",
			AuthorID:   4,
			TargetID:   2,
			TargetType: "question",
			Likes:      12,
		},
		{
			Content:    "收藏了，慢慢看",
			AuthorID:   1,
			TargetID:   1,
			TargetType: "note",
			Likes:      8,
		},
		{
			Content:    "写得太好了，学到了很多",
			AuthorID:   5,
			TargetID:   2,
			TargetType: "note",
			Likes:      34,
		},
	}

	db.Create(&comments)
	fmt.Printf("已创建 %d 个评论\n", len(comments))
}

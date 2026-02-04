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
	users := initUsers(db)
	questions := initQuestions(db, users)
	initAnswers(db, users, questions)
	initQuestionLikes(db, users, questions)
	notes := initNotes(db, users)
	initNoteLikes(db, users, notes)
	villages := initVillages(db)
	posts := initPosts(db, users, villages)
	initPostLikes(db, users, posts)
	initPostReplies(db, users, posts)
	initComments(db, users, questions, notes)
	initCommentLikes(db)
	initChats(db, users)

	fmt.Println("数据初始化完成！")
}

func cleanData(db *gorm.DB) {
	db.Exec("SET FOREIGN_KEY_CHECKS = 0")
	db.Exec("TRUNCATE TABLE comment_likes")
	db.Exec("TRUNCATE TABLE comments")
	db.Exec("TRUNCATE TABLE post_likes")
	db.Exec("TRUNCATE TABLE messages")
	db.Exec("TRUNCATE TABLE chats")
	db.Exec("TRUNCATE TABLE posts")
	db.Exec("TRUNCATE TABLE village_members")
	db.Exec("TRUNCATE TABLE villages")
	db.Exec("TRUNCATE TABLE note_likes")
	db.Exec("TRUNCATE TABLE notes")
	db.Exec("TRUNCATE TABLE question_likes")
	db.Exec("TRUNCATE TABLE answers")
	db.Exec("TRUNCATE TABLE questions")
	db.Exec("TRUNCATE TABLE users")
	db.Exec("SET FOREIGN_KEY_CHECKS = 1")
	fmt.Println("已清空现有数据")
}

// 专业领域定义
type UserProfile struct {
	Username    string
	Email       string
	Bio         string
	Avatar      string
	Field       string
	Specialties []string
}

func initUsers(db *gorm.DB) []model.User {
	profiles := []UserProfile{
		// 科技领域
		{Username: "tech_guru", Email: "tech@guru.com", Bio: "资深软件架构师，10年+开发经验，专注云原生和微服务", Field: "科技", Specialties: []string{"Go", "Kubernetes", "微服务"}},
		{Username: "ai_researcher", Email: "ai@research.com", Bio: "AI研究员，深耕机器学习和深度学习领域", Field: "科技", Specialties: []string{"机器学习", "深度学习", "NLP"}},
		{Username: "frontend_master", Email: "frontend@master.com", Bio: "前端专家，Vue/React双栈开发者", Field: "科技", Specialties: []string{"Vue", "React", "TypeScript"}},
		{Username: "security_expert", Email: "security@expert.com", Bio: "网络安全专家，白帽黑客，漏洞挖掘", Field: "科技", Specialties: []string{"网络安全", "渗透测试", "密码学"}},
		{Username: "data_engineer", Email: "data@engineer.com", Bio: "大数据工程师，数据仓库和实时计算", Field: "科技", Specialties: []string{"大数据", "Spark", "Flink"}},

		// 财经领域
		{Username: "finance_wiz", Email: "finance@wiz.com", Bio: "CFA持证人，资深投资分析师，专注价值投资", Field: "财经", Specialties: []string{"股票", "基金", "财报分析"}},
		{Username: "crypto_trader", Email: "crypto@trader.com", Bio: "加密货币交易员，区块链早期投资者", Field: "财经", Specialties: []string{"比特币", "DeFi", "NFT"}},
		{Username: "realestate_pro", Email: "realestate@pro.com", Bio: "房地产投资专家，帮助100+家庭实现资产配置", Field: "财经", Specialties: []string{"房产", "资产配置", "税务规划"}},
		{Username: "tax_advisor", Email: "tax@advisor.com", Bio: "注册税务师，个人和企业税务规划", Field: "财经", Specialties: []string{"税务", "保险", "退休规划"}},

		// 医疗健康
		{Username: "dr_health", Email: "dr@health.com", Bio: "三甲医院主治医师，专注内科疾病诊疗", Field: "医疗", Specialties: []string{"内科", "健康管理", "慢病"}},
		{Username: "nutritionist", Email: "nutrition@ist.com", Bio: "注册营养师，科学饮食指导", Field: "医疗", Specialties: []string{"营养", "减重", "运动营养"}},
		{Username: "psychologist", Email: "psych@ologist.com", Bio: "国家二级心理咨询师，认知行为疗法", Field: "医疗", Specialties: []string{"心理咨询", "情绪管理", "亲子关系"}},
		{Username: "fitness_coach", Email: "fitness@coach.com", Bio: "ACE认证健身教练，体能训练专家", Field: "医疗", Specialties: []string{"增肌", "减脂", "体能训练"}},
		{Username: "yoga_instructor", Email: "yoga@instructor.com", Bio: "瑜伽导师，RYT500认证，身心平衡", Field: "医疗", Specialties: []string{"瑜伽", "冥想", "普拉提"}},

		// 教育学习
		{Username: "english_teacher", Email: "english@teacher.com", Bio: "雅思8.5分，10年英语教学经验", Field: "教育", Specialties: []string{"雅思", "托福", "商务英语"}},
		{Username: "math_tutor", Email: "math@tutor.com", Bio: "数学竞赛教练，奥赛金牌导师", Field: "教育", Specialties: []string{"数学", "物理", "竞赛"}},
		{Username: "career_mentor", Email: "career@mentor.com", Bio: "职业规划师，帮助500+职场人转型", Field: "教育", Specialties: []string{"职业规划", "面试", "简历"}},
		{Username: "study_guru", Email: "study@guru.com", Bio: "学习方法论专家，高效学习技巧", Field: "教育", Specialties: []string{"记忆法", "时间管理", "考试技巧"}},

		// 生活方式
		{Username: "life_artist", Email: "life@artist.com", Bio: "生活美学博主，极简主义践行者", Field: "生活", Specialties: []string{"收纳", "家居", "极简"}},
		{Username: "traveler", Email: "travel@er.com", Bio: "环球旅行家，已到访50+国家", Field: "生活", Specialties: []string{"自由行", "摄影", "签证攻略"}},
		{Username: "foodie", Email: "food@ie.com", Bio: "米其林餐厅品鉴师，美食评论家", Field: "生活", Specialties: []string{"中餐", "西餐", "烘焙"}},
		{Username: "fashion_blogger", Email: "fashion@blogger.com", Bio: "时尚博主，穿搭灵感分享", Field: "生活", Specialties: []string{"穿搭", "美妆", "护肤"}},
		{Username: "diy_crafter", Email: "diy@crafter.com", Bio: "手工达人，DIY创意制作", Field: "生活", Specialties: []string{"手工", "木工", "园艺"}},

		// 职场发展
		{Username: "hr_director", Email: "hr@director.com", Bio: "500强企业HRD，人才发展专家", Field: "职场", Specialties: []string{"招聘", "绩效", "团队管理"}},
		{Username: "product_manager", Email: "pm@manager.com", Bio: "资深产品经理，多款百万级产品负责人", Field: "职场", Specialties: []string{"产品设计", "数据分析", "用户研究"}},
		{Username: "sales_champion", Email: "sales@champion.com", Bio: "销冠导师，销售心理学专家", Field: "职场", Specialties: []string{"销售技巧", "谈判", "客户管理"}},
		{Username: "startup_founder", Email: "startup@founder.com", Bio: "连续创业者，两家独角兽公司联合创始人", Field: "职场", Specialties: []string{"创业", "融资", "商业模式"}},

		// 情感关系
		{Username: "emotion_healer", Email: "emotion@healer.com", Bio: "婚姻家庭咨询师，情感修复专家", Field: "情感", Specialties: []string{"婚姻", "恋爱", "沟通技巧"}},
		{Username: "parenting_expert", Email: "parenting@expert.com", Bio: "亲子教育专家，正面管教讲师", Field: "情感", Specialties: []string{"育儿", "早教", "亲子关系"}},
		{Username: "relationship_coach", Email: "relationship@coach.com", Bio: "人际关系导师，社交技巧提升", Field: "情感", Specialties: []string{"社交", "情商", "人际沟通"}},

		// 法律
		{Username: "lawyer_pro", Email: "lawyer@pro.com", Bio: "执业律师，专注民商事诉讼", Field: "法律", Specialties: []string{"合同法", "劳动法", "知识产权"}},
		{Username: "ip_attorney", Email: "ip@attorney.com", Bio: "知识产权律师，专利代理师", Field: "法律", Specialties: []string{"专利", "商标", "著作权"}},

		// 设计创意
		{Username: "ui_designer", Email: "ui@designer.com", Bio: "UI/UX设计师，设计系统专家", Field: "设计", Specialties: []string{"UI设计", "交互设计", "Figma"}},
		{Username: "photography_pro", Email: "photo@pro.com", Bio: "专业摄影师，国家地理供稿人", Field: "设计", Specialties: []string{"风光摄影", "人像", "后期"}},
		{Username: "video_creator", Email: "video@creator.com", Bio: "视频创作者，B站百大UP主", Field: "设计", Specialties: []string{"剪辑", "调色", "内容创作"}},
	}

	var users []model.User
	for _, profile := range profiles {
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)
		user := model.User{
			Username:     profile.Username,
			Email:        profile.Email,
			Bio:          profile.Bio,
			PasswordHash: string(hashedPassword),
			Avatar:       fmt.Sprintf("https://api.dicebear.com/7.x/avataaars/svg?seed=%s", profile.Username),
		}
		users = append(users, user)
	}

	db.Create(&users)
	fmt.Printf("已创建 %d 个专业领域用户\n", len(users))
	return users
}

// 问题模板
type QuestionTemplate struct {
	Title   string
	Content string
	Tags    string
	Field   string
}

func initQuestions(db *gorm.DB, users []model.User) []model.Question {
	templates := []QuestionTemplate{
		// 科技类
		{Title: "2024年最值得学习的编程语言是什么？", Content: "随着技术的不断发展，新的编程语言层出不穷。对于想要进入IT行业或者提升技能的人来说，2024年最值得投资的编程语言有哪些？请从就业前景、薪资水平、学习难度等方面分析。", Tags: "科技,编程,职业发展", Field: "科技"},
		{Title: "Kubernetes和Docker如何配合使用？", Content: "刚接触容器化技术，对K8s和Docker的关系有些困惑。能否详细解释一下两者的区别和联系？在实际项目中应该如何配合使用？", Tags: "科技,云原生,DevOps", Field: "科技"},
		{Title: "AI会取代程序员吗？", Content: "最近ChatGPT等AI工具越来越强大，可以写代码、debug。作为程序员有些焦虑，想听听大家的看法，AI真的会取代我们吗？", Tags: "科技,AI,职业", Field: "科技"},
		{Title: "前端性能优化有哪些最佳实践？", Content: "网站加载速度很慢，用户体验不好。请问前端性能优化应该从哪些方面入手？有没有推荐的工具和方法？", Tags: "科技,前端,性能", Field: "科技"},
		{Title: "如何防范SQL注入攻击？", Content: "公司最近做安全审计，发现了一些SQL注入漏洞。想请教一下防范SQL注入的最佳实践，以及常用的安全测试工具。", Tags: "科技,安全,数据库", Field: "科技"},
		{Title: "大数据处理选Spark还是Flink？", Content: "公司要做实时数据处理，技术选型在Spark和Flink之间犹豫。两者的优缺点是什么？分别适合什么场景？", Tags: "科技,大数据,实时计算", Field: "科技"},

		// 财经类
		{Title: "如何培养良好的理财习惯？", Content: "刚工作不久，想要开始理财但不知道从何入手。请问有哪些适合新手的理财方法？如何制定合理的预算和储蓄计划？", Tags: "财经,理财,生活", Field: "财经"},
		{Title: "基金定投真的能赚钱吗？", Content: "听说基金定投是懒人理财法，但市场波动这么大，定投真的能赚钱吗？应该怎么选择基金和制定定投策略？", Tags: "财经,基金,投资", Field: "财经"},
		{Title: "比特币还值得投资吗？", Content: "看着比特币价格起起伏伏，想了解一下加密货币的投资价值。现在入场还来得及吗？风险有多大？", Tags: "财经,加密货币,投资", Field: "财经"},
		{Title: "2024年是买房的好时机吗？", Content: "房价走势不明朗，刚需购房者很纠结。从投资角度分析，现在适合买房吗？应该关注哪些因素？", Tags: "财经,房产,投资", Field: "财经"},
		{Title: "个人所得税如何合理避税？", Content: "年收入增加了，税负也越来越重。想了解一下合法合规的税务优化方法，有哪些专项附加扣除可以用？", Tags: "财经,税务,规划", Field: "财经"},

		// 医疗类
		{Title: "如何改善睡眠质量？", Content: "经常失眠，晚上睡不着早上起不来。有什么科学的方法可以改善睡眠质量吗？褪黑素可以吃吗？", Tags: "医疗,睡眠,健康", Field: "医疗"},
		{Title: "减肥期间应该怎么吃？", Content: "想减肥但不想节食，怕反弹。请问科学的减脂饮食应该怎么安排？三大营养素的比例是多少？", Tags: "医疗,营养,减重", Field: "医疗"},
		{Title: "焦虑症如何自我调节？", Content: "最近工作压力大，总是焦虑不安，有时候还会心悸。不想吃药，有什么自我调节的方法吗？", Tags: "医疗,心理,焦虑", Field: "医疗"},
		{Title: "健身新手应该如何制定训练计划？", Content: "办了健身卡三个月了，但去健身房的次数屈指可数。不知道如何制定适合自己的训练计划，求指导！", Tags: "医疗,健身,健康", Field: "医疗"},
		{Title: "瑜伽和普拉提有什么区别？", Content: "想开始练习，在瑜伽和普拉提之间纠结。两者的区别是什么？分别适合什么人群？", Tags: "医疗,瑜伽,运动", Field: "医疗"},

		// 教育类
		{Title: "雅思口语如何快速提分？", Content: "雅思口语一直是弱项，考试的时候总是紧张。有什么练习方法和应试技巧可以快速提高口语分数？", Tags: "教育,雅思,英语", Field: "教育"},
		{Title: "孩子数学成绩不好怎么办？", Content: "孩子小学三年级，数学总是跟不上。作为家长应该怎么辅导？要不要报补习班？", Tags: "教育,数学,育儿", Field: "教育"},
		{Title: "想转行做产品经理，应该怎么准备？", Content: "做了3年开发，想转行做产品经理。需要学习什么技能？如何准备面试？", Tags: "教育,职业,转型", Field: "教育"},
		{Title: "如何提高记忆力？", Content: "感觉记忆力越来越差，刚看过的书就忘了。有什么科学的记忆方法可以提高学习效率？", Tags: "教育,记忆,学习", Field: "教育"},

		// 生活类
		{Title: "推荐几个适合周末短途旅行的地方", Content: "坐标北京，周末想出去走走，preferably 2-3小时车程。喜欢自然风光和古镇，有什么推荐吗？", Tags: "生活,旅行,推荐", Field: "生活"},
		{Title: "如何做出完美的牛排？", Content: "在家尝试做牛排总是失败，要么太老要么太生。求大神分享煎牛排的技巧，包括选材、腌制、火候控制等。", Tags: "生活,美食,烹饪", Field: "生活"},
		{Title: "小户型如何最大化利用空间？", Content: "买了套60平的小两居，感觉空间不够用。有什么收纳技巧和家具推荐可以让空间显得更大？", Tags: "生活,家居,收纳", Field: "生活"},
		{Title: "有哪些提升工作效率的工具推荐？", Content: "感觉每天工作效率很低，经常被打断。大家有什么好用的时间管理或效率工具推荐吗？", Tags: "生活,效率,工具", Field: "生活"},
		{Title: "新手如何入门摄影？", Content: "刚买了相机，想学习摄影。应该从哪些方面入手？有什么推荐的教程或书籍？", Tags: "生活,摄影,入门", Field: "生活"},

		// 职场类
		{Title: "面试时如何谈薪资？", Content: "拿到了offer但薪资不满意，想谈高一些。有什么谈判技巧和话术可以参考？", Tags: "职场,面试,薪资", Field: "职场"},
		{Title: "如何管理团队中的刺头员工？", Content: "团队里有个能力不错但态度很差的员工，经常不服从管理。作为管理者应该怎么处理？", Tags: "职场,管理,团队", Field: "职场"},
		{Title: "销售新人如何快速开单？", Content: "刚做销售一个月，还没有开单，压力很大。有什么快速成交的技巧和方法？", Tags: "职场,销售,技巧", Field: "职场"},
		{Title: "创业需要准备什么？", Content: "有个创业想法，但不知道从何开始。创业前需要做哪些准备？如何找到合伙人？", Tags: "职场,创业,融资", Field: "职场"},

		// 情感类
		{Title: "异地恋如何维持感情？", Content: "和女朋友异地两年了，最近感觉沟通越来越少，感情有些变淡。有没有成功维持异地恋的朋友分享一下经验？", Tags: "情感,恋爱,生活", Field: "情感"},
		{Title: "如何跟青春期的孩子沟通？", Content: "孩子14岁，越来越叛逆，说啥都不听。作为父母应该怎么跟青春期的孩子有效沟通？", Tags: "情感,育儿,亲子", Field: "情感"},
		{Title: "社交恐惧症怎么克服？", Content: "很害怕社交场合，一到人多的地方就紧张。有什么方法可以克服社恐，提升社交能力？", Tags: "情感,社交,心理", Field: "情感"},
		{Title: "婚后如何保持新鲜感？", Content: "结婚5年了，感觉生活越来越平淡。夫妻之间应该如何保持新鲜感和激情？", Tags: "情感,婚姻,关系", Field: "情感"},

		// 法律类
		{Title: "租房合同要注意哪些条款？", Content: "第一次租房，看了合同有很多不懂的地方。租房合同中有哪些坑需要注意？", Tags: "法律,租房,合同", Field: "法律"},
		{Title: "被公司违法辞退怎么维权？", Content: "公司以业绩不达标为由要辞退我，但没有提前通知也没有赔偿。这种情况应该怎么维权？", Tags: "法律,劳动,维权", Field: "法律"},
		{Title: "如何保护自己的知识产权？", Content: "开发了一个小程序，担心被人抄袭。应该如何保护自己的知识产权？需要申请专利吗？", Tags: "法律,知识产权,专利", Field: "法律"},

		// 设计类
		{Title: "UI设计师如何提升审美？", Content: "做UI设计两年了，感觉遇到瓶颈，作品缺乏美感。应该如何提升审美和设计水平？", Tags: "设计,UI,审美", Field: "设计"},
		{Title: "风光摄影如何构图？", Content: "喜欢拍风景，但照片总是平平无奇。风光摄影有哪些经典的构图技巧？", Tags: "设计,摄影,构图", Field: "设计"},
		{Title: "短视频剪辑有什么技巧？", Content: "开始做短视频，但剪辑出来的视频不够吸引人。有什么剪辑技巧和节奏把控的方法？", Tags: "设计,视频,剪辑", Field: "设计"},
	}

	var questions []model.Question
	for _, template := range templates {
		// 找到对应领域的用户
		var authorID uint
		for _, user := range users {
			if contains(getUserField(user.Username), template.Field) {
				authorID = user.ID
				break
			}
		}
		if authorID == 0 {
			authorID = users[rand.Intn(len(users))].ID
		}

		question := model.Question{
			Title:    template.Title,
			Content:  template.Content,
			AuthorID: authorID,
			Tags:     template.Tags,
			Likes:    rand.Intn(500) + 50,
			Views:    rand.Intn(5000) + 500,
			Status:   1,
		}
		question.CreatedAt = time.Now().Add(-time.Duration(rand.Intn(60)) * 24 * time.Hour)
		questions = append(questions, question)
	}

	db.Create(&questions)
	fmt.Printf("已创建 %d 个问题\n", len(questions))
	return questions
}

func initAnswers(db *gorm.DB, users []model.User, questions []model.Question) {
	answerTemplates := []string{
		"作为一个有多年经验的专业人士，我认为这个问题需要从多个角度来看。首先...其次...最后...",
		"这个问题问得很好，我来分享一下我的看法。根据我的经验...",
		"我遇到过类似的情况，当时的解决方案是...希望对你有帮助！",
		"从专业角度分析，主要有以下几点：1. ... 2. ... 3. ...",
		"补充一下前面几位答主的观点，还有一点很重要...",
		"这个问题比较复杂，我尽量说得详细一些...",
		"个人观点，仅供参考。我认为关键在于...",
		"实战经验丰富的人来答一波。我们团队的做法是...",
	}

	var answers []model.Answer
	for _, question := range questions {
		// 每个问题2-5个回答
		answerCount := rand.Intn(4) + 2
		for i := 0; i < answerCount; i++ {
			content := answerTemplates[rand.Intn(len(answerTemplates))]
			authorID := users[rand.Intn(len(users))].ID
			// 避免作者自己回答
			if authorID == question.AuthorID {
				authorID = users[(rand.Intn(len(users)-1)+1)%len(users)].ID
			}

			answer := model.Answer{
				QuestionID: question.ID,
				Content:    content,
				AuthorID:   authorID,
				Likes:      rand.Intn(100) + 10,
				IsAI:       rand.Float32() < 0.1, // 10%是AI回答
				Status:     1,
			}
			answer.CreatedAt = time.Now().Add(-time.Duration(rand.Intn(30)) * 24 * time.Hour)
			answers = append(answers, answer)
		}
	}

	db.Create(&answers)
	fmt.Printf("已创建 %d 个回答\n", len(answers))
}

func initQuestionLikes(db *gorm.DB, users []model.User, questions []model.Question) {
	var likes []model.QuestionLike
	likeMap := make(map[string]bool)

	for _, question := range questions {
		// 每个问题10-50个点赞
		likeCount := rand.Intn(40) + 10
		for i := 0; i < likeCount; i++ {
			userID := users[rand.Intn(len(users))].ID
			key := fmt.Sprintf("%d-%d", question.ID, userID)
			if !likeMap[key] && userID != question.AuthorID {
				likeMap[key] = true
				likes = append(likes, model.QuestionLike{
					QuestionID: question.ID,
					UserID:     userID,
				})
			}
		}
	}

	db.Create(&likes)
	fmt.Printf("已创建 %d 个问题点赞\n", len(likes))
}

func initNotes(db *gorm.DB, users []model.User) []model.Note {
	noteTemplates := []struct {
		Title    string
		Content  string
		Category string
		Tags     string
		Field    string
	}{
		{Title: "Go语言并发编程笔记", Content: "Goroutine是Go语言的轻量级线程，由Go运行时管理。使用go关键字即可启动一个新的goroutine。Channel用于goroutine之间的通信，遵循CSP模型...\n\n关键概念：\n1. Goroutine - 轻量级线程\n2. Channel - 线程安全通信\n3. Select - 多路复用\n4. Mutex - 互斥锁\n5. WaitGroup - 等待组", Category: "学习", Tags: "Go,并发,编程", Field: "科技"},
		{Title: "Kubernetes入门指南", Content: "K8s核心概念解析：\n\nPod - 最小部署单元\nService - 服务发现和负载均衡\nDeployment - 声明式部署\nConfigMap/Secret - 配置管理\nIngress - 外部访问\n\n常用命令：\nkubectl get pods\nkubectl apply -f deployment.yaml\nkubectl logs <pod-name>", Category: "学习", Tags: "K8s,云原生,DevOps", Field: "科技"},
		{Title: "机器学习算法总结", Content: "监督学习：\n- 线性回归\n- 逻辑回归\n- 决策树\n- 随机森林\n- SVM\n- 神经网络\n\n无监督学习：\n- K-means\n- 层次聚类\n- PCA\n\n强化学习：\n- Q-learning\n- Policy Gradient", Category: "学习", Tags: "机器学习,AI,算法", Field: "科技"},
		{Title: "React Hooks最佳实践", Content: "useState - 状态管理\nuseEffect - 副作用处理\nuseContext - 上下文\nuseReducer - 复杂状态\nuseMemo - 性能优化\nuseCallback - 回调优化\nuseRef - DOM引用\n\n自定义Hooks封装技巧...", Category: "学习", Tags: "React,前端,Hooks", Field: "科技"},
		{Title: "Web安全攻防手册", Content: "常见攻击类型：\n1. SQL注入\n2. XSS跨站脚本\n3. CSRF跨站请求伪造\n4. 文件上传漏洞\n5. 越权访问\n\n防御措施：\n- 参数化查询\n- 输入验证\n- 输出编码\n- CSRF Token\n- 权限控制", Category: "学习", Tags: "安全,Web,攻防", Field: "科技"},

		{Title: "我的2024年阅读清单", Content: "1. 《人类简史》- 尤瓦尔·赫拉利\n2. 《原则》- 瑞·达利欧\n3. 《深度工作》- 卡尔·纽波特\n4. 《思考，快与慢》- 丹尼尔·卡尼曼\n5. 《Atomic Habits》- James Clear\n6. 《纳瓦尔宝典》\n7. 《穷查理宝典》\n8. 《反脆弱》", Category: "生活", Tags: "阅读,书单,成长", Field: "生活"},
		{Title: "日本关西7日游攻略", Content: "Day1: 大阪城公园→心斋桥→道顿堀\nDay2: 环球影城\nDay3: 奈良公园→东大寺\nDay4: 京都清水寺→二年坂三年坂\nDay5: 伏见稻荷大社→岚山\nDay6: 神户牛肉→有马温泉\nDay7: 大阪购物返程\n\n预算：人均8000元\n交通：ICOCA卡+JR Pass", Category: "旅行", Tags: "日本,旅行,攻略", Field: "生活"},
		{Title: "ChatGPT使用技巧总结", Content: "1. 角色扮演：让AI扮演特定角色回答问题\n2. 分步骤提问：复杂问题拆解成多个小问题\n3. 提供示例：给AI参考示例提高回答质量\n4. 迭代优化：根据回答不断调整提示词\n5. 限定格式：要求特定输出格式\n6. 设定约束：明确限制条件\n\nPrompt工程实践...", Category: "科技", Tags: "AI,ChatGPT,效率", Field: "科技"},
		{Title: "30天减脂计划", Content: "Week1: 适应期，每天30分钟有氧\nWeek2: 增加力量训练，每周3次\nWeek3: 提高强度，加入HIIT\nWeek4: 综合训练，注意饮食控制\n\n饮食原则：\n- 高蛋白（每公斤体重1.5-2g）\n- 低碳水（选择低GI食物）\n- 适量健康脂肪\n- 多喝水（每天2L+）", Category: "健康", Tags: "健身,减脂,计划", Field: "医疗"},
		{Title: "摄影入门：构图技巧", Content: "三分法则：将画面分成九宫格，主体放在交点处\n对称构图：适合建筑、倒影等\n引导线：利用道路、河流等引导视线\n框架构图：利用门窗等作为前景框架\n留白：适当留白增加意境\n\n实战练习建议...", Category: "兴趣", Tags: "摄影,构图,技巧", Field: "设计"},
		{Title: "情绪管理的5个方法", Content: "1. 觉察情绪：识别自己当下的情绪状态\n2. 接纳情绪：允许自己有负面情绪，不抗拒\n3. 表达情绪：通过写日记、倾诉等方式释放\n4. 转移注意力：运动、听音乐等\n5. 认知重构：换个角度看问题\n\n推荐书籍：《情绪急救》", Category: "心理", Tags: "情绪,心理,健康", Field: "医疗"},

		{Title: "投资理财入门指南", Content: "理财的第一步是建立应急基金，建议储备3-6个月的生活费。然后是保险配置，包括重疾险、医疗险、意外险等。接下来可以考虑定投指数基金...\n\n资产配置建议：\n- 现金：10-20%\n- 债券：30-40%\n- 股票：40-50%\n- 其他：5-10%", Category: "财经", Tags: "理财,投资,基金", Field: "财经"},
		{Title: "基金定投策略分享", Content: "定投的优势：\n1. 分散风险，摊平成本\n2. 强制储蓄，积少成多\n3. 无需择时，简单易行\n\n选择基金的要点：\n- 长期业绩优秀\n- 基金经理稳定\n- 费率合理\n- 规模适中", Category: "财经", Tags: "基金,定投,策略", Field: "财经"},
		{Title: "买房避坑指南", Content: "看房注意事项：\n1. 地段：交通、配套、学区\n2. 户型：朝向、采光、通风\n3. 开发商：品牌、口碑、资金\n4. 合同：条款、交付标准、违约责任\n\n贷款计算技巧...", Category: "财经", Tags: "房产,买房,攻略", Field: "财经"},

		{Title: "雅思备考经验贴", Content: "听力：每天精听1小时，推荐剑桥真题\n阅读：掌握同义替换，限时练习\n写作：背诵高分范文，积累句型\n口语：找语伴练习，录音复盘\n\n备考时间规划：\n- 基础期（1个月）\n- 强化期（1个月）\n- 冲刺期（2周）", Category: "学习", Tags: "雅思,英语,备考", Field: "教育"},
		{Title: "转行产品经理路线图", Content: "阶段一：基础知识（1-2个月）\n- 学习产品方法论\n- 掌握Axure/Figma\n- 了解数据分析\n\n阶段二：项目实战（2-3个月）\n- 做竞品分析\n- 写PRD文档\n- 画原型图\n\n阶段三：求职准备\n- 准备作品集\n- 模拟面试\n- 内推投递", Category: "职场", Tags: "产品经理,转行,路线", Field: "职场"},

		{Title: "极简生活实践", Content: "物品断舍离原则：\n1. 一年未用即丢弃\n2. 一进一出原则\n3. 购买前问自己：真的需要吗？\n\n数字极简：\n- 清理手机App\n- 取消不必要订阅\n- 减少社交媒体使用\n\n时间极简...", Category: "生活", Tags: "极简,断舍离,生活", Field: "生活"},
		{Title: "完美牛排制作秘籍", Content: "选材：\n- 部位：眼肉、西冷、菲力\n- 等级：M3-M5适合煎制\n- 厚度：2-3cm最佳\n\n制作步骤：\n1. 室温回温30分钟\n2. 厨房纸吸干水分\n3. 撒盐和黑胡椒\n4. 大火热锅，每面煎1-2分钟\n5. 加入黄油、大蒜、迷迭香\n6. 淋油，醒肉5分钟", Category: "生活", Tags: "美食,牛排,烹饪", Field: "生活"},
		{Title: "小户型收纳技巧", Content: "垂直收纳：利用墙面空间\n隐藏收纳：床底、沙发底\n多功能家具：储物床、折叠桌\n分类整理：标签化管理\n定期清理：每季度断舍离\n\n推荐好物：\n- 真空压缩袋\n- 收纳箱\n- 门后挂钩\n- 分层置物架", Category: "生活", Tags: "收纳,家居,小户型", Field: "生活"},

		{Title: "面试谈薪技巧", Content: "谈薪前的准备：\n1. 调研市场薪资水平\n2. 明确自己的底线和期望\n3. 准备好谈判理由\n\n谈薪话术：\n\"基于我的经验和能力，以及市场行情，我期望的薪资是XX。我相信我能为公司创造更大的价值。\"\n\n福利也要谈：\n- 年终奖\n- 股票期权\n- 带薪假期\n- 培训机会", Category: "职场", Tags: "面试,薪资,谈判", Field: "职场"},
		{Title: "团队管理心得", Content: "管理的核心是人。几点经验：\n1. 以身作则，树立榜样\n2. 充分授权，信任团队\n3. 及时反馈，正向激励\n4. 关注成长，培养人才\n5. 营造氛围，快乐工作\n\n处理冲突的方法...", Category: "职场", Tags: "管理,团队,领导力", Field: "职场"},

		{Title: "异地恋保鲜秘诀", Content: "1. 固定沟通时间，视频通话\n2. 分享日常，保持参与感\n3. 共同计划，期待见面\n4. 信任对方，避免猜疑\n5. 制造惊喜，寄送礼物\n6. 共同目标，规划未来\n\n最重要的是：有结束异地的计划", Category: "情感", Tags: "异地恋,恋爱,感情", Field: "情感"},
		{Title: "青春期亲子沟通", Content: "沟通原则：\n1. 倾听而非说教\n2. 尊重隐私，给予空间\n3. 平等对话，避免命令\n4. 关注情绪，理解需求\n5. 以身作则，言传身教\n\n实用技巧：\n- 选择合适时机\n- 从孩子感兴趣的话题切入\n- 多用开放式问题", Category: "情感", Tags: "育儿,青春期,沟通", Field: "情感"},
	}

	var notes []model.Note
	for _, template := range noteTemplates {
		// 找到对应领域的用户
		var authorID uint
		for _, user := range users {
			if contains(getUserField(user.Username), template.Field) {
				authorID = user.ID
				break
			}
		}
		if authorID == 0 {
			authorID = users[rand.Intn(len(users))].ID
		}

		note := model.Note{
			Title:    template.Title,
			Content:  template.Content,
			Category: template.Category,
			Tags:     template.Tags,
			AuthorID: authorID,
			Status:   1,
		}
		note.CreatedAt = time.Now().Add(-time.Duration(rand.Intn(90)) * 24 * time.Hour)
		notes = append(notes, note)
	}

	db.Create(&notes)
	fmt.Printf("已创建 %d 个笔记\n", len(notes))
	return notes
}

func initNoteLikes(db *gorm.DB, users []model.User, notes []model.Note) {
	var likes []model.NoteLike
	likeMap := make(map[string]bool)

	for _, note := range notes {
		// 每个笔记5-30个点赞
		likeCount := rand.Intn(25) + 5
		for i := 0; i < likeCount; i++ {
			userID := users[rand.Intn(len(users))].ID
			key := fmt.Sprintf("%d-%d", note.ID, userID)
			if !likeMap[key] && userID != note.AuthorID {
				likeMap[key] = true
				likes = append(likes, model.NoteLike{
					NoteID: note.ID,
					UserID:   userID,
				})
			}
		}
	}

	db.Create(&likes)
	fmt.Printf("已创建 %d 个笔记点赞\n", len(likes))
}

func initVillages(db *gorm.DB) []model.Village {
	villages := []model.Village{
		{Name: "科技前沿", Description: "探讨最新科技趋势，分享技术干货", MemberCount: 12580},
		{Name: "理财投资", Description: "交流投资理财心得，实现财富增长", MemberCount: 8932},
		{Name: "情感树洞", Description: "倾诉情感困惑，分享恋爱经验", MemberCount: 15670},
		{Name: "旅行日记", Description: "分享旅行故事，推荐旅游目的地", MemberCount: 6789},
		{Name: "美食天地", Description: "分享美食制作，推荐餐厅美食", MemberCount: 23450},
		{Name: "健身打卡", Description: "记录健身历程，分享运动经验", MemberCount: 9876},
		{Name: "职场成长", Description: "职场经验分享，职业发展交流", MemberCount: 11234},
		{Name: "学习成长", Description: "学习方法交流，知识分享互助", MemberCount: 15432},
		{Name: "摄影爱好者", Description: "摄影作品分享，技巧交流", MemberCount: 5678},
		{Name: "读书分享", Description: "好书推荐，阅读心得交流", MemberCount: 8901},
	}

	db.Create(&villages)
	fmt.Printf("已创建 %d 个村落\n", len(villages))
	return villages
}

func initPosts(db *gorm.DB, users []model.User, villages []model.Village) []model.Post {
	postTemplates := []string{
		"今天学到了很多新知识，分享给大家！",
		"有个问题想请教大家，希望能得到帮助。",
		"最近的一些心得体会，记录一下。",
		"推荐一个很好用的工具/方法。",
		"大家对这个话题怎么看？欢迎讨论。",
		"打卡第%d天，坚持就是胜利！",
		"分享一个有趣的经历...",
		"求助！遇到这个问题怎么解决？",
		"干货分享，建议收藏！",
		"今天的心情有点复杂...",
	}

	var posts []model.Post
	for _, village := range villages {
		// 每个村落10-20个帖子
		postCount := rand.Intn(11) + 10
		for i := 0; i < postCount; i++ {
			content := postTemplates[rand.Intn(len(postTemplates))]
			if i < 5 {
				content = fmt.Sprintf(content, i+1)
			}

			post := model.Post{
				Content:   content,
				AuthorID:  users[rand.Intn(len(users))].ID,
				VillageID: village.ID,
				Likes:     rand.Intn(100),
				Comments:  rand.Intn(20),
				Status:    1,
			}
			post.CreatedAt = time.Now().Add(-time.Duration(rand.Intn(30)) * 24 * time.Hour)
			posts = append(posts, post)
		}
	}

	db.Create(&posts)
	fmt.Printf("已创建 %d 个帖子\n", len(posts))
	return posts
}

func initPostLikes(db *gorm.DB, users []model.User, posts []model.Post) {
	var likes []model.PostLike
	likeMap := make(map[string]bool)

	for _, post := range posts {
		// 每个帖子5-20个点赞
		likeCount := rand.Intn(15) + 5
		for i := 0; i < likeCount; i++ {
			userID := users[rand.Intn(len(users))].ID
			key := fmt.Sprintf("%d-%d", post.ID, userID)
			if !likeMap[key] && userID != post.AuthorID {
				likeMap[key] = true
				likes = append(likes, model.PostLike{
					PostID: post.ID,
					UserID:   userID,
				})
			}
		}
	}

	db.Create(&likes)
	fmt.Printf("已创建 %d 个帖子点赞\n", len(likes))
}

func initPostReplies(db *gorm.DB, users []model.User, posts []model.Post) {
	// Post模型没有ParentID字段，暂时不创建回复
	// 后续如果需要可以添加Reply模型
	fmt.Println("帖子回复功能待实现（需要添加Reply模型）")
}

func initComments(db *gorm.DB, users []model.User, questions []model.Question, notes []model.Note) {
	commentTemplates := []string{
		"非常有用的分享，感谢！",
		"我也遇到过类似的问题，后来通过...解决了",
		"说得很对，补充一点...",
		"收藏了，慢慢看",
		"写得太好了，学到了很多",
		"这个观点很有意思，我补充一下...",
		"感谢分享，已收藏！",
		"确实是这样，深有体会",
		"请问有更详细的教程吗？",
		"受益匪浅，谢谢！",
	}

	var comments []model.Comment

	// 为问题添加评论
	for _, question := range questions {
		// 每个问题2-8个评论
		commentCount := rand.Intn(7) + 2
		for i := 0; i < commentCount; i++ {
			comment := model.Comment{
				Content:    commentTemplates[rand.Intn(len(commentTemplates))],
				AuthorID:   users[rand.Intn(len(users))].ID,
				TargetID:   question.ID,
				TargetType: "question",
				Likes:      rand.Intn(50) + 5,
				Status:     1,
			}
			comment.CreatedAt = time.Now().Add(-time.Duration(rand.Intn(20)) * 24 * time.Hour)
			comments = append(comments, comment)
		}
	}

	// 为笔记添加评论
	for _, note := range notes {
		// 每个笔记1-5个评论
		commentCount := rand.Intn(5) + 1
		for i := 0; i < commentCount; i++ {
			comment := model.Comment{
				Content:    commentTemplates[rand.Intn(len(commentTemplates))],
				AuthorID:   users[rand.Intn(len(users))].ID,
				TargetID:   note.ID,
				TargetType: "note",
				Likes:      rand.Intn(30) + 3,
				Status:     1,
			}
			comment.CreatedAt = time.Now().Add(-time.Duration(rand.Intn(25)) * 24 * time.Hour)
			comments = append(comments, comment)
		}
	}

	db.Create(&comments)
	fmt.Printf("已创建 %d 个评论\n", len(comments))
}

func initCommentLikes(db *gorm.DB) {
	// 评论点赞通过Comment模型的Likes字段统计，不需要单独表
	fmt.Println("评论点赞已集成到评论模型")
}

func initChats(db *gorm.DB, users []model.User) {
	chatTemplates := []string{
		"你好，想请教一个问题",
		"在吗？有个事想咨询一下",
		"你好，看了你的分享很有收获",
		"能请教一下关于...的问题吗？",
		"你好，想和你交流一下",
	}

	replyTemplates := []string{
		"你好，什么问题？",
		"在的，请说",
		"谢谢认可，有什么问题尽管问",
		"可以的，请讲",
		"你好，很高兴认识你",
	}

	var chats []model.Chat
	var messages []model.Message

	// 创建20个聊天会话
	for i := 0; i < 20; i++ {
		user1 := users[rand.Intn(len(users))].ID
		user2 := users[rand.Intn(len(users))].ID
		if user1 == user2 {
			user2 = users[(i+1)%len(users)].ID
		}

		chat := model.Chat{
			UserID:       user1,
			ReceiverID:   user2,
			Type:         "user",
			LastMessage:  chatTemplates[rand.Intn(len(chatTemplates))],
			UnreadCount:  rand.Intn(5),
		}
		chats = append(chats, chat)
	}

	db.Create(&chats)

	// 为每个聊天创建消息
	for _, chat := range chats {
		msgCount := rand.Intn(8) + 2
		for j := 0; j < msgCount; j++ {
			var content string
			var senderID uint
			if j%2 == 0 {
				content = chatTemplates[rand.Intn(len(chatTemplates))]
				senderID = chat.UserID
			} else {
				content = replyTemplates[rand.Intn(len(replyTemplates))]
				senderID = chat.ReceiverID
			}

			msg := model.Message{
				ChatID:   chat.ID,
				SenderID: senderID,
				Content:  content,
				Type:     "text",
				Status:   1,
			}
			msg.CreatedAt = time.Now().Add(-time.Duration(rand.Intn(7)) * 24 * time.Hour)
			messages = append(messages, msg)
		}
	}

	db.Create(&messages)
	fmt.Printf("已创建 %d 个聊天会话和 %d 条消息\n", len(chats), len(messages))
}

// 辅助函数
func contains(field string, item string) bool {
	return field == item
}

func getUserField(username string) string {
	fieldMap := map[string]string{
		"tech_guru": "科技", "ai_researcher": "科技", "frontend_master": "科技",
		"security_expert": "科技", "data_engineer": "科技",
		"finance_wiz": "财经", "crypto_trader": "财经", "realestate_pro": "财经", "tax_advisor": "财经",
		"dr_health": "医疗", "nutritionist": "医疗", "psychologist": "医疗",
		"fitness_coach": "医疗", "yoga_instructor": "医疗",
		"english_teacher": "教育", "math_tutor": "教育", "career_mentor": "教育", "study_guru": "教育",
		"life_artist": "生活", "traveler": "生活", "foodie": "生活",
		"fashion_blogger": "生活", "diy_crafter": "生活",
		"hr_director": "职场", "product_manager": "职场", "sales_champion": "职场", "startup_founder": "职场",
		"emotion_healer": "情感", "parenting_expert": "情感", "relationship_coach": "情感",
		"lawyer_pro": "法律", "ip_attorney": "法律",
		"ui_designer": "设计", "photography_pro": "设计", "video_creator": "设计",
	}
	if field, ok := fieldMap[username]; ok {
		return field
	}
	return ""
}

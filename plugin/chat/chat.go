// Package chat 对话插件
package chat

import (
	"math/rand"
	"strconv"
	"time"

	ctrl "github.com/FloatTech/zbpctrl"
	"github.com/FloatTech/zbputils/control"
	zero "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/extension/rate"
	"github.com/wdvxdr1123/ZeroBot/message"
)

var (
	poke   = rate.NewManager[int64](time.Minute*5, 8) // 戳一戳
	engine = control.Register("chat", &ctrl.Options[*zero.Ctx]{
		DisableOnDefault: false,
		Help:             "chat\n- [BOT名字]\n- [戳一戳BOT]\n- 空调开\n- 空调关\n- 群温度\n- 设置温度[正整数]",
	})
)

func init() { // 插件主体
	// 被喊名字
	engine.OnFullMatch("", zero.OnlyToMe).SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			var nickname = zero.BotConfig.NickName[0]
			time.Sleep(time.Second * 1)
			ctx.SendChain(message.Text(
				[]string{
					nickname + "在此，有何贵干~",
					"(っ●ω●)っ在~",
					"这里是" + nickname + "(っ●ω●)っ",
					nickname + "不在呢~",
				}[rand.Intn(4)],
			))
		})
	// 戳一戳
	engine.On("notice/notify/poke", zero.OnlyToMe).SetBlock(false).
		Handle(func(ctx *zero.Ctx) {
			var nickname = zero.BotConfig.NickName[0]
			switch {
			case poke.Load(ctx.Event.GroupID).AcquireN(0):
				// 5分钟共8块命令牌 一次消耗3块命令牌
				time.Sleep(time.Second * 1)
				ctx.SendChain(message.Text(
					[]string{
						"系内！",
					        "喵呜~",
					        "In the pipe, five by five.",
					        "寄！",
					        "八嘎！Hentai！无路赛！！！",
					        "呜啊！",
					        "不许色色！",
					        "铁咩！",
					        "达咩跌斯！",
					    	"投降降~",
						"跟你这只猪讲不通",
						"红豆泥跌斯噶？",
						"哈？",
						"我这就去摇人！",
						"吃老子一拳！",
						"啊？",
						"少儿不宜",
						"这样啊，还挺牛批的",
						"把你饭叼走",
						"呐，多洗爹~,
						"再您马的见",
						"你总是说些别人听不懂的东西",
						"你们不吃饭吗？",
						"猫猫也不是好惹的！",
						"yue~",
						"我的快乐呢~",
						"你小子行",
						"心甘情愿落入你的圈套",
						"小孩塞牙",
						"奇迹和魔法都在下水道",
						"奇迹和魔法需要收费",
						"giao",
						"拿去看病吧",
						"被有机体气晕",
						"你不要拉我被子",
						"今天就到这吧，我下线了",
						"说什么炼铜不开心，肯定是谎话",
						"听我说谢谢你~",
						"被活活气死在床上",
					}[rand.Intn(39)],
				))
			case poke.Load(ctx.Event.GroupID).Acquire():
				// 5分钟共8块命令牌 一次消耗1块命令牌
				time.Sleep(time.Second * 1)
				ctx.SendChain(message.Text("黑化程度：100%"))
			default:
				// 频繁触发，不回复
			}
		})
	// 群空调
	var AirConditTemp = map[int64]int{}
	var AirConditSwitch = map[int64]bool{}
	engine.OnFullMatch("空调开").SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			AirConditSwitch[ctx.Event.GroupID] = true
			ctx.SendChain(message.Text("❄️哔~"))
		})
	engine.OnFullMatch("空调关").SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			AirConditSwitch[ctx.Event.GroupID] = false
			delete(AirConditTemp, ctx.Event.GroupID)
			ctx.SendChain(message.Text("💤哔~"))
		})
	engine.OnRegex(`设置温度(\d+)`).SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			if _, exist := AirConditTemp[ctx.Event.GroupID]; !exist {
				AirConditTemp[ctx.Event.GroupID] = 26
			}
			if AirConditSwitch[ctx.Event.GroupID] {
				temp := ctx.State["regex_matched"].([]string)[1]
				AirConditTemp[ctx.Event.GroupID], _ = strconv.Atoi(temp)
				ctx.SendChain(message.Text(
					"❄️风速中", "\n",
					"群温度 ", AirConditTemp[ctx.Event.GroupID], "℃",
				))
			} else {
				ctx.SendChain(message.Text(
					"💤", "\n",
					"群温度 ", AirConditTemp[ctx.Event.GroupID], "℃",
				))
			}
		})
	engine.OnFullMatch(`群温度`).SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			if _, exist := AirConditTemp[ctx.Event.GroupID]; !exist {
				AirConditTemp[ctx.Event.GroupID] = 26
			}
			if AirConditSwitch[ctx.Event.GroupID] {
				ctx.SendChain(message.Text(
					"❄️风速中", "\n",
					"群温度 ", AirConditTemp[ctx.Event.GroupID], "℃",
				))
			} else {
				ctx.SendChain(message.Text(
					"💤", "\n",
					"群温度 ", AirConditTemp[ctx.Event.GroupID], "℃",
				))
			}
		})
}

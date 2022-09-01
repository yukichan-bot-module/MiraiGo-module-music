package music

import (
	"strings"
	"sync"

	"github.com/Logiase/MiraiGo-Template/bot"
	"github.com/Logiase/MiraiGo-Template/utils"
	"github.com/Mrs4s/MiraiGo/client"
	"github.com/Mrs4s/MiraiGo/message"
	"github.com/yukichan-bot-module/MiraiGo-module-music/internal/service"
)

var instance *music
var logger = utils.GetModuleLogger("com.aimerneige.music")

type music struct {
}

func init() {
	instance = &music{}
	bot.RegisterModule(instance)
}

func (m *music) MiraiGoModule() bot.ModuleInfo {
	return bot.ModuleInfo{
		ID:       "com.aimerneige.music",
		Instance: instance,
	}
}

// Init 初始化过程
// 在此处可以进行 Module 的初始化配置
// 如配置读取
func (m *music) Init() {
}

// PostInit 第二次初始化
// 再次过程中可以进行跨 Module 的动作
// 如通用数据库等等
func (m *music) PostInit() {
}

// Serve 注册服务函数部分
func (m *music) Serve(b *bot.Bot) {
	b.GroupMessageEvent.Subscribe(func(c *client.QQClient, msg *message.GroupMessage) {
		msgString := msg.ToString()
		if strings.HasPrefix(msgString, "点歌") {
			songName := msgString[6:]
			songName = strings.TrimSpace(songName)
			if songName != "" {
				song, errMsg := service.SearchCloudMusic(songName)
				if errMsg == "" {
					c.SendGroupMusicShare(msg.GroupCode, song)
				} else {
					c.SendGroupMessage(msg.GroupCode, simpleText(errMsg))
				}
			}
			return
		}
	})
}

// Start 此函数会新开携程进行调用
// ```go
//
//	go exampleModule.Start()
//
// ```
// 可以利用此部分进行后台操作
// 如 http 服务器等等
func (m *music) Start(b *bot.Bot) {
}

// Stop 结束部分
// 一般调用此函数时，程序接收到 os.Interrupt 信号
// 即将退出
// 在此处应该释放相应的资源或者对状态进行保存
func (m *music) Stop(b *bot.Bot, wg *sync.WaitGroup) {
	// 别忘了解锁
	defer wg.Done()
}

func simpleText(msg string) *message.SendingMessage {
	return message.NewSendingMessage().Append(message.NewText(msg))
}

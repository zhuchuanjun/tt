package main

import (
	"context"
	"fmt"
	"github.com/baidubce/bce-qianfan-sdk/go/qianfan"
)

func main() {
	// 使用安全认证AK/SK鉴权，替换下列示例中参数，安全认证Access Key替换your_iam_ak，Secret Key替换your_iam_sk
	qianfan.GetConfig().AccessKey = "ALTAKsvjl8ZHYCU1f7NxJEoLEK"
	qianfan.GetConfig().SecretKey = "7ae1408ab55342c6ac48a45eb8c1b559"

	// 调用对话Chat，可以通过 WithModel 指定模型，例如指定ERNIE-3.5-8K，参数对应ERNIE-Bot
	chat := qianfan.NewChatCompletion(
		qianfan.WithModel("ERNIE-Bot"),
		//qianfan.WithModel("ERNIE-Bot-4"),
	)

	// 发起对话，例如介绍下北京
	resp, err := chat.Do(
		context.TODO(),
		&qianfan.ChatCompletionRequest{
			Messages: []qianfan.ChatCompletionMessage{
				qianfan.ChatCompletionUserMessage("标题：掌握波段行情，学习炒股的艺术\n\n在金融市场的波澜壮阔中，股票交易成为了人们寻求财富增长的重要途径之一。然而，要在股市的起伏波动中寻找机会并非易事，需要深入理解和精准把握波段行情。本文将探讨如何运用波段行情的知识来炒股。\n\n首先，什么是波段行情？简单来说，波段行情是指股票价格在一段时间内的变动趋势，可以是上升、下降或者横盘。股市投资者要做的就是研究这些波段行情，寻找买入和卖出的最佳时机。\n\n学习波段行情首先要理解股市是由无数个波段组成的。每个波段都有其起点和终点，这就是我们常说的“低买高卖”的原则。但是，识别每个波段的高点和低点并不容易，需要通过对历史数据的分析，结合技术指标的使用，来预测股票的走势。\n\n技术分析是研究波段行情的重要工具。其中，常见的有趋势线、支撑线和阻力线、移动平均线等。通过这些技术分析工具，投资者可以对股票的走势有更清晰的了解，从而做出更合理的投资决策。\n\n然而，依靠技术分析并不能保证成功。股市是一个复杂的系统，受到许多因素的影响，包括宏观经济环境、公司的业绩报告、政策变化等。因此，投资者在炒股时，不仅要掌握技术分析，还要关注这些因素，做出全面的判断。\n\n总的来说，利用波段行情炒股是一种需要耐心和技巧的活动。投资者需要具备良好的知识储备，积累丰富的实践经验，才能在股市中找到属于自己的机会。同时，我们必须明白，无论股市的波段如何变化，理性和冷静始终是成功的关键。\n\n股市如战场，只有充分准备，才能在战斗中取得胜利。希望每一位投资者都能在炒股的道路上，通过理解和把握波段行情，找到属于自己的投资策略，最终实现财富的增长。。对以上内容, 提取2个标签, 每个标签一行, 不要有多余的文字"),
			},
		},
	)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Println(resp.Result)
}

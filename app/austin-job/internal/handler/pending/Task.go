package pending

import (
	"austin-go/app/austin-common/types"
	"austin-go/app/austin-job/internal/handler/handlers"
	"austin-go/app/austin-job/internal/handler/services"
	"austin-go/app/austin-job/internal/handler/services/deduplication"
	"austin-go/app/austin-job/internal/svc"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
)

type Task struct {
	TaskInfo types.TaskInfo
	svcCtx   *svc.ServiceContext
}

func NewTask(taskInfo types.TaskInfo, svcCtx *svc.ServiceContext) *Task {
	return &Task{TaskInfo: taskInfo, svcCtx: svcCtx}
}

func (t Task) Run(ctx context.Context) {
	// 0. 丢弃消息 根据redis配置丢弃某个模板的所有消息
	if services.NewDiscardMessageService(t.svcCtx).IsDiscard(ctx, &t.TaskInfo) {
		logx.WithContext(ctx).Infow("消息被丢弃", logx.Field("task_info", t.TaskInfo))
		return
	}
	// 1. 屏蔽消息 比如夜间屏蔽,夜间屏蔽次日八点发送
	services.NewShieldService(t.svcCtx).Shield(ctx, &t.TaskInfo)
	// 2.平台通用去重 1. N分钟相同内容去重, 2. 一天内N次相同渠道去重
	if len(t.TaskInfo.Receiver) > 0 {
		deduplication.NewDeduplicationRuleService(t.svcCtx).Duplication(ctx, &t.TaskInfo)
	}
	// 3. 真正发送消息
	if len(t.TaskInfo.Receiver) > 0 {
		err := handlers.GetHandler(t.TaskInfo.SendChannel).DoHandler(ctx, t.TaskInfo)
		if err != nil {
			logx.Error(err)
		}
	}
}

type TaskRun interface {
	Run(ctx context.Context)
}

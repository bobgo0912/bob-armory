package model

import "github.com/bobgo0912/bob-armory/pkg/game/constant"

type GameRoundEvent struct {
	GameRoundId   string             `json:"game_round_id" db:"game_round_id"`     // 游戏期号
	GameCode      string             `json:"game_code" db:"game_code"`             // 游戏编码 BGM:MeGA  BGR:Rush
	PrizeNumber   string             `json:"prize_number" db:"prize_number"`       // 开奖号码（全量发送）
	GameEvent     constant.GameEvent `json:"game_event" db:"game_event"`           // begin 开始 prize_on_going 开奖中  finish开奖结束
	EventTime     int64              `db:"event_time" json:"event_time"`           //事件事件
	NextGameRound string             `json:"next_game_round" db:"next_game_round"` //下一期奖旗ID
}

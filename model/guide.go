package model

type GuideResq struct {
	Target     string `form:"target`
	StartTime  string `form:"start_time"`
	StartStage string `form:"start_stage"`
	EndTime    string `form:"end_time"`
	EndStage   string `form:"end_stage"`
}

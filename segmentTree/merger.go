package segmentTree

// merger接口
type Merger interface {
	// 具体的merge逻辑
	Merege(leftEle, rightEle interface{}) interface{}
}

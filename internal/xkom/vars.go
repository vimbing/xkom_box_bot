package xkom

var (
	GrayBox   BoxType = 1
	BlueBox   BoxType = 2
	PurpleBox BoxType = 3
)

func (b BoxType) GetWebhookColor() int {
	switch b {
	case 1:
		return 10197915
	case 2:
		return 33247
	case 3:
		return 8388747
	default:
		return 10197915
	}
}

package enum

var (
	PSA = PayStatus{0}
	PSB = PayStatus{1}
	PSC = PayStatus{2}
)

type PayStatus struct {
	status int
}

func Usage(s *PayStatus) {
}

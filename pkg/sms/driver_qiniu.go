package sms

// Qiniu 实现 sms.Driver interface
type Qiniu struct {
}

func (s *Qiniu) Send(phone string, message Message, config map[string]string) bool {

}

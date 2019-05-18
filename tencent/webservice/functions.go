package webservice

type (
	Config struct {
		Key    string
		Secret string
	}
	Tencent struct {
		Config Config
		Path   string
		Params map[string]string
		Sign   string
	}
)

func (T *Tencent) GetSign() (sign string) {
	T.Sign = sign
	return
}

func Helloword() string {
	return "Halo, My name is Li Haijun."
}

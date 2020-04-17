package go-abstraction

import "fmt"

type SloganSayer interface {
    Slogan()
}

// Campaign은 자기 자신의 인스턴스에서 SloganSayer를 받을 수 있음
// Campaign은 또한 SloganSayer 인터페이스를 구현하고 있으므로 SloganSayer이기도 함
// 이는 체이닝시 유용하다
type Campaign struct{
    SloganSayer
}

// SaySlogan은 파라미터로 Campaign 또한 받을 수 있음
func SaySlogan(s SloganSayer) {
    s.Slogan()
}

// Hillary는 SloganSayer 인터페이스를 구현함
// Hillary는 SloganSayer임
type Hillary struct{}

func (h *Hillary) Slogan() {
    fmt.Println("Stronger together.")
}

// Trump는 SloganSayer 인터페이스를 구현함
// Trump는 SloganSayer임
type Trump struct{}

func (t *Trump) Slogan() {
    fmt.Println("Make American great again.")
}

package person

// Person 人のデータ構造
type Person struct {
	Name   string
	Gender string
	Age    int
	Height int
	Weight int
}

// IsMale 男性ならtrueを返す
func (p *Person) IsMale() bool {
	if p.Gender == "male" {
		return true
	}
	return false
}

// IsLess160cm 160cm未満ならtrueを返す
func (p *Person) IsLess160cm() bool {
	if p.Height < 160 {
		return true
	}
	return false
}

// IsMore50kg 50kg以上ならtrueを返す
func (p *Person) IsMore50kg() bool {
	if p.Weight >= 50 {
		return true
	}
	return false
}

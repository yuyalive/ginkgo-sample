package person

// Person 人のデータ構造
type Person struct {
	name   string
	gender string
	age    int
	height int
	weight int
}

// IsMale 男性ならtrueを返す
func (p *Person) IsMale() bool {
	if p.gender == "male" {
		return true
	}
	return false
}

// IsLess160cm 160cm未満ならtrueを返す
func (p *Person) IsLess160cm() bool {
	if p.age < 160 {
		return true
	}
	return false
}

// IsMore50kg 50kg以上ならtrueを返す
func (p *Person) IsMore50kg() bool {
	if p.weight >= 50 {
		return true
	}
	return false
}

package services

type Test struct {
	ID       uint   `gorm:"id"`
	Name     string `gorm:"name"`
	Lastname string `gorm:"lastname"`
	Age      uint   `gorm:"age"`
}

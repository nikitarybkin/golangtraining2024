package main

func main() {

	testStruct := new(Gun)
	_ = testStruct
	/*
	 * Экземпляр созданной вами структуры необходимо передать в качестве
	 * аргумента функции testStruct, которая выполнит проверку соблюдения
	 * всех условий задания/
	 */

}

type Gun struct {
	On    bool
	Ammo  int
	Power int
}

func (g *Gun) Shoot() bool {
	if g.On && g.Ammo > 0 {
		g.Ammo--
		return true
	} else {
		return false
	}
}

func (g *Gun) RideBike() bool {
	if g.On && g.Power > 0 {
		g.Power--
		return true
	} else {
		return false
	}

}

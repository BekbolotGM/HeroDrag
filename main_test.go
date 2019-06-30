package main
import (
    "testing"
    "github.com/stretchr/testify/assert"
    //"strings"
    //"fmt"
)

var tests = []struct {
	name string
}{
    {"tc greater then"},
    {"tc less then"},
}

func TestDragDamage(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Greater(t, DragDamage(), -1, "Число должно быть больше -1.")
            assert.Less(t, DragDamage(), 21,"Число должно быть меньше 21.")
		})
	}
}

func TestHeroDamage(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Greater(t, HeroDamage(), -1,"Число должно быть больше -1.")
            assert.Less(t, HeroDamage(), 21,"Число должно быть меньше 21.")
		})
	}
}


func TestFatiqueDragUp(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Greater(t, FatiqueDragUp(), -1, "Число должно быть больше -1.")
            assert.Less(t, FatiqueDragUp(), 31,"Число должно быть меньше 31.")
		})
	}
}
func TestFatiqueHeroUp(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Greater(t, FatiqueHeroUp(), -1,"Число должно быть больше -1.")
            assert.Less(t, FatiqueHeroUp(), 31,"Число должно быть меньше 31.")
		})
	}
}


func TestAngerDragDamage(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Greater(t, AngerDragDamage(), -1, "Число должно быть больше -1.")
            assert.Less(t, AngerDragDamage(), 41,"Число должно быть меньше 41.")
		})
	}
}
func TestCheckAngerDrag(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Greater(t, CheckAngerDrag(), -1, "Число должно быть больше -1.")
            assert.Less(t, CheckAngerDrag(), 31,"Число должно быть меньше 31.")
		})
	}
}
func TestHeroAngerUp(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Greater(t, HeroAngerUp(), -1, "Число должно быть больше -1.")
		//assert.Less(t, HeroAngerUp(), 31,"Число должно быть меньше 31.")
		})
	}
}


func TestAngerHeroDamage(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Greater(t, AngerHeroDamage(), -1, "Число должно быть больше -1.")
            assert.Less(t, AngerHeroDamage(), 41,"Число должно быть меньше 41.")
		})
	}
}
func TestCheckAngerHero(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Greater(t, CheckAngerHero(), -1, "Число должно быть больше -1.")
			//assert.Less(t, CheckAngerHero(), 31,"Число должно быть меньше 31.")
		})
	}
}
func TestDragAngerUp(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Greater(t, DragAngerUp(), -1, "Число должно быть больше -1.")
			//assert.Less(t, DragAngerUp(), 31,"Число должно быть меньше 31.")
		})
	}
}

func TestArmor(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Less(t, Armor(), 101, "Число должно быть меньше 101")
		})
	}
}

func TestRandomEvent2(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
            assert.Greater(t, RandomEvent2(), -1 , "Число должно быть больше -1.")
            assert.Less(t, RandomEvent2(), 21,"Число должно быть меньше 21.")
		})
	}
}
func TestRandomEvent(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
            assert.Greater(t, RandomEvent(), -1 , "Число должно быть больше -1.")
            assert.Less(t, RandomEvent(), 21,"Число должно быть меньше 21.")
		})
	}
}

func TestSaveGame(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
            assert.FileExists(t,"/home/bekbolot/Desktop/heroDrag/HeroDrag/savingData.json","Файл не существует" )
		})
	}
}

func TestRandomQuote(t *testing.T) {
	assert.NotEmpty(t, RandomQuote(),"Не должно быть пусто")
}

/* func TestHeroDead(t *testing.T) {
	//for _, tt := range tests {
		//t.Run(t, func(t *testing.T) {
			assert.Equal(t, HeroDead(), 0, "Число должно быть равно 0.")
		//})
	//}
} */
/* func TestDragonDead(t *testing.T) {
	//for _, tt := range tests {
		//t.Run(t, func(t *testing.T) {
			assert.Equal(t,  0, DragonDead(), "Число должно быть равно 0.")
		//})
	//}
} */
func TestShowHeroHealth(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Greater(t, ShowHeroHealth(), -1, "Число должно быть больше -1.")
			assert.Less(t, ShowHeroHealth(), 101, "Число должно меньше больше 101.")
		})
	}
}
func TestShowDragonHealth(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Greater(t, ShowDragonHealth(), -1, "Число должно быть больше -1.")
			assert.Less(t, ShowDragonHealth(), 101, "Число должно меньше больше 101.")
		})
	}
}
func TestShowHeroActQuantity(t *testing.T) {
	     assert.Greater(t, ShowHeroActQuantity(), -1, "Число должно быть больше -1.")
}
func TestShowDragonActQuantity(t *testing.T) {
           assert.Greater(t, ShowDragonActQuantity(), -1, "Число должно быть больше -1.")
}
func TestDragonHealthUp(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Greater(t, DragonHealthUp(), 0, "Число должно быть больше 0.")
			assert.Less(t, DragonHealthUp(), 101, "Число должно меньше больше 101.")
		})
	}
}
func TestDragonAtack(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Greater(t, DragonAtack(), -1, "Число должно быть больше -1.")
			assert.Less(t, DragonAtack(), 21, "Число должно меньше больше 21.")
		})
	}
}

func TestHeroAtack(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Greater(t, HeroAtack(), -1, "Число должно быть больше -1.")
			assert.Less(t, HeroAtack(), 21, "Число должно меньше больше 21.")
		})
	}
}
/* func TestActionHero(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Less(t, ActionHero("2"),"4", "должно быть меньше 4")
	assert.Greater(t, ActionHero("1"),"0", "должно быть больше 0")
	assert.Greater(t, ActionHero("3"),"0", "должно быть больше 0")
		})
	}
} */

/* func TestSetTuningDragonName(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
	    assert.Less(t, SetTuningDragonName("1"),"3", "должно быть меньше 3")
	    assert.Greater(t, SetTuningDragonName("2"),"0", "должно быть больше 0")
		})
	}
} */
/* func TestPrintHeroName(t *testing.T) {
	assert.NotEmpty(t, PrintHeroName("Герой"), "не должно быть пусто")
	assert.Empty(t, PrintHeroName(""), " должно быть пусто")
} */

/* func TestGetAndOutputRandomNameOfDragon(t *testing.T) {
	assert.NotEmpty(t, GetAndOutputRandomNameOfDragon(),("Герой"), "не должно быть пусто")
} */
/* func TestPrintDragName(t *testing.T) {
	assert.NotEmpty(t, PrintDragName("Дракон"), "не должно быть пусто")
	assert.Empty(t, PrintDragName(""), " должно быть пусто")
} */

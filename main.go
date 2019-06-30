package main
import (
 "fmt"
 "os"
 "bufio"
 "io/ioutil"
 "log"
 "net/http"
 "strings"
 "encoding/json"
 "time"
 "math/rand"
)

var timer1 *time.Timer
type Menu struct{
	itemMainMenu, itemMenuSelectWeapon,itemMenuLevelOfGame, itemMenuActionHero string
}
var ( menu = &Menu {})
type Hero struct {
	Name string
	Health, Armor, Damage,Fatique, Anger, Count,RandomNumofHero int
	Weapon string
	hAct, hHealthUp bool
}
type Drag struct{
	Name string
	Health, Damage, Fatique, Anger, Count,RandomNumofDrag int
	setAutoDragName bool
}
type Save struct {
	HeroName, DragName, Weaponn string
	HeroHealth, DragHealth , HeroFatique, DragFatique, HeroAnger, DragAnger, HeroCount, DragCount, Level int
	SaveORnot bool
}
var ( dInfo = &Drag { "", 100, 20, 0, 0, 0,0, false,})
  var ( hInfo = &Hero { 
	"", 100, 30, 20, 0, 0, 0, 0,"", false, false,
    }
)
var (savedInfo = &Save {})

func main(){
  
	fmt.Println("============\nГлавное меню\n=========")
	fmt.Println("1.Начать новую игру\n2.Настройки \n3.Продолжить сохраненную игру\n============")
	//inputItemMainMenu()
	CallingSpecificFunction(menu.itemMainMenu)
	if savedInfo.SaveORnot != true {
		SelectItemMenuOfWeapon()
  MenuLevelOfGame()
	}
}

func CallingSpecificFunction(n string) string{
	switch n{
		case "1":
			if dInfo.setAutoDragName ==false{
				InputNameOfHero()
				InputNameOfDrag()
			}else if dInfo.setAutoDragName == true{
				InputNameOfHero()
				GetAndOutputRandomNameOfDragon()
		}
	  case "2": 
			InputItemMenuOfTuningDragonName()
	  case "3":
			ContinueGame()
		}
		return n
}

func inputItemMainMenu()string{
	fmt.Println("Нажмите цифру 1 для начала игры или цифру, 2 для настройки игры, 3 для продолжения сохраненной игры") 
	menu.itemMainMenu = AcceptInput()
	for{
		if menu.itemMainMenu == ""{
			fmt.Println("Нажмите на 1 или 2 или 3")
			menu.itemMainMenu = AcceptInput()
			}else{
			break
			}
		}
	return menu.itemMainMenu
}

func InputItemMenuOfTuningDragonName()string{
		fmt.Println("Нажмите цифру 1 для выбора имени дракона пользователем, цифру 2 для случайного имени \n1.Выбор пользователя \n2.Случайное имя дракона")
		itemTuningDragonName := AcceptInput()
		for{
			if itemTuningDragonName == ""{
				fmt.Println("Введите 1 или 2")
				itemTuningDragonName = AcceptInput()
				}else{
				break
				}
		}
		return itemTuningDragonName
	}
func SetTuningDragonName(n string) string{
	switch n{
		case "1": 
			//fmt.Println("Вы выбрали функцию ввода имени Дракона пользователем")
			 main()
			dInfo.setAutoDragName = false
		case "2":
			dInfo.setAutoDragName = true
			//fmt.Println("Вы выбрали функцию случайного ввода имени Дракона")
		 main()
	}	
	return n
}

func GetAndOutputRandomNameOfDragon()string {
	fmt.Println("Идет загрузка случайного имени дракона.....Ждите....")
	dragName := map[string]string{}
	response, err := http.Get("https://uinames.com/api/")
    if err != nil {
			  dInfo.Name = "Смауг"
				fmt.Println("Имя дракона загрузить не удалось.Имя по умолчанию: ", dInfo.Name)
        return dInfo.Name
    }
  responseData, err := ioutil.ReadAll(response.Body)
    if err != nil {
        log.Fatal(err)
		}
	json.Unmarshal(responseData, &dragName)
	dInfo.Name = string(dragName["name"])
	fmt.Println("Случайное имя дракона: ",string(dragName["name"]))
	return dInfo.Name
}


//Данную функцию не стал тестить так это ф-ии ЯП
func InputNameOfHero()string{
 for{
	fmt.Println("Введите имя Героя: ")
	hInfo.Name = AcceptInput()
	PrintHeroName(hInfo.Name)
	if len(hInfo.Name)>0{
		break
	}
 }
 return hInfo.Name
}

func PrintHeroName(n string) bool{
	if n != ""{
		fmt.Println("Имя героя: ",n)
		return true
	}else{
		return false
	}
	
}
func PrintDragName(n string)bool{
	if n != ""{
		fmt.Println("Имя дракона: ",n)
		return true
	}else{
		return false
	}
}
func InputNameOfDrag(){
	for {
		fmt.Print("Введите имя Дракона: ")
		dInfo.Name = AcceptInput()
		PrintDragName(dInfo.Name)
	  if len(dInfo.Name)>0{
			break
		}
	}
}

func MenuLevelOfGame() {
	fmt.Println("Меню выбора уровня сложности игры\n===========\n1.Легкий \n2.Средний\n3.Сложный\nНажмите цифру 1 для выбора легкого уровня, 2 для среднего уровня, 3 для сложного уровня")
	  iMenu := AcceptInput()
		switch iMenu {
			case "1": 
			  savedInfo.Level = 1
				Battle()
			case "2": 
			  savedInfo.Level = 2
				Battle()
      case "3":
				savedInfo.Level = 3
			  go Battle()
				Time1()
			default:
				MenuLevelOfGame()
			}
}

func SelectItemMenuOfWeapon() string{
	fmt.Println("Оружейный арсенал героя\n===================\n1.Меч \n2.Топор \n3.Коса \n4.Серп\nНажмите цифру 1 для выбора Меча, 2 для выбора Топора, 3 для выбора Косы, 4 для выбора Серпа")
	  iMenu := AcceptInput()
		switch iMenu {
    case "1":
			hInfo.Weapon = "Меч"
		case "2":
			hInfo.Weapon = "Топор"
		case "3":
			hInfo.Weapon = "Коса"
		case "4":
			hInfo.Weapon = "Серп"
		default:
			//fmt.Println("Нажмите цифру 1 для выбора Меча, 2 для выбора Топора, 3 для выбора Косы, 4 для выбора Серпа")
			SelectItemMenuOfWeapon()
		}
		return hInfo.Weapon
}

func ActionHero(n string)string{
    switch n{
		case "1":
		  hInfo.hAct = true
			hInfo.hHealthUp = false
		case "2":
		 hInfo.hAct = false
		if hInfo.Health == 100{
			fmt.Println("Доктор сказал, что Вы симулянт и с Вами все в порядке!!!")
	  }else if hInfo.Health < 100{
			hInfo.hHealthUp = true
		}
		fmt.Println(hInfo.hHealthUp)
	  case "3":
		SaveGame()
		//os.Exit(1)
	}
	return n
}
func inputItemMenuActionHero()string{
  fmt.Println("=========\n1.Атаковать\n2.Лечиться\n3.Сохранить и выйти\nВыберите действие героя - 1 для атаки, 2 для лечения, 3 сохранения игры\n==============")
  menu.itemMenuActionHero=AcceptInput()
  for{
	if menu.itemMenuActionHero == ""{
		fmt.Println("Ge!!!")
		menu.itemMenuActionHero = AcceptInput()
	  }else{
		break
	  }
  }
 return menu.itemMenuActionHero
}


func Battle()string{
	var winner string
	for{
		if savedInfo.Level == 3{
    hInfo.RandomNumofHero = random(1, 10)	
		}
		 HeroDamage()
		 DragDamage()
		 inputItemMenuActionHero()
		 ActionHero(menu.itemMenuActionHero)
    if savedInfo.Level == 2{
			RandomEvent()
		}
		//Hero ataack
		if hInfo.hAct == true {
			if savedInfo.Level == 3{
				timer1.Stop()
				Armor()
			}
			HeroAtack()
			if hInfo.Health <= 0{
			  winner =hInfo.Name
				break
			}
			DragonHealthUp()
			DragonAtack()
			if dInfo.Health <= 0{
				winner = dInfo.Name
				break
			}
			 
			HeroHealthUp()
		 }

		if savedInfo.Level == 3{
			if hInfo.hAct == false || hInfo.hHealthUp == false{
				go Time1()
			}
		}
	}
	return winner
}

func HeroAtack()int{
	if hInfo.Health > 0{ //герой наносит удар if true
		hInfo.Count += 1 //cчетчик
						if savedInfo.Level == 3{
			AngerHeroDamage()
		}
		fmt.Println("Герой наносит  удар!")//сколько ударов нанес герой
		if hInfo.RandomNumofHero == 0 {
			fmt.Println("Герой помахнулсяяя")
		}
		ShowHeroActQuantity()
		if savedInfo.Level ==  2{
			weaponGObad:=1
			hInfo.RandomNumofHero -=weaponGObad // оружие тупится
			FatiqueHeroUp()
		}
		dInfo.Health = dInfo.Health - hInfo.RandomNumofHero//минус жизни дракона
	}
	DragonDead()
	ShowDragonHealth()
	return hInfo.RandomNumofHero
}

//Протестил
func HeroHealthUp() int{
	if hInfo.hHealthUp == true{
		hInfo.Health +=10
		if hInfo.Health > 100{
			hInfo.Health = 100
		}
		fmt.Println("Герой подлечился ", ShowHeroHealth(), ShowDragonHealth())
		fmt.Println("===========")
		fmt.Println("Дракон наносит Вам ответный удар",ShowDragonActQuantity())
		hInfo.Health = hInfo.Health - dInfo.RandomNumofDrag
		ShowHeroHealth()
	 }
	 return hInfo.Health
}

func DragonAtack() int{
	if dInfo.Health > 0 && dInfo.RandomNumofDrag%2 ==0 {
		dInfo.Count += 1
			if savedInfo.Level == 3{
				AngerDragDamage()
			}
			fmt.Println("Дракон наносит Вам ответный удар")
			if dInfo.RandomNumofDrag == 0{
        fmt.Println("Дракон промахнулсяяяя!")
			}
			ShowDragonActQuantity()
		if savedInfo.Level == 2{
			FatiqueDragUp()
		}
		hInfo.Health = hInfo.Health - dInfo.RandomNumofDrag
		}
	HeroDead()
	ShowHeroHealth()
	return dInfo.RandomNumofDrag
}

//Протестил
func DragonHealthUp() int{
	if dInfo.RandomNumofDrag%2 != 0{
		dInfo.Health += 10 //Плюс для жизни дракона
		if dInfo.Health > 100 {
			dInfo.Health = 100
		}
		fmt.Println("Дракон решил зализать раны")
		ShowDragonHealth()
		//ShowHeroHealth(hInfo.Health) //инф-ия о жизни персонажей
	}
	return dInfo.Health
}
//Протестил
func ShowDragonHealth() int{
	if dInfo.Health >= 0 {
		fmt.Println("Уровень жизни дракона: ", dInfo.Health)
	}
	return dInfo.Health
}

//Протестил
func DragonDead() int{
	if dInfo.Health < 0 {
		dInfo.Health = 0
		 fmt.Println("Урааааа!!! Поздравляем, Вы выиграли!!! В качестве бонуса Вы получите крутую цитату...Ждите")
		 RandomQuote()
	}
	return dInfo.Health
}
 
//Протестил
func ShowDragonActQuantity() int{
	fmt.Println("Кол-во ходов дракона: ", dInfo.Count)
	return dInfo.Count
}
//Протестил
func ShowHeroActQuantity() int{
	fmt.Println("Кол-во ходов героя: ", hInfo.Count)
	return hInfo.Count
}

//Протестил
func ShowHeroHealth() int{
	if hInfo.Health >= 0{
		fmt.Println("Уровень жизн героя: ",hInfo.Health)//показывает уровень жизни героя
	}
	return hInfo.Health
}


//Протестил
func HeroDead() int{
	if hInfo.Health < 0 { //если уровень жизни героя ушел в минус
		hInfo.Health = 0   //то присвоем ноль 
		fmt.Println("Неудача!!! Вы проиграли!!!")
	}
	return hInfo.Health
}



type Q struct{
	Quotes map[string]interface{} `json:"quotes"`
}
//Протестил
func RandomQuote()string{
	response, err := http.Get("https://aitorp6.herokuapp.com/quotes/api/random")
	d:= "Если Ты не можешь дать имя своей функции, то Ты не понимаешь ее логики"
	if err != nil {
		 
		fmt.Println("Не удалось загрузить цитату. Цитата по умолчанию", d)
		return d
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal("Ошибка интернет соединения",err)
	}
	var quotes Q
  json.Unmarshal([]byte(responseData), &quotes)
	fmt.Println("Ваш бонус: ",quotes.Quotes["quote"])
	d=quotes.Quotes["quote"].(string)
	return d
}

//не стал тестить
func Time1(){
	timer1 = time.NewTimer(5 * time.Second)
	<-timer1.C
	fmt.Println("Время вышло! Вы проиграли.")
	os.Exit(1)
}
// протестил
func SaveGame(){
 data := &Save{ HeroName: hInfo.Name ,DragName: dInfo.Name, Weaponn: hInfo.Weapon,
							 HeroHealth: hInfo.Health, DragHealth:dInfo.Health,
							 HeroFatique: hInfo.Fatique, DragFatique: dInfo.Fatique,
							 HeroAnger: hInfo.Anger, DragAnger: dInfo.Anger,
							 HeroCount: hInfo.Count, DragCount: dInfo.Count,
							 Level: savedInfo.Level,
							 SaveORnot: true,
             }
 b, err := json.Marshal(data)
 if err != nil {
 fmt.Println(err)
 return
 }
 file, err := os.Create("savingData.json")
 if err != nil{
 fmt.Println("Невозможно создать файл:", err) 
 os.Exit(1) 
 }else{
	fmt.Println("Игра сохранена.")
 }
 defer file.Close() 
 file.WriteString(string(b))
}

func ContinueGame() {
	plan, _ := ioutil.ReadFile("savingData.json")
	var data interface{}
	err := json.Unmarshal([]byte(plan), &data)
	if err != nil{
		fmt.Println(err) 
		os.Exit(1) 
 }
		md := data.(map[string]interface{})
		hInfo.Name    = md["HeroName"].(string)
		hInfo.Health  = int(md["HeroHealth"].(float64))
		hInfo.Fatique = int(md["HeroFatique"].(float64))
	  hInfo.Anger   = int(md["HeroAnger"].(float64))
		hInfo.Count   = int(md["HeroCount"].(float64))
		hInfo.Weapon  = md["Weaponn"].(string)
		dInfo.Name    = md["DragName"].(string)
		dInfo.Health  = int(md["DragHealth"].(float64))
		dInfo.Fatique = int(md["DragFatique"].(float64))
		dInfo.Anger   = int(md["DragAnger"].(float64))
		dInfo.Count   = int(md["DragCount"].(float64))
		savedInfo.Level = int(md["Level"].(float64))
		savedInfo.SaveORnot = md["SaveORnot"].(bool)
		fmt.Println("Данные сохраненной игры:\nУровень игры: ", savedInfo.Level)
		fmt.Println("Имя героя: ",  hInfo.Name, "Оружие героя: ", hInfo.Weapon, "Уровень жизни: ",hInfo.Health, "Кол-во ходов: ",hInfo.Count)
		fmt.Println("Имя дракона: ",dInfo.Name, "Уровень жизни: ",dInfo.Health, "Кол-во ходов: ",dInfo.Count)
	  switch savedInfo.Level{
		case 1:
			Battle()
		case 2:
			Battle()
		case 3:
			fmt.Println("СЛОЖНЫЙ УРОВЕНЬ. ВЫ ДОЛЖНЫ СДЕЛАТЬ ХОД В ТЕЧЕНИИ 5 СЕКУНД!!!")
			go Battle()
			Time1()
		default:
			ContinueGame()
		}
}


func random(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}
// протестил
func RandomEvent() int{
	  HeroDamage()
		if hInfo.RandomNumofHero< 3 && hInfo.RandomNumofHero%2 != 0{
		  fmt.Println("Случайное событие! Молния ударила Героя и нанесла ему урон")
		  hInfo.Health -=hInfo.RandomNumofHero
		  if hInfo.Health > 0{
			fmt.Println("Уровень жизни Героя: ",hInfo.Health)
		  }
		}
		return hInfo.RandomNumofHero
	}
// протестил
func RandomEvent2() int{
	DragDamage()
	if dInfo.RandomNumofDrag < 3 && dInfo.RandomNumofDrag%2 == 0{
		fmt.Println("Случайное событие! Молния ударила Дракона и нанесла ему урон")
		dInfo.Health -= dInfo.RandomNumofDrag
		if dInfo.Health > 0 {
		fmt.Println ( "Уровень жизни Дракона: ", dInfo.Health )
	 }
	}
	return dInfo.RandomNumofDrag
}

//протестил
func HeroDamage() int{
	hInfo.RandomNumofHero = random(1,20)
	return hInfo.RandomNumofHero
}
//протестил
func FatiqueHeroUp() int{
	HeroDamage()
	hInfo.Fatique += 5
	if hInfo.Fatique  == 30 && hInfo.RandomNumofHero%2 !=0 { 
	  hInfo.RandomNumofHero = 0
	}
	if hInfo.Fatique == 30{
		hInfo.Fatique = 0
	}
	return hInfo.Fatique
}

//протестил
func DragDamage() int{
	dInfo.RandomNumofDrag = random(1,20)
	return dInfo.RandomNumofDrag
}
//протестил
func FatiqueDragUp() int{
	dInfo.Fatique += 5//если дракон нанес удар, то усталость +5
	if dInfo.Fatique == 30 && dInfo.RandomNumofDrag !=0 { //если усталость равна 30 и остаток от деления на 2 не равен 0, то есть шанс промахнуться
	  fmt.Println("Дракон промахнулся!")
		dInfo.RandomNumofDrag = 0
		}
	if dInfo.Fatique == 30 {
		dInfo.Fatique = 0
	}
	return dInfo.Fatique
}

//протестил
func AngerHeroDamage() int {
	rand.Seed(time.Now().Unix())
	missORnot := []int{ 15, 17, 2, 3, 5, 7, 9, 11, 13, }
	n := rand.Int() % len(missORnot)
	nn := missORnot[n]

	CheckAngerHero()

  if hInfo.Anger == 30 && nn%2 == 0 { //если уровень злости равен 30 и остаток от дел
		hInfo.RandomNumofHero +=20
	}else if hInfo.Anger == 30 && nn%2 != 0{
		hInfo.RandomNumofHero = 0
	}
	if  hInfo.Anger == 30{
		hInfo.Anger = 0
	}
	DragAngerUp()
  return hInfo.RandomNumofHero
}
//протестил
func CheckAngerHero() int{
	if hInfo.Anger == 30{
		fmt.Println("Герой зол, +20 к Его урону, +больше шансов промахнуться")
	}
 return hInfo.Anger
}
//протестил
func DragAngerUp()int{
	dInfo.Anger += 5
	return dInfo.Anger
}

//протестил
func AngerDragDamage() int {
	rand.Seed(time.Now().Unix())
	missORnot := []int{ 15, 17, 2, 3, 5, 7, 9, 11, 13, }
	n := rand.Int() % len(missORnot)
	nn := missORnot[n]
  CheckAngerDrag()
  if dInfo.Anger == 30 && nn%2 == 0 { //если уровень злости равен 30 и остаток от дел
		dInfo.RandomNumofDrag +=20
	}else if dInfo.Anger == 30 && nn%2 != 0{
		dInfo.RandomNumofDrag = 0
	}
	if dInfo.Anger ==30{
		dInfo.Anger = 0
	}
  HeroAngerUp()
  return dInfo.RandomNumofDrag
}
//протестил
func CheckAngerDrag() int{
	if dInfo.Anger == 30{
		fmt.Println("Дракон зол, +20 к Его урону, +больше шансов промахнуться")
	}
 return dInfo.Anger
}

//протестил
func HeroAngerUp()int{
	hInfo.Anger += 5
	return hInfo.Anger
}


//протестил
func Armor() int{
	if hInfo.Count == 5{
		fmt.Println("Активирована броня +",hInfo.Armor,"hp")
		hInfo.Health +=hInfo.Armor
	}
	if hInfo.Health > 100{
		hInfo.Health = 100
	}
	return hInfo.Health
}
//
func AcceptInput()string{
	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()
	input := strings.TrimSpace(scan.Text())
	return input
}
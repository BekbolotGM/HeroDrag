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

var ( dInfo = &Drag { "", 100, 20, 0, 0, 0, 0, false,})
var ( hInfo = &Hero { 
	"", 100, 30, 20, 0, 0, 0, 0,"", false, false,
    }
)
var (savedInfo = &Save {})

func main(){
	fmt.Println("Добро пожаловать в игру Герой против Дракона!!!")
	MainMenu()
	SelectItemMenu()
	if savedInfo.SaveORnot != true {
   SelectWeapon()
	 MenuLevelOfGame()
	}
}

//function of Selecting item of menu
func SelectItemMenu(){
	    fmt.Println("Нажмите цифру 1 для начала игры или цифру, 2 для настройки игры, 3 для продолжения сохраненной игры")
		  itemMenu :=bufio.NewScanner(os.Stdin)
		  itemMenu.Scan()
		  iMenu := strings.TrimSpace(itemMenu.Text())
		  switch iMenu{
			case "1": 
				if dInfo.setAutoDragName ==false{
					NameOfHero()
		      NameOfDrag()
					break
				}else if dInfo.setAutoDragName == true{
					//fmt.Println("Задайте имя Вашего героя")
					NameOfHero()
					RandomNameOfDrag()
				}
			case "2": 
				ConfigNameOfDragMenu()
			case "3":
				ContinueGame()
			default:
			SelectItemMenu()
		}
}

//func main menu
func MainMenu(){
	fmt.Println("============\nГлавное меню\n=========")
	fmt.Println("1.Начать новую игру \n2.Настройки \n3.Продолжить сохраненную игру\n============")
}

//function of config dragon name
func ConfigNameOfDragMenu(){
	fmt.Println("Нажмите цифру 1 для выбора имени дракона пользователем, цифру 2 для случайного имени \n1.Выбор пользователя \n2.Случайное имя дракона")
	  selectSetting :=bufio.NewScanner(os.Stdin)
		selectSetting.Scan()
		sSetting := strings.TrimSpace(selectSetting.Text())
	    switch sSetting{
			case "1": 
				fmt.Println("Вы выбрали функцию ввода имени Дракона пользователем")
				MainMenu()
				SelectItemMenu()
			case "2":
				dInfo.setAutoDragName = true
				fmt.Println("Вы выбрали функцию случайного ввода имени Дракона")
				MainMenu()
				SelectItemMenu()
			default:
				ConfigNameOfDragMenu()
		}
}

func RandomNameOfDrag() string{
	fmt.Println("Идет загрузка случайного имени дракона.....Ждите....")
    dragName := map[string]string{}
    response, err := http.Get("https://uinames.com/api/")
    if err != nil {
        fmt.Print(err.Error())
        os.Exit(1)
    }
  responseData, err := ioutil.ReadAll(response.Body)
    if err != nil {
        log.Fatal(err)
    }
  json.Unmarshal(responseData, &dragName)
	fmt.Println("Случайное имя дракона: ",string(dragName["name"]))
	dInfo.Name = string(dragName["name"])
	return dInfo.Name
}

func NameOfHero() string{
	  heroName :=bufio.NewScanner(os.Stdin)
		fmt.Println("Введите имя Героя: ")
		heroName.Scan()
		hInfo.Name = strings.TrimSpace(heroName.Text())
		if len(hInfo.Name)>0{
			fmt.Println("Имя героя: ", hInfo.Name)
	  }
		if hInfo.Name == ""{
			fmt.Println("Имя Героя не может быть пустым")
			NameOfHero()
			return ""
		}
	return hInfo.Name
}

func NameOfDrag() {
	for {
		dragName := bufio.NewScanner(os.Stdin)
		fmt.Print("Введите имя Дракона: ")
		dragName.Scan()
		dInfo.Name = strings.TrimSpace(dragName.Text())
		if dInfo.Name == ""{
			fmt.Println("Имя Дракона не может быть пустым")
		}
	    if len(dInfo.Name)>0{
		fmt.Println("Имя Дракона: ", dInfo.Name)
		break
	    }
	}
}

func MenuLevelOfGame(){
	fmt.Println("Меню выбора уровня сложности игры")
	fmt.Println("===========")
	fmt.Println("1.Легкий \n2.Средний\n3.Сложный")
	fmt.Println("Нажмите цифру 1 для выбора легкого уровня, 2 для среднего уровня, 3 для сложного уровня")
	  itemMenu :=bufio.NewScanner(os.Stdin)
		itemMenu.Scan()
		iMenu := itemMenu.Text()
		switch iMenu {
			case "1": 
			  savedInfo.Level = 1
				fmt.Println("ЛЕГКИЙ УРОВЕНЬ")
				Level1()
			case "2": 
			  savedInfo.Level = 2
				fmt.Println("СРЕДНИЙ УРОВЕНЬ")
				Level1()
            case "3":
				savedInfo.Level = 3
				fmt.Println("СЛОЖНЫЙ УРОВЕНЬ. ВЫ ДОЛЖНЫ СДЕЛАТЬ ХОД В ТЕЧЕНИИ 5 СЕКУНД!!!")
			    go Level1()
				Time1()
			default:
				MenuLevelOfGame()
			}
}

func SelectWeapon(){
	fmt.Println("Оружейный арсенал героя\n===================")
	fmt.Println("1.Меч \n2.Топор \n3.Коса \n4.Серп\nНажмите цифру 1 для выбора Меча, 2 для выбора Топора, 3 для выбора Косы, 4 для выбора Серпа")
  fmt.Println("")
	  itemMenu :=bufio.NewScanner(os.Stdin)
		itemMenu.Scan()
		iMenu := itemMenu.Text()
		switch iMenu {
        case "1":
			fmt.Println("Вы выбрали Меч")
			hInfo.Weapon = "Меч"
		case "2":
			fmt.Println("Вы выбрали Топор")
			hInfo.Weapon = "Топор"
		case "3":
			fmt.Println("Вы выбрали Косу")
			hInfo.Weapon = "Коса"
		case "4":
			fmt.Println("Вы выбрали Серп")
			hInfo.Weapon = "Серп"
		default:
			fmt.Println("Нажмите цифру 1 для выбора Меча, 2 для выбора Топора, 3 для выбора Косы, 4 для выбора Серпа")
			SelectWeapon()
		}
}

func Level1() {
	heroHealthMsg := "Уровень жизни героя: "
	dragHealthMsg := "Уровень жизни дракона: "
	weaponGObad:=1
	for {
		rand.Seed(time.Now().UnixNano())
		if savedInfo.Level == 3{
		hInfo.RandomNumofHero = random(1, 10)	
		}else{
		hInfo.RandomNumofHero = random(1, 20)
		}
		dInfo.RandomNumofDrag = random(1, 20)
		
		ActionHero()

		if savedInfo.Level == 2{
			RandomEvent()
		}
		if hInfo.hAct == true {
			if savedInfo.Level == 3{
				timer1.Stop()
				Armor()
			}
			if hInfo.Health > 0{ //герой наносит удар if true
				hInfo.Count += 1 //cчетчик
                if savedInfo.Level == 3{
					AngerHero()
				}
				fmt.Println("Герой наносит  удар!","Кол-во ходов Героя: ",hInfo.Count)//сколько ударов нанес герой
				if savedInfo.Level ==  2{
					hInfo.RandomNumofHero -=weaponGObad // оружие тупится
					FatiqueHero()
				}
				dInfo.Health = dInfo.Health - hInfo.RandomNumofHero//минус жизни дракона
				
				if dInfo.Health > 0 {
					fmt.Println ( dragHealthMsg, dInfo.Health )
				}
				if dInfo.Health <= 0 {
					dInfo.Health = 0
					fmt.Println(heroHealthMsg,hInfo.Health)// показывает уровень жизни героя
					fmt.Println(dragHealthMsg,dInfo.Health," Количество ходов Дракона: ",dInfo.Count)//уровень жизни героя и кол-во ходов дракона
					fmt.Println("Урааааа!!! Поздравляем, Вы выиграли!!! В качестве бонуса Вы получите крутую цитату...Ждите")
					RandomQuote()
					break
				}
			}
			
	        //fmt.Println("\n")

			if dInfo.Health > 0 && dInfo.RandomNumofDrag%2 !=0 {
				dInfo.Count += 1
					if savedInfo.Level == 3{
						AngerDrag()
					}
					fmt.Println("Дракон наносит Вам ответный удар","Количество ходов Дракона: ",dInfo.Count)
				if savedInfo.Level == 2{
					FatiqueDrag()
				}
				hInfo.Health = hInfo.Health - dInfo.RandomNumofDrag
				if hInfo.Health > 0{
					fmt.Println(heroHealthMsg,hInfo.Health)//показывает уровень жизни герояы
				}
			} else if dInfo.RandomNumofDrag%2 == 0{
				dInfo.Health += 10 //Плюс для жизни дракона
				if dInfo.Health > 100 {
				   dInfo.Health = 100
				}
				fmt.Println("Дракон решил зализать раны")
				fmt.Println(dragHealthMsg,dInfo.Health, heroHealthMsg,hInfo.Health)//инф-ия о жизни персонажей
			}
		 }
		 if hInfo.hHealthUp == true{
			hInfo.Health +=10
			if hInfo.Health > 100{
				hInfo.Health = 100
			}
			fmt.Println("Герой подлечился ", heroHealthMsg, hInfo.Health, dragHealthMsg,dInfo.Health)
			fmt.Println("===========")
			fmt.Println("Дракон наносит Вам ответный удар","Количество ходов Дракона: ",dInfo.Count)
			hInfo.Health = hInfo.Health - dInfo.RandomNumofDrag
			fmt.Println(heroHealthMsg,hInfo.Health) 
		 }
		 if hInfo.Health <= 0 { //если уровень жизни героя ушел в минус
			hInfo.Health = 0   //то присвоем ноль 
			fmt.Println(heroHealthMsg,hInfo.Health, "Кол-во ходов Героя: ", hInfo.Count)//инфа о жизни героя и о кол-ве ходов
			fmt.Println(dragHealthMsg,dInfo.Health) //инф-ия о жизни дракона
			fmt.Println("Неудача!!! Вы проиграли!!!")
			break
		}
		 if hInfo.hAct == false || hInfo.hHealthUp == false{
			 go Time1()
		 }
    }
}

func random(min int, max int) int {
	return rand.Intn(max-min) + min
}

func ActionHero() string {
	fmt.Println("=========\n1.Атаковать\n2.Лечиться\n3.Сохранить и выйти\nВыберите действие героя - 1 для атаки, 2 для лечения, 3 сохранения игры\n==============")
    action :=bufio.NewScanner(os.Stdin)
		action.Scan()
		heroAction := action.Text()
		switch heroAction{
		case "1":
		  hInfo.hAct = true
		  hInfo.hHealthUp = false
		case "2":
		 hInfo.hAct = false
		if hInfo.Health == 100{
			fmt.Println("Доктор сказал, что Вы симулянт и с Вами все в порядке!!!")
	  } else {
			hInfo.hHealthUp = true
		}
	  case "3":
		SaveGame()
		os.Exit(1)
	default:
		ActionHero()
	}
		return heroAction
}

type Q struct{
	Quotes map[string]interface{} `json:"quotes"`
}

func RandomQuote(){
	response, err := http.Get("https://aitorp6.herokuapp.com/quotes/api/random")
	if err != nil {
			fmt.Print(err.Error())
			os.Exit(1)
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
			log.Fatal("Ошибка интернет соединения",err)
	}
	var quotes Q
  json.Unmarshal([]byte(responseData), &quotes)
	fmt.Println("Ваш бонус: ",quotes.Quotes["quote"],"==Автор цитаты: ", quotes.Quotes["author"])
}

func Time1(){
	timer1 = time.NewTimer(5 * time.Second)
	<-timer1.C
	fmt.Println("Время вышло! Вы проиграли.")
	os.Exit(1)
}

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
	file, err := os.Open("savingData.json")
    if err != nil{
        fmt.Println(err) 
        os.Exit(1) 
    }
    defer file.Close() 
			
				plan, _ := ioutil.ReadFile("savingData.json")
				var data interface{}
				
				err = json.Unmarshal([]byte(plan), &data)
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
		//fmt.Println("\n")
	  switch savedInfo.Level{
		case 1:
			Level1()
		case 2:
			Level1()
		case 3:
			fmt.Println("СЛОЖНЫЙ УРОВЕНЬ. ВЫ ДОЛЖНЫ СДЕЛАТЬ ХОД В ТЕЧЕНИИ 5 СЕКУНД!!!")
			go Level1()
			Time1()
		default:
			ContinueGame()
		}
			
}

func RandomEvent() int{
	randomEvent := dInfo.RandomNumofDrag + hInfo.RandomNumofHero
	var health int
		if randomEvent < 6 && randomEvent%2 != 0{
		  fmt.Println("Случайное событие! Молния ударила Героя и отняла 20 hp")
		  hInfo.Health -=20
		  health = hInfo.Health
		  if hInfo.Health > 0{ //это условие для того чтобы, если уровень жизни уйдет в минус то инф-ия о жизни героя не отобразится
			fmt.Println("Уровень жизни Героя: ",hInfo.Health)//показывает уровень жизни герояы
		  }
		}else if randomEvent < 6 && randomEvent%2 == 0{
		  fmt.Println("Случайное событие! Молния ударила Дракона и отняла 20 hp")
		  dInfo.Health -=20
		  health = dInfo.Health
		  if dInfo.Health > 0 { //это условие для того чтобы, если уровень жизни уйдет в минус то инф-ия о жизни дракона не отобразится
			fmt.Println ( "уровень жизни Дракона: ", dInfo.Health )//инфа о уровни жизни дракона
		  }
		}
	return health
}

func FatiqueHero() int{
	hInfo.Fatique += 5//если герой нанес удар, то усталость +5
	if hInfo.Fatique == 30 && hInfo.RandomNumofHero%2 !=0 { //если усталость равна 30 и остаток от деления на 2 не равен 0, то есть шанс промахнуться
        fmt.Println("Герой промахнулся!")
		hInfo.RandomNumofHero = 0
		hInfo.Fatique = 0//обнуляем усталость героя
	}
	return hInfo.RandomNumofHero
}

func FatiqueDrag() int{
	dInfo.Fatique += 5//если дракон нанес удар, то усталость +5
	if dInfo.Fatique == 30 && dInfo.RandomNumofDrag%2 !=0 { //если усталость равна 30 и остаток от деления на 2 не равен 0, то есть шанс промахнуться
	    fmt.Println("Дракон промахнулся!")
		dInfo.RandomNumofDrag = 0
		dInfo.Fatique = 0//обнуляем усталость дракона
	}
	return dInfo.RandomNumofDrag
}

func AngerHero() int {
	rand.Seed(time.Now().Unix())
	missORnot := []int{ 15, 17, 2, 3, 5, 7, 9, 11, 13, }
	h := rand.Int() % len(missORnot)
	hh := missORnot[h]
	 if hInfo.Anger == 30{
		fmt.Println("Герой зол, +20 к Его урону, +больше шансов промахнуться")
	 }
	if hInfo.Anger == 30 && hh%2 == 0{
		hInfo.RandomNumofHero +=20
	}else if hInfo.Anger == 30 && hh%2 != 0 {
		hInfo.RandomNumofHero = 0
	    fmt.Println("Герой промахнулсяяяяяя!!!")
	}
	dInfo.Anger +=5
	return  hInfo.RandomNumofHero
}

func AngerDrag() int {
	rand.Seed(time.Now().Unix())
	missORnot := []int{ 15, 17, 2, 3, 5, 7, 9, 11, 13, }
	n := rand.Int() % len(missORnot)
	nn := missORnot[n]
	if dInfo.Anger == 30{
		fmt.Println("Дракон зол, +20 к Его урону, +больше шансов промахнуться")
	}
	if dInfo.Anger == 30 && nn%2 == 0 { //если уровень злости равен 30 и остаток от дел
		dInfo.RandomNumofDrag +=20
	}else if dInfo.Anger == 30 && nn%2 != 0{
		dInfo.RandomNumofDrag = 0
	    fmt.Println("Дракон промахнулсяяяяя!!!")
	}
	hInfo.Anger += 5
	return dInfo.RandomNumofDrag
}

func Armor() int{
	rand.Seed(time.Now().Unix())
	missORnot := []int{ 15, 17, 2, 3, 5, 7, 9, 11, 13, 19,21,31,33,37,39, }
	n := rand.Int() % len(missORnot)
	nn := missORnot[n]
	if hInfo.Health < 80 && nn%2 == 0 {
		fmt.Println("Активирована броня +",hInfo.Armor,"hp")
		hInfo.Health +=hInfo.Armor
	}
	if hInfo.Health > 100{
		hInfo.Health = 100
	}
	return hInfo.Health
}
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
var setAutoDragName bool = false
var dName,hName, ddName string 
var hAct,hHealthUp bool = false,false
type Hero struct {
	Health, Armor, Damage,Fatique, Anger   int
	Weapons []string
}
type Drag struct{
	Health, Damage, Fatique, Anger int
}

func main(){
	fmt.Println("Добро пожаловать в игру Герой против Дракона!!!")
	mainMenu()
	selectItemMenu()
	selectWeapon()
	menuLevelOfGame()
}

//function of Selecting item of menu
func selectItemMenu(){
	    fmt.Println("Нажмите цифру 1 для начала игры или цифру 2 для настройки игры")
		itemMenu :=bufio.NewScanner(os.Stdin)
		itemMenu.Scan()
		iMenu := strings.TrimSpace(itemMenu.Text())
		switch iMenu{
			case "1": 
				if setAutoDragName ==false{
					nameforperson()
					break
				}else if setAutoDragName == true{
					//fmt.Println("Задайте имя Вашего героя")
					nameOfHero()
					randomNameOfDrag()
				}
			case "2": 
				configNameOfDragMenu()
			default:
			selectItemMenu()
		}
}


//func main menu
func mainMenu(){
	fmt.Println("============")
	fmt.Println("Главное меню")
	fmt.Println("============")
	fmt.Println("1.Начать игру")
	fmt.Println("2.Настройки")
	fmt.Println("====================")
}


//function of config dragon name
func configNameOfDragMenu(){
	fmt.Println("Нажмите цифру 1 для выбора имени дракона пользователем, цифру 2 для случайного имени")
	fmt.Println("1.Выбор пользователя")
	fmt.Println("2.Случайное имя дракона")
	
		selectSetting :=bufio.NewScanner(os.Stdin)
		selectSetting.Scan()
		sSetting := strings.TrimSpace(selectSetting.Text())
	    switch sSetting{
			case "1": 
				fmt.Println("Вы выбрали функцию ввода имени Дракона пользователем")
				mainMenu()
				selectItemMenu()
			case "2":
				setAutoDragName = true
				fmt.Println("Вы выбрали функцию случайного ввода имени Дракона")
				mainMenu()
				selectItemMenu()
			default:
				configNameOfDragMenu()
		}
}



func randomNameOfDrag(){
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
	ddName = string(dragName["name"])
}


//function of selecting person name
func nameforperson(){
	if setAutoDragName == false{
		nameOfHero()
		nameOfDrag()
	}else if setAutoDragName == true{
		nameOfHero()
	}
}


func nameOfHero(){
	for {
		heroName :=bufio.NewScanner(os.Stdin)
		fmt.Print("Введите имя Героя: ")
		heroName.Scan()
		hName = strings.TrimSpace(heroName.Text())
		if hName == ""{
			fmt.Println("Имя Героя не может быть пустым")
		}
	    if len(hName)>0{
		  fmt.Println("Имя героя: ", hName)
		 break
	    }
	}
}


func nameOfDrag() {
	for {
		dragName := bufio.NewScanner(os.Stdin)
		fmt.Print("Введите имя Дракона: ")
		dragName.Scan()
		dName = strings.TrimSpace(dragName.Text())
		
		
		if dName == ""{
			fmt.Println("Имя Дракона не может быть пустым")
		}
	    if len(dName)>0{
		fmt.Println("Имя Дракона: ", dName)
		break
	    }
	}
}


func menuLevelOfGame(){
	fmt.Println("Меню выбора уровня сложности игры")
	fmt.Println("===========")
	fmt.Println("1.Легкий")
	fmt.Println("2.Средний")
	fmt.Println("3.Сложный")
	fmt.Println("Нажмите цифру 1 для выбора легкого уровня, 2 для среднего уровня, 3 для сложного уровня")
	
		itemMenu :=bufio.NewScanner(os.Stdin)
		itemMenu.Scan()
		iMenu := itemMenu.Text()
		switch iMenu {
			case "1": 
				fmt.Println("ЛЕГКИЙ УРОВЕНЬ")
				level1()
			case "2": 
				fmt.Println("СРЕДНИЙ УРОВЕНЬ")
				level2()
			case "3":
				fmt.Println("СЛОЖНЫЙ УРОВЕНЬ")
				level3()
			default:
				menuLevelOfGame()
			}
}


func selectWeapon(){
	fmt.Println("Оружейный арсенал героя")
	fmt.Println("=======================")
	fmt.Println("1.Меч")
	fmt.Println("2.Топор")
	fmt.Println("3.Коса")
	fmt.Println("4.Серп")

	fmt.Println("Нажмите цифру 1 для выбора Меча, 2 для выбора Топора, 3 для выбора Косы, 4 для выбора Серпа")
	
		itemMenu :=bufio.NewScanner(os.Stdin)
		itemMenu.Scan()
		iMenu := itemMenu.Text()
		switch iMenu {
        case "1":
			fmt.Println("Вы выбрали Меч")
		case "2":
			fmt.Println("Вы выбрали Топор")
		case "3":
			fmt.Println("Вы выбрали Косу")
		case "4":
			fmt.Println("Вы выбрали Серп")
		default:
			fmt.Println("Нажмите цифру 1 для выбора Меча, 2 для выбора Топора, 3 для выбора Косы, 4 для выбора Серпа")
			selectWeapon()
		}
}


var ( dInfo = &Drag { 100, 20, 0, 0,})
var ( hInfo = &Hero { 
	100, 30, 20, 0, 0,
	[]string {
	"Меч",
	"Топор",
	"Коса",
	"Серп",
          },
    }
)
func level1() {
	heroHealthMsg := "Уровень жизни героя: "
	dragHealthMsg := "Уровень жизни дракона: "
	countHero := 0 //нужно для счета ходов Героя
	countDrag := 0 //нужно для счета ходов Дракона
	for {
		rand.Seed(time.Now().UnixNano())
		randomNumofHero := random(1, 20)
		randomNumofDrag := random(1, 20)
		actionHero()
		
		if hAct == true {
			if hInfo.Health > 0{ //герой наносит удар if true
				countHero += 1 //cчетчик
				fmt.Println("Герой наносит  удар!","Кол-во ходов Героя: ",countHero)//сколько ударов нанес герой
				
				dInfo.Health = dInfo.Health - randomNumofHero//минус жизни дракона
				if dInfo.Health > 0 {
					fmt.Println ( dragHealthMsg, dInfo.Health )
				}
				if dInfo.Health <= 0 {
					dInfo.Health = 0
					fmt.Println(heroHealthMsg,hInfo.Health)// показывает уровень жизни героя
					fmt.Println(dragHealthMsg,dInfo.Health," Количество ходов Дракона: ",countDrag)//уровень жизни героя и кол-во ходов дракона
					fmt.Println("Урааааа!!! Поздравляем, Вы выиграли!!! В качестве бонуса Вы получите крутую цитату...Ждите")
					randomQuote()
					break
				}
			}
			
		

			fmt.Println("\n")

			if dInfo.Health > 0 && randomNumofDrag%2 !=0 {
				countDrag += 1
				fmt.Println("Дракон наносит Вам ответный удар","Количество ходов Дракона: ",countDrag)
				
				hInfo.Health = hInfo.Health - randomNumofDrag
				if hInfo.Health > 0{
					fmt.Println(heroHealthMsg,hInfo.Health)//показывает уровень жизни герояы
				}
				if hInfo.Health <= 0 { //если уровень жизни героя ушел в минус
					hInfo.Health = 0   //то присвоем ноль 
					fmt.Println(heroHealthMsg,hInfo.Health, "Кол-во ходов Героя: ", countHero)//инфа о жизни героя и о кол-ве ходов
					fmt.Println(dragHealthMsg,dInfo.Health) //инф-ия о жизни дракона
					fmt.Println("Неудача!!! Вы проиграли!!!")
					break
				}
			} else if randomNumofDrag%2 == 0{
				dInfo.Health += 10 //Плюс для жизни дракона
				if dInfo.Health > 100 {
				   dInfo.Health = 100
				}
				fmt.Println("Дракон решил зализать раны")
				fmt.Println(dragHealthMsg,dInfo.Health, heroHealthMsg,hInfo.Health)//инф-ия о жизни персонажей
			}
		 }
		 if hHealthUp == true{
			hInfo.Health +=10
			if hInfo.Health > 100{
				hInfo.Health = 100
			}
			fmt.Println("Герой подлечился ", heroHealthMsg, hInfo.Health, dragHealthMsg,dInfo.Health) 
		 }

		 
	}
}

func level2() {
	heroHealthMsg := "Уровень жизни героя: "
	dragHealthMsg := "Уровень жизни дракона: "
	countHero := 0 //нужно для счета ходов Героя
	countDrag := 0 //нужно для счета ходов Дракона
	for {
		rand.Seed(time.Now().UnixNano())
		randomNumofHero := random(1, 20)
		randomNumofDrag := random(1, 20)
		randomEvent := randomNumofDrag + randomNumofHero
		if randomEvent < 6 && randomEvent%2 != 0{
		  fmt.Println("Случайное событие! Молния ударила Героя и отняла 20 hp")
		  hInfo.Health -=20
		  if hInfo.Health > 0{ //это условие для того чтобы, если уровень жизни уйдет в минус то инф-ия о жизни героя не отобразится
			fmt.Println(heroHealthMsg,hInfo.Health)//показывает уровень жизни герояы
		  }
		}else if randomEvent < 6 && randomEvent%2 == 0{
		  fmt.Println("Случайное событие! Молния ударила Дракона и отняла 20 hp")
		  dInfo.Health -=20
		  if dInfo.Health > 0 { //это условие для того чтобы, если уровень жизни уйдет в минус то инф-ия о жизни дракона не отобразится
			fmt.Println ( dragHealthMsg, dInfo.Health )//инфа о уровни жизни дракона
		  }
		}
		actionHero()
		
		if hAct == true {
			rand.Seed(time.Now().Unix())
	        missORnot := []int{ 15, 17, 2, 3, 5, 7, 9, 11, 13, }
	        h := rand.Int() % len(missORnot)
	        hh := missORnot[h]
	       
			if hInfo.Health > 0{ //герой наносит удар if true
				countHero += 1 //cчетчик ходов
				if hInfo.Anger == 30{
					fmt.Println("Герой зол, +20 к Его урону, +больше шансов промахнуться")
				}
				if hInfo.Anger == 30 && hh%2 == 0{
					randomNumofHero +=20
				}else if hInfo.Anger == 30 && hh%2 != 0 {
					randomNumofHero = 0
				}
				fmt.Println("Герой наносит  удар!","Кол-во ходов Героя: ",countHero)//сколько ударов нанес герой
				if hInfo.Anger == 30 && hh%2 != 0{
					fmt.Println("Герой промахнулся!!!")
				}
				hInfo.Fatique += 5//если герой нанес удар, то усталость +5
				dInfo.Anger +=5 //после каждого удара героя, злость дракона растет на 5
				if hInfo.Fatique == 30 && randomNumofHero%2 !=0 { //если усталость равна 30 и остаток от деления на 2 не равен 0, то есть шанс промахнуться
					fmt.Println("Герой промахнулся!")
					randomNumofHero = 0
					hInfo.Fatique = 0//обнуляем усталость героя
				}
				dInfo.Health = dInfo.Health - randomNumofHero//минус жизни дракона
				if dInfo.Health > 0 {
					fmt.Println ( dragHealthMsg, dInfo.Health )//инфа о уровни жизни дракона
				}
				
	        
				if dInfo.Health <= 0 { //проверка если уровень жизни дракона ноль, то герой выигрывает
					dInfo.Health = 0 //здесь присваиваем ноль чтобы жизнь не уходила в минус
					fmt.Println(heroHealthMsg,hInfo.Health)// показывает уровень жизни героя
					fmt.Println(dragHealthMsg,dInfo.Health," Количество ходов Дракона: ",countDrag)//уровень жизни героя и кол-во ходов дракона
					fmt.Println("Урааааа!!! Поздравляем, Вы выиграли!!! В качестве бонуса Вы получите крутую цитату...Ждите")
					randomQuote()
					break
				}
			}
			
		

			fmt.Println("\n")
			rand.Seed(time.Now().Unix())
	        missORnot = []int{ 15, 17, 2, 3, 5, 7, 9, 11, 13, }
	        n := rand.Int() % len(missORnot)
	        nn := missORnot[n]
	       

			if dInfo.Health > 0 && randomNumofDrag%2 !=0 {
				countDrag += 1//счтечик хода дракона
				if dInfo.Anger == 30{
					fmt.Println("Дракон зол, +20 к Его урону, +больше шансов промахнуться")
				}
				if dInfo.Anger == 30 && nn%2 == 0 { //если уровень злости равен 30 и остаток от деления равен нулю то +20 к урону
					randomNumofDrag +=20
				}else if dInfo.Anger == 30 && nn%2 != 0{
					randomNumofDrag = 0
				}
				fmt.Println("Дракон наносит Вам ответный удар","Количество ходов Дракона: ",countDrag)
				if dInfo.Anger == 30 && nn%2 != 0{
					fmt.Println("Дракон промахнулся!!!")
				}
				hInfo.Anger += 5
				dInfo.Fatique += 5//если дракон нанес удар, то усталость +5
				if dInfo.Fatique == 30 && randomNumofDrag%2 !=0 { //если усталость равна 30 и остаток от деления на 2 не равен 0, то есть шанс промахнуться
					fmt.Println("Дракон промахнулся!")
					randomNumofDrag = 0
					dInfo.Fatique = 0//обнуляем усталость дракона
				}
				hInfo.Health = hInfo.Health - randomNumofDrag
				if hInfo.Health > 0{
				  fmt.Println(heroHealthMsg,hInfo.Health)//показывает уровень жизни герояы
				}
				if hInfo.Health <= 0 { //проверака если уровень жизни героя ноль то дракон побеждает 
					hInfo.Health = 0   //здесь присваиваем ноль чтобы жизнь не уходила в минус
					fmt.Println(heroHealthMsg,hInfo.Health, "Кол-во ходов Героя: ", countHero)//инфа о жизни героя и о кол-ве ходов
					fmt.Println(dragHealthMsg,dInfo.Health) //инф-ия о жизни дракона
					fmt.Println("Неудача!!! Вы проиграли!!!")
					break
				}
			} else if randomNumofDrag%2 == 0{
				dInfo.Health += 10 //Плюс для жизни дракона
				if dInfo.Health > 100 {
				   dInfo.Health = 100 //присваиваем 100 чтобы жизнь не превышала 100 hp
				}
				fmt.Println("Дракон решил зализать раны")
				fmt.Println(dragHealthMsg,dInfo.Health, heroHealthMsg,hInfo.Health)//инф-ия о жизни персонажей
			}
		 }
		 if hHealthUp == true{
			hInfo.Health +=10 //плюс для жизни героя
			if hInfo.Health > 100{
				hInfo.Health = 100 //присваиваем 100 чтобы жизнь не превышала 100 hp
			}
			fmt.Println("Герой подлечился ", heroHealthMsg, hInfo.Health, dragHealthMsg,dInfo.Health) // инф-ия о жизни героя и дракона
		  }
	}
}

func level3() {
	heroHealthMsg := "Уровень жизни героя: "
	dragHealthMsg := "Уровень жизни дракона: "
	countHero := 0 //нужно для счета ходов Героя 
	countDrag := 0 //нужно для счета ходов Дракона
	weaponGObad := 1
	
	for {
		
	  
		rand.Seed(time.Now().UnixNano())
		randomNumofHero := random(1, 10)
		randomNumofDrag := random(1, 20)

	  
	actionHero()
	
    
		if hAct == true {
			
			if hInfo.Health > 0{ //герой наносит удар if true
				countHero += 1 //cчетчик
				
					
				randomNumofHero -= weaponGObad //оружие тупится с каждым ударом на 1
				if randomNumofHero < 0 {
					randomNumofHero = 0
				}
				rand.Seed(time.Now().Unix())
				missORnot := []int{ 15, 17, 2, 3, 5, 7, 9, 11, 13, 19,21,31,33,37,39, }
				n := rand.Int() % len(missORnot)
				nn := missORnot[n]
				if hInfo.Health < 80 && nn%2 == 0 {
					fmt.Println("Активирована броня +",hInfo.Armor,"hp")
					hInfo.Health +=hInfo.Armor
				}
				fmt.Println("Герой наносит  удар!","Кол-во ходов Героя: ",countHero)//сколько ударов нанес герой
				
				if nn%2 == 0 { //еслти тру, то +40 к урону героя
					randomNumofHero +=40
				}
				
				dInfo.Health = dInfo.Health - randomNumofHero//минус жизни дракона
				if dInfo.Health > 0 {
					fmt.Println ( dragHealthMsg, dInfo.Health )
				}
				if dInfo.Health <= 0 {
					dInfo.Health = 0
					fmt.Println(heroHealthMsg,hInfo.Health)// показывает уровень жизни героя
					fmt.Println(dragHealthMsg,dInfo.Health," Количество ходов Дракона: ",countDrag)//уровень жизни героя и кол-во ходов дракона
					fmt.Println("Урааааа!!! Поздравляем, Вы выиграли!!! В качестве бонуса Вы получите крутую цитату...Ждите")
					randomQuote()
					break
				}
			}
			
		  fmt.Println("\n")

			if dInfo.Health > 0 && randomNumofDrag%2 !=0 {
				countDrag += 1
				fmt.Println("Дракон наносит Вам ответный удар","Количество ходов Дракона: ",countDrag)
				
				hInfo.Health = hInfo.Health - randomNumofDrag
				if hInfo.Health > 0{
					fmt.Println(heroHealthMsg,hInfo.Health)//показывает уровень жизни герояы
				}
				if hInfo.Health <= 0 { //если уровень жизни героя ушел в минус
					hInfo.Health = 0   //то присвоем ноль 
					fmt.Println(heroHealthMsg,hInfo.Health, "Кол-во ходов Героя: ", countHero)//инфа о жизни героя и о кол-ве ходов
					fmt.Println(dragHealthMsg,dInfo.Health) //инф-ия о жизни дракона
					fmt.Println("Неудача!!! Вы проиграли!!!")
					break
				}
			} else if randomNumofDrag%2 == 0{
				dInfo.Health += 10 //Плюс для жизни дракона
				if dInfo.Health > 100 {
				   dInfo.Health = 100
				}
				fmt.Println("Дракон решил зализать раны")
				fmt.Println(dragHealthMsg,dInfo.Health, heroHealthMsg,hInfo.Health)//инф-ия о жизни персонажей
			}
		 }
		 if hHealthUp == true{
			hInfo.Health +=10
			if hInfo.Health > 100{
				hInfo.Health = 100
			}
			fmt.Println("Герой подлечился ", heroHealthMsg, hInfo.Health, dragHealthMsg,dInfo.Health) 
		 }
  }
}
func random(min int, max int) int {
	return rand.Intn(max-min) + min
}

func actionHero() string {
	
	fmt.Println("===========================================================")
	fmt.Println("1.Атаковать")
	fmt.Println("2.Лечиться")
	fmt.Println("Выберите действие героя - 1 для атаки, 2 для лечения")
	fmt.Println("===========================================================")
	   
	    action :=bufio.NewScanner(os.Stdin)
		action.Scan()
		heroAction := action.Text()
		switch heroAction{
		case "1":
		  hAct = true
		  hHealthUp = false
		case "2":
		 
		  hAct = false
		 
		if hInfo.Health == 100{
			fmt.Println("Доктор сказал, что Вы симулянт и с Вами все в порядке!!!")
	
		} else {
			hHealthUp = true
		}
	default:
		actionHero()
	
		}
		return heroAction
}
type Q struct{
	Quotes map[string]interface{} `json:"quotes"`
}
func randomQuote(){
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
package main
import (
 "fmt"
 "os"
 "bufio"
 "regexp"
 "math/rand"
 "time"
)
var setAutoDragName bool = false
var hName string // var HeroName
var dName, ddName string // var DragName
var dragHP, heroHP int
func main(){
	greeting()
	mainMenu()
	instruction()
	selectItemMenu()
	menuLevelOfGame()
	selectWeapon()
	fight1()
}

//fucntion greeting
func greeting(){
	fmt.Println("Добро пожаловать в игру Герой против Дракона!")
}

//function of Selecting item of menu
func selectItemMenu(){
	for {
		itemMenu :=bufio.NewScanner(os.Stdin)
		itemMenu.Scan()
		iMenu := itemMenu.Text()
		if iMenu == "1" {
			if setAutoDragName ==false{
				nameforperson()
				break
			}else if setAutoDragName == true{
				//fmt.Println("Задайте имя Вашего героя")
				nameOfHero()
				randomNameOfDrag()
				break
			}
			
		}else if iMenu == "2" {
			configNameOfDragMenu()
			break
		}else if (iMenu != "1" || iMenu != "2"){
			instruction()
		}
	}
}


//func main menu
func mainMenu(){
	fmt.Println("Главное меню")
	fmt.Println("============")
	fmt.Println("1.Начать игру")
	fmt.Println("2.Настройки")
}


//function of config dragon name
func configNameOfDragMenu(){
	fmt.Println("Нажмите цифру 1 для выбора имени дракона пользователем, цифру 2 для случайного имени")
	fmt.Println("1.Выбор пользователя")
	fmt.Println("2.Случайное имя дракона")
	
	for {
		selectSetting :=bufio.NewScanner(os.Stdin)
		selectSetting.Scan()
		sSetting := selectSetting.Text()
	
		if sSetting == "1" {
			fmt.Println("Вы выбрали функцию ввода имени Дракона пользователем")
			mainMenu()
			instruction()
			selectItemMenu()
			break
		}else if sSetting == "2" {
			setAutoDragName = true
			fmt.Println("Вы выбрали функцию случайного ввода имени Дракона")
			mainMenu()
			instruction()
			selectItemMenu()
			break
		}else if (sSetting != "1" || sSetting != "2"){
			fmt.Println("Нажмите цифру 1 для выбора имени Дракона пользователем, цифру 2 для случайного имени")
		}
	}
}


//function instruction
func instruction(){
	fmt.Println("Нажмите цифру 1 для начала игры или цифру 2 для настройки игры")
}
func randomNameOfDrag(){
    rand.Seed(time.Now().Unix())
    dragonsName := []string{
    "BlackDragon",
    "WhiteDragon",
    "NorthDragon",
    "SourthDragon",
    }
	n := rand.Int() % len(dragonsName)
	ddName = dragonsName[n]
    fmt.Println("Случайное имя дракона: ", dragonsName[n])
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
		hName = heroName.Text()
		space := regexp.MustCompile(`\s+`)
		s := space.ReplaceAllString(hName, "")
		
		if s == "" || len(s) == 0{
			fmt.Println("Имя Героя не может быть пустым")
		}
	if len(s)>0{
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
		dName = dragName.Text()
		space := regexp.MustCompile(`\s+`)
		d := space.ReplaceAllString(dName, "")
		
		if d == "" || len(d) == 0{
			fmt.Println("Имя Дракона не может быть пустым")
		}
	    if len(d)>0{
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
	for {
		itemMenu :=bufio.NewScanner(os.Stdin)
		itemMenu.Scan()
		iMenu := itemMenu.Text()
		if iMenu == "1" {
            fmt.Println("Легкий уровень")
		    break
		}else if iMenu == "2" {
			fmt.Println("Средний уровень")
			break
		}else if iMenu =="3"{
			fmt.Println("Сложный уровень")
			break
		}else if iMenu != "1" || iMenu != "2" || iMenu != "3" {
			fmt.Println("Нажмите цифру 1 для выбора легкого уровня, 2 для среднего уровня, 3 для сложного уровня")
		}
	}
}


func selectWeapon(){
	fmt.Println("Оружейный арсенал героя")
	fmt.Println("=======================")
	fmt.Println("1.Молот Тора")
	fmt.Println("2.Лук Соколиного глаза")
	fmt.Println("3.Меч Фродо Беггинса")
	fmt.Println("4.Молоток Вашего Соседа")

	fmt.Println("Нажмите цифру 1 для выбора Молота, 2 для выбора Лука, 3 для выбора меча, 4 для выбора Молотка")
	for {
		itemMenu :=bufio.NewScanner(os.Stdin)
		itemMenu.Scan()
		iMenu := itemMenu.Text()
		if iMenu == "1" {
			fmt.Println("Вы выбрали Молот Тора, и нахрена!? Вы же не сможете его поднять, выберите лучше Молоток!")
			fmt.Println("Нажмите цифру 1 для выбора Молота, 2 для выбора Лука, 3 для выбора меча, 4 для выбора Молотка")
		}else if iMenu == "2" {
			fmt.Println("Вы выбрали Лук Соколиного Глаза")
			break
		}else if iMenu =="3"{
			fmt.Println("Вы выбрали Меч Фродо Бегинса")
			break
		}else if iMenu == "4"{
			fmt.Println("Вы выбрали Молоток, теперь Ваш сосед не будет стучать им в 8 утра по выходным")
			break
		}else if iMenu != "1" || iMenu != "2" || iMenu != "3" {
			fmt.Println("Нажмите цифру 1 для выбора Молота, 2 для выбора Лука, 3 для выбора Меча, 4 для выбора Молотка")
		}
	}
}


func fight1(){
	fmt.Println("Погода ясная, осадки не ожидаются. Судья дает старт игре!!!")
	fmt.Println("Битва началась!!! Крики с трибун: ДРАКА,ДРАКА,ДРАКА!!!")
	if setAutoDragName == true {
		fmt.Println("Уровень жизни ",hName ," || ","Уровень жизни ",ddName)
	}else{
		fmt.Println("Уровень жизни ",hName ," || ","Уровень жизни ",dName)
	actionHero()
	}
}
func actionHero() {
	fmt.Println("===========================================================")
	fmt.Println("1.Атаковать")
	fmt.Println("2.Лечиться")
	fmt.Println("Выберите действие героя - 1 для атаки, 2 для лечения")
	fmt.Println("===========================================================")
	    check := 0
	    action :=bufio.NewScanner(os.Stdin)
		action.Scan()
		heroAction := action.Text()
		switch heroAction{
		case "1":
		  fmt.Println("Герой атакует")
		  
		case "2":
		  check = 1
		  if check == 1{
			  fmt.Println("Доктор сказал, что Вы симулянт и с Вами все в порядке!!!")
		  }else {
			  fmt.Println("Герой лечится") }
		  
		default:
			fmt.Println("Выберите действие героя - 1 для атаки, 2 для обороны")
			actionHero()
		}
		
}
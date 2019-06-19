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
)
var setAutoDragName bool = false
var hName string // var HeroName
var dName, ddName string // var DragName

func main(){
	greeting()
	mainMenu()
	selectItemMenu()
	menuLevelOfGame()
	selectWeapon()
	fight1()
}

//fucntion greeting
func greeting(){
	fmt.Println("Добро пожаловать в игру Герой против Дракона!!!")
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
				fmt.Println("Легкий уровень")
				
			case "2": 
				fmt.Println("Средний уровень")
				
			case "3":
				fmt.Println("Сложный уровень")
				
			default:
				menuLevelOfGame()
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
	
		itemMenu :=bufio.NewScanner(os.Stdin)
		itemMenu.Scan()
		iMenu := itemMenu.Text()
		switch iMenu {
        case "1":
			fmt.Println("Вы выбрали Молот Тора, и нахрена!? Вы же не сможете его поднять, выберите лучше Молоток!")
			fmt.Println("Нажмите цифру 1 для выбора Молота, 2 для выбора Лука, 3 для выбора меча, 4 для выбора Молотка")
		case "2":
			fmt.Println("Вы выбрали Лук Соколиного Глаза")
		case "3":
			fmt.Println("Вы выбрали Меч Фродо Бегинса")
		case "4":
			fmt.Println("Вы выбрали Молоток, теперь Ваш сосед не будет стучать им в 8 утра по выходным")
		default:
			fmt.Println("Нажмите цифру 1 для выбора Молота, 2 для выбора Лука, 3 для выбора Меча, 4 для выбора Молотка")
			selectWeapon()
		}
}


func fight1() {
	fmt.Println("Погода ясная, осадки не ожидаются. Судья дает старт игре!!!")
	fmt.Println("Битва началась!!! Крики с трибун: ДРАКА,ДРАКА,ДРАКА!!!")
	if setAutoDragName == true {
		fmt.Println("Уровень жизни ",hName ," || ","Уровень жизни ",ddName)
		actionHero()
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
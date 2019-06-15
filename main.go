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
func main(){
	greeting()
	mainMenu()
	instruction()
	selectItemMenu()
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
				fmt.Println("Задайте имя Вашего героя")
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
		hName := heroName.Text()
		space := regexp.MustCompile(`\s+`)
		s := space.ReplaceAllString(hName, "")
		
		if s == "" || len(s) == 0{
			fmt.Println("Имя Героя не может быть пустым")
		}
	if len(s)>0{
		fmt.Println("Имя героя: ", s)
		break
	}
	}
}


func nameOfDrag() {
	for {
		dragName := bufio.NewScanner(os.Stdin)
		fmt.Print("Введите имя Дракона: ")
		dragName.Scan()
		dName := dragName.Text()
		space := regexp.MustCompile(`\s+`)
		s := space.ReplaceAllString(dName, "")
		
		if s == "" || len(s) == 0{
			fmt.Println("Имя Дракона не может быть пустым")
		}
	    if len(s)>0{
		fmt.Println("Имя Дракона: ", dName)
		break
	    }
	}
}
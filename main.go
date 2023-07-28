package main

import (
	"fmt"
	"github.com/eiannone/keyboard"

	"math/rand"
	"time"
)

const (
	Street     string = "street"
	Prison            = "prison"
	Railway           = "railway"
	GoToPrison        = "go to prison"
	Start             = "start"
	Communal          = "communal apartment"
	Parking           = "parking"
	Coffer            = "coffer"
	Chance            = "chance"
	Tax               = "tax"
)

type card struct {
	name                 string
	cardType             string
	color                string
	price                int
	cardsCounter         int
	cardsForMonopoly     int
	railwayCounter       int
	rent                 int
	rentWithMonopoly     int
	rentWithOneHouse     int
	rentWithTwoHouses    int
	rentWithThreeHouses  int
	rentWithFourHouses   int
	rentWithHotel        int
	rentWithTwoRailway   int
	rentWithThreeRailway int
	rentWithFourRailway  int
	housePrice           int
	hotelPrice           int
	isBuyable            bool
	isBuyed              bool
	owner                player
	ownerID              int
	isMonopoly           bool
}

type player struct {
	playerID            int
	name                string
	balance             int
	brownForMonopoly    []card
	blueForMonopoly     []card
	pinkForMonopoly     []card
	orangeForMonopoly   []card
	redForMonopoly      []card
	yellowForMonopoly   []card
	greenForMonopoly    []card
	darkBlueForMonopoly []card
	railways            []card
	communals           []card
	cards               []card
	position            int
	communalCards       int
	railWayCards        int
	isPrisoner          bool
	prisonCounter       int
}

var field = []card{
	{
		name:      "Вперед",
		cardType:  Start,
		color:     "Белый",
		isBuyable: false,
	},
	{
		name:                "Житная ул.",
		cardType:            Street,
		color:               "Коричневый",
		price:               60,
		rent:                2,
		cardsForMonopoly:    2,
		rentWithMonopoly:    4,
		rentWithOneHouse:    10,
		rentWithTwoHouses:   30,
		rentWithThreeHouses: 90,
		rentWithFourHouses:  160,
		rentWithHotel:       250,
		housePrice:          50,
		hotelPrice:          50,
		isBuyable:           true,
		isBuyed:             false,
	},
	{
		name:      "Общественная казна",
		cardType:  Coffer,
		color:     "Белый",
		isBuyable: false,
	},
	{
		name:                "Нагатинская ул.",
		cardType:            Street,
		color:               "Коричневый",
		price:               60,
		rent:                4,
		cardsForMonopoly:    2,
		rentWithMonopoly:    8,
		rentWithOneHouse:    20,
		rentWithTwoHouses:   60,
		rentWithThreeHouses: 180,
		rentWithFourHouses:  320,
		rentWithHotel:       450,
		housePrice:          50,
		hotelPrice:          50,
		isBuyable:           true,
		isBuyed:             false,
	},
	{
		name:      "Подоходный налог",
		cardType:  Tax,
		color:     "Белый",
		price:     200,
		isBuyable: false,
	},
	{
		name:                 "Рижская железная дорога",
		cardType:             Railway,
		color:                "Белый",
		price:                200,
		rent:                 25,
		railwayCounter:       0,
		rentWithTwoRailway:   50,
		rentWithThreeRailway: 100,
		rentWithFourRailway:  200,
		isBuyable:            true,
		isBuyed:              false,
	},
	{
		name:                "Варшавское шоссе",
		cardType:            Street,
		color:               "Голубой",
		price:               100,
		rent:                6,
		cardsForMonopoly:    3,
		rentWithMonopoly:    12,
		rentWithOneHouse:    30,
		rentWithTwoHouses:   90,
		rentWithThreeHouses: 270,
		rentWithFourHouses:  400,
		rentWithHotel:       550,
		housePrice:          50,
		hotelPrice:          50,
		isBuyable:           true,
		isBuyed:             false,
	},
	{
		name:      "Шанс",
		cardType:  Chance,
		color:     "Белый",
		isBuyable: false,
	},
	{
		name:                "Ул. Огарева",
		cardType:            Street,
		color:               "Голубой",
		price:               100,
		rent:                6,
		cardsForMonopoly:    3,
		rentWithMonopoly:    12,
		rentWithOneHouse:    30,
		rentWithTwoHouses:   90,
		rentWithThreeHouses: 270,
		rentWithFourHouses:  400,
		rentWithHotel:       550,
		housePrice:          50,
		hotelPrice:          50,
		isBuyable:           true,
		isBuyed:             false,
	},
	{
		name:                "Первая парковая ул.",
		cardType:            Street,
		color:               "Голубой",
		price:               120,
		rent:                8,
		cardsForMonopoly:    3,
		rentWithMonopoly:    16,
		rentWithOneHouse:    40,
		rentWithTwoHouses:   100,
		rentWithThreeHouses: 300,
		rentWithFourHouses:  450,
		rentWithHotel:       600,
		housePrice:          50,
		hotelPrice:          50,
		isBuyable:           true,
		isBuyed:             false,
	},
	{
		name:      "Тюрьма",
		cardType:  Prison,
		color:     "Белый",
		isBuyable: false,
	},
	{
		name:                "Ул. Полянка",
		cardType:            Street,
		color:               "Розовый",
		price:               140,
		rent:                10,
		cardsForMonopoly:    3,
		rentWithMonopoly:    20,
		rentWithOneHouse:    50,
		rentWithTwoHouses:   150,
		rentWithThreeHouses: 450,
		rentWithFourHouses:  625,
		rentWithHotel:       750,
		housePrice:          100,
		hotelPrice:          100,
		isBuyable:           true,
		isBuyed:             false,
	},
	{
		name:             "Электростанция",
		cardType:         Communal,
		color:            "Белый",
		price:            150,
		rent:             Roll() * 4,
		cardsForMonopoly: 2,
		rentWithMonopoly: Roll() * 10,
		isBuyable:        true,
		isBuyed:          false,
	},
	{
		name:                "Ул. Сретенка",
		cardType:            Street,
		color:               "Розовый",
		price:               140,
		rent:                10,
		cardsForMonopoly:    3,
		rentWithMonopoly:    20,
		rentWithOneHouse:    50,
		rentWithTwoHouses:   150,
		rentWithThreeHouses: 450,
		rentWithFourHouses:  625,
		rentWithHotel:       750,
		housePrice:          100,
		hotelPrice:          100,
		isBuyable:           true,
		isBuyed:             false,
	},
	{
		name:                "Ростовская наб.",
		cardType:            Street,
		color:               "Розовый",
		price:               160,
		rent:                12,
		cardsForMonopoly:    3,
		rentWithMonopoly:    24,
		rentWithOneHouse:    60,
		rentWithTwoHouses:   180,
		rentWithThreeHouses: 500,
		rentWithFourHouses:  700,
		rentWithHotel:       900,
		housePrice:          100,
		hotelPrice:          100,
		isBuyable:           true,
		isBuyed:             false,
	},
	{
		name:                 "Курская железная дорога",
		cardType:             Railway,
		color:                "Белый",
		price:                200,
		rent:                 25,
		rentWithTwoRailway:   50,
		rentWithThreeRailway: 100,
		rentWithFourRailway:  200,
		isBuyable:            true,
		isBuyed:              false,
	},
	{
		name:                "Рязанский проспект",
		cardType:            Street,
		color:               "Оранжевый",
		price:               180,
		rent:                14,
		cardsForMonopoly:    3,
		rentWithMonopoly:    28,
		rentWithOneHouse:    70,
		rentWithTwoHouses:   200,
		rentWithThreeHouses: 550,
		rentWithFourHouses:  750,
		rentWithHotel:       950,
		housePrice:          100,
		hotelPrice:          100,
		isBuyable:           true,
		isBuyed:             false,
	},
	{
		name:      "Общественная казна",
		cardType:  Coffer,
		color:     "Белый",
		isBuyable: false,
	},
	{
		name:                "Вывилова",
		cardType:            Street,
		color:               "Оранжевый",
		price:               180,
		rent:                14,
		cardsForMonopoly:    3,
		rentWithMonopoly:    28,
		rentWithOneHouse:    70,
		rentWithTwoHouses:   200,
		rentWithThreeHouses: 550,
		rentWithFourHouses:  750,
		rentWithHotel:       950,
		housePrice:          100,
		hotelPrice:          100,
		isBuyable:           true,
		isBuyed:             false,
	},
	{
		name:                "Рублевское шоссе",
		cardType:            Street,
		color:               "Оранжевый",
		price:               200,
		rent:                16,
		cardsForMonopoly:    3,
		rentWithMonopoly:    32,
		rentWithOneHouse:    80,
		rentWithTwoHouses:   220,
		rentWithThreeHouses: 600,
		rentWithFourHouses:  800,
		rentWithHotel:       1000,
		housePrice:          100,
		hotelPrice:          100,
		isBuyable:           true,
		isBuyed:             false,
	},
	{
		name:      "Бесплатная парковка",
		cardType:  Parking,
		color:     "Белый",
		isBuyable: false,
	},
	{
		name:                "Тверская",
		cardType:            Street,
		color:               "Красный",
		price:               220,
		rent:                18,
		cardsForMonopoly:    3,
		rentWithMonopoly:    36,
		rentWithOneHouse:    90,
		rentWithTwoHouses:   250,
		rentWithThreeHouses: 700,
		rentWithFourHouses:  875,
		rentWithHotel:       1050,
		housePrice:          150,
		hotelPrice:          150,
		isBuyable:           true,
		isBuyed:             false,
	},
	{
		name:      "Шанс",
		cardType:  Chance,
		color:     "Белый",
		isBuyable: false,
	},
	{
		name:                "Пушкинская",
		cardType:            Street,
		color:               "Красный",
		price:               220,
		rent:                18,
		cardsForMonopoly:    3,
		rentWithMonopoly:    36,
		rentWithOneHouse:    90,
		rentWithTwoHouses:   250,
		rentWithThreeHouses: 700,
		rentWithFourHouses:  875,
		rentWithHotel:       1050,
		housePrice:          150,
		hotelPrice:          150,
		isBuyable:           true,
		isBuyed:             false,
	},
	{
		name:                "Площадь Маяковского",
		cardType:            Street,
		color:               "Красный",
		price:               240,
		rent:                20,
		cardsForMonopoly:    3,
		rentWithMonopoly:    40,
		rentWithOneHouse:    100,
		rentWithTwoHouses:   300,
		rentWithThreeHouses: 750,
		rentWithFourHouses:  925,
		rentWithHotel:       1100,
		housePrice:          150,
		hotelPrice:          150,
		isBuyable:           true,
		isBuyed:             false,
	},
	{
		name:                 "Казанская железная дорога",
		cardType:             Railway,
		color:                "Белый",
		price:                200,
		rent:                 25,
		rentWithTwoRailway:   50,
		rentWithThreeRailway: 100,
		rentWithFourRailway:  200,
		isBuyable:            true,
		isBuyed:              false,
	},
	{
		name:                "Ул. Грузинский Вал ",
		cardType:            Street,
		color:               "Желтый",
		price:               260,
		rent:                22,
		cardsForMonopoly:    3,
		rentWithMonopoly:    44,
		rentWithOneHouse:    110,
		rentWithTwoHouses:   330,
		rentWithThreeHouses: 800,
		rentWithFourHouses:  975,
		rentWithHotel:       1150,
		housePrice:          150,
		hotelPrice:          150,
		isBuyable:           true,
		isBuyed:             false,
	},
	{
		name:                "Новинский бульвар",
		cardType:            Street,
		color:               "Желтый",
		price:               260,
		rent:                22,
		cardsForMonopoly:    3,
		rentWithMonopoly:    44,
		rentWithOneHouse:    110,
		rentWithTwoHouses:   330,
		rentWithThreeHouses: 800,
		rentWithFourHouses:  975,
		rentWithHotel:       1150,
		housePrice:          150,
		hotelPrice:          150,
		isBuyable:           true,
		isBuyed:             false,
	},
	{
		name:             "Водопровод",
		cardType:         Communal,
		color:            "Белый",
		price:            150,
		rent:             Roll() * 4,
		cardsForMonopoly: 2,
		rentWithMonopoly: Roll() * 10,
		isBuyable:        true,
		isBuyed:          false,
	},
	{
		name:                "Смоленская площадь",
		cardType:            Street,
		color:               "Желтый",
		price:               280,
		rent:                24,
		cardsForMonopoly:    3,
		rentWithMonopoly:    48,
		rentWithOneHouse:    120,
		rentWithTwoHouses:   360,
		rentWithThreeHouses: 850,
		rentWithFourHouses:  1025,
		rentWithHotel:       1200,
		housePrice:          150,
		hotelPrice:          150,
		isBuyable:           true,
		isBuyed:             false,
	},
	{
		name:      "Отправляйтесь в тюрьму",
		cardType:  GoToPrison,
		color:     "Белый",
		isBuyable: false,
	},
	{
		name:                "Ул. Щусева",
		cardType:            Street,
		color:               "Зеленый",
		price:               300,
		rent:                26,
		cardsForMonopoly:    3,
		rentWithMonopoly:    52,
		rentWithOneHouse:    130,
		rentWithTwoHouses:   390,
		rentWithThreeHouses: 900,
		rentWithFourHouses:  1100,
		rentWithHotel:       1275,
		housePrice:          200,
		hotelPrice:          200,
		isBuyable:           true,
		isBuyed:             false,
	},
	{
		name:                "Гоголевский бульвар",
		cardType:            Street,
		color:               "Зеленый",
		price:               300,
		rent:                26,
		cardsForMonopoly:    3,
		rentWithMonopoly:    52,
		rentWithOneHouse:    130,
		rentWithTwoHouses:   390,
		rentWithThreeHouses: 900,
		rentWithFourHouses:  1100,
		rentWithHotel:       1275,
		housePrice:          200,
		hotelPrice:          200,
		isBuyable:           true,
		isBuyed:             false,
	},
	{
		name:      "Общественная казна",
		cardType:  Coffer,
		color:     "Белый",
		isBuyable: false,
	},
	{
		name:                "Кутузовский проспект",
		cardType:            Street,
		color:               "Зеленый",
		price:               320,
		rent:                26,
		cardsForMonopoly:    3,
		rentWithMonopoly:    56,
		rentWithOneHouse:    150,
		rentWithTwoHouses:   450,
		rentWithThreeHouses: 1000,
		rentWithFourHouses:  1200,
		rentWithHotel:       1400,
		housePrice:          200,
		hotelPrice:          200,
		isBuyable:           true,
		isBuyed:             false,
	},
	{
		name:                 "Ленинградская железная дорога",
		cardType:             Railway,
		color:                "Белый",
		price:                200,
		rent:                 25,
		rentWithTwoRailway:   50,
		rentWithThreeRailway: 100,
		rentWithFourRailway:  200,
		isBuyable:            true,
		isBuyed:              false,
	},
	{
		name:      "Шанс",
		cardType:  Chance,
		color:     "Белый",
		isBuyable: false,
	},
	{
		name:                "Ул. Малая бронная",
		cardType:            Street,
		color:               "Синий",
		price:               350,
		rent:                35,
		cardsForMonopoly:    2,
		rentWithMonopoly:    70,
		rentWithOneHouse:    175,
		rentWithTwoHouses:   500,
		rentWithThreeHouses: 1100,
		rentWithFourHouses:  1300,
		rentWithHotel:       1500,
		housePrice:          200,
		hotelPrice:          200,
		isBuyable:           true,
		isBuyed:             false,
	},
	{
		name:      "Сверхналог",
		cardType:  Tax,
		color:     "Белый",
		price:     100,
		isBuyable: false,
	},
	{
		name:                "Ул. Арбат",
		cardType:            Street,
		color:               "Синий",
		price:               400,
		rent:                50,
		cardsForMonopoly:    3,
		rentWithMonopoly:    100,
		rentWithOneHouse:    200,
		rentWithTwoHouses:   600,
		rentWithThreeHouses: 1400,
		rentWithFourHouses:  1700,
		rentWithHotel:       2000,
		housePrice:          200,
		hotelPrice:          200,
		isBuyable:           true,
		isBuyed:             false,
	},
}
var numOfPlayers int
var counter = 0
var players []player
var playerName string

func main() {

	rand.Seed(time.Now().UnixNano())
	fmt.Println("Вас приветствует игра Монополия")

	fmt.Println("Выберите кол-во игроков от 2 до 4")
	fmt.Scan(&numOfPlayers)
	if numOfPlayers < 2 || numOfPlayers > 4 {
		for {
			if numOfPlayers < 2 || numOfPlayers > 4 {
				fmt.Println("Некорректный ввод ")
				fmt.Scan(&numOfPlayers)
				continue
			}
			break
		}
	}

	CreatePlayer(numOfPlayers)
	fmt.Println("Теперь с помощью кубиков определим, кто будет начинать игру")
	WhoFirst(numOfPlayers)

	for {
		Game()
		fmt.Println("Нажмите любую копку для продолжения ")
		keyboard.GetSingleKey()

	}

}

func WhoFirst(a int) {
	var arrOfPlayers []int
	var max, maxind int
	switch a {
	case 2:

		for i := 0; i < 2; i++ {
			fmt.Println("______________________________________________________________")
			fmt.Println("Введите имя игрока ", i+1)
			fmt.Scan(&playerName)
			players[i].name = playerName
			players[i].playerID = i + 1

			fmt.Printf("игрок %s кидает кубики\n", players[i].name)
			fmt.Println("______________________________________________________________")
			arrOfPlayers = append(arrOfPlayers, rand.Intn(12)+1)
			time.Sleep(1 * time.Second)
			fmt.Printf("игрок %s выкинул %d\n", players[i].name, arrOfPlayers[i])
			max = arrOfPlayers[i]

		}

		for ind, v := range arrOfPlayers {
			if v >= max {
				max = v
				maxind = ind
			}
		}

		time.Sleep(1 * time.Second)
		fmt.Println("______________________________________________________________")
		fmt.Printf("Игру начинает игрок %s\n", players[maxind].name)

		counter = maxind

	case 3:
		for i := 0; i < 3; i++ {
			fmt.Println("______________________________________________________________")
			fmt.Println("Введите имя игрока ", i+1)
			fmt.Scan(&playerName)
			players[i].name = playerName
			players[i].playerID = i + 1

			fmt.Printf("игрок %s кидает кубики\n", players[i].name)
			fmt.Println("______________________________________________________________")
			arrOfPlayers = append(arrOfPlayers, rand.Intn(12)+1)
			time.Sleep(1 * time.Second)
			fmt.Printf("игрок %s выкинул %d\n", players[i].name, arrOfPlayers[i])
			max = arrOfPlayers[i]
		}

		for ind, v := range arrOfPlayers {
			if v >= max {
				max = v
				maxind = ind

			}
		}
		time.Sleep(1 * time.Second)
		fmt.Println("______________________________________________________________")
		fmt.Printf("Игру начинает игрок %s\n", players[maxind].name)

		counter = maxind

	case 4:
		for i := 0; i < 4; i++ {
			fmt.Println("______________________________________________________________")
			fmt.Println("Введите имя игрока ", i+1)
			fmt.Scan(&playerName)
			players[i].name = playerName
			players[i].playerID = i + 1

			fmt.Printf("игрок %s кидает кубики\n", players[i].name)
			fmt.Println("______________________________________________________________")
			arrOfPlayers = append(arrOfPlayers, rand.Intn(12)+1)
			time.Sleep(1 * time.Second)
			fmt.Printf("игрок %s выкинул %d\n", players[i].name, arrOfPlayers[i])
			max = arrOfPlayers[i]
		}

		for ind, v := range arrOfPlayers {
			if v >= max {
				max = v
				maxind = ind

			}
		}

		time.Sleep(1 * time.Second)
		fmt.Println("______________________________________________________________")
		fmt.Printf("Игру начинает игрок %s\n", players[maxind].name)

		counter = maxind
	}
}

func CounterSum(numOfPlayers int, d *int) {

	if *d+1 == numOfPlayers {
		*d = 0
	} else {
		*d++
	}

}

func CreatePlayer(numOfPlayers int) {
	var gamer = player{balance: 1500}
	for i := 0; i < numOfPlayers; i++ {
		players = append(players, gamer)
	}
}

func FirstCube() int {
	return rand.Intn(6) + 1

}
func SecondCube() int {
	return rand.Intn(6) + 1

}
func Roll() int {
	return FirstCube() + SecondCube()

}

var choose int

func Game() {

	move := 0
	move = Roll()
	players[counter].position += move
	if players[counter].position > 39 {
		players[counter].position = players[counter].position - 39
	}
	if players[counter].isPrisoner == true {
		players[counter].prisonCounter--
		InPrison()
	} else {
		if field[players[counter].position].price == 0 {
			fmt.Println("______________________________________________________________")
			fmt.Print("Игрок ", players[counter].name, " выкинул ", move, " и переместился на поле ", field[players[counter].position].name, "\n")
		} else if field[players[counter].position].cardType == Tax {
			taxStanding(move)
		} else {
			fieldBuying(move)
		}
		if players[counter].position == 30 {
			goToPrison(move)

		}

	}

	CounterSum(numOfPlayers, &counter)
	fmt.Println("______________________________________________________________")
	fmt.Print("Ход игрока ", players[counter].name, "\n", "\n")

}

func Monopoly() {

	if players[counter].cards[len(players[counter].cards)-1].color == "Коричневый" && players[counter].cards[len(players[counter].cards)-1].cardsForMonopoly == len(players[counter].brownForMonopoly) {
		fmt.Println("Вы собрали монополию коричневого цвета ")
		j := 0

		for i := 0; i < len(field); i++ {
			for k := 0; k < len(players[counter].brownForMonopoly); k++ {
				if players[counter].brownForMonopoly[k].name == field[i].name {
					field[i].rent = field[i].rentWithMonopoly
					field[i].isMonopoly = true
					j++
					if j == field[i].cardsForMonopoly {
						break
					}
				}
			}

		}

	} else if players[counter].cards[len(players[counter].cards)-1].color == "Голубой" && players[counter].cards[len(players[counter].cards)-1].cardsForMonopoly == len(players[counter].blueForMonopoly) {
		fmt.Println("Вы собрали монополию голубого цвета ")
		j := 0

		for i := 0; i < len(field); i++ {
			for k := 0; k < len(players[counter].blueForMonopoly); k++ {
				if players[counter].blueForMonopoly[k].name == field[i].name {
					field[i].rent = field[i].rentWithMonopoly
					field[i].isMonopoly = true
					j++
					if j == field[i].cardsForMonopoly {
						break
					}
				}
			}

		}

	} else if players[counter].cards[len(players[counter].cards)-1].color == "Розовый" && players[counter].cards[len(players[counter].cards)-1].cardsForMonopoly == len(players[counter].pinkForMonopoly) {
		fmt.Println("Вы собрали монополию розового цвета ")
		j := 0

		for i := 0; i < len(field); i++ {
			for k := 0; k < len(players[counter].pinkForMonopoly); k++ {
				if players[counter].pinkForMonopoly[k].name == field[i].name {
					field[i].rent = field[i].rentWithMonopoly
					field[i].isMonopoly = true
					j++
					if j == field[i].cardsForMonopoly {
						break
					}
				}
			}

		}

	} else if players[counter].cards[len(players[counter].cards)-1].color == "Оранжевый" && players[counter].cards[len(players[counter].cards)-1].cardsForMonopoly == len(players[counter].orangeForMonopoly) {
		fmt.Println("Вы собрали монополию оранжевого цвета ")
		j := 0

		for i := 0; i < len(field); i++ {
			for k := 0; k < len(players[counter].orangeForMonopoly); k++ {
				if players[counter].orangeForMonopoly[k].name == field[i].name {
					field[i].rent = field[i].rentWithMonopoly
					field[i].isMonopoly = true
					j++
					if j == field[i].cardsForMonopoly {
						break
					}
				}
			}

		}

	} else if players[counter].cards[len(players[counter].cards)-1].color == "Красный" && players[counter].cards[len(players[counter].cards)-1].cardsForMonopoly == len(players[counter].redForMonopoly) {
		fmt.Println("Вы собрали монополию красного цвета ")
		j := 0

		for i := 0; i < len(field); i++ {
			for k := 0; k < len(players[counter].redForMonopoly); k++ {
				if players[counter].redForMonopoly[k].name == field[i].name {
					field[i].rent = field[i].rentWithMonopoly
					field[i].isMonopoly = true
					j++
					if j == field[i].cardsForMonopoly {
						break
					}
				}
			}

		}

	} else if players[counter].cards[len(players[counter].cards)-1].color == "Желтый" && players[counter].cards[len(players[counter].cards)-1].cardsForMonopoly == len(players[counter].yellowForMonopoly) {
		fmt.Println("Вы собрали монополию желтого цвета ")
		j := 0

		for i := 0; i < len(field); i++ {
			for k := 0; k < len(players[counter].yellowForMonopoly); k++ {
				if players[counter].yellowForMonopoly[k].name == field[i].name {
					field[i].rent = field[i].rentWithMonopoly
					field[i].isMonopoly = true
					j++
					if j == field[i].cardsForMonopoly {
						break
					}
				}
			}

		}

	} else if players[counter].cards[len(players[counter].cards)-1].color == "Зеленый" && players[counter].cards[len(players[counter].cards)-1].cardsForMonopoly == len(players[counter].greenForMonopoly) {
		fmt.Println("Вы собрали монополию зеленого цвета ")
		j := 0

		for i := 0; i < len(field); i++ {
			for k := 0; k < len(players[counter].greenForMonopoly); k++ {
				if players[counter].greenForMonopoly[k].name == field[i].name {
					field[i].rent = field[i].rentWithMonopoly
					field[i].isMonopoly = true
					j++
					if j == field[i].cardsForMonopoly {
						break
					}
				}
			}

		}

	} else if players[counter].cards[len(players[counter].cards)-1].color == "Синий" && players[counter].cards[len(players[counter].cards)-1].cardsForMonopoly == len(players[counter].darkBlueForMonopoly) {
		fmt.Println("Вы собрали монополию синего цвета ")
		j := 0

		for i := 0; i < len(field); i++ {
			for k := 0; k < len(players[counter].darkBlueForMonopoly); k++ {
				if players[counter].darkBlueForMonopoly[k].name == field[i].name {
					field[i].rent = field[i].rentWithMonopoly
					field[i].isMonopoly = true
					j++
					if j == field[i].cardsForMonopoly {
						break
					}
				}
			}

		}

	} else if players[counter].cards[len(players[counter].cards)-1].cardType == Communal {
		j := 0
		switch len(players[counter].communals) {
		case 2:

			for i := 0; i < len(field); i++ {
				for k := 0; k < len(players[counter].communals); k++ {
					if players[counter].communals[k].name == field[i].name {
						field[i].rent = field[i].rentWithMonopoly
						field[i].isMonopoly = true
						j++
						if j == field[i].cardsForMonopoly {
							break
						}
					}
				}

			}

		}
	} else if players[counter].cards[len(players[counter].cards)-1].cardType == Railway {
		j := 0
		switch len(players[counter].railways) {
		case 2:
			for i := 0; i < len(field); i++ {
				for k := 0; k < len(players[counter].railways); k++ {
					if players[counter].railways[k].name == field[i].name {
						field[i].rent = field[i].rentWithTwoRailway
						field[i].isMonopoly = true
						j++
						if j == 2 {
							break
						}
					}
				}
			}
		case 3:
			for i := 0; i < len(field); i++ {
				for k := 0; k < len(players[counter].railways); k++ {
					if players[counter].railways[k].name == field[i].name {
						field[i].rent = field[i].rentWithThreeRailway
						field[i].isMonopoly = true
						j++
						if j == 3 {
							break
						}
					}
				}
			}
		case 4:
			for i := 0; i < len(field); i++ {
				for k := 0; k < len(players[counter].railways); k++ {
					if players[counter].railways[k].name == field[i].name {
						field[i].rent = field[i].rentWithFourRailway
						field[i].isMonopoly = true
						j++
						if j == 4 {
							break
						}
					}
				}
			}
		}
	}
}

func Auction(fieldPrice int) {

	var choice, upPrice int
	var arrOfVoters []player

	for i := 0; i < len(players); i++ {
		arrOfVoters = append(arrOfVoters, players[i])
	}

	arrOfVoters = Remove(arrOfVoters, counter)

	i := 0

	for noCounter := 0; noCounter <= len(players)-1; {

		if fieldPrice != field[players[counter].position].price && len(arrOfVoters) == 1 {

			fmt.Println("Игрок ", arrOfVoters[i].name, " купил поле за ", fieldPrice)
			fmt.Println("______________________________________________________________")
			players[i].balance -= fieldPrice
			players[i].cards = append(players[i].cards, field[players[counter].position])
			field[players[counter].position].isBuyed = true

			field[players[counter].position].ownerID = arrOfVoters[i].playerID
			ColoredCardsOwn()
			Monopoly()
			return

		} else if len(arrOfVoters) == 1 {
			fmt.Println("______________________________________________________________")
			fmt.Printf("%s Хотите ли вы купить это поле за %d?\n", arrOfVoters[i].name, fieldPrice)
			fmt.Println("1.Да    2.Нет")

			fmt.Scan(&choice)
			switch choice {
			case 1:
				fmt.Println("______________________________________________________________")
				fmt.Println("Игрок ", arrOfVoters[i].name, " купил поле за ", fieldPrice)
				players[counter].balance -= fieldPrice
				players[counter].cards = append(players[counter].cards, field[players[counter].position])
				field[players[counter].position].isBuyed = true
				field[players[counter].position].ownerID = players[i].playerID
				Monopoly()
				ColoredCardsOwn()
				fmt.Println("______________________________________________________________")
				return
			case 2:
				fmt.Println("______________________________________________________________")
				fmt.Println("Поле осталось некупленным ")
				return
				fmt.Println("______________________________________________________________")
			}
		}
		fmt.Println("______________________________________________________________")
		fmt.Println(arrOfVoters[i].name, " Хотите ли вы поднять цену? Сейчас торги ведутся за ", fieldPrice, "\nВаш баланс: ", arrOfVoters[i].balance)
		fmt.Println("1.Да    2.Нет")

		fmt.Scan(&choice)
		switch choice {
		case 1:
			fmt.Println("______________________________________________________________")
			fmt.Println("На сколько вы хотите поднять цену? Ваш баланс: ", arrOfVoters[i].balance)
			fmt.Scan(&upPrice)
			for upPrice+fieldPrice > arrOfVoters[i].balance {
				fmt.Println("______________________________________________________________")
				fmt.Println("Введите сумму, которую позволяет ваш баланс")
				fmt.Scan(&upPrice)
			}
			fieldPrice += upPrice
		case 2:
			noCounter++
			arrOfVoters = Remove(arrOfVoters, i)

		}
		i++
		if i > len(arrOfVoters)-1 {
			i = 0
		}
		if arrOfVoters[i].balance <= fieldPrice {
			arrOfVoters = Remove(arrOfVoters, i)
		}

	}

}

func ColoredCardsOwn() {
	switch field[players[counter].position].color {
	case "Коричневый":
		players[counter].brownForMonopoly = append(players[counter].brownForMonopoly, field[players[counter].position])
	case "Голубой":
		players[counter].blueForMonopoly = append(players[counter].blueForMonopoly, field[players[counter].position])
	case "Розовый":
		players[counter].pinkForMonopoly = append(players[counter].pinkForMonopoly, field[players[counter].position])
	case "Оранжевый":
		players[counter].orangeForMonopoly = append(players[counter].orangeForMonopoly, field[players[counter].position])
	case "Красный":
		players[counter].redForMonopoly = append(players[counter].redForMonopoly, field[players[counter].position])
	case "Желтый":
		players[counter].yellowForMonopoly = append(players[counter].yellowForMonopoly, field[players[counter].position])
	case "Зеленый":
		players[counter].greenForMonopoly = append(players[counter].greenForMonopoly, field[players[counter].position])
	case "Синий":
		players[counter].darkBlueForMonopoly = append(players[counter].darkBlueForMonopoly, field[players[counter].position])
	}

	switch field[players[counter].position].cardType {
	case Railway:
		players[counter].railways = append(players[counter].railways, field[players[counter].position])
	case Communal:
		players[counter].communals = append(players[counter].communals, field[players[counter].position])
	}

}

func fieldBuying(move int) {
	if field[players[counter].position].isBuyed == true && field[players[counter].position].ownerID != players[counter].playerID {
		fmt.Println("______________________________________________________________")
		fmt.Println("Вы встали на владения игрока ", players[field[players[counter].position].ownerID-1].name, "\n", "Заплатите ему ренту в размере ", field[players[counter].position].rent)
		fmt.Println("______________________________________________________________")
		fmt.Scan()
		if players[counter].balance >= field[players[counter].position].rent {
			players[counter].balance -= field[players[counter].position].rent
			players[field[players[counter].position].ownerID-1].balance += field[players[counter].position].rent
			fmt.Println("Вы заплатили ренту в размере ", field[players[counter].position].rent, ", теперь ваш баланс: ", players[counter].balance)
		} else {
			fmt.Println("______________________________________________________________")
			fmt.Println("У вас недостаточно средств для оплаты ренты. Ваш баланс:", players[counter].balance, "Цена ренты", field[players[counter].position].rent)

		}
	} else {
		if field[players[counter].position].ownerID == players[counter].playerID {
			fmt.Println("Вы встали на свои владения ", field[players[counter].position].name)
			fmt.Println(players[counter].playerID)
			fmt.Println(field[players[counter].position].ownerID)
		} else {
			fmt.Println("______________________________________________________________")
			fmt.Print("Игрок ", players[counter].name, " выкинул ", move, " и переместился на поле ", field[players[counter].position].name, "\nЦена этого поля: ", field[players[counter].position].price, "\n\n")

			fmt.Println("Хотите ли вы купить это поле? \nВаш баланс: ", players[counter].balance, "\n")
			fmt.Println("1. Да        2. Нет", "\n")
			fmt.Scan(&choose)
			if choose == 1 {
				players[counter].balance -= field[players[counter].position].price
				players[counter].cards = append(players[counter].cards, field[players[counter].position])
				field[players[counter].position].isBuyed = true
				ColoredCardsOwn()
				Monopoly()
				field[players[counter].position].ownerID = players[counter].playerID
				fmt.Println(players[counter].playerID)
				fmt.Println("______________________________________________________________")
				fmt.Println("Поздравляем игрока ", players[counter].name, "с покупкой! ")
				fmt.Println("Ваш баланс теперь: ", players[counter].balance, "\n")
				fmt.Println("______________________________________________________________")
			} else {
				fmt.Println("Поскольку игрок ", players[counter].name, "отказался покупать клетку, объявляется аукцион")
				fmt.Println("______________________________________________________________")
				Auction(field[players[counter].position].price)
			}
		}

	}
}

func taxStanding(move int) {
	fmt.Println("______________________________________________________________")
	fmt.Print("Игрок ", players[counter].name, " выкинул ", move, " и переместился на поле ", field[players[counter].position].name, "\nОн должен заплатить: ", field[players[counter].position].price, "\n")
	fmt.Println("Заплалтить?")
	fmt.Println("1. Да")
	fmt.Println("______________________________________________________________")
	fmt.Scan(&choose)
	if choose == 1 {
		players[counter].balance -= field[players[counter].position].price
		fmt.Println("Ваш баланс теперь: ", players[counter].balance, "\n")

	}
}

func Remove(slice []player, s int) []player {

	// 1. Выполнить сдвиг a[i+1:] влево на один индекс.
	copy(slice[s:], slice[s+1:])

	// 2. Удалить последний элемент (записать нулевое значение).
	slice[len(slice)-1] = player{}

	// 3. Усечь срез.
	slice = slice[:len(slice)-1]
	return slice
}

var choosePrison int

func goToPrison(move int) {
	players[counter].prisonCounter = 2
	players[counter].position = 9
	players[counter].isPrisoner = true
	fmt.Println("Игрок ", players[counter].name, "Попадает в тюрьму ")
	if players[counter].isPrisoner == true {
		move = 0
		fmt.Println("Чтобы выйти, вы можете заплатить 50 или выбросить дубль на кубиках ")
		fmt.Println("1. Бросить кубики\n2. Заплатить 50 ")
		fmt.Scan(&choosePrison)
		switch choosePrison {
		case 1:
			fmt.Println("Выпало:", FirstCube(), " ", SecondCube())
			if FirstCube() == SecondCube() {
				fmt.Println("Вы выходите из тюрьмы ")
				players[counter].isPrisoner = false
			} else {
				fmt.Println("Не повезло")
			}
		case 2:
			players[counter].balance -= 50
			fmt.Println("Вы свободны ")
			players[counter].isPrisoner = false
		}

	}

}
func InPrison() {

	fmt.Println(players[counter].prisonCounter)
	if players[counter].prisonCounter == 0 {
		players[counter].isPrisoner = false
		fmt.Println("Вы выходите из тюрьмы ")
		return
	} else {
		fmt.Println("Вы в тюрье ")
		fmt.Println("Чтобы выйти, вы можете заплатить 50 или выбросить дубль на кубиках ")
		fmt.Println("1. Бросить кубики\n2. Заплатить 50 ")
		fmt.Scan(&choosePrison)
		switch choosePrison {
		case 1:
			fmt.Println("Выпало:", FirstCube(), " ", SecondCube())
			if FirstCube() == SecondCube() {
				fmt.Println("Вы выходите из тюрьмы ")
				players[counter].isPrisoner = false
			} else {
				fmt.Println("Не повезло")
			}
		case 2:
			players[counter].balance -= 50
			fmt.Println("Вы свободны ")
			players[counter].isPrisoner = false
		}
	}
}

package main

import "fmt"

// пакет используется для проверки ответа, не удаляйте его

// пакет используется для проверки ответа, не удаляйте его



func main() {
	value1, value2, operation := readTask() // исходные данные получаются с помощью этой функции
    										// все полученные значения имеют тип пустого интерфейса
	switch value1.(type){//1 уровень
	case float64:
		switch value2.(type){//2 уровень
			case float64:
				switch operation.(type){//3 уровень
					case string:
						switch operation{//3.5 уровень
							case "+":
							answer:=value1.(float64)+value2.(float64)
							fmt.Printf("%.4f",answer)
							case "-":
							answer:=value1.(float64)-value2.(float64)
							fmt.Printf("%.4f",answer)
							case "/":
							answer:=value1.(float64)/value2.(float64)
							fmt.Printf("%.4f",answer)
							case "*":
							answer:=value1.(float64)*value2.(float64)
							fmt.Printf("%.4f",answer)
							default:
							fmt.Println("неизвестная операция")
						}

					default:
						fmt.Println("неизвестная операция")
				}
			default:
				fmt.Printf("value:=%v: %T",value2,value2)
		}
	default:
		fmt.Printf("value=%v: %T",value1,value1)
	}									  
}

func readTask() (value1, value2, operation interface{}) {

	return 5.3, 5.2, "0939" //тут играемся с параметрами
	
	}

/*
Анонимные функции:
fmt.Println(strings.Map(func(r rune) rune{
		if unicode.IsDigit(r){
			return '*'
		}else{
			return r
		}
	},"123123привет1233"))

То есть мы можем вызывать функцию непосредственно при вызове другой, просто прописывать внутри что делает эта функция
*/

/*
Пример передачи функции как аргумента для другой:
func main() {
	fmt.Println(strings.Map(digitToStar,"привет1пока2я3тут4посижу5"))
}

func digitToStar(x rune)rune{
	if unicode.IsDigit(x){
		return '*'
	}else{
		return x
	}
}
*/

/*
Преобразование типов данных:
1) целочисленные типы между собой:
var number int8 = 15
var numberHigh int64 = int64(number)
Так же можно и в более легкий тип, но важно учитывать, что при больших значениях изначального числа может возникнуть ошибка
2) целочисленные и числа с плавающей точкой:
При приведении можно так же, как и в п.1, но важно учитывать, что данные будут теряться:
var x float64 = 25.7
var y int = int(x)
//y=25

Так же, если в делении участвует хотя бы одно число с плавающей точкой - результат будет числом с плавающей точкой

3) целочисленное в строку
a) strconv.Itoa(int x) => на выходе x типа "string"
b) strconv.FormatInt(переменная int, система счистления(2,4,8,10,16..)) => на выходе переменная в нужно системе счистления
4) целые беззнаковые числа в строку
var a uint64 = 10101
res := strconv.FormatUint(a, 10)
fmt.Println(res) // 10101
5) Числа с плавающей запятой в строку
var a float64 = 1.0123456789

	// 1 параметр - число для конвертации
	// fmt - форматирование
	// prec - точность (кол-во знаков после запятой)
	// bitSize - 32 или 64 (32 для float32, 64 для float64)
	fmt.Println(strconv.FormatFloat(a, 'f', 2, 64)) // 1.01

	// если мы хотим учесть все цифры после запятой, то можем в prec передать -1
	fmt.Println(strconv.FormatFloat(a, 'f', -1, 64)) // 1.0123456789

	// Возможные форматы fmt:
	// 'f' (-ddd.dddd, no exponent),
	// 'b' (-ddddp±ddd, a binary exponent),
	// 'e' (-d.dddde±dd, a decimal exponent),
	// 'E' (-d.ddddE±dd, a decimal exponent),
	// 'g' ('e' for large exponents, 'f' otherwise),
	// 'G' ('E' for large exponents, 'f' otherwise),
	// 'x' (-0xd.ddddp±ddd, a hexadecimal fraction and binary exponent), or
	// 'X' (-0Xd.ddddP±ddd, a hexadecimal fraction and binary exponent).
	var b float64 = 2222 * 1023 * 245 * 2 * 52
	fmt.Println(strconv.FormatFloat(b, 'e', -1, 64)) // 5.791874088e+10
6)bool в строку
var a = true
res := strconv.FormatBool(a)
fmt.Println(res)     	// true
fmt.Printf("%T", res)   // string

7) Строки в целые числа
функция strconv.Atoi(x int)
Вызывать нужно вместе с ошибкой:
x, err := strconv.Atoi(chysloVStroke string)
if err!=nil{
	panic(err)
}
fmt.Println(x) //int

8) string в float
функция: strconv.ParseFloat(s string, разрядность (32||64))

*/

/*
Создание отображений (правильные способы)
1) m1 := make(map[int]int) - int как пример, может быть любой другой тип
2) m2 := map[int]int{
	12:2,
	1:5,
}
ВАЖНО
Не стоит задавать мап просто через var m0 map[int]int
При последующей попытке добавления пары возникнет ошибка, которая вызовет панику

Функции для работы с map:
1) delete(имя_карты, ключ)
*/

/*
Ввод с клавиатуры:
Вводить можно просто через пробел - IDE поймет, что это разные элементы.
Пример:
for i:=0; i < 10; i++{
	fmt.Scan(&text) //если вводелить через проблем 10 чисел - будет проведено 10 итераций
}
*/

/*
Что бы связать два файла нужно подвязать их под один пакет
при этом если в файле main2.go будет использоваться и импортироваться пакет math - при вызове функции второго файла из первого
мы можем не импортировать в первый файл пакеты, использованные во втором файле.
*/

/*
Синтаксис функции:
func имя_функции (Список параметров) (список возвращаемых значений){
	Выполняемые параметры
}
*/

/*
Пример функции:
func hello(x int, y int){
	fmt.Println("Hello")
}
*/

/*
Что бы изучить какую-то функцию нужно вконсоли прописать go doc fmt.Println
*/

/*
Инициализация среза
var a []int
var b []int = []int{1, 2, 3}
c := []int{1, 2, 3}
d := []int{1: 12}
*/

/*
Советы по выгрузке:

1)
Если вы хотите сделать все правильно, то алгоритм ваших действий  примерно таков:
Создать проект в папке /src/ваш_любой_ник/имя_проекта.
Если у вас уже есть github аккаунт то можете создать проект так: /src/github.com/username/имя_проекта.
2)

*/
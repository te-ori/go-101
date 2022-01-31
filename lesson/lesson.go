package lesson

import "fmt"

// Normalde `Score` bilgisini doğrudan `uint8` tipinde de tanımlayablirdik.
// `Score` dediğimiz şey şimdi de `uint8`'den farklı bir şey değil. Yine
// 0 ile 255 arasındaki pozitif tamsayı değerleri alabilir. Fakat `re-type`
// yaparak daha *anlamlı* ve *anlaşılabilir* bir tip tanımı yaptık.
type Score uint8

/*
Eğer tipimiz sadece primitif bir tip olsaydı o tipteki değişkenler üzerinde  çağrılabilecek, bu tip bir fonksiyon tanımlayamayacaktık. Bunun yerine

 func Grade(score uint8) string {
	 //...
 }

şeklinde de bir fonksiyon kullanabilirdik. Fakat `Score` u tanımlayarak *bağlamı* daha da kuvvetlendirdik. *level of abstraction*'ı yani *genellemenin seviyesini* azaltarak anlaşılabilirliği arttırdık. Böylece kodu okuyan kişi açısından `uint8` gibi çok daha genel bir veri yerine, `Score` gibi `Lesson` bağlamında çok daha açık bir anlama sahip tipimiz oldu. Ayrıca doğrudan `Score` tipi üzerine tanımlayabildiğimiz metodlarımız sayesinde yine amacını daha iyi ifade eden fonksiyonlarımız oldu. Son olarak da intellisense'e sahip IDE'lerde `Score` için tanımlanmış tüm metodları görebiliriz. Diğer primitif tiplerde böyle bir şansımız olmayacak.
*/
func (s Score) Grade() string {
	if s >= 85 {
		return "A"
	} else if s >= 60 {
		return "B"
	} else if s >= 45 {
		return "C"
	} else if s > 20 {
		return "D"
	} else {
		return "F"
	}
}

type Lesson struct {
	Name string
	/* Burada `scores` field'ını private yaparak enkapsülasyon yaptık. Böylece `Lesson` `struct`'ını kullanacak kişilerin onu yanlış kullanma ihtimalini sınırlamış olduk. Bu uygulama kapsmaında derslerden alınacak puanların 0 - 100 arasında olmasını istiyoruz. Fakat bu değerleri kullanabileceğimiz en küçük tipimiz `uint8`. Bu tip ile 0'dan küçük bir değer almasını engelleyebiliyoruz fakat 100'den büyük bir değer almasını engelleyemiyoruz. Çünkü `uint8` 0-255 arasındaki değerleri alabilir. Yani eğer `scores` public olsaydı, bu package dışında tanımlanmış, örneğin, `Lesson` tipindeki bir `mathematic` değişkeni üzerinde şunu yapabilecektik:

	mathematic.Scores[0] = 150

	Fakat bu alanı private yaparak ve `scores` üzerinde değişiklik yapma fırsatını kullanıcılara yalnızca, aşağıda, `Lesson` üzerinde tanımladığımız `SetScoreOf(nthExam uint8, score Score) (bool, error)` metodu aracılığıyla verdiğimiz için `scrose` içerisinde istediğimiz aralık dışında değerler olmasını engellemiş olduk.
	*/
	scores [3]Score
}

func NewLesson(name string) Lesson {
	return Lesson{
		Name: name,
	}
}

func (l *Lesson) SetScoreOf(nthExam uint8, score Score) (bool, error) {
	if score > 100 {
		return false, fmt.Errorf("Not 100'den büyük olamaz")
	}

	if nthExam > 2 {
		return false, fmt.Errorf("Bir dersin üç sınvaı olabilir")
	}

	l.scores[nthExam] = score

	return true, nil
}

func (l *Lesson) GetScoreOf(nthExam uint8) (Score, error) {
	if nthExam > 2 {
		return 0, fmt.Errorf("Bir dersin üç sınvaı olabilir")
	}

	return l.scores[nthExam], nil
}

/*
Burada `scores` bilgisinin enkapsüllemesini desteklliyoruz. `Score` bilgisi `Lesson` üzerinde private olduğu için zaten bu paket dışından erişilemeyecek. Bu bilgiye ulaşılmak istendiğinde dizinin `slice`'ını ya da  `pointer`'ını dönmek yerine istenden değeri primitif olarak yani kopyasını dönüyoruz.
*/
func (l *Lesson) Notes() [3]Score {
	return l.scores
}

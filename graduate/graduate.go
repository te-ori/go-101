package graduate

import (
	"fmt"
	"scholl/lesson"
	"scholl/student"
)

type GraduateResult struct {
	StudentName string
	Final       lesson.Score
	IsGraduate  bool
}

/* 'Re-typeing' ile gittikçe karmaşıklaşan ve ilk bakışta çok fazla anlam ifade etmeyen tiplerin daha anlamlı hale gelmesini sağladık. Bu kodun okunabilirliğini de arttırıyor. Ayrıca, her zaman aynı pratiklikte olamsada, örneğin tipimizde bir değişiklik yapmak istediğimizde -mesela dönüş tipini değiştirmek istersek- kolaylık sağlayacaktır
 */
type Students []*student.Student
type GraduateCalcFunc func([]lesson.Score) (float32, bool)

/*
	Diyelim ki uygulamamaızda öğrencilerimizin bir dersten yıl sonu ortalamalarını ve o dersten başarılı olup olmadığını hesaplamak istiyoruz. Bu durumda geliştireceğimiz çözüm aşağı yukarı şu adımlardan oluşacaktır:

	* Bir öğrenci listesi al
	* Bu listedeki her bir öğrenci için
	* * Bu öğrencinin üzerinde istenen ders var mı bak
	* * * Yoksa bu öğrenciyi atla
	* * * Varsa dersten aldığı notlar üzerinden başarı notunu hesapla
	* * * Başarı notunun hedeften yüksek olup olmadığına bak
	* Öğrencilerin yıl sonu başarı sonuçlarını içeren listeyi dön.

	Peki başarı durumlarının farklı yöntemlerle hesaplanmasını istersek ne yapmamız gerekiyor? Diyelim ki iki yöntemimiz var birinde öğrencinin genel ortalamasına göre birinde de apırlıklı ortalamasına göre hesaplayacağız. Bu seferde şöyle bir yol izleyebiliriz:
	* Bir öğrenci listesi al
	* Bu listedeki her bir öğrenci için
	* * Bu öğrencinin üzerinde istenen ders var mı bak
	* * * Yoksa bu öğrenciyi atla
	* * * Varsa ve  birinci yöntemle hesapşamak sityorsa
	* * * * Birinci yönteme göre ağırlık hesapla
	* * * Varsa ve  ikinci yöntemle hesapşamak sityorsa
	* * * * Birinci yönteme göre ağırlık hesapla
	* * * Başarı notunun hedeften yüksek olup olmadığına bak
	* Öğrencilerin yıl sonu başarı sonuçlarını içeren listeyi dön.

	Peki yöntemlerimiz artarsa? Burada durup şunu düşünmeliyiz. Bu yapılan işlerin hangileri tercih ettiğimiz yönteme bağlı, hangileri değil? Buna karar verdikten sonra bütün yöntemler için sabit olan işleri ayırıp değişen ksımları da dışardab parametre olarak alırsak esnek bir yöntem bulabiliriz. Nasıl ki fonksiyonlarımızda yapılan işe göre değişen değerleri dışardan parametre olarak alıyorsak ve böylece genel geçer bir çözüm yöntemi hazırlıyorsak bu sefer de değişkenlik gösteren adımlarımızı dışarıdan almak istiyoruz. Parametre olarak fonksiyon alan ya da parametre olarak fonksiyon dönen fonksiyonlara 'Higer Order Function' deniyor.

	Burada hof'lardan şu şekilde faydalanbiliriz: öğrenci listesi üzerinde gezinmemiz, öğrenci üzerinde ilgli dersi aramamız ve ders varsa notlarını almamız, hesaplanmış başarı sonuçlarını sonuç listesine dönmemiz seçeceğimiz hesaplama yönteminden bağımsız. Hesaplama yöntemimiz için ise ders notlarına ihtiyacımız var. Yani hesaplama yöntemimizi ders notlarını alıp, üzerinde istediği hesapları yaptıktan sonra bir final notu ve başarı durumu bilgisi dönmek genelleyebiliriz. Bu genellememizi de kodumuzda `func([]lesson.Score) (float32, bool)` tipinde bir parametre tanımlayarak sağlıyoruz.
*/
// retyping olmasaydı bu fonksiyon tanımı böyle olacakt:
// func GetGraduateList(lessonName string, students []*student.Student, graduateCalc func([]lesson.Score) (uint8, bool)) []*GraduateResult {
func GetGraduateList(lessonName string, students Students, graduateCalc GraduateCalcFunc) []*GraduateResult {
	var result []*GraduateResult

	// Bu kısımlar hesap yönteminden bağımsız ve her hesap yöntemi için neredeyse aynı kalacak ksımlar. Öyleyse bunları her yöntem için terkar etmemize gerek yok.
	for _, student := range students {
		scores := student.ScoresOfLesson(lessonName)

		if scores == nil {
			continue
		}

		// `graduateCalc` bizim parametremiz. Artık fonksiyonu çağırdığımız yerde parametre olarak `[]lesson.Score` tipinde bir değer alan ve `(float32, bool) ` tipinte dönüş yapan herhangi bir fonksiyonu verebiliriz.
		finalNote, isGraduate := graduateCalc(scores)
		result = append(result, &GraduateResult{student.Name, lesson.Score(finalNote), isGraduate})
	}

	return result
}

func GraduateCalculationByStandartAverage(threshold uint8) GraduateCalcFunc {
	return func(scores []lesson.Score) (float32, bool) {
		fmt.Println(scores[0], scores[1], scores[2])
		avg := (float32(scores[0]) + float32(scores[2]) + float32(scores[0])) / 3

		return avg, uint8(avg) >= threshold
	}
}

func GraduateCalculationByWeightedAverage(lessonName string, students Students, threshold uint8) GraduateCalcFunc {
	return func(scores []lesson.Score) (float32, bool) {
		firstExam := float32(scores[0]) * 0.2
		secondExam := float32(scores[1]) * 0.2
		thirdExam := float32(scores[2]) * 0.4
		avg := firstExam + secondExam + thirdExam

		return avg, uint8(avg) >= threshold
	}
}

func GraduateCalculationByTotalAverage(lessonName string, students Students) GraduateCalcFunc {
	var total, count uint16
	for _, student := range students {
		scores := student.ScoresOfLesson(lessonName)

		if scores == nil {
			continue
		}
		count++
		total = total + uint16(scores[0]+scores[1]+scores[2])
	}

	threshold := uint8(total / (count * 3))
	return func(scores []lesson.Score) (float32, bool) {
		firstExam := float32(scores[0]) * 0.2
		secondExam := float32(scores[1]) * 0.2
		thirdExam := float32(scores[2]) * 0.4
		avg := firstExam + secondExam + thirdExam

		return avg, uint8(avg) >= threshold
	}
}

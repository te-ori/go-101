package student

import (
	"fmt"
	"scholl/lesson"
	"sort"
)

type Student struct {
	Name string
	/*
		Öğrencinin bir veya daha fazla dersi olabileceğini ve her öğrecinin farklı sayıda dersi seçebileceğini ve yine de bir öğrencinin alabileceği ders sayısının bir üst sınırı olduğunu -ve şimdilik 100 olduğunu- düşünelim. Bunu yapabileceğimiz en az üç çeşit veri yapımız var: `array`, `slice` ve `map`. Bunarın arasından neden `map`'i seçtik? Temelde iki sebebi var:
		* Birincisi bellek kullanımını optimize etmek. `array`'lerin kaç eleman tutabileceği tanımlandıkları anda belirtilmeli ve bu noktadan sonra `array`ler daha fazla veya daha az kapasiteye sahip olamaz. Eğer derslerimiz için `array` kullanmayı tercih etseydik, öğrencinin gerçekte kaç ders aldığından bağımsız olarak tümü için 10 derslik kapasite ayrılacaktı. Yani öğrenci bir ders alıyorsa, dokuz derslik alan bellekte israf edilecek. Eğer öğrencilerin büyük çoğunluğunun sekiz ve üzeri ders alacağını garanti edemiyorsak, öğrenci sayısı arttıkça bellek israfı da o oranda artacak.
		* `slice`'lar ile bellek israfının önüne geçebiliriz. Çünkü `array`'lerin aksine `slice`'ların kapasitesi ihtiyaca göre değişebiliyor. Peki o zaman niye `map`? Bir diğer önemli kriter de veriye erişim performansı. Bunu ise uygulamamızın kullanım amacı belirleyecek. Uygulamamızdan beklentimiz öğrencinin tüm derslerini hiç bir koşul olmadan listelemesi olabilir. Ya da 'fizik dersinden aldığı notları', 'matematik dersinin ortalamasını', 'edebiyat dersinden barajı geçip geçmediğin' de bekleyebiliriz uygulamamızdan. İkinci durumda veriye ulaşım konusu oldukça önem kazandı. Peki biz bu verilere nasıl ulaşacağız? Örneğin '0. dersin notlarını görmek' anlamlı bir iş olur mu? Bu haliyle oldukça amaçsız bir iş olacak çünkü '0. ders' herhangi bir ders olabilir. Ama 'öğrencinin fizik dersi notlarını' getirmek kulağa daha anlamlı geliyor. Yani istediğimiz verinin kapsını açacak anahtar o dersin adı. Eğer dersleri `slice` içerisinde tutsaydık istediğimiz dersin bilgilerini almak için aşağı yukarı şöyle bir yöntem kullanmamız gerekecekti:

		func (s *Student) GetLessonByName(name string) *lesson.Lesson {
			for index, lesson := range s.lesson {
				if lesson.Name == name {
					return lesson
				}
			}

			return nil
		}

		Bu da şu demek: ders sayısı arttıkça o dersi adına göre bulma işi daha fazla kaynak gerektirecekti (O(n)). Ama `map`'ler içerisindeki eleman sayısından bağımsız olarka her türlü sorgulama için sabit süre garantisi vermekte (O(1)). Buradaki O(1)'in gerçek değeri birkaç eleman için `slice`'ı tek tek gezmekten daha maliyetli olabilir. Fakat eleman sayısı arttıkça `map`'in daha verimli olacağı bir durum ortaya çıkacaktır.
	*/
	lessons map[string]*lesson.Lesson
}

func NewStudent(name string) Student {
	return Student{
		Name:    name,
		lessons: make(map[string]*lesson.Lesson),
	}
}

/*
`Student` üzerindeki `lesson`'ları private yaparak ve aşağıdaki metodlarımız sayesinde burada da enkapsülleme yaptık. Aşağıdaki işlemlerin tamamını `lesson`'ı public yaparak da sağlayabilirdik. Bu belki kullanıldıkları yerlerde daha esneklik bile sağlayabilir. Buradaki enkapsüllemenin getirisi verilerimizin tutarlılığını ve doğruluğunu sağlamanın da ötesinde hatta kod tekrarından da fazlası. Eğer `lesson`'ı public yapsaydık verilerimiz üzerinde işlem yapmak isteyen kişi aynı zamanda `map` üzerinde de işlem yapmak zorunda kalacaktı. Tek arzusu `student` üzerinde işlem yapmak isteyen birisi için gereksiz bir maaliyet. Buradaki enkapsülasyon sayesinde hem kodu kullanacak kişinin gereksiz mental yüklerle uğraşmamasını sağlıyoruz hem de ilerde `lessons`'ı saklama biçimimizde olacak bir değişiklikten de etkilenmemesini sağlıyoruz. Mesela ilerde `map`'ten daha faydalı bir veri tipi bulduk ve onu kullanacağız. Eğer enkapsüllemeyi yeterince yapmasaydık ders bilgilerinin herhangi bir şekilde kullanıldığı tüm yerlerde değişiklik yapmak zorunda kalackatık. Fakat bu haliyle sadece bu `package` içerisinde değişiklik yapmamız yeterli olacak. Paket dışında `lesson`'lara ulaşan kişiler yine aynı metodu aynı parametre ile çağırmaya devam edecek. Bu durumun YAGNI (you aint gonna need it, buna asla ihtiyacın olmayacak) ile çeliştiği düşünülebilir. Burada belkide hiç kullanılmayacak bir ihtiyaca çözüm üretmedik. Aksine zaten var olan 'veriye erişim ihtiyacımızın' gelecekte sorunsuz bir şekilde değiştirilebilmesini sağladık.
*/
func (s *Student) InsertLesson(lesson *lesson.Lesson) (bool, error) {
	_, isExist := s.lessons[lesson.Name]
	if isExist {
		return false, fmt.Errorf("Bir öğrenci aynı dersi birden fazla defa alamaz")
	}

	s.lessons[lesson.Name] = lesson

	return true, nil
}

func (s *Student) ListOfLessons() []string {
	names := make([]string, len(s.lessons))
	j := 0
	for key := range s.lessons {
		names[j] = key
		j++
	}

	sort.Strings(names)
	return names
}

func (s *Student) ScoresOfLesson(name string) []lesson.Score {
	lesson, isExist := s.lessons[name]

	if !isExist {
		return nil
	}

	notes := lesson.Notes()
	return notes[:]
}

# Recursion



### Recursion nedir, ne işe yarar ve neden kullanırız?

Recursion basitçe bir fonksiyonun kendi kendini çağırmasıdır. Çoğu zaman performans olarak loop kullanmaktan daha verimsiz olsa bile eğer öncelikli amacımız kodu daha anlaşılır ve temiz yazmak ise recursion çok daha anlaşılır kod yazmamızı sağlar. Hatta bununla ilgili şöyle bir sözü de eklemek istiyorum.

> “Loops may achieve a performance gain for your program. Recursion may achieve a performance gain for your programmer. Choose which is more important in your situation!
>
>  *Leigh Caldwell* 



### Base Case ve Recursive Case

Recursive fonksiyonlarımız iki kısımdan kısımdan oluşur. İlk kısım Base Case kısmında fonksiyonun nerede duracağını belirtiriz. Recursive Case kısmında ise fonksiyonumuzun hangi pattern ile tekrardan çağıralacağını belirtiriz. Bununla ilgili basit örnekleri görmek için kaynak kodlarına bakabilirsiniz. Eğer recursive bir fonksiyonda Base Case'i kullanmazsak fonksiyonumuzun sınırsız döngüye girmiş gibi davranır ve bellek boyutumuza göre belirli bir yerde program hata verir.



### Stack

Stack en temel veri yapılarından bir diğeridir. Stack'i üst üst birbirine iğnelediğimiz notlar olarak düşünebilirsiniz. Notları sürekli en üstten ekleriz ve okuyacağımız zaman da yine en üstten okur ve atarız. Stack veri yapısı aynı bu mantıkta çalışır. Veriler üst üste dizilir ve bir veri stackten çıkartılacağı zaman önce en üstteki yani en son eklenen eleman çıkartılır. Buna Bilgisayar Bilimleri'nde **Last in First Out (LIFO)** adı verilir. 

### Call Stack

Bilgisayarlarımız da stack veri yapılarını aktif olarak kullanır. Bu stack **call stack** olarak adlandırılır. Çağırdığımız tüm fonksiyonlar bu stack üzerine eklenir. Şimdi basit bir örnekle açıklamaya çalışacağım.



```
def greet(name):
	greet2(name)
	bye(name)
	return True

greet("Furkan")
```



Yukarıda bir fonksiyonun içinde çağırılmış 2 fonksiyon görüyoruz. En dışarıdaki greet fonksiyonunu çağırdığımız zaman greet fonksiyonu bilgisayarımızdaki call stack üzerinde oluşturulur ve içerisindeki name değişkeni de memoryde tutulur. Ardından greet2 fonksiyonu çağırılır ve call stackte greet fonksiyonun üzerine eklenir ve name değişkenine onun da erişimi vardır. Ardından greet2 fonksiyonu çalışır ve return eder. Return ettiği için call stackten çıkartılır. Son olarak bye fonksiyonu çağırılır ve call stackte en tepeye eklenir. Bu fonksiyon da return ettikten sonra aynı şekilde stackten çıkartılır ve en sona greet fonksiyonu kalır o da True değerini return edip stackten çıkartılır. Örnekten de anlayabileceğiniz gibi call stacke fonksiyonlar üst üste eklendi ve son eklenen hep ilk çıkan oldu. 



İşte recursive fonksiyonlarımızda base case yazmazsak her fonksiyon çağrısında o fonksiyon call stack üzerine ekleneceği için bellek boyutumuza bağlı olarak call stack bir yerden sonra şişecek ve daha yeni eleman ekleyemediği için programımız hata verecektir.

Base case yazdığımız durumda ise fonksiyonumuz call stacke base case'e erişene kadar eklenecek sonra en son base case eklendiğinde yukarıda bahsettiğim Last in First Out mantığıyla fonksiyonları return edip call stackten çıkartmaya başlayacaktır.

Örneğin recursive bir faktöriyel fonksiyonu yazdığımızı düşünelim. (Örnek kodlar mevcuttur.) fac(3) çağırdığımızda call stack şu şekilde olacaktır.

| Call Stack |
| ---------- |
| fact(1)    |
| fact(2)    |
| fact(3)    |

 Base case burada n değerinin 1'e eşit olması olduğu için base case'e ulaştığımızda fonksiyonumuz sırasıyla aşağıya doğru değerleri return etmeye başlayacaktır.



Recursion konusuyla ilgili edindiğim bilgiler ve aldığım notlar bu şekildedir. Anlaması ilk başlarda biraz zor olsa da Base Case ve Recursive Case mantığını kafanıza oturtursanız anlamanızı kolaylaştıracağına inanıyorum. Örnek bazı recursive fonksiyonları örnek kodlar içerisinde bulabilirsiniz.





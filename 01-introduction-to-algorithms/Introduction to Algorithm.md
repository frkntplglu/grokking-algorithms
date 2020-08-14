# Introduction to Algorithm

Algoritmaları basitçe herhangi bir işlemi yaptırmak için sıralı şekilde verdiğimiz komutlar listesi olarak düşünebiliriz. Farklı işlemler için farklı farklı algoritmalar kullanabiliriz. Bu repo üzerinde en temel olarak görülen searching, sorting, selection gibi algoritmalara ait detayları ve kod örneklerini paylaşmaya çalışacağım.



İlk olarak bir arama örneğiyle başlayacağız. Diyelim ki sıralı şekilde listelenmiş 100 elemanlı bir sayı listemiz olsun. Bu liste içerisinden bir sayıyı aramak istediğimizde nasıl bir yöntem izleriz? Listenin başından başlayıp tek tek sayıları kontrol ederek bu işlemi yapabiliriz. Bu işe yarayan bir yöntem olsa da her zaman en doğru yöntem değildir. En kötü senaryoyu düşündüğümüz zaman aradığımız sayı listenin son elemanı olabilir. Böyle bir durumda tüm listeyi dolaşmak zorunda kalırız. Listemizin eleman sayısı arttıkça işlem süresi de artacaktır. Peki bu işlemi daha efektif şekilde nasıl yapabiliriz?



### Binary Search

Binary Search en bilindik arama algoritmalarından biridir. Sıralanmış listeler üzerinde çalışır. Binary Search ile arama yaptığımızda bize aradığımız elemanı bulabilirse konumunu bulamazsa ise NULL değer döner. 



Binary Search arama yapmaya listenin tam ortasından başlar. Ortadaki eleman ile aramak istediğimiz elemanı karşılaştırır ve aradığımız elemanın, ortadaki elemandan büyük küçük olma durumuna göre listenin diğer yarısını bir daha asla kontrol etmez. Her seferinde listenin ortasına giderek bu işlemi, aradığımız elemanı bulana kadar devam ettirir. 



Örneğin 100 elemanlı bir listede arama yaptığımız zaman ilk başta bahsettiğimiz arama yöntemi (simple search) en kötü senaryoda 100 işlem yaparken Binary Search ise en fazla 7 işlemde aramayı gerçekleştirir.



Buradan yola çıkarak şunu söyleyebiliriz.

**n** elemanlı bir listede arama yapmak için simple search **n** işlem yaparken, binary search **log n** işlem yaparak sonuca ulaşır. 



100 gibi küçük elemanlı listelerde bu çok büyük sorunlar yaratmasa da gelin farklı bir senaryo düşünelim. 



1.000.000.000 elemanın olduğu bir veriyi düşünelim ve her işlemin 1 ms sürdüğünü varsayalım.

| Simple Search                             | Binary Search            |
| ----------------------------------------- | ------------------------ |
| n = 1.000.000.000<br />1x1.000.000.000 ms | log n = 32 <br />1x32 ms |

Yani bu demek oluyor ki yapacağımız işlem simple search ile 11 gün sürerken, binary search ile aynı işlem sadece 32ms içerisinde yapılabilir. 

Bu basit örnekle bile doğru algoritma kullanmanın ne kadar faydalı olduğunu görebiliyoruz.

Binary Search için yazılmış Python kodunu repo içerisinde bulabilirsiniz.



### Big O Notasyonu

Big O notasyonu bize algoritmanın ne kadar hızlı çalıştığını gösteren özel bir gösterim türüdür.  Yaptığımız işlemin ne kadar sürdüğü kullandığımız verinin büyüklüğünü de bağlı olduğu için kesin bir yargıya varamayız. Bu yüzden big o notasyonu bize işlemin süresini değil, bu işlemi gerçekleştirmek kaç operasyon yapılması gerektiğini söyler. 



Big O notasyonu adından da anlaşılabileceği gibi büyük "O" harfi ve işlem sayısının yazılmasından oluşur.



**Simple Search :** O(n)

**Binary Search :** O(log n)



Bir değer gözardı etmememiz gereken bilgi ise, big o notasyonu her zaman worst-case'e göre hesap yapar.  Yani O(n) ve O(log n) algoritmalarımızın en kötü senaryodaki işlem sayısıdır. 



### Yaygın Algoritmaların Big O notasyonları

| Binary Search             | O(log n)         |
| ------------------------- | ---------------- |
| **Simple Search**         | O(n)             |
| **Quick Sort**            | O(n * log n)     |
| **Selection Sort**        | O(n<sup>2</sup>) |
| **Traveling Salesperson** | O(n!)            |



### Traveling Salesperson

Bu problem yaygın olarak bilinen ve çözülememiş bir Bilgisayar Bilimleri problemidir. Probleme göre bir satış elemanımız var ve gitmesi gereken 5 farklı şehir. Bu satış elemanı gidebileceği en kısa nasıl bulur? 5 farklı şehir olduğu için 5! kadar permütasyon mevcuttur ve dolasıyıla toplamda 720 işlem yapması gerekir. Şehir sayısı arttığında permütasyon sayısı da devasa şekilde artacaktır. Örneğin 100+ şehrimiz olduğunu düşünürsek bu pek de hesaplanabilir olmayacaktır. 



İlk chapterdan edindiğim bilgi ve aldığım notlar bu şekildedir. Hatalı gördüğünüz herhangi bir kısım için PR oluşturmak lütfen çekinmeyin. Okuduğunuz için teşekkür ederim.
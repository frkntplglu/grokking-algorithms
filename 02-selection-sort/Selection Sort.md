# Selection Sort



### Bellek Nasıl Çalışır?

Bilgisayarımızın belleğini çekmeceleri olan bir dolaba benzetebiliriz. Her çekmece bellekteki bir alanı temsil eder ve kendine ait adresi bulunmaktadır ve her alan tek bir değer tutar. Eğer birden fazla değer depolamak istiyorsak **Array** ya da **Linked List** kullanmamız gerekiyor.



### Array

Arrayler elemanların ardışık sırayla tutulduğu bir veri türüdür. Arrayin tüm elemanları bellekte sıralı şekilde depolanır. Örneğin 3 elemanlı bir arrayimiz olduğunu var sayalım, bu 3 eleman bellekte ard arda tutulur ve 3. elemandan sonraki bellek alanı başka bir program tarafından kullanılıyor olabilir.  Bu sebepten ötürü arraye yeni bir eleman eklemek istediğimiz zaman eğer bellek müsait değilse tüm array yeni eleman sayısına göre bellekte uygun bir yere taşınmak zorundadır. Her eleman eklemek istediğimizde bu işlemi yapacak olmamız pek kulağa hoş gelen bir yöntem değil maalesef. Bu sorunu çözebilmek için bir diğer veri türü olan **Linked List**'i kullanabiliriz. Ona birazdan değineceğim. Bunun dışında arrayin her bir elemanının kendine ait bir indexi bulunur ve bu indexler 0'dan başlar. Örneğin 4 elemanlı bir arrayin ilk elemanının indexi 0 iken son elemanınki 3 olur. 



### Linked List

Linked List'te tüm elemanlar bellekte rastgele depolanır ve her bir eleman bir sonraki elemanın bellek adresini tutar. Bu adres sayesinde bir sonraki elemana erişim sağlanır. Bu rastgele depolanma özelliği sayesinde linked listlere eleman eklerken herhangi bir taşıma işlemi yapmamıza gerek yoktur. Linked listlerin de dezavantajı okuma işlemleridir. Linked listteki elemanların adresini bilemeyeceğimiz için bir elemanı okuyabilmek için tüm listeyi dolaşmamız gerekir.



### Array ve Linked List Karşılaştırması

Eleman ekleme ve silme işlemlerinde linked list performans olarak çok daha üstündür. Çünkü arraylerde bu işlemleri yapmak için hem taşıma yapmak hem de sıralı depolandığı için diğer elemanları kaydırmak gerekmektedir. Okuma işleminde ise arraylerin random access (rastgele erişim) özelliği olduğu için istediğimiz elemana çok daha rahat ulaşabiliriz. Okuma işleminde linked list sequential access (sıralı erişim) özelliğine sahip olduğu için yukarıda da bahsettiğim gibi tüm listeyi dolaşmamız gerekir.



#### Çalışma Süreleri

|            | Array | Linked List |
| ---------- | ----- | ----------- |
| **Read**   | O(1)  | O(n)        |
| **Insert** | O(n)  | O(1)        |
| **Delete** | O(n)  | O(1)        |



### Selection Sort

Selection Sort aslında çok hızlı çalışmasa da **(O(n<sup>2</sup>))** yaygın kullanılan bir sıralama algoritmasıdır. Bildiğim kadarıyla 2 farklı şekilde implement edilebiliyor. Birinde en büyük veya en küçük elemanı bulup mevcut array içerisinde swap yaparken, diğer yöntemde de sırayla yine en büyük veya en küçük eleman bulunup bu eleman farklı bir arraye eklenerek yeni ve sıralanmış bir array oluşturulur. Kitap 2. yöntemi kullandığı ve ben de bunu daha anlaşılır bulduğum için örnek kodu da bu şekilde yazdım. Ayrıca selection sort ve diğer bir çok algoritmanın görsel olarak nasıl çalıştığını görmek isterseniz  aşağıdaki siteye göz atabilirsiniz. Okuduğunuz için teşekkür ederim bir sonraki bölümde görüşmek üzere.

[https://visualgo.net/en]: 





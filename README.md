# TARJANS-13521159

## Disusun Oleh
 - Sulthan Dzaky Alfaro
 - 13521159

## Cara Penggunaan Program

 - Pertama-tama siapkan 2 terminal (untuk backend dan frontend)
 - Untuk terminal pertama, kita gunakan untuk backend. Pertama pastikan path terminal sudah berada pada .../TARJANS-13521159/BE. Lalu tulis ```go run Tarjans.go``` pada terminal tersebut. Apabila sudah terlihat ```Listening and serving HTTP on localhost:8080```, maka backend sudah aktif dan siap digunakan.
 - Untuk terminal kedua, kita gunakan untuk frontend. Pertama pastikan path terminal sudah berada pada .../TARJANS-13521159/FE/my-app. Lalu tulis ```npm start``` pada terminal tersebut. Tunggu beberapa saat sampai terminal sudah tertulis ```webpack compiled successfully```(tandanya web sudah siap digunakan). Jika web tidak muncul secara otomatis pada web browser anda, anda bisa menulis ``` http://localhost:3000``` pada browser kalian setelah web sudah siap digunakan. Web sudah siap digunakan dengan cara mengisi simpul.
 - Untuk mengisi simpul anda dapat memilih mengisi simpul secara manual atau menginputkan file txt. Isi simpul pada text box/file txt dengan format seperti contoh berikut.
    ```
    A B
    B C
    C A
    B D
    D E
    E F
    F E
    ```
    Maksud dari input diatas adalah A -> B, B ->
    C , dan seterusnya(pengisian input berhati hati karena salah spasi saja maka akan diperingati oleh web).
 - Setelah itu, pencet start untuk memulai program dan hasil akan ditampilkan dibawahnya.
## Algoritma Tarjan's
Algoritma Tarjan adalah algoritma dalam teori graf yang digunakan untuk menemukan komponen-komponen terhubung dalam sebuah graf berarah (directed graph) atau graf tak berarah (undirected graph). Tujuan utama dari algoritma Tarjan adalah mengidentifikasi dan memisahkan kelompok simpul-simpul yang saling terhubung dalam graf. Salah satu penggunaan algoritma Tarjans ini adalah pencarian SCC (Strongly Connected Component) dan bridge pada sebuah graph.
### Kompleksitas Algoritma Tarjan's
Algoritma Tarjan's sendiri bekerja dengan menggunakan teknik rekursif yang biasa digunakan dalam graph, yaitu DFS (Depth First Search). DFS sendiri merupakan algoritma pencarian yang mendalam sebuah graph/pohon, mengikuti salah satu cabang graph/pohon hingga menemukan solusi. Karena algoritma Tarjan's ini menggunakan DFS, sehingga kompleksitas dari algoritma ini adalah O(E + V) dimana E adalah edges dan V adalah vertices(simpul).
### Modifikasi Algoritma Tarjan's pada Pencarian Bridges
Untuk pencarian bridges pada sebuah graph dengan menggunakan algoritma Tarjan's sebenarnya hampir sama dengan penggunaan algoritma Tarjan's dalam pencarian SCC. Kita tetap menggunakan variabel low-link untuk mengetahui apakah ada bridge atau tidak. Namun ada perbedaan sedikit pada pencarian bridges. Perbedaan tersebut adalah yang pertama tentunya tidak memakai stack, tetapi kita memakai acuan parent untuk membatasi rekursif atau apabila node yang dicek tidak memiliki tetangga lagi. Apabila node yang kita cek bukan parent dari tetangga dari node yang kita cek (back edge), maka kita update nilai low-link dari node yang kita cek dengan nilai minimum dari low-link dari node yang dicek dengan id tetangga. Untuk menentukan apakah edge merupakan bridge dengan cara melihat apakah nilai low-link tetangga dari node yang kita cek bernilai lebih dari id dari node yang kita cek atau tidak. Jika iya, maka edge tersebut adalah bridge. Jika tidak, maka edge tersebut bukan bridge.
### Jenis Edges Yang Ada Pada Graph
 - Back Edge

   Back edge adalah edge berarah yang menghubungkan suatu simpul dengan nenek moyangnya dalam lintasan rekursif DFS (Depth-First Search) pada graf. Dalam graf berarah, jika terdapat tepi dari simpul u ke simpul v, dan simpul v adalah nenek moyang dari simpul u dalam lintasan DFS, maka tepi tersebut disebut sebagai back edge. Back edge bisa diibaratkan sebagai sirkuit tertutup dalam graf. Back edge tidak ada pada graf tak berarah. Dalam graf tak berarah, konsep back edge tidak berlaku, karena setiap tepi dapat dianggap sebagai back edge.
 - Cross Edge

   Cross edge adalah tepi berarah yang menghubungkan simpul-simpul yang tidak memiliki hubungan hierarki (bukan parent-child) dalam lintasan rekursif selama proses DFS. Dalam graf berarah, jika terdapat tepi dari simpul u ke simpul v, dan simpul v tidak berada dalam jalur ancestral simpul u dalam DFS, maka tepi tersebut disebut sebagai cross edge. Cross edge dapat menghubungkan simpul-simpul pada level yang sama atau berbeda dalam DFS tree.

## Referensi 
 - Algoritma Tarjans SCC

   https://youtu.be/wUgWX0nc4NY

 - Algoritma Tarjans Bridge

   https://www.geeksforgeeks.org/bridge-in-a-graph/

   https://youtu.be/Rhxs4k6DyMM

 - Framework React

   https://reactjs.org/

 - Library
   - MUI React

      Library ini digunakan untuk mempermudah dalam pembuatan UI pada frontend.
   
     https://mui.com/getting-started/usage/
   - Graphviz
   
      Library ini digunakan untuk mempermudah dalam pembuatan graph pada frontend.

      https://www.npmjs.com/package/react-graph-vis
   
   - Axios

      Library ini digunakan untuk menghubungkan API dengan frontend.

      https://www.npmjs.com/package/axios

   - Go gin-gonic
   
      Library ini digunakan untuk mempermudah dalam pembuatan API pada backend.

      https://gin-gonic.com/



     
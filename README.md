# MST-Based-Clustering-API

## Deskripsi Program

MST-Based Clustering merupakan salah satu algoritma unsupervised pada machine learning yang banyak digunakan. Jika diberikan sebuah dataset dengan n buah titik random, algoritma ini akan membangun sebuah Minimum Spanning Tree (MST), kemudian melakukan pengelompokkan data dengan cara memotong sisi MST mulai dari sisi dengan bobot terbesar. Jumlah sisi yang dipotong menentukan jumlah cluster yang akan dibuat, untuk setiap n cluster akan ada n-1 pemotongan sisi mulai dari sisi dengan bobot terbesar.

## Teknologi dan Framework

Server pada website MST-Based CLustering dibuat menggunakan:
- Golang
- Docker
- CI/CD (Unit test Algoritma Kruskal)

## Penjelasan Algoritma Kruskal

Algoritma Kruskal adalah algoritma yang digunakan untuk membentuk Pohon Merentang Minimum (<i>Minimum Spanning Tree</i>) dengan menghubungkan semua titik dengan bobot terendah. Algoritma Kruskal mirip seperti Algoritma A* yang memiliki Algoritma Greedy dan Teknik Heuristik di dalamnya namun Algoritma Kruskal hanya bertujuan untuk menghubungkan semua titik, berbeda halnya dengan Algoritma A* yang bertujuan untuk mencari jalur yang tersambung dari awal sampai akhir. Algoritma Greedy pada Algoritma Kruskal yaitu pengambilan kedua titik dengan bobot sisi terendah dan Teknik Heuristik yang diterapkan adalah tidak ada sisi dari pohon sehingga membentuk lintasan sirkuler di dalamnya. Apabila sisi dengan bobot terendah yang diambil membentuk lintasan sirkuler, maka Algoritma Kruskal akan mencari sisi lain dengan bobot terendah sehingga <i>Minimum Spanning Tree</i> dapat terbentuk.

## Referensi Belajar

- https://bibliografi.my.id/contoh-soal-algoritma-prim-dan-kruskal/
- https://qiita.com/osk_kamui/items/1539ade3c23f58b89f80
- https://dev.to/gopher/build-ci-cd-pipelines-in-go-with-github-actions-and-dockers-1ko7

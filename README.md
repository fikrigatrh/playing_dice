# playing_dice

II. Soal Praktek:
Buatlah sebuah script permainan dadu yang menerima input N jumlah pemain dan M jumlah
dadu, dengan peraturan sebagai berikut:
1. Pada awal permainan, setiap pemain mendapatkan dadu sejumlah M unit.
2. Semua pemain akan melemparkan dadu mereka masing-masing secara bersamaan
3. Setiap pemain akan mengecek hasil dadu lemparan mereka dan melakukan evaluasi
seperti berikut:
a. Dadu angka 6 akan dikeluarkan dari permainan dan ditambahkan sebagai poin
bagi pemain tersebut.
b. Dadu angka 1 akan diberikan kepada pemain yang duduk disampingnya.
Contoh, pemain pertama akan memberikan dadu angka 1 nya ke pemain kedua.
c. Dadu angka 2,3,4 dan 5 akan tetap dimainkan oleh pemain.
4. Setelah evaluasi, pemain yang masih memiliki dadu akan mengulangi step yang ke-2
sampai tinggal 1 pemain yang tersisa.
a. Untuk pemain yang tidak memiliki dadu lagi dianggap telah selesai bermain.
5. Pemain yang memiliki poin terbanyak lah yang menang.

Contoh:
Pemain = 3, Dadu = 4
==================
Giliran 1 lempar dadu:
Pemain #1 (0): 3,6,1,3
Pemain #2 (0): 2,4,5,5
Pemain #3 (0): 1,2,5,6
Setelah evaluasi:
Pemain #1 (1): 3,3,1
Pemain #2 (0): 2,4,5,5,1
Pemain #3 (1): 2,5
==================
Giliran 2 lempar dadu:
Pemain #1 (1): 1,2,6
Pemain #2 (0): 4,3,1,3,3
Pemain #3 (1): 1,6
Setelah evaluasi:
Pemain #1 (2): 2,1
Pemain #2 (0): 4,3,3,3,1
Pemain #3 (2): 1
==================
Giliran 3 lempar dadu:
Pemain #1 (2): 6,1
Pemain #2 (0): 2,5,6,4,6
Pemain #3 (2): 1
Setelah evaluasi:
Pemain #1 (3): 1
Pemain #2 (2): 2,5,4,1
Pemain #3 (2): _ (Berhenti bermain karena tidak memiliki dadu)
==================
Giliran 4 lempar dadu:
Pemain #1 (3): 1
Pemain #2 (2): 3,4,5,5
Pemain #3 (2): _ (Berhenti bermain karena tidak memiliki dadu)
Setelah evaluasi:
Pemain #1 (3): _ (Berhenti bermain karena tidak memiliki dadu)
Pemain #2 (2): 3,4,5,5
Pemain #3 (2): _ (Berhenti bermain karena tidak memiliki dadu)
==================
Game berakhir karena hanya pemain #2 yang memiliki dadu.
Game dimenangkan oleh pemain #1 karena memiliki poin lebih banyak dari pemain lainnya

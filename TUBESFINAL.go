package main

import (
	"fmt"
	"time"
)

const NMAX = 100

type rating struct {
	jumlah float64
	total  float64
	hasil  float64
}

type film struct {
	ID        int
	Judul     string
	Genre     string
	Tahun     int
	Sutradara string
	Sinopsis  string
	Rating    rating
}

var movie [NMAX]film

var terurut bool

var count int = 0

func main() {
	var m string
	fmt.Println("===================================================")
	fmt.Printf("**\tSelamat datang di Perpustakaan Film\t**\n")
	fmt.Println("===================================================")
	time.Sleep(1 * time.Second)
	for {
		fmt.Println("===================================================")
		fmt.Printf("**\t\t\tLogin\t\t\t**\n")
		fmt.Println("===================================================")
		fmt.Println("1. User")
		fmt.Println("2. Admin")
		fmt.Println("3. Exit")
		fmt.Println("===================================================")
		fmt.Print("Pilihanmu: ")
		fmt.Scan(&m)
		if m == "1" {
			user()

		} else if m == "2" {
			menuAdmin()

		} else if m == "3" {
			return
		} else {

			fmt.Println("Invalid")
			time.Sleep(1 * time.Second)
		}
	}
}

func menuAdmin() {
	// F.S.: Menampilkan menu Admin dengan opsi untuk menambah, mengedit film, atau logout.
	// I.S.: Tidak ada. Fungsi ini dipanggil saat aplikasi dimulai.
	for {
		var pilihan int
		fmt.Println("===================================================")
		fmt.Printf("**\t\t\tAdmin\t\t\t**\n")
		fmt.Println("===================================================")
		fmt.Println("1. Tambah Film")
		fmt.Println("2. Edit Film")
		fmt.Println("3. Logout")
		fmt.Println("===================================================")
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			tambahFilm()
		case 2:
			editFilm()
		case 3:
			fmt.Println("Logout...")
			time.Sleep(1 * time.Second)
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}
func tambahFilm() {
	// F.S.: Menambahkan film baru ke dalam daftar film.
	// I.S.: Daftar film awal telah terinisialisasi, dan pengguna memilih untuk menambah film baru.
	if count >= NMAX {
		fmt.Println("Kapasitas maksimum film telah tercapai.")
		return
	}

	var id, jumlah, i int
	var genre string
	fmt.Print("Masukan jumlah film yang ingin dimasukan : ")
	fmt.Scan(&jumlah)
	for i = 0; i < jumlah; i++ {

		fmt.Print("Masukkan ID Film: ")
		fmt.Scan(&id)

		// Cek apakah ID sudah digunakan sebelumnya
		if cariFilmByID(id) != -1 {
			fmt.Println("ID film sudah digunakan sebelumnya.")
			return
		} else {
			movie[count].ID = id
		}

		fmt.Print("Masukkan Judul Film: ")
		fmt.Scan(&movie[count].Judul)

		// Memilih genre film
		fmt.Println("Pilih Genre Film:")
		fmt.Println("1. Drama")
		fmt.Println("2. Action")
		fmt.Println("3. Komedi")
		fmt.Println("4. Horor")
		fmt.Println("5. Romance")
		fmt.Print("Pilihan Anda: ")
		fmt.Scan(&genre)

		// Mengonversi genre dari angka ke string menggunakan if-else
		if genre == "1" {
			movie[count].Genre = "Drama"
		} else if genre == "2" {
			movie[count].Genre = "Action"
		} else if genre == "3" {
			movie[count].Genre = "Komedi"
		} else if genre == "4" {
			movie[count].Genre = "Horor"
		} else if genre == "5" {
			movie[count].Genre = "Romance"
		} else {
			fmt.Println("Genre tidak valid.")
			return
		}

		fmt.Print("Masukkan Sutradara Film: ")
		fmt.Scan(&movie[count].Sutradara)
		fmt.Print("Masukkan Tahun Rilis Film: ")
		fmt.Scan(&movie[count].Tahun)
		fmt.Print("Masukkan Sinopsis Film: ")
		fmt.Scan(&movie[count].Sinopsis)
		movie[count].Rating.hasil = 0.0

		// Cetak data film yang baru ditambahkan
		fmt.Println("Data film yang baru ditambahkan:")
		fmt.Println("ID: ", id)
		fmt.Println("Judul: ", movie[count].Judul)
		fmt.Println("Genre: ", movie[count].Genre)
		fmt.Println("Sutradara: ", movie[count].Sutradara)
		fmt.Println("Tahun: ", movie[count].Tahun)
		fmt.Println("Sinopsis: ", movie[count].Sinopsis)

		fmt.Println("Film berhasil ditambahkan.")
		terurut = false
		count++

	}
}

func editFilm() {
	// F.S.: Memberikan opsi kepada pengguna untuk menghapus atau mengubah film yang ada.
	// I.S.: Daftar film telah terinisialisasi, dan pengguna memilih untuk mengedit film.
	for {
		var pilihan int
		fmt.Println("Menu Edit Film:")
		fmt.Println("1. Hapus Film")
		fmt.Println("2. Ubah Film")
		fmt.Println("3. Kembali ke Menu Utama")
		fmt.Print("Pilih menu: ")
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			hapusFilm()
		case 2:
			ubahFilm()
		case 3:
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}
func hapusFilm() {
	// I.S: ID film yang akan dihapus belum dimasukkan.
	// F.S: Jika ID film ditemukan, film dihapus dari array films dan count berkurang 1. Jika tidak ditemukan, tampilkan pesan kesalahan.
	var id int
	fmt.Print("Masukkan ID Film yang akan dihapus: ")
	fmt.Scan(&id)

	index := cariFilmByID(id)
	if index == -1 {
		fmt.Println("Film tidak ditemukan.")
		return
	}

	// Geser elemen-elemen setelah film yang dihapus
	for i := index; i < count-1; i++ {
		movie[i] = movie[i+1]
	}
	count--
	fmt.Println("Film berhasil dihapus.")
}

// Fungsi untuk mengubah film
func ubahFilm() {
	// I.S: ID film yang akan diubah belum dimasukkan.
	// F.S: Jika ID film ditemukan, data film di array films diubah sesuai input pengguna. Jika tidak ditemukan, tampilkan pesan kesalahan.

	var id int
	var genre int

	fmt.Print("Masukkan ID Film yang akan diubah: ")
	fmt.Scan(&id)

	index := cariFilmByID(id)
	if index == -1 {
		fmt.Println("Film tidak ditemukan.")
		return
	}

	fmt.Printf("Film saat ini: %+v\n", movie[index])

	fmt.Print("Masukkan Judul Film baru: ")
	fmt.Scan(&movie[index].Judul)

	// Memilih genre film
	fmt.Println("Pilih Genre Film:")
	fmt.Println("1. Drama")
	fmt.Println("2. Action")
	fmt.Println("3. Komedi")
	fmt.Println("4. Horor")
	fmt.Println("5. Romance")
	fmt.Print("Pilihan Anda: ")
	fmt.Scan(&genre)

	// Mengonversi genre dari angka ke string menggunakan if-else
	if genre == 1 {
		movie[index].Genre = "drama"
	} else if genre == 2 {
		movie[index].Genre = "Action"
	} else if genre == 3 {
		movie[index].Genre = "Komedi"
	} else if genre == 4 {
		movie[index].Genre = "Horor"
	} else if genre == 5 {
		movie[index].Genre = "Romance"
	} else {
		fmt.Println("Genre tidak valid.")
		return
	}

	fmt.Print("Masukkan Sutradara Film: ")
	fmt.Scan(&movie[index].Sutradara)
	fmt.Print("Masukkan Tahun Rilis Film: ")
	fmt.Scan(&movie[index].Tahun)
	fmt.Print("Masukkan Sinopsis Film: ")
	fmt.Scan(&movie[index].Sinopsis)

	// Cetak data film yang baru ditambahkan
	fmt.Println("Data film yang baru ditambahkan:")
	fmt.Println("ID: ", id)
	fmt.Println("Judul: ", movie[index].Judul)
	fmt.Println("Genre: ", movie[index].Genre)
	fmt.Println("Sutradara: ", movie[index].Sutradara)
	fmt.Println("Tahun: ", movie[index].Tahun)
	fmt.Println("Sinopsis: ", movie[index].Sinopsis)

	fmt.Println("Film berhasil diubah.")

	// Simpan perubahan ke array films
	count++
}

// Fungsi untuk melakukan sequential search dan Binary search pada array film berdasarkan ID
func cariFilmByID(id int) int {
	// I.S: ID dari film yang sedang dicari
	// F.S: Kembalikan indeks film yang memiliki ID yang dicari jika ditemukan, jika tidak kembalikan -1.
	var mid, left, right int
	if terurut == false {
		for i := 0; i < count; i++ {
			if movie[i].ID == id {
				return i // Film ditemukan, kembalikan indeksnya
			}
		}
	} else if terurut == true {
		left, right = 0, count-1

		for left <= right {
			mid = left + (right-left)/2

			if movie[mid].ID == id {
				return mid // Film ditemukan, kembalikan indeksnya
			}
			if movie[mid].ID < id {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}

	}
	return -1 // Film tidak ditemukan
}

func user() {
	// F.S.: Memungkinkan pengguna untuk mencari film dalam daftar film yang tersedia atau melakukan logout.
	// I.S.: Tidak ada. Fungsi ini dipanggil saat pengguna masuk ke mode pengguna.
	var m string
	for {
		fmt.Println("===================================================")
		fmt.Printf("**\t\t\tUser\t\t\t**\n")
		fmt.Println("===================================================")
		fmt.Println("1. Cari Film")
		fmt.Println("2. Logout")
		fmt.Println("===================================================")
		fmt.Println("Pilihanmu : ")

		fmt.Scan(&m)
		if m == "1" {
			cari_film()

		} else if m == "2" {
			fmt.Println("Logout...")
			time.Sleep(1 * time.Second)
			return
		}
	}
}

func cari_film() {
	// F.S.: Memungkinkan pengguna untuk mencari film berdasarkan ID, menampilkan daftar film, memilih film, mencari lanjutan, atau mengurutkan.
	// I.S.: Daftar film telah terinisialisasi, dan pengguna memilih untuk mencari film.
	var i int
	var m string
	m = "0"
	for {
		for i = 0; i < count; i++ {
			fmt.Println("===================================================")
			fmt.Printf("ID(%v)\tJudul\t\t:\t%v\n", movie[i].ID, movie[i].Judul)
			fmt.Printf("\tGenre\t\t:\t%v\n", movie[i].Genre)
			fmt.Printf("\tRilis tahun\t:\t%v\n", movie[i].Tahun)
			fmt.Printf("\tRating\t\t:\t%v(%v)\n", movie[i].Rating.hasil, movie[i].Rating.jumlah)
			fmt.Println("===================================================")
		}
		fmt.Println("Cari Film")
		fmt.Println("1. Pilih Film")
		fmt.Println("2. Cari lanjutan")
		fmt.Println("3. Urutkan")
		fmt.Println("4. Kembali")
		fmt.Println("Pilih")
		fmt.Scan(&m)
		if m == "1" {
			pilihfilm()

		} else if m == "2" {
			carilanjut()

		} else if m == "3" {
			urutkan()

		} else if m == "4" {
			return
		} else {
			fmt.Println("Invalid")
			time.Sleep(1 * time.Second)

		}
	}

}

func pilihfilm() {
	// F.S.: Memungkinkan pengguna untuk memilih film dari daftar film yang ditampilkan dan menambahkan rating ke film tersebut.
	// I.S.: Daftar film telah terinisialisasi, dan pengguna memilih untuk memilih film.

	var id, panjang int
	var m string
	for id != -1 {
		fmt.Println("masukan Id film yang ingin dipilih")
		fmt.Scan(&id)

		id = cariFilmByID(id)
		if id != -1 {
			panjang = 0
			fmt.Println("===================================================")
			fmt.Printf("Judul\t\t:\t%v\n", movie[id].Judul)
			fmt.Printf("Genre\t\t:\t%v\n", movie[id].Genre)
			fmt.Printf("Sutradara\t:\t%v\n", movie[id].Sutradara)
			fmt.Printf("Rilis tahun\t:\t%v\n", movie[id].Tahun)
			fmt.Printf("Rating\t\t:\t%v(%v)\n", movie[id].Rating.hasil, movie[id].Rating.jumlah)
			fmt.Println("===================================================")
			fmt.Println("Sinopsis :")
			for panjang < len(movie[id].Sinopsis) {
				fmt.Printf("%.52s\n", movie[id].Sinopsis[panjang:])
				panjang = panjang + 52
			}
			fmt.Println("===================================================")
			//fmt.Println(movie[id].Sinopsis)
			fmt.Println("1. Tambah rating")
			fmt.Println("2. Kembali")
			fmt.Scan(&m)
			if m == "1" {
				tambahrating(id)
			} else if m == "2" {
				return
			}
			id = -1
		} else {
			fmt.Println("Tidak Terdapat Film Dengan ID Tersebut")

		}

	}
}

func tambahrating(id int) {
	// F.S.: Menambahkan rating baru untuk film yang dipilih oleh pengguna.
	// I.S.: Pengguna telah memilih film untuk menambahkan rating.
	var rating float64
	//kasih batasan rating 1-5
	fmt.Println("masukan rating 1-5")
	fmt.Scan(&rating)
	for rating < 0 || rating > 5 {
		fmt.Println("masukan invalid")
		time.Sleep(1 * time.Second)
		fmt.Println("masukan rating 1-5")
		fmt.Scan(&rating)
	}
	movie[id].Rating.jumlah++
	movie[id].Rating.total = movie[id].Rating.total + rating
	movie[id].Rating.hasil = movie[id].Rating.total / movie[id].Rating.jumlah
}

func carilanjut() {
	// F.S.: Memungkinkan pengguna untuk melakukan pencarian lanjutan berdasarkan genre, sutradara, atau tahun rilis film.
	// I.S.: Pengguna telah memilih untuk melakukan pencarian lanjutan.
	var m string
	for {
		fmt.Println("===================================================")
		fmt.Printf("**\t\t\tCari Lanjut\t\t\t**\n")
		fmt.Println("===================================================")
		fmt.Println("Berdasarkan apa?")
		fmt.Println("1. Genre")
		fmt.Println("2. Sutradara")
		fmt.Println("3. Tahun Rilis")
		fmt.Println("4. Kembali")
		fmt.Println("===================================================")
		fmt.Println("Pilih")
		fmt.Scan(&m)
		if m == "1" {
			carilanjutgenre()

		} else if m == "2" {
			carilanjutsutradara()

		} else if m == "3" {
			carilanjuttahun()

		} else if m == "4" {
			return
		} else {
			fmt.Println("Invalid")
			time.Sleep(1 * time.Second)

		}
	}

}

func carilanjutsutradara() {
	// F.S.: Menampilkan daftar sutradara yang tersedia kemudian meminta input nama sutradara yang ingin dicari. Setelah itu, menampilkan daftar film yang disutradarai oleh sutradara yang dicari.
	// I.S.: Daftar film telah terinisialisasi, dan pengguna memilih untuk mencari film berdasarkan sutradara.

	var m, Sutra string
	var i, j, n int
	var ada bool
	var sut [NMAX]string
	n = 1
	fmt.Println("Yang Tersedia")
	for i = 0; i < count; i++ {
		ada = false
		j = 0
		for j < n && ada == false {
			if sut[j] == movie[i].Sutradara {
				ada = true
			}
			j++
		}
		if ada == false {
			sut[n] = movie[i].Sutradara
			n++
			fmt.Println(movie[i].Sutradara)
		}
	}
	fmt.Print("Masukan Nama yang Ingin Dicari : ")
	fmt.Scan(&Sutra)
	for i = 0; i < count; i++ {
		if Sutra == movie[i].Sutradara {
			fmt.Println("===================================================")
			fmt.Printf("ID(%v)\tJudul\t\t:\t%v\n", movie[i].ID, movie[i].Judul)
			fmt.Printf("\tGenre\t\t:\t%v\n", movie[i].Genre)
			fmt.Printf("\tRilis tahun\t:\t%v\n", movie[i].Tahun)
			fmt.Printf("\tRating\t\t:\t%v(%v)\n", movie[i].Rating.hasil, movie[i].Rating.jumlah)
			fmt.Println("===================================================")
		}
	}
	for {
		fmt.Println("1. Pilih Film")
		fmt.Println("2. Kembali")
		fmt.Println("Pilih")
		fmt.Scan(&m)
		if m == "1" {
			pilihfilm()

		} else if m == "2" {
			return
		} else {
			fmt.Println("Invalid")

		}
	}
}

func carilanjutgenre() {
	// F.S.: Meminta input genre film yang ingin dicari, kemudian menampilkan daftar film berdasarkan genre yang dipilih oleh pengguna.
	// I.S.: Daftar film telah terinisialisasi, dan pengguna memilih untuk mencari film berdasarkan genre.

	var m, genre string
	var i, tot int
	tot = 0
	fmt.Println("Pilihan Genre Film:")
	fmt.Println("1. Drama")
	fmt.Println("2. Action")
	fmt.Println("3. Komedi")
	fmt.Println("4. Horor")
	fmt.Println("5. Romance")
	fmt.Print("Pilihan Anda: ")
	fmt.Scan(&genre)
	// Mengonversi genre dari angka ke string menggunakan if-else
	if genre == "1" {
		for i = 0; i < count; i++ {
			if movie[i].Genre == "Drama" {
				fmt.Println("===================================================")
				fmt.Printf("ID(%v)\tJudul\t\t:\t%v\n", movie[i].ID, movie[i].Judul)
				fmt.Printf("\tGenre\t\t:\t%v\n", movie[i].Genre)
				fmt.Printf("\tRilis tahun\t:\t%v\n", movie[i].Tahun)
				fmt.Printf("\tRating\t\t:\t%v(%v)\n", movie[i].Rating.hasil, movie[i].Rating.jumlah)
				fmt.Println("===================================================")
				tot++
			}
		}
		if tot != 0 {
			for {
				fmt.Println("Cari Film")
				fmt.Println("1. Pilih Film")
				fmt.Println("2. Kembali")
				fmt.Println("Pilih")
				fmt.Scan(&m)
				if m == "1" {
					pilihfilm()

				} else if m == "2" {
					return
				} else {
					fmt.Println("Invalid")
				}
			}
		} else {
			fmt.Println("Tidak Terdapat Film dengan Genre Tersebut")
		}

	} else if genre == "2" {
		for i = 0; i < count; i++ {
			if movie[i].Genre == "Action" {
				fmt.Println("===================================================")
				fmt.Printf("ID(%v)\tJudul\t\t:\t%v\n", movie[i].ID, movie[i].Judul)
				fmt.Printf("\tGenre\t\t:\t%v\n", movie[i].Genre)
				fmt.Printf("\tRilis tahun\t:\t%v\n", movie[i].Tahun)
				fmt.Printf("\tRating\t\t:\t%v(%v)\n", movie[i].Rating.hasil, movie[i].Rating.jumlah)
				fmt.Println("===================================================")
				tot++
			}
		}
		if tot != 0 {
			for {
				fmt.Println("Cari Film")
				fmt.Println("1. Pilih Film")
				fmt.Println("2. Kembali")
				fmt.Println("Pilih")
				fmt.Scan(&m)
				if m == "1" {
					pilihfilm()

				} else if m == "2" {
					return
				} else {
					fmt.Println("Invalid")
				}
			}
		} else {
			fmt.Println("Tidak Terdapat Film dengan Genre Tersebut")
		}
	} else if genre == "3" {
		for i = 0; i < count; i++ {
			if movie[i].Genre == "Comedy" {
				fmt.Println("===================================================")
				fmt.Printf("ID(%v)\tJudul\t\t:\t%v\n", movie[i].ID, movie[i].Judul)
				fmt.Printf("\tGenre\t\t:\t%v\n", movie[i].Genre)
				fmt.Printf("\tRilis tahun\t:\t%v\n", movie[i].Tahun)
				fmt.Printf("\tRating\t\t:\t%v(%v)\n", movie[i].Rating.hasil, movie[i].Rating.jumlah)
				fmt.Println("===================================================")
				tot++
			}
		}
		if tot != 0 {
			for {
				fmt.Println("Cari Film")
				fmt.Println("1. Pilih Film")
				fmt.Println("2. Kembali")
				fmt.Println("Pilih")
				fmt.Scan(&m)
				if m == "1" {
					pilihfilm()

				} else if m == "2" {
					return
				} else {
					fmt.Println("Invalid")
				}
			}
		} else {
			fmt.Println("Tidak Terdapat Film dengan Genre Tersebut")
		}
	} else if genre == "4" {
		for i = 0; i < count; i++ {
			if movie[i].Genre == "Horror" {
				fmt.Println("===================================================")
				fmt.Printf("ID(%v)\tJudul\t\t:\t%v\n", movie[i].ID, movie[i].Judul)
				fmt.Printf("\tGenre\t\t:\t%v\n", movie[i].Genre)
				fmt.Printf("\tRilis tahun\t:\t%v\n", movie[i].Tahun)
				fmt.Printf("\tRating\t\t:\t%v(%v)\n", movie[i].Rating.hasil, movie[i].Rating.jumlah)
				fmt.Println("===================================================")
				tot++
			}
		}
		if tot != 0 {
			for {
				fmt.Println("Cari Film")
				fmt.Println("1. Pilih Film")
				fmt.Println("2. Kembali")
				fmt.Println("Pilih")
				fmt.Scan(&m)
				if m == "1" {
					pilihfilm()

				} else if m == "2" {
					return
				} else {
					fmt.Println("Invalid")
				}
			}
		} else {
			fmt.Println("Tidak Terdapat Film dengan Genre Tersebut")
		}
	} else if genre == "5" {
		for i = 0; i < count; i++ {
			if movie[i].Genre == "Romance" {
				fmt.Println("===================================================")
				fmt.Printf("ID(%v)\tJudul\t\t:\t%v\n", movie[i].ID, movie[i].Judul)
				fmt.Printf("\tGenre\t\t:\t%v\n", movie[i].Genre)
				fmt.Printf("\tRilis tahun\t:\t%v\n", movie[i].Tahun)
				fmt.Printf("\tRating\t\t:\t%v(%v)\n", movie[i].Rating.hasil, movie[i].Rating.jumlah)
				fmt.Println("===================================================")
				tot++
			}
		}
		if tot != 0 {
			for {
				fmt.Println("Cari Film")
				fmt.Println("1. Pilih Film")
				fmt.Println("2. Kembali")
				fmt.Println("Pilih")
				fmt.Scan(&m)
				if m == "1" {
					pilihfilm()

				} else if m == "2" {
					return
				} else {
					fmt.Println("Invalid")

				}
			}
		} else {
			fmt.Println("Tidak Terdapat Film dengan Genre Tersebut")
		}
	} else {
		fmt.Println("Genre tidak valid.")
	}
}

func carilanjuttahun() {
	// F.S.: Menampilkan daftar tahun rilis yang tersedia kemudian meminta input tahun rilis yang ingin dicari. Setelah itu, menampilkan daftar film yang dirilis pada tahun yang dicari.
	// I.S.: Daftar film telah terinisialisasi, dan pengguna memilih untuk mencari film berdasarkan tahun rilis.

	var m string
	var i, j, n, th int
	var ada bool
	var tr [NMAX]int
	n = 1
	fmt.Println("Yang Tersedia")
	for i = 0; i < count; i++ {
		ada = false
		j = 0
		for j < n && ada == false {
			if tr[j] == movie[i].Tahun {
				ada = true
			}
			j++
		}
		if ada == false {
			tr[n] = movie[i].Tahun
			n++
			fmt.Println(movie[i].Tahun)
		}
	}
	fmt.Print("Masukan Tahun yang Ingin Dicari : ")
	fmt.Scan(&th)
	for i = 0; i < count; i++ {
		if th == movie[i].Tahun {
			fmt.Println("===================================================")
			fmt.Printf("ID(%v)\tJudul\t\t:\t%v\n", movie[i].ID, movie[i].Judul)
			fmt.Printf("\tGenre\t\t:\t%v\n", movie[i].Genre)
			fmt.Printf("\tRilis tahun\t:\t%v\n", movie[i].Tahun)
			fmt.Printf("\tRating\t\t:\t%v(%v)\n", movie[i].Rating.hasil, movie[i].Rating.jumlah)
			fmt.Println("===================================================")
		}
	}
	for {
		fmt.Println("1. Pilih Film")
		fmt.Println("2. Kembali")
		fmt.Println("Pilih")
		fmt.Scan(&m)
		if m == "1" {
			pilihfilm()

		} else if m == "2" {
			return
		} else {
			fmt.Println("Invalid")
		}
	}
}

func urutkan() {
	// F.S.: Memungkinkan pengguna untuk mengurutkan daftar film berdasarkan rating, tahun rilis, nama, atau ID.
	// I.S.: Daftar film telah terinisialisasi, dan pengguna memilih untuk mengurutkan daftar film.

	var m string
	for {
		fmt.Println("===================================================")
		fmt.Printf("**\t\t\tUrutkan\t\t\t**\n")
		fmt.Println("===================================================")
		fmt.Println("Berdasarkan apa?")
		fmt.Println("1. Rating")
		fmt.Println("2. Tahun Rilis")
		fmt.Println("3. Nama")
		fmt.Println("4. ID")
		fmt.Println("5. Kembali")
		fmt.Println("===================================================")
		fmt.Println("Pilih")
		fmt.Scan(&m)
		if m == "1" {
			urutkanrating()

		} else if m == "2" {
			urutkantahun()

		} else if m == "3" {
			urutkannama()

		} else if m == "4" {
			urutkanid()

		} else if m == "5" {
			return
		} else {
			fmt.Println("Invalid")
			time.Sleep(1 * time.Second)

		}
	}
}

func urutkanrating() {
	// F.S.: Memungkinkan pengguna untuk mengurutkan daftar film berdasarkan rating film secara ascending atau descending.
	// I.S.: Daftar film telah terinisialisasi, dan pengguna memilih untuk mengurutkan daftar film berdasarkan rating.

	var m string
	var j, i, idx int
	var temp film
	for {
		fmt.Println("Menu Sorted by Rating")
		fmt.Println("1. Ascending")
		fmt.Println("2. Descending")
		fmt.Println("3. Kembali")
		fmt.Println("Pilih")
		fmt.Scan(&m)
		if m == "1" {
			for i = 0; i < count; i++ {
				idx = i
				for j = i + 1; j < count; j++ {
					if movie[j].Rating.hasil < movie[idx].Rating.hasil {
						idx = j
					}
				}
				temp = movie[i]
				movie[i] = movie[idx]
				movie[idx] = temp
			}
			fmt.Println("Domne")
			cari_film()

		} else if m == "2" {
			for i = 0; i < count; i++ {
				idx = i
				for j = i + 1; j < count; j++ {
					if movie[j].Rating.hasil > movie[idx].Rating.hasil {
						idx = j
					}
				}
				temp = movie[i]
				movie[i] = movie[idx]
				movie[idx] = temp
			}
			fmt.Println("Domne")
			cari_film()

		} else if m == "3" {
			return
		} else {
			fmt.Println("Invalid")
			time.Sleep(1 * time.Second)

		}
	}
}
func urutkantahun() {
	// F.S.: Memungkinkan pengguna untuk mengurutkan daftar film berdasarkan tahun rilis film secara ascending atau descending.
	// I.S.: Daftar film telah terinisialisasi, dan pengguna memilih untuk mengurutkan daftar film berdasarkan tahun rilis.

	var temp film
	var m string
	var j, i int
	for {
		fmt.Println("Menu Sorted by Year")
		fmt.Println("1. Ascending")
		fmt.Println("2. Descending")
		fmt.Println("3. Kembali")
		fmt.Println("Pilih")
		fmt.Scan(&m)
		if m == "1" {
			for i = 1; i < count; i++ {
				for j = i; j > 0 && movie[j-1].Tahun > movie[j].Tahun; j-- {
					temp = movie[j]
					movie[j] = movie[j-1]
					movie[j-1] = temp
				}
			}
			fmt.Println("Domne")
			cari_film()

		} else if m == "2" {
			for i = 1; i < count; i++ {
				for j = i; j > 0 && movie[j-1].Tahun < movie[j].Tahun; j-- {
					temp = movie[j]
					movie[j] = movie[j-1]
					movie[j-1] = temp
				}
			}
			fmt.Println("Domne")
			cari_film()

		} else if m == "3" {
			return
		} else {
			fmt.Println("Invalid")
			time.Sleep(1 * time.Second)

		}
	}
}

func urutkannama() {
	// F.S.: Mengurutkan daftar film berdasarkan judul film secara ascending atau descending.
	// I.S.: Daftar film telah terinisialisasi, dan pengguna memilih untuk mengurutkan daftar film berdasarkan judul.

	var temp film
	var m string
	var j, i int
	for {
		fmt.Println("Menu Sorted by Name")
		fmt.Println("1. Ascending")
		fmt.Println("2. Descending")
		fmt.Println("3. Kembali")
		fmt.Println("Pilih")
		fmt.Scan(&m)
		if m == "1" {
			for i = 1; i < count; i++ {
				for j = i; j > 0 && movie[j-1].Judul > movie[j].Judul; j-- {
					temp = movie[j]
					movie[j] = movie[j-1]
					movie[j-1] = temp
				}
			}
			fmt.Println("Domne")
			cari_film()

		} else if m == "2" {
			for i = 1; i < count; i++ {
				for j = i; j > 0 && movie[j-1].Judul < movie[j].Judul; j-- {
					temp = movie[j]
					movie[j] = movie[j-1]
					movie[j-1] = temp
				}
			}
			fmt.Println("Domne")
			cari_film()

		} else if m == "3" {
			return
		} else {
			fmt.Println("Invalid")
			time.Sleep(1 * time.Second)

		}
	}
}

func urutkanid() {
	// F.S.: Mengurutkan daftar film berdasarkan ID film secara ascending atau descending.
	// I.S.: Daftar film telah terinisialisasi, dan pengguna memilih untuk mengurutkan daftar film berdasarkan ID.

	var m string
	var j, i, idx int
	var temp film
	for {
		fmt.Println("Menu Sorted by ID")
		fmt.Println("1. Ascending")
		fmt.Println("2. Descending")
		fmt.Println("3. Kembali")
		fmt.Println("Pilih")
		fmt.Scan(&m)
		if m == "1" {
			for i = 0; i < count; i++ {
				idx = i
				for j = i + 1; j < count; j++ {
					if movie[j].ID < movie[idx].ID {
						idx = j
					}
				}
				temp = movie[i]
				movie[i] = movie[idx]
				movie[idx] = temp
			}
			fmt.Println("Domne")
			terurut = true
			cari_film()

		} else if m == "2" {
			for i = 0; i < count; i++ {
				idx = i
				for j = i + 1; j < count; j++ {
					if movie[j].ID > movie[idx].ID {
						idx = j
					}
				}
				temp = movie[i]
				movie[i] = movie[idx]
				movie[idx] = temp
			}
			fmt.Println("Domne")
			terurut = true
			cari_film()
		} else if m == "3" {
			return
		} else {
			fmt.Println("Invalid")
			time.Sleep(1 * time.Second)

		}
	}
}

package main

import (
	"fmt"
	"os"
	"log"
	"github.com/rwcarlsen/goexif/exif"
)

func main(){
	if len(os.Args) < 2{
		log.Fatal("Kullanım: go run main.go <gorsel_yolu.jpg>")
	}

	fileName := os.Args[1]
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal("Dosya açılamadı:", err)
	}
	defer f.Close()

	x, err := exif.Decode(f)
	if err != nil {
		log.Fatal("EXIF verisi bulunamadı veya okunamadı:", err)
	}

	//Thumbnail decoding
	thumbDataü err := x.JpegThumbnail()
	if err != nil {
		fmt.Println("[-] Thumbnail bulunamadı veya silinmiş.")
	} else {
		baseName := filepath.Base(fileName)
		thumbName := "thumb" + baseName

		err = os.WriteFile(thumbName, thumbData, 0644)
		if err != nil {
			fmt.Printf("[-] Thumbnail kaydedilirken bir hata oluştur: %v\n", err)
		} else {
			fmt.Printf("[+] Veri bulundu. Gömülü orijinal thumbnail kaydedildi: %s\n", thumbName)
		}
	}

	camModel, _ := x.Get(exif.Model)
	fmt.Println("[*] Kamera Modeli:", camModel)
	
	tm, _ := x.DateTime()
	fmt.Println("[*] Çekim Tarihi:", tm)

	lat, long, err := x.LatLong()
	if err != nil {
		fmt.Println("[-] GPS Verisi Bulunamadı.")
	} else{
		fmt.Println("[+] GPS Koordinatları:\nEnlem: %f\n,Boylam: %f\n", lat, long)
		fmt.Println("[+] Google Maps Linki: https://www.google.com/maps/place/%f,%f\n", lat, long)
	}
}
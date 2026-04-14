package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"sync"
	"time"
)

// 1. Veri Modelleri (Structs)
// JSON dosyamızdaki mimariyi Go'nun statik tiplerine mapliyoruz.
type Project struct {
	Name    string   `json:"name"`
	RepoURL string   `json:"repo_url"`
	Tags    []string `json:"tags"`
}

type Category struct {
	ID       string    `json:"id"`
	Projects []Project `json:"projects"`
}

type Data struct {
	LastUpdated string     `json:"last_updated"`
	Categories  []Category `json:"categories"`
}

// 2. Eşzamanlı (Concurrent) Link Kontrol Fonksiyonu
func checkLink(p Project, wg *sync.WaitGroup, client *http.Client) {
	// Goroutine bittiğinde WaitGroup'a haber ver (Defer kuralı)
	defer wg.Done()

	req, err := http.NewRequest("GET", p.RepoURL, nil)
	if err != nil {
		fmt.Printf("❌ [HATA] %s: İstek oluşturulamadı (%s)\n", p.Name, p.RepoURL)
		return
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("❌ [ULAŞILAMIYOR] %s: Bağlantı hatası (%s)\n", p.Name, p.RepoURL)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		fmt.Printf("✅ [OK] %s\n", p.Name)
	} else {
		// 404 Not Found veya başka bir hata varsa yakalarız
		fmt.Printf("⚠️ [UYARI] %s: HTTP %d döndürdü (%s)\n", p.Name, resp.StatusCode, p.RepoURL)
	}
}

func main() {
	fmt.Println("🚀 Awesome Turkiye Validator Başlatılıyor...")

	// Proje ana dizininden scripts/cmd/validator/ klasörüne olan göreceli yol
	filePath := "../../../data/data.json"

	// Dosyayı belleğe al
	fileBytes, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Kritik Hata: data.json okunamadı. Yolu kontrol et: %v\n", err)
		os.Exit(1)
	}

	// JSON'ı Struct'a Unmarshal et
	var data Data
	err = json.Unmarshal(fileBytes, &data)
	if err != nil {
		fmt.Printf("Kritik Hata: JSON formatı bozuk: %v\n", err)
		os.Exit(1)
	}

	// 3. İş Parçacığı (Goroutine) Yönetimi
	var wg sync.WaitGroup

	// Senior Dokunuşu: Asla varsayılan HTTP Client kullanma, daima Timeout belirle.
	// Aksi takdirde cevap vermeyen bir sunucu scripti sonsuza dek askıda bırakır.
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	fmt.Printf("Veritabanında bulunan kategoriler taranıyor...\n")
	fmt.Println("--------------------------------------------------")

	// Tüm projeleri döngüye al ve her biri için ayrı bir Goroutine başlat
	for _, category := range data.Categories {
		for _, project := range category.Projects {
			wg.Add(1)                          // Bekleme grubuna 1 görev ekle
			go checkLink(project, &wg, client) // Eşzamanlı çalıştır
		}
	}

	// Tüm Goroutine'lerin bitmesini bekle
	wg.Wait()

	fmt.Println("--------------------------------------------------")
	fmt.Println("🏁 Tarama tamamlandı.")
}

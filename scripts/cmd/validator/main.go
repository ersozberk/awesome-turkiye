package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"sync"
	"sync/atomic" // SENIOR DOKUNUŞU: Eşzamanlı işlemlerde güvenli bayrak yönetimi
	"time"
)

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

// Global hata takipçisi (herhangi bir goroutine hata bulursa bunu true yapacak)
var hasError atomic.Bool

func checkLink(p Project, wg *sync.WaitGroup, client *http.Client) {
	defer wg.Done()

	req, err := http.NewRequest("GET", p.RepoURL, nil)
	if err != nil {
		fmt.Printf("❌ [HATA] %s: İstek oluşturulamadı (%s)\n", p.Name, p.RepoURL)
		hasError.Store(true) // HATA YAKALANDI
		return
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("❌ [ULAŞILAMIYOR] %s: Bağlantı hatası (%s)\n", p.Name, p.RepoURL)
		hasError.Store(true) // HATA YAKALANDI
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		fmt.Printf("✅ [OK] %s\n", p.Name)
	} else {
		// UYARI yerine artık KIRIK LİNK diyoruz ve hata olarak işaretliyoruz
		fmt.Printf("❌ [KIRIK LİNK] %s: HTTP %d döndürdü (%s)\n", p.Name, resp.StatusCode, p.RepoURL)
		hasError.Store(true) // HATA YAKALANDI
	}
}

func main() {
	fmt.Println("🚀 Awesome Turkiye Validator Başlatılıyor...")

	filePath := "../../../data/data.json"
	fileBytes, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Kritik Hata: data.json okunamadı: %v\n", err)
		os.Exit(1)
	}

	var data Data
	if err := json.Unmarshal(fileBytes, &data); err != nil {
		fmt.Printf("Kritik Hata: JSON formatı bozuk: %v\n", err)
		os.Exit(1)
	}

	var wg sync.WaitGroup
	client := &http.Client{Timeout: 10 * time.Second}

	fmt.Printf("Veritabanı taranıyor...\n")
	fmt.Println("--------------------------------------------------")

	for _, category := range data.Categories {
		for _, project := range category.Projects {
			wg.Add(1)
			go checkLink(project, &wg, client)
		}
	}

	wg.Wait()

	fmt.Println("--------------------------------------------------")

	// FİNAL KONTROLÜ: GitHub'ın kırmızı veya yeşil yakacağına burada karar veriyoruz
	if hasError.Load() {
		fmt.Println("🛑 Tarama tamamlandı ancak KIRIK LİNKLER bulundu! PR reddediliyor.")
		os.Exit(1) // Sistemi çökerterek CI pipeline'ı durdur
	} else {
		fmt.Println("🏁 Tarama tamamlandı. Tüm linkler sağlıklı.")
		os.Exit(0) // Başarılı çıkış
	}
}

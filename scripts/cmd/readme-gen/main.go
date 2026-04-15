package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

// JSON Şemamıza uygun veri modelleri
type Project struct {
	Name        string            `json:"name"`
	RepoURL     string            `json:"repo_url"`
	Description map[string]string `json:"description"`
	Tags        []string          `json:"tags"`
}

type Category struct {
	ID       string            `json:"id"`
	Title    map[string]string `json:"title"`
	Projects []Project         `json:"projects"`
}

type Data struct {
	LastUpdated string     `json:"last_updated"`
	Categories  []Category `json:"categories"`
}

func main() {
	fmt.Println("📝 README Jeneratörü Başlatılıyor...")

	// JSON'ı oku (Yine kök dizine 3 adım çıkıyoruz)
	filePath := "../../../data/data.json"
	fileBytes, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Printf("❌ Hata: data.json okunamadı: %v\n", err)
		os.Exit(1)
	}

	var data Data
	if err := json.Unmarshal(fileBytes, &data); err != nil {
		fmt.Printf("❌ Hata: JSON ayrıştırılamadı: %v\n", err)
		os.Exit(1)
	}

	// İngilizce README oluştur
	enContent := generateMarkdown(data, "en")
	os.WriteFile("../../../README.md", []byte(enContent), 0644)
	fmt.Println("✅ README.md (İngilizce) başarıyla oluşturuldu.")

	// Türkçe README oluştur
	trContent := generateMarkdown(data, "tr")
	os.WriteFile("../../../README.tr.md", []byte(trContent), 0644)
	fmt.Println("✅ README.tr.md (Türkçe) başarıyla oluşturuldu.")
}

func generateMarkdown(data Data, lang string) string {
	var sb strings.Builder

	// ==========================================
	// 1. HERO SECTION (BÜYÜK GÖRSEL VE ORTALANMIŞ BAŞLIK)
	// ==========================================
	sb.WriteString("<div align=\"center\">\n\n")

	// Dinamik, hareketli ve projenin renklerine (Kırmızı/Siyah) uygun GitHub afişi
	sb.WriteString("<img src=\"https://capsule-render.vercel.app/api?type=waving&color=ef4444&height=200&section=header&text=Awesome%20Türkiye&fontSize=60&fontAlignY=35&animation=twinkling&fontColor=ffffff\" alt=\"Awesome Turkiye Banner\" />\n\n")

	if lang == "en" {
		sb.WriteString("### 🌍 The Digital Map of Turkey's Open Source & Tech Ecosystem\n\n")
		sb.WriteString("A curated, community-driven database of Turkish open-source projects, tech communities, and digital nomad guides.\n\n")
	} else {
		sb.WriteString("### 🌍 Türkiye'nin Açık Kaynak ve Teknoloji Ekosistemi Haritası\n\n")
		sb.WriteString("Türkiye'den çıkan açık kaynak projeler, teknoloji toplulukları ve dijital yaşam rehberleri için topluluk odaklı veritabanı.\n\n")
	}

	// Rozetler (Badges)
	sb.WriteString("![Validator Status](https://github.com/ersozberk/awesome-turkiye/actions/workflows/validate.yml/badge.svg) ")
	sb.WriteString("![Generator Status](https://github.com/ersozberk/awesome-turkiye/actions/workflows/generate-readme.yml/badge.svg) ")
	sb.WriteString("![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)\n\n")

	// Hızlı Linkler
	if lang == "en" {
		sb.WriteString("🌐 **[Explore the Live Website](https://awesome-turkiye.vercel.app)** • 🤝 **[How to Contribute](CONTRIBUTING.md)** • 🇹🇷 **[Türkçe](README.tr.md)**\n\n")
	} else {
		sb.WriteString("🌐 **[Canlı Web Sitesini Keşfet](https://awesome-turkiye.vercel.app)** • 🤝 **[Nasıl Katkıda Bulunurum?](CONTRIBUTING.md)** • 🇬🇧 **[English](README.md)**\n\n")
	}
	sb.WriteString("</div>\n\n---\n\n")

	// ==========================================
	// 2. İÇİNDEKİLER TABLOSU (TABLE OF CONTENTS)
	// ==========================================
	if lang == "en" {
		sb.WriteString("## 📋 Table of Contents\n\n")
	} else {
		sb.WriteString("## 📋 İçindekiler\n\n")
	}

	for _, cat := range data.Categories {
		// DÜZELTİLEN KISIM: Harita (Map) erişimi
		title := cat.Title["en"]
		if lang == "tr" {
			title = cat.Title["tr"]
		}
		// Tıklanabilir içindekiler linkleri
		sb.WriteString(fmt.Sprintf("- [%s](#%s)\n", title, cat.ID))
	}
	sb.WriteString("\n---\n\n")

	// ==========================================
	// 3. KATEGORİLER VE PROJELER
	// ==========================================
	for _, cat := range data.Categories {
		// DÜZELTİLEN KISIM: Harita (Map) erişimi
		title := cat.Title["en"]
		if lang == "tr" {
			title = cat.Title["tr"]
		}

		// Kategori Başlığı (HTML anchor ile linklemeyi garantiye alıyoruz)
		sb.WriteString(fmt.Sprintf("## <a name=\"%s\"></a>%s\n\n", cat.ID, title))

		if len(cat.Projects) == 0 {
			if lang == "en" {
				sb.WriteString("*No projects in this category yet. Be the first to [add one](../CONTRIBUTING.md)!*\n\n")
			} else {
				sb.WriteString("*Bu kategoride henüz proje yok. İlk ekleyen sen ol: [PR Gönder](../CONTRIBUTING.md)!*\n\n")
			}
			continue
		}

		for _, proj := range cat.Projects {
			// DÜZELTİLEN KISIM: Harita (Map) erişimi
			desc := proj.Description["en"]
			if lang == "tr" {
				desc = proj.Description["tr"]
			}

			// Proje satırı (Bold isim, link ve italik açıklama)
			sb.WriteString(fmt.Sprintf("- **[%s](%s)** - *%s*\n", proj.Name, proj.RepoURL, desc))
		}

		// UX Dokunuşu: Her kategori sonuna "Başa Dön" linki
		if lang == "en" {
			sb.WriteString("\n[⬆️ Back to Top](#-table-of-contents)\n\n")
		} else {
			sb.WriteString("\n[⬆️ Başa Dön](#-içindekiler)\n\n")
		}
	}

	// ==========================================
	// 4. GAMIFICATION (CONTRIBUTORS) & FOOTER
	// ==========================================
	sb.WriteString("---\n\n")
	if lang == "en" {
		sb.WriteString("<div align=\"center\">\n\n")
		sb.WriteString("## 💖 Contributors\n\n")
		sb.WriteString("Thanks to everyone who has contributed! Submit a PR to join the hall of fame.\n\n")
	} else {
		sb.WriteString("<div align=\"center\">\n\n")
		sb.WriteString("## 💖 Katkıda Bulunanlar\n\n")
		sb.WriteString("Bu projeyi büyüten herkese teşekkürler! Yüzünüzü burada görmek için bir PR gönderin.\n\n")
	}

	sb.WriteString("[![Contributors](https://contrib.rocks/image?repo=ersozberk/awesome-turkiye)](https://github.com/ersozberk/awesome-turkiye/graphs/contributors)\n\n")

	if lang == "en" {
		sb.WriteString("<br/><p>Built with ❤️ using Go and GitHub Actions. MIT License.</p>\n")
	} else {
		sb.WriteString("<br/><p>Go ve GitHub Actions kullanılarak ❤️ ile geliştirildi. MIT Lisansı.</p>\n")
	}
	sb.WriteString("</div>\n")

	return sb.String()
}

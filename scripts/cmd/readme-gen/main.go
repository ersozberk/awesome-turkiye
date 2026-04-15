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
	// 1. HEADER (MİNİMALİST VE KURUMSAL)
	// ==========================================
	sb.WriteString("<div align=\"center\">\n\n")
	sb.WriteString("# 🇹🇷 Awesome Türkiye\n\n")

	if lang == "en" {
		sb.WriteString("**The Digital Map of Turkey's Open Source & Tech Ecosystem**\n\n")
		sb.WriteString("A curated, community-driven database of Turkish open-source projects, tech communities, and digital nomad guides.\n\n")
	} else {
		sb.WriteString("**Türkiye'nin Açık Kaynak ve Teknoloji Ekosistemi Haritası**\n\n")
		sb.WriteString("Türkiye'den çıkan açık kaynak projeler, teknoloji toplulukları ve dijital yaşam rehberleri için topluluk odaklı veritabanı.\n\n")
	}

	// Rozetler (Yan yana, düzenli)
	sb.WriteString("[![Validator](https://github.com/ersozberk/awesome-turkiye/actions/workflows/validate.yml/badge.svg)](https://github.com/ersozberk/awesome-turkiye/actions) ")
	sb.WriteString("[![Generator](https://github.com/ersozberk/awesome-turkiye/actions/workflows/generate-readme.yml/badge.svg)](https://github.com/ersozberk/awesome-turkiye/actions) ")
	sb.WriteString("[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)\n\n")

	if lang == "en" {
		sb.WriteString("🌐 **[Live Engine](https://awesome-turkiye.vercel.app)** &nbsp;•&nbsp; 🤝 **[Contribute](CONTRIBUTING.md)** &nbsp;•&nbsp; 🇹🇷 **[Türkçe](README.tr.md)**\n\n")
	} else {
		sb.WriteString("🌐 **[Canlı Keşif Motoru](https://awesome-turkiye.vercel.app)** &nbsp;•&nbsp; 🤝 **[Katkıda Bulun](CONTRIBUTING.md)** &nbsp;•&nbsp; 🇬🇧 **[English](README.md)**\n\n")
	}
	sb.WriteString("</div>\n\n---\n\n")

	// ==========================================
	// 2. MİMARİ ÖZETİ (ARCHITECTURE OVERVIEW)
	// ==========================================
	if lang == "en" {
		sb.WriteString("## ⚙️ Architecture\n")
		sb.WriteString("This is not just a static list. It's a self-maintaining ecosystem:\n")
		sb.WriteString("- **Data Layer:** Centralized `data.json` acts as the single source of truth.\n")
		sb.WriteString("- **CI/CD Pipeline:** Custom **Go** scripts automatically validate broken links and generate bilingual READMEs via GitHub Actions.\n")
		sb.WriteString("- **Discovery Engine:** A lightning-fast **Next.js** web interface built for exploration.\n\n")
	} else {
		sb.WriteString("## ⚙️ Mimari Altyapı\n")
		sb.WriteString("Bu proje sadece statik bir liste değil, kendi kendini yöneten bir sistemdir:\n")
		sb.WriteString("- **Veri Katmanı:** Tüm içerik merkezi bir `data.json` dosyasından yönetilir.\n")
		sb.WriteString("- **CI/CD Boru Hattı:** Özel yazılmış **Go** botları, kırık linkleri tespit eder ve GitHub Actions üzerinden iki dilde README üretir.\n")
		sb.WriteString("- **Keşif Motoru:** Keşfedilebilirlik için optimize edilmiş, ışık hızında bir **Next.js** web arayüzü.\n\n")
	}

	// ==========================================
	// 3. İÇİNDEKİLER TABLOSU
	// ==========================================
	if lang == "en" {
		sb.WriteString("## 📋 Table of Contents\n\n")
	} else {
		sb.WriteString("## 📋 İçindekiler\n\n")
	}

	for _, cat := range data.Categories {
		title := cat.Title["en"]
		if lang == "tr" {
			title = cat.Title["tr"]
		}
		sb.WriteString(fmt.Sprintf("- [%s](#%s)\n", title, cat.ID))
	}
	sb.WriteString("\n---\n\n")

	// ==========================================
	// 4. KATEGORİLER VE PROJELER (TABLO GÖRÜNÜMÜ)
	// ==========================================
	for _, cat := range data.Categories {
		title := cat.Title["en"]
		if lang == "tr" {
			title = cat.Title["tr"]
		}

		sb.WriteString(fmt.Sprintf("## <a name=\"%s\"></a>%s\n\n", cat.ID, title))

		if len(cat.Projects) == 0 {
			if lang == "en" {
				sb.WriteString("> *No projects here yet. Be the first to [submit a PR](../CONTRIBUTING.md)!*\n\n")
			} else {
				sb.WriteString("> *Bu kategoride henüz veri yok. İlk ekleyen olmak için [PR Gönder](../CONTRIBUTING.md)!*\n\n")
			}
			continue
		}

		// TABLO BAŞLIKLARI
		if lang == "en" {
			sb.WriteString("| 🚀 Project | 📝 Description | 🏷️ Tags |\n")
		} else {
			sb.WriteString("| 🚀 Proje | 📝 Açıklama | 🏷️ Etiketler |\n")
		}
		sb.WriteString("| :--- | :--- | :--- |\n")

		// TABLO İÇERİĞİ
		for _, proj := range cat.Projects {
			desc := proj.Description["en"]
			if lang == "tr" {
				desc = proj.Description["tr"]
			}

			// Etiketleri yan yana kod bloğu formatında birleştir (Örn: `react`, `go`)
			tagsFormatted := ""
			if len(proj.Tags) > 0 {
				tagsFormatted = "`" + strings.Join(proj.Tags, "`, `") + "`"
			}

			sb.WriteString(fmt.Sprintf("| **[%s](%s)** | %s | %s |\n", proj.Name, proj.RepoURL, desc, tagsFormatted))
		}

		if lang == "en" {
			sb.WriteString("\n<div align=\"right\"><a href=\"#-table-of-contents\">⬆️ Back to Top</a></div>\n\n")
		} else {
			sb.WriteString("\n<div align=\"right\"><a href=\"#-içindekiler\">⬆️ Başa Dön</a></div>\n\n")
		}
	}

	// ==========================================
	// 5. CONTRIBUTORS & FOOTER
	// ==========================================
	sb.WriteString("---\n\n")
	if lang == "en" {
		sb.WriteString("## 💖 Contributors\n\n")
	} else {
		sb.WriteString("## 💖 Katkıda Bulunanlar\n\n")
	}

	sb.WriteString("<a href=\"https://github.com/ersozberk/awesome-turkiye/graphs/contributors\">\n")
	sb.WriteString("  <img src=\"https://contrib.rocks/image?repo=ersozberk/awesome-turkiye\" alt=\"Contributors\" />\n")
	sb.WriteString("</a>\n\n")

	if lang == "en" {
		sb.WriteString("> Built with using Go and Next.js. Distributed under the MIT License.\n")
	} else {
		sb.WriteString("> Go ve Next.js mimarisiyle geliştirildi. MIT Lisansı ile dağıtılmaktadır.\n")
	}

	return sb.String()
}

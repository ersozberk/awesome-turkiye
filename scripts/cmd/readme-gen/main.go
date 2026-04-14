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

// Markdown metnini inşa eden ana fonksiyon
func generateMarkdown(data Data, lang string) string {
	var sb strings.Builder

	// 1. Başlık ve Dil Geçişi (Header)
	if lang == "en" {
		sb.WriteString("# Awesome Turkiye 🇹🇷\n\n")
		// CI/CD ve Lisans Rozetleri
		sb.WriteString("![Validator Status](https://github.com/ersozberk/awesome-turkiye/actions/workflows/validate.yml/badge.svg) ")
		sb.WriteString("![Generator Status](https://github.com/ersozberk/awesome-turkiye/actions/workflows/generate-readme.yml/badge.svg) ")
		sb.WriteString("![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)\n\n")

		sb.WriteString("A curated database of the Turkish Open Source Ecosystem, Communities, and Digital Life.\n\n")
		sb.WriteString("🌍 **[Türkçe versiyon için tıklayın](README.tr.md)** | 🌐 **[Website & Discovery Engine](https://awesome-turkiye-something.vercel.app)**\n\n") // Vercel linkini buraya ekleyebilirsin
	} else {
		sb.WriteString("# Harika Türkiye 🇹🇷\n\n")
		sb.WriteString("![Validator Status](https://github.com/ersozberk/awesome-turkiye/actions/workflows/validate.yml/badge.svg) ")
		sb.WriteString("![Generator Status](https://github.com/ersozberk/awesome-turkiye/actions/workflows/generate-readme.yml/badge.svg) ")
		sb.WriteString("![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)\n\n")

		sb.WriteString("Türkiye Açık Kaynak Ekosistemi, Toplulukları ve Dijital Yaşamı için küratörlü veritabanı.\n\n")
		sb.WriteString("🌍 **[Click here for the English version](README.md)** | 🌐 **[Web Sitesi ve Keşif Motoru](https://awesome-turkiye-something.vercel.app)**\n\n")
	}

	// 2. İçindekiler Tablosu (Table of Contents)
	sb.WriteString("## ")
	if lang == "en" {
		sb.WriteString("Contents\n")
	} else {
		sb.WriteString("İçindekiler\n")
	}

	for _, category := range data.Categories {
		title := category.Title[lang]
		// Başlığı URL formatına çevir (Boşlukları tire yap, küçük harfe çevir)
		anchor := strings.ToLower(strings.ReplaceAll(title, " ", "-"))
		sb.WriteString(fmt.Sprintf("- [%s](#%s)\n", title, anchor))
	}
	sb.WriteString("\n---\n\n")

	// 3. Kategoriler ve Projeler Listesi
	for _, category := range data.Categories {
		sb.WriteString(fmt.Sprintf("## %s\n\n", category.Title[lang]))

		for _, project := range category.Projects {
			desc := project.Description[lang]
			// Etiketleri formatla
			tags := ""
			for _, tag := range project.Tags {
				tags += fmt.Sprintf("`#%s` ", tag)
			}

			// Proje satırını yazdır
			sb.WriteString(fmt.Sprintf("- [%s](%s) - %s %s\n", project.Name, project.RepoURL, desc, tags))
		}
		sb.WriteString("\n")
	}

	// 4. Lisans Altbilgisi
	sb.WriteString("---\n\n")
	if lang == "en" {
		sb.WriteString("## 💖 Contributors\n\n")
		sb.WriteString("Thanks to everyone who has contributed to this project! Submit a PR to see your face here.\n\n")
	} else {
		sb.WriteString("## 💖 Katkıda Bulunanlar\n\n")
		sb.WriteString("Bu projeye katkı sağlayan herkese teşekkürler! Yüzünüzü burada görmek için bir PR gönderin.\n\n")
	}

	// Sihirli satır: GitHub API'sinden projeye PR atan herkesin resmini otomatik çeker
	sb.WriteString("[![Contributors](https://contrib.rocks/image?repo=ersozberk/awesome-turkiye)](https://github.com/ersozberk/awesome-turkiye/graphs/contributors)\n\n")

	sb.WriteString("---\n")
	if lang == "en" {
		sb.WriteString("Built with ❤️ using Go and GitHub Actions. MIT License.")
	} else {
		sb.WriteString("Go ve GitHub Actions kullanılarak ❤️ ile geliştirildi. MIT Lisansı.")
	}

	return sb.String()
}

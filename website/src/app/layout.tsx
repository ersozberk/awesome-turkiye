import type { Metadata } from 'next';
import { Inter } from 'next/font/google';
import './globals.css';

// Performans için Google fontlarını sunucu tarafında (build time) optimize ediyoruz
const inter = Inter({ subsets: ['latin'], display: 'swap' });

/**
 * Global Metadata ve SEO Konfigürasyonu
 * @description Bu obje, Next.js tarafından tüm sayfalara otomatik olarak <head> etiketleri (meta, title, og) olarak eklenir.
 * @see https://nextjs.org/docs/app/building-your-application/optimizing/metadata
 */
export const metadata: Metadata = {
  metadataBase: new URL('https://awesome-turkiye.vercel.app'),
  title: {
    default: 'Awesome Türkiye 🇹🇷 | Açık Kaynak Ekosistemi',
    template: '%s | Awesome Türkiye', // Alt sayfalarda otomatik başlık formatı
  },
  description: 'Türkiye\'nin açık kaynak yazılım ekosistemi, teknoloji toplulukları ve dijital yaşam rehberi. GitHub projelerini keşfedin, yeteneklerinizi sergileyin.',
  keywords: [
    'açık kaynak', 'open source', 'türkiye', 'yazılım', 'github', 
    'awesome list', 'teknoloji', 'frontend', 'backend', 'developer'
  ],
  authors: [{ name: 'Berk Ersöz', url: 'https://github.com/ersozberk' }],
  creator: 'Berk Ersöz',
  publisher: 'Awesome Türkiye Topluluğu',
  
  // Sosyal medyada (LinkedIn, X, Discord) link paylaşıldığında çıkacak zengin kart
  openGraph: {
    type: 'website',
    locale: 'tr_TR',
    url: 'https://awesome-turkiye.vercel.app', // Vercel canlı linkini buraya eklemelisin
    title: 'Awesome Türkiye 🇹🇷',
    description: 'Türkiye\'nin teknoloji ve açık kaynak haritası. Projeleri keşfet, ekosisteme katkı sağla.',
    siteName: 'Awesome Türkiye',
  },
  
  // Twitter/X için özel kart yapısı (Büyük resimli önizleme sağlar)
  twitter: {
    card: 'summary_large_image',
    title: 'Awesome Türkiye 🇹🇷',
    description: 'Türkiye\'nin açık kaynak yazılım ve teknoloji ekosistemi.',
    creator: '@senin_twitter_kullanici_adin', // Varsa ekle, yoksa bu satırı silebilirsin
  },
  
  // Arama motoru botlarına siteyi indeksleme izni veriyoruz
  robots: {
    index: true,
    follow: true,
    googleBot: {
      index: true,
      follow: true,
      'max-video-preview': -1,
      'max-image-preview': 'large',
      'max-snippet': -1,
    },
  },
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="tr" className="scroll-smooth">
      {/* antialiased: Yazı tiplerini işletim sistemi seviyesinde pürüzsüzleştirir (özellikle macOS için)
        selection: Kullanıcı metin seçtiğinde oluşan mavi rengi projenin kırmızı/turuncu temasına uydurur
      */}
      <body className={`${inter.className} bg-[#0a0a0a] text-gray-100 antialiased selection:bg-red-500/30 selection:text-red-200`}>
        {children}
      </body>
    </html>
  );
}
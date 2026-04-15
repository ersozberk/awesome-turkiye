import { ImageResponse } from 'next/og';

// Vercel'de bu dosyanın en hızlı Edge Node'larda çalışmasını sağlar
export const runtime = 'edge';

// Resmin meta verileri
export const alt = 'Awesome Türkiye Ekosistemi';
export const size = {
  width: 1200,
  height: 630,
};
export const contentType = 'image/png';

/**
 * Dinamik OG Image Jeneratörü
 * @description Bu fonksiyon, sitemizin linki paylaşıldığında gösterilecek resmi HTML/CSS kullanarak anında üretir.
 */
export default async function Image() {
  return new ImageResponse(
    (
      <div
        style={{
          background: 'linear-gradient(to bottom right, #0a0a0a, #1a1a1a)',
          width: '100%',
          height: '100%',
          display: 'flex',
          flexDirection: 'column',
          alignItems: 'center',
          justifyContent: 'center',
          fontFamily: 'sans-serif',
          border: '4px solid #ef4444', // Tailwind red-500 rengi
        }}
      >
        <div
          style={{
            display: 'flex',
            alignItems: 'center',
            justifyContent: 'center',
            marginBottom: '40px',
          }}
        >
          {/* Geometrik şık bir logo temsili */}
          <div style={{ width: '60px', height: '60px', borderRadius: '50%', background: '#ef4444', marginRight: '20px' }} />
          <h1
            style={{
              fontSize: '80px',
              fontWeight: '900',
              color: 'white',
              margin: 0,
              letterSpacing: '-2px',
            }}
          >
            Awesome <span style={{ color: '#ef4444' }}>Türkiye</span>
          </h1>
        </div>

        <p
          style={{
            fontSize: '32px',
            color: '#a3a3a3', // Tailwind gray-400
            textAlign: 'center',
            maxWidth: '800px',
            lineHeight: 1.4,
            margin: 0,
          }}
        >
          Türkiye'nin açık kaynak yazılım ekosistemi, teknoloji toplulukları ve dijital yaşam rehberi.
        </p>
        
        <div style={{ display: 'flex', marginTop: '60px', gap: '20px' }}>
          <span style={{ fontSize: '24px', color: '#ef4444', border: '2px solid #ef4444', padding: '10px 20px', borderRadius: '12px' }}>
            github.com/ersozberk/awesome-turkiye
          </span>
        </div>
      </div>
    ),
    {
      ...size,
    }
  );
}
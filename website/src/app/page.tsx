import { getAwesomeData } from './utils/getData';
import Dashboard from '@/components/Dashboard';

export default function Home() {
  // Veriyi sunucu tarafında (build anında) alıyoruz
  const data = getAwesomeData();

  return (
    <main className="min-h-screen bg-[#0a0a0a] text-white p-6 md:p-12 font-sans selection:bg-red-500 selection:text-white">
      
      <header className="max-w-5xl mx-auto mb-12 text-center md:text-left">
        <h1 className="text-5xl md:text-6xl font-extrabold tracking-tight">
          Awesome <span className="text-transparent bg-clip-text bg-gradient-to-r from-red-500 to-orange-400">Türkiye</span>
        </h1>
        <p className="mt-4 text-gray-400 text-lg md:text-xl max-w-2xl">
          Türkiye'nin açık kaynak yazılım ekosistemi, teknoloji toplulukları ve dijital yaşam rehberi. 
        </p>
      </header>

      {/* Etkileşimli Arama ve Filtreleme Bileşeni */}
      <Dashboard initialData={data} />
      
    </main>
  );
}
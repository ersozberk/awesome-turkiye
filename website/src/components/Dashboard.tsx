"use client"; // Bu bir istemci (tarayıcı) bileşenidir

import { useState } from "react";

export default function Dashboard({ initialData }: { initialData: any }) {
  const [searchTerm, setSearchTerm] = useState("");
  const [activeCategory, setActiveCategory] = useState("all");

  // Arama ve kategoriye göre projeleri filtreleme mantığı
  const filteredCategories = initialData.categories.map((category: any) => {
    const filteredProjects = category.projects.filter((project: any) => {
      const matchesSearch =
        project.name.toLowerCase().includes(searchTerm.toLowerCase()) ||
        project.description.tr.toLowerCase().includes(searchTerm.toLowerCase());
      
      const matchesCategory = activeCategory === "all" || category.id === activeCategory;

      return matchesSearch && matchesCategory;
    });

    return { ...category, projects: filteredProjects };
  }).filter((category: any) => category.projects.length > 0); // Boş kategorileri gizle

  return (
    <div className="max-w-5xl mx-auto space-y-8">
      
     
      {/* Arama ve Filtreleme Çubuğu */}
      <div className="bg-gray-800 p-6 rounded-2xl border border-gray-700 shadow-xl flex flex-col md:flex-row gap-4">
        <input
          type="text"
          aria-label="Proje veya teknoloji ara"
          placeholder="Proje veya teknoloji ara... (Örn: nextjs)"
          className="flex-1 bg-gray-900 text-white px-4 py-3 rounded-xl border border-gray-600 focus:outline-none focus:border-red-500 transition"
          value={searchTerm}
          onChange={(e) => setSearchTerm(e.target.value)}
        />
        
        <select 
          aria-label="Kategori seçin"
          className="bg-gray-900 text-white px-4 py-3 rounded-xl border border-gray-600 focus:outline-none focus:border-red-500"
          value={activeCategory}
          onChange={(e) => setActiveCategory(e.target.value)}
        >
          <option value="all">Tüm Kategoriler</option>
          {initialData.categories.map((cat: any) => (
            <option key={cat.id} value={cat.id}>{cat.title.tr}</option>
          ))}
        </select>
      </div>

      {/* Projelerin Listelenmesi */}
      <div className="space-y-12">
        {filteredCategories.length === 0 ? (
          <div className="text-center text-gray-400 py-10">Aradığınız kritere uygun proje bulunamadı.</div>
        ) : (
          filteredCategories.map((category: any) => (
            <div key={category.id} className="animate-fade-in">
              <h2 className="text-2xl font-bold mb-6 text-gray-100 flex items-center gap-3">
                <span className="w-2 h-8 bg-red-500 rounded-full"></span>
                {category.title.tr}
              </h2>
              
              <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
                {category.projects.map((project: any, index: number) => (
                  <div key={index} className="bg-gray-800 p-6 rounded-2xl border border-gray-700 hover:border-red-500 hover:shadow-[0_0_20px_rgba(239,68,68,0.2)] transition-all duration-300 group">
                    <a href={project.repo_url} target="_blank" rel="noopener noreferrer" className="text-xl font-bold text-white group-hover:text-red-400 transition-colors">
                      {project.name}
                    </a>
                    <p className="text-gray-400 mt-3 text-sm leading-relaxed min-h-[40px]">
                      {project.description.tr}
                    </p>
                    
                    <div className="mt-6 flex flex-wrap gap-2">
                      {project.tags.map((tag: string) => (
                        <span key={tag} className="bg-gray-900 border border-gray-700 text-xs px-3 py-1.5 rounded-lg text-gray-300 font-mono">
                          #{tag}
                        </span>
                      ))}
                    </div>
                  </div>
                ))}
              </div>
            </div>
          ))
        )}
      </div>
    </div>
  );
}
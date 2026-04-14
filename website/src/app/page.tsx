import { getAwesomeData } from './utils/getData';

export default function Home() {
  const data = getAwesomeData();

  return (
    <main className="min-h-screen bg-gray-900 text-white p-10 font-sans">
      <div className="max-w-5xl mx-auto">
        
        <header className="mb-12 border-b border-gray-700 pb-6">
          <h1 className="text-5xl font-extrabold text-transparent bg-clip-text bg-gradient-to-r from-red-500 to-white">
            Awesome Turkiye
          </h1>
          <p className="mt-4 text-gray-400 text-lg">
            Türkiye'nin açık kaynak ve teknoloji ekosistemi.
          </p>
        </header>

        <section className="space-y-12">
          {data.categories.map((category: any) => (
            <div key={category.id}>
              <h2 className="text-3xl font-bold mb-6 text-gray-100 border-l-4 border-red-500 pl-4">
                {category.title.tr} <span className="text-gray-500 text-sm ml-2">({category.title.en})</span>
              </h2>
              
              <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
                {category.projects.map((project: any, index: number) => (
                  <div key={index} className="bg-gray-800 p-6 rounded-xl border border-gray-700 hover:border-red-500 transition-colors">
                    <a href={project.repo_url} target="_blank" rel="noopener noreferrer" className="text-xl font-bold text-blue-400 hover:underline">
                      {project.name}
                    </a>
                    <p className="text-gray-300 mt-2">{project.description.tr}</p>
                    
                    <div className="mt-4 flex flex-wrap gap-2">
                      {project.tags.map((tag: string) => (
                        <span key={tag} className="bg-gray-700 text-xs px-2 py-1 rounded-md text-gray-300">
                          #{tag}
                        </span>
                      ))}
                    </div>
                  </div>
                ))}
              </div>
            </div>
          ))}
        </section>

      </div>
    </main>
  );
}
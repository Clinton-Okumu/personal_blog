import { useState, useMemo } from 'react';
import ArticleCard from '../components/ArticleCard';

interface Article {
  slug: string;
  title: string;
  description: string;
  publishedAt: string;
  readTime: number;
  category: string;
  featured?: boolean;
}

const HomeContent = () => {
  const [searchTerm, setSearchTerm] = useState('');

  // Enhanced sample article data
  const articles: Article[] = [
    {
      slug: 'my-first-article',
      title: 'My First Article',
      description:
        'This is a short description of my very first article. It talks about exciting new things and fresh perspectives on modern development.',
      publishedAt: '2025-01-15',
      readTime: 5,
      category: 'Getting Started',
      featured: true,
    },
    {
      slug: 'another-great-post',
      title: 'Another Great Post',
      description:
        'Here is another article, discussing various topics and insights that will help you grow as a developer.',
      publishedAt: '2025-01-20',
      readTime: 8,
      category: 'Development',
    },
    {
      slug: 'tech-trends-2025',
      title: 'Tech Trends in 2025',
      description:
        'An in-depth look at the technology trends expected in the year 2025, including AI, web development, and emerging frameworks.',
      publishedAt: '2025-02-01',
      readTime: 12,
      category: 'Technology',
      featured: true,
    },
  ];

  // Filter articles based on search term
  const filteredArticles = useMemo(() => {
    if (!searchTerm) return articles;
    return articles.filter(
      (article) =>
        article.title.toLowerCase().includes(searchTerm.toLowerCase()) ||
        article.description.toLowerCase().includes(searchTerm.toLowerCase()) ||
        article.category.toLowerCase().includes(searchTerm.toLowerCase()),
    );
  }, [searchTerm, articles]);

  const featuredArticles = filteredArticles.filter((article) => article.featured);
  const regularArticles = filteredArticles.filter((article) => !article.featured);

  return (
    <div className="bg-gradient-to-br from-slate-50 to-blue-50 min-h-screen">
      <div className="container mx-auto px-4 py-12">
        {/* Header Section */}
        <div className="text-center mb-12">
          <h1 className="text-5xl font-bold bg-gradient-to-r from-blue-600 to-purple-600 bg-clip-text text-transparent mb-4">
            Latest Articles
          </h1>
          <p className="text-xl text-gray-600 max-w-2xl mx-auto mb-8">
            Discover insights, tutorials, and thoughts on modern web development and technology
            trends.
          </p>

          {/* Search Bar */}
          <div className="max-w-md mx-auto relative">
            <input
              type="text"
              placeholder="Search articles..."
              value={searchTerm}
              onChange={(e) => setSearchTerm(e.target.value)}
              className="w-full px-4 py-3 pr-12 rounded-full border-2 border-gray-200 focus:border-blue-500 focus:outline-none focus:ring-2 focus:ring-blue-200 transition-all duration-300"
              aria-label="Search articles"
            />
            <svg
              className="absolute right-4 top-1/2 transform -translate-y-1/2 w-5 h-5 text-gray-400"
              fill="none"
              stroke="currentColor"
              viewBox="0 0 24 24"
            >
              <path
                strokeLinecap="round"
                strokeLinejoin="round"
                strokeWidth={2}
                d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"
              />
            </svg>
          </div>
        </div>

        {/* Featured Articles */}
        {featuredArticles.length > 0 && (
          <section className="mb-16">
            <h2 className="text-3xl font-bold text-gray-800 mb-8 flex items-center">
              <span className="w-2 h-8 bg-gradient-to-b from-yellow-400 to-orange-500 rounded-full mr-3"></span>
              Featured Articles
            </h2>
            <div className="grid grid-cols-1 lg:grid-cols-2 gap-8">
              {featuredArticles.map((article) => (
                <ArticleCard key={article.slug} {...article} featured={true} />
              ))}
            </div>
          </section>
        )}

        {/* Regular Articles */}
        {regularArticles.length > 0 && (
          <section>
            <h2 className="text-3xl font-bold text-gray-800 mb-8">All Articles</h2>
            <div className="grid grid-cols-1 md:grid-cols-2 xl:grid-cols-3 gap-8">
              {regularArticles.map((article) => (
                <ArticleCard key={article.slug} {...article} />
              ))}
            </div>
          </section>
        )}

        {/* No Results Message */}
        {filteredArticles.length === 0 && searchTerm && (
          <div className="text-center py-16">
            <div className="w-24 h-24 mx-auto mb-6 bg-gray-100 rounded-full flex items-center justify-center">
              <svg
                className="w-12 h-12 text-gray-400"
                fill="none"
                stroke="currentColor"
                viewBox="0 0 24 24"
              >
                <path
                  strokeLinecap="round"
                  strokeLinejoin="round"
                  strokeWidth={2}
                  d="M9.172 16.172a4 4 0 015.656 0M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"
                />
              </svg>
            </div>
            <h3 className="text-2xl font-semibold text-gray-800 mb-2">No articles found</h3>
            <p className="text-gray-600 mb-4">
              Try adjusting your search terms or browse all articles.
            </p>
            <button
              onClick={() => setSearchTerm('')}
              className="px-6 py-2 bg-blue-500 text-white rounded-full hover:bg-blue-600 transition-colors duration-300"
            >
              Clear Search
            </button>
          </div>
        )}
      </div>
    </div>
  );
};

export default HomeContent;

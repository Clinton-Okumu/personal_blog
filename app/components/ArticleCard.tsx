import { Star } from 'lucide-react';
import { Link } from 'react-router-dom';

interface ArticleCardProps {
  title: string;
  description: string;
  slug: string;
  publishedAt: string;
  readTime: number;
  category: string;
  featured?: boolean;
}

const ArticleCard = ({
  title,
  description,
  slug,
  publishedAt,
  readTime,
  category,
  featured = false,
}: ArticleCardProps) => {
  const formatDate = (dateString: string) => {
    return new Date(dateString).toLocaleDateString('en-US', {
      month: 'short',
      day: 'numeric',
      year: 'numeric',
    });
  };

  return (
    <article
      className={`
        group relative bg-white rounded-lg border transition-all duration-200 hover:shadow-lg
        ${featured ? 'border-blue-200 shadow-md' : 'border-gray-200 shadow-sm'}
      `}
    >
      {/* Featured Badge - Simpler */}
      {featured && (
        <div className="absolute -top-2 -right-2">
          <div className="bg-blue-500 text-white px-2 py-1 rounded-full text-xs font-medium flex items-center shadow-sm">
            <Star className="w-3 h-3 mr-1" />
            Featured
          </div>
        </div>
      )}

      <div className="p-6">
        {/* Header */}
        <div className="flex items-center justify-between mb-3 text-sm text-gray-500">
          <span className="bg-gray-100 px-2 py-1 rounded text-gray-700 font-medium">
            {category}
          </span>
          <span>{readTime} min read</span>
        </div>

        {/* Title */}
        <h3 className="text-xl font-semibold text-gray-900 mb-3 leading-tight">
          <Link to={`/article/${slug}`} className="hover:text-blue-600 transition-colors">
            {title}
          </Link>
        </h3>

        {/* Description */}
        <p className="text-gray-600 text-sm leading-relaxed mb-4">{description}</p>

        {/* Footer */}
        <div className="flex items-center justify-between pt-3 border-t border-gray-100">
          <time dateTime={publishedAt} className="text-xs text-gray-500">
            {formatDate(publishedAt)}
          </time>

          <Link
            to={`/article/${slug}`}
            className="text-blue-500 hover:text-blue-600 text-sm font-medium transition-colors"
          >
            Read more →
          </Link>
        </div>
      </div>
    </article>
  );
};

export default ArticleCard;

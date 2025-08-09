import { Link } from 'react-router-dom';

const Footer = () => {
  return (
    <footer className="bg-gray-900 text-gray-300 py-6">
      <div className="container mx-auto px-4 flex flex-col md:flex-row justify-between items-center">
        <p className="text-sm">&copy; 2025 Clint Okumu. All rights reserved.</p>
        <nav className="mt-3 md:mt-0 space-x-4">
          <Link to="/" className="hover:underline">
            Home
          </Link>
          <Link to="/blog" className="hover:underline">
            Blog
          </Link>
          <Link to="/contact" className="hover:underline">
            Contact
          </Link>
        </nav>
      </div>
    </footer>
  );
};

export default Footer;

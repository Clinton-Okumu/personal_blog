import React from 'react';
import { Link, useLocation } from 'react-router-dom';

const AdminSidebar: React.FC = () => {
  const location = useLocation();

  const navItems = [
    { name: 'Dashboard', path: '/admin/dashboard' },
    { name: 'Add Article', path: '/admin/add-article' },
    { name: 'Edit Articles', path: '/admin/edit-article' },
  ];

  const isActiveLink = (path: string) => {
    return location.pathname === path;
  };

  return (
    <nav className="w-64 bg-gray-800 text-white p-4 space-y-4 min-h-screen">
      <h2 className="text-2xl font-bold mb-6">Admin Panel</h2>
      <ul>
        {navItems.map((item) => (
          <li key={item.name}>
            <Link
              to={item.path}
              className={`block py-2 px-4 rounded-md transition-colors duration-200 ${
                isActiveLink(item.path)
                  ? 'bg-gray-700 text-blue-400'
                  : 'hover:bg-gray-700 hover:text-blue-300'
              }`}
            >
              {item.name}
            </Link>
          </li>
        ))}
      </ul>
    </nav>
  );
};

export default AdminSidebar;

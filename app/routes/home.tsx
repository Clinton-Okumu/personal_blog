import Footer from '~/components/Footer';
import Navbar from '~/components/Navbar';
import HomeContent from '~/pages/Home';
import ErrorBoundary from '~/components/ErrorBoundary';
import type { Route } from './+types/home';

export function meta({}: Route.MetaArgs) {
  return [
    { title: 'New React Router App' },
    { name: 'description', content: 'Welcome to React Router!' },
  ];
}

export default function Home() {
  return (
    <main className="min-h-screen flex flex-col">
      <Navbar />
      <ErrorBoundary>
        <div className="flex-1">
          <HomeContent />
        </div>
      </ErrorBoundary>
      <Footer />
    </main>
  );
}

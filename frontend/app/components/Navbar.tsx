'use client';

import Link from 'next/link';
import { usePathname } from 'next/navigation';
import { Calendar, PieChart, Settings, Home } from 'lucide-react';

export default function Navbar() {
  const pathname = usePathname();

  const isActive = (path: string) => {
    return pathname === path ? 'text-indigo-600' : 'text-gray-400 hover:text-gray-600';
  };

  return (
    <nav className="fixed bottom-0 left-0 right-0 bg-white/80 backdrop-blur-md border-t border-gray-200 z-50 shadow-[0_-4px_6px_-1px_rgba(0,0,0,0.1)]">
      <div className="container mx-auto px-4 max-w-7xl h-16 flex items-center justify-around md:justify-center md:gap-16">
        <Link
          href="/"
          className={`flex flex-col items-center justify-center w-16 h-full transition-colors ${isActive('/')}`}
        >
          <Home size={24} className="mb-1" />
          <span className="text-[10px] font-medium">ホーム</span>
        </Link>
        <Link
          href="/notes"
          className={`flex flex-col items-center justify-center w-16 h-full transition-colors ${isActive('/notes')}`}
        >
          <Calendar size={24} className="mb-1" />
          <span className="text-[10px] font-medium">記録</span>
        </Link>
        <Link
          href="/report"
          className={`flex flex-col items-center justify-center w-16 h-full transition-colors ${isActive('/report')}`}
        >
          <PieChart size={24} className="mb-1" />
          <span className="text-[10px] font-medium">分析</span>
        </Link>
        <Link
          href="/setting"
          className={`flex flex-col items-center justify-center w-16 h-full transition-colors ${isActive('/setting')}`}
        >
          <Settings size={24} className="mb-1" />
          <span className="text-[10px] font-medium">設定</span>
        </Link>
      </div>
    </nav>
  );
}

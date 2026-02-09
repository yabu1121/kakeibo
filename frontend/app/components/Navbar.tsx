'use client';

import Link from 'next/link';
import { usePathname } from 'next/navigation';
import { Calendar, PieChart, Settings, Home, User } from 'lucide-react';

export default function Navbar() {
  const pathname = usePathname();

  const isActive = (path: string) => {
    return pathname === path ? 'text-indigo-600 bg-indigo-50' : 'text-gray-600 hover:bg-gray-50';
  };

  return (
    <nav className="bg-white shadow-sm border-b border-gray-200 mb-8 sticky top-0 z-50">
      <div className="container mx-auto px-4 max-w-7xl">
        <div className="flex justify-between items-center h-16">
          <Link href="/" className="flex items-center gap-2 font-bold text-xl text-indigo-600">
            <span className="bg-gradient-to-r from-indigo-600 to-purple-600 bg-clip-text text-transparent">
              家計簿アプリ
            </span>
          </Link>

          {/* Desktop Menu */}
          <div className="hidden md:flex gap-4">
            <Link
              href="/"
              className={`flex items-center px-4 py-2 rounded-lg font-medium transition-colors ${isActive('/')}`}
            >
              <Home className="mr-2" size={20} />
              ホーム
            </Link>
            <Link
              href="/notes"
              className={`flex items-center px-4 py-2 rounded-lg font-medium transition-colors ${isActive('/notes')}`}
            >
              <Calendar className="mr-2" size={20} />
              記録
            </Link>
            <Link
              href="/report"
              className={`flex items-center px-4 py-2 rounded-lg font-medium transition-colors ${isActive('/report')}`}
            >
              <PieChart className="mr-2" size={20} />
              分析
            </Link>
            <Link
              href="/setting"
              className={`flex items-center px-4 py-2 rounded-lg font-medium transition-colors ${isActive('/setting')}`}
            >
              <Settings className="mr-2" size={20} />
              設定・サブスク
            </Link>
          </div>

          {/* Mobile Menu (Simplified for now) */}
          <div className="md:hidden flex gap-2">
            <Link href="/notes" className={`p-2 rounded-lg ${isActive('/notes')}`}>
              <Calendar size={24} />
            </Link>
            <Link href="/report" className={`p-2 rounded-lg ${isActive('/report')}`}>
              <PieChart size={24} />
            </Link>
            <Link href="/setting" className={`p-2 rounded-lg ${isActive('/setting')}`}>
              <Settings size={24} />
            </Link>
          </div>
        </div>
      </div>
    </nav>
  );
}

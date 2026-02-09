'use client';

import { useState, useEffect } from 'react';
import { User, Mail, FileText, Lock, LogOut, ChevronRight, Edit2 } from 'lucide-react';
import { apiClient } from '@/lib/api';
import { User as UserType } from '@/types/user';

export default function SettingPage() {
  const [user, setUser] = useState<UserType | null>(null);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const loadUser = async () => {
      try {
        setLoading(true);
        // 仮実装: ユーザー一覧の先頭を取得（ログイン機能がないため）
        const users = await apiClient.getUsers();
        if (users.length > 0) {
          setUser(users[0]);
        }
      } catch (error) {
        console.error('Failed to load user:', error);
      } finally {
        setLoading(false);
      }
    };

    loadUser();
  }, []);

  if (loading) {
    return (
      <div className="flex items-center justify-center min-h-[50vh]">
        <div className="animate-spin rounded-full h-16 w-16 border-t-4 border-b-4 border-indigo-600"></div>
      </div>
    );
  }

  return (
    <div className="pb-8 space-y-6">
      <h1 className="text-2xl font-bold text-gray-800 mb-6 px-2">設定</h1>

      {/* プロフィールカード */}
      <div className="bg-white rounded-2xl shadow-sm border border-gray-100 p-6">
        <div className="flex items-center justify-between mb-6">
          <h2 className="text-lg font-semibold text-gray-800">プロフィール</h2>
          <button className="p-2 text-indigo-600 hover:bg-indigo-50 rounded-full transition-colors">
            <Edit2 size={20} />
          </button>
        </div>

        <div className="flex items-center gap-4 mb-6">
          <div className="w-16 h-16 bg-indigo-100 rounded-full flex items-center justify-center text-indigo-600">
            {user?.icon ? (
              <img src={user.icon} alt={user.name} className="w-full h-full rounded-full object-cover" />
            ) : (
              <User size={32} />
            )}
          </div>
          <div>
            <h3 className="text-xl font-bold text-gray-900">{user?.real_name || user?.name || 'ゲストユーザー'}</h3>
            <p className="text-gray-500 text-sm">@{user?.name || 'guest'}</p>
          </div>
        </div>

        <div className="space-y-4">
          <div className="flex items-center gap-3 text-gray-600">
            <Mail size={20} />
            <span>{user?.email || 'email@example.com'}</span>
          </div>
          <div className="flex items-start gap-3 text-gray-600">
            <FileText size={20} className="mt-1" />
            <p className="text-sm leading-relaxed">
              {user?.profile_memo || 'プロフィールメモは未設定です。'}
            </p>
          </div>
        </div>
      </div>

      {/* アプリ情報・規約 */}
      <div className="bg-white rounded-2xl shadow-sm border border-gray-100 overflow-hidden">
        <div className="p-4 border-b border-gray-100 hover:bg-gray-50 transition-colors cursor-pointer flex items-center justify-between">
          <div className="flex items-center gap-3 text-gray-700">
            <FileText size={20} className="text-gray-400" />
            <span>利用規約</span>
          </div>
          <ChevronRight size={20} className="text-gray-400" />
        </div>
        <div className="p-4 border-b border-gray-100 hover:bg-gray-50 transition-colors cursor-pointer flex items-center justify-between">
          <div className="flex items-center gap-3 text-gray-700">
            <Lock size={20} className="text-gray-400" />
            <span>プライバシーポリシー</span>
          </div>
          <ChevronRight size={20} className="text-gray-400" />
        </div>
        <div className="p-4 hover:bg-gray-50 transition-colors cursor-pointer flex items-center justify-between group">
          <div className="flex items-center gap-3 text-red-600">
            <LogOut size={20} />
            <span className="font-medium">ログアウト</span>
          </div>
          <ChevronRight size={20} className="text-gray-400 group-hover:text-red-400 transition-colors" />
        </div>
      </div>

      <div className="text-center text-gray-400 text-sm mt-8">
        v1.0.0
      </div>
    </div>
  );
}

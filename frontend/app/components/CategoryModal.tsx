'use client';

import { useState } from 'react';
import { X } from 'lucide-react';
import { apiClient } from '@/lib/api';
import { Category } from '@/types/category';

interface CategoryModalProps {
  onClose: () => void;
  onCategoryAdded: (categories: Category[]) => void;
}

export default function CategoryModal({ onClose, onCategoryAdded }: CategoryModalProps) {
  const [newCategoryName, setNewCategoryName] = useState('');

  const handleCreateCategory = async (e: React.FormEvent) => {
    e.preventDefault();
    if (!newCategoryName.trim()) return;
    try {
      await apiClient.createCategory(newCategoryName);
      const categoriesData = await apiClient.getCategories();
      onCategoryAdded(categoriesData);
      setNewCategoryName('');
      onClose();
    } catch (error) {
      console.error('カテゴリーの追加に失敗しました:', error);
      alert('カテゴリーの追加に失敗しました');
    }
  };

  return (
    <div className="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
      <div className="bg-white rounded-xl p-6 w-96 shadow-2xl">
        <div className="flex justify-between items-center mb-4">
          <h3 className="text-xl font-bold text-gray-800">カテゴリー追加</h3>
          <button onClick={onClose} className="text-gray-500 hover:text-gray-700" aria-label="閉じる">
            <X size={24} />
          </button>
        </div>
        <form onSubmit={handleCreateCategory}>
          <div className="mb-4">
            <label htmlFor="new-category-name" className="block text-sm font-medium text-gray-700 mb-2">
              カテゴリー名
            </label>
            <input
              id="new-category-name"
              type="text"
              value={newCategoryName}
              onChange={(e) => setNewCategoryName(e.target.value)}
              className="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-indigo-500 focus:border-transparent"
              placeholder="例: 食費、交通費"
              autoFocus
              required
            />
          </div>
          <div className="flex justify-end gap-3">
            <button
              type="button"
              onClick={onClose}
              className="px-4 py-2 text-gray-700 bg-gray-100 rounded-lg hover:bg-gray-200 transition-colors"
            >
              キャンセル
            </button>
            <button
              type="submit"
              className="px-4 py-2 text-white bg-indigo-600 rounded-lg hover:bg-indigo-700 transition-colors"
            >
              追加
            </button>
          </div>
        </form>
      </div>
    </div>
  );
}

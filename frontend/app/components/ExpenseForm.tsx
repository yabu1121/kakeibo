'use client';

import { useState, useEffect } from 'react';
import { PlusCircle } from 'lucide-react';
import { apiClient } from '@/lib/api';
import { Category } from '@/types/category';

interface ExpenseFormProps {
  categories: Category[];
  onExpenseAdded: () => void;
  onOpenCategoryModal: () => void;
}

export default function ExpenseForm({ categories, onExpenseAdded, onOpenCategoryModal }: ExpenseFormProps) {
  const todayStr = new Date().toISOString().split('T')[0];
  const [newExpense, setNewExpense] = useState({
    user_id: 1,
    amount: 0,
    date: todayStr,
    category_id: categories[0]?.id || 1,
    memo: '',
  });

  useEffect(() => {
    if (categories.length > 0 && newExpense.category_id === 1) {
      setNewExpense(prev => ({ ...prev, category_id: categories[0].id }));
    }
  }, [categories]);

  const handleCreateExpense = async (e: React.FormEvent) => {
    e.preventDefault();
    try {
      await apiClient.createExpense(newExpense);
      onExpenseAdded();
      setNewExpense({
        user_id: 1,
        amount: 0,
        date: todayStr,
        category_id: categories[0]?.id || 1,
        memo: '',
      });
    } catch (error) {
      console.error('消費の追加に失敗しました:', error);
      alert('消費の追加に失敗しました');
    }
  };

  return (
    <form onSubmit={handleCreateExpense} className="mb-8 p-6 bg-gray-50 rounded-xl">
      <div className="grid grid-cols-1 md:grid-cols-3 gap-4">
        <div>
          <label htmlFor="expense-amount" className="block text-sm font-medium text-gray-700 mb-2">金額</label>
          <input
            id="expense-amount"
            type="number"
            value={newExpense.amount}
            onChange={(e) => setNewExpense({ ...newExpense, amount: parseInt(e.target.value) })}
            className="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-indigo-500 focus:border-transparent"
            required
          />
        </div>
        <div>
          <label htmlFor="expense-date" className="block text-sm font-medium text-gray-700 mb-2">日付</label>
          <input
            id="expense-date"
            type="date"
            value={newExpense.date}
            onChange={(e) => setNewExpense({ ...newExpense, date: e.target.value })}
            className="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-indigo-500 focus:border-transparent"
            required
          />
        </div>
        <div>
          <label htmlFor="expense-category" className="block text-sm font-medium text-gray-700 mb-2">カテゴリー</label>
          <div className="flex gap-2">
            <select
              id="expense-category"
              value={newExpense.category_id}
              onChange={(e) => setNewExpense({ ...newExpense, category_id: parseInt(e.target.value) })}
              className="flex-1 px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-indigo-500 focus:border-transparent"
              title="カテゴリー選択"
            >
              {categories.map((cat) => (
                <option key={cat.id} value={cat.id}>{cat.name}</option>
              ))}
            </select>
            <button
              type="button"
              onClick={onOpenCategoryModal}
              className="p-2 bg-indigo-100 text-indigo-600 rounded-lg hover:bg-indigo-200 transition-colors"
              title="カテゴリーを追加"
            >
              <PlusCircle size={20} />
            </button>
          </div>
        </div>
      </div>
      <div className="flex gap-4 mt-4">
        <button type="submit" className="bg-indigo-600 text-white px-6 py-2 rounded-lg hover:bg-indigo-700 transition-colors">追加</button>
      </div>
    </form>
  );
}

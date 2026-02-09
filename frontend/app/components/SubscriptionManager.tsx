'use client';

import { useState, useEffect } from 'react';
import { PlusCircle, CreditCard, X } from 'lucide-react';
import { apiClient } from '@/lib/api';
import { Subscription } from '@/types/subscription';
import { Category } from '@/types/category';

interface SubscriptionManagerProps {
  initialSubscriptions: Subscription[];
  categories: Category[];
  onOpenCategoryModal: () => void;
}

export default function SubscriptionManager({ initialSubscriptions, categories, onOpenCategoryModal }: SubscriptionManagerProps) {
  const [subscriptions, setSubscriptions] = useState<Subscription[]>(initialSubscriptions);
  const [showSubscriptionForm, setShowSubscriptionForm] = useState(false);
  const [newSubscription, setNewSubscription] = useState({
    user_id: 1,
    product_name: '',
    category_id: categories[0]?.id || 1,
    frequency: 'monthly',
  });

  useEffect(() => {
    setSubscriptions(initialSubscriptions);
  }, [initialSubscriptions]);

  useEffect(() => {
    if (categories.length > 0 && newSubscription.category_id === 1) {
      setNewSubscription(prev => ({ ...prev, category_id: categories[0].id }));
    }
  }, [categories]);

  const handleCreateSubscription = async (e: React.FormEvent) => {
    e.preventDefault();
    try {
      await apiClient.createSubscription(newSubscription);
      const updatedSubs = await apiClient.getSubscriptions();
      setSubscriptions(updatedSubs);
      setShowSubscriptionForm(false);
      setNewSubscription({
        user_id: 1,
        product_name: '',
        category_id: categories[0]?.id || 1,
        frequency: 'monthly',
      });
    } catch (error) {
      console.error('サブスクの追加に失敗しました:', error);
      alert('サブスクの追加に失敗しました');
    }
  };

  return (
    <div className="bg-white rounded-2xl shadow-lg p-8">
      <div className="flex justify-between items-center mb-6">
        <h2 className="text-2xl font-bold text-gray-800">サブスク管理</h2>
        <button
          onClick={() => setShowSubscriptionForm(!showSubscriptionForm)}
          className="flex items-center gap-2 bg-indigo-600 text-white px-6 py-3 rounded-xl hover:bg-indigo-700 transition-colors shadow-md"
        >
          {showSubscriptionForm ? <X size={20} /> : <PlusCircle size={20} />}
          {showSubscriptionForm ? 'キャンセル' : '新規追加'}
        </button>
      </div>

      {showSubscriptionForm && (
        <form onSubmit={handleCreateSubscription} className="mb-8 p-6 bg-gray-50 rounded-xl">
          <div className="grid grid-cols-1 md:grid-cols-3 gap-4">
            <div>
              <label htmlFor="sub-name" className="block text-sm font-medium text-gray-700 mb-2">商品名</label>
              <input
                id="sub-name"
                type="text"
                value={newSubscription.product_name}
                onChange={(e) => setNewSubscription({ ...newSubscription, product_name: e.target.value })}
                className="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-indigo-500 focus:border-transparent"
                required
              />
            </div>
            <div>
              <label htmlFor="sub-freq" className="block text-sm font-medium text-gray-700 mb-2">頻度</label>
              <select
                id="sub-freq"
                value={newSubscription.frequency}
                onChange={(e) => setNewSubscription({ ...newSubscription, frequency: e.target.value })}
                className="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-indigo-500 focus:border-transparent"
                title="頻度選択"
              >
                <option value="monthly">月額</option>
                <option value="yearly">年額</option>
              </select>
            </div>
            <div>
              <label htmlFor="sub-category" className="block text-sm font-medium text-gray-700 mb-2">カテゴリー</label>
              <div className="flex gap-2">
                <select
                  id="sub-category"
                  value={newSubscription.category_id}
                  onChange={(e) => setNewSubscription({ ...newSubscription, category_id: parseInt(e.target.value) })}
                  className="flex-1 px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-indigo-500 focus:border-transparent"
                  title="カテゴリー選択"
                >
                  {categories.map((cat) => (
                    <option key={cat.id} value={cat.id}>
                      {cat.name}
                    </option>
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
            <button
              type="submit"
              className="bg-indigo-600 text-white px-6 py-2 rounded-lg hover:bg-indigo-700 transition-colors"
            >
              追加
            </button>
          </div>
        </form>
      )}

      <div className="space-y-4">
        {subscriptions.length === 0 ? (
          <p className="text-center text-gray-500 py-8">まだサブスクがありません</p>
        ) : (
          subscriptions.map((sub) => (
            <div
              key={sub.id}
              className="flex justify-between items-center p-4 bg-gray-50 rounded-xl hover:bg-gray-100 transition-colors"
            >
              <div>
                <p className="font-semibold text-gray-800">{sub.product_name}</p>
                <p className="text-sm text-gray-500">{sub.frequency === 'monthly' ? '月額' : '年額'}</p>
              </div>
              <div className="text-right">
                <span className="inline-block px-3 py-1 bg-purple-100 text-purple-700 rounded-full text-sm font-medium">
                  {sub.category?.name || categories.find((c) => c.id === sub.category_id)?.name || 'カテゴリー不明'}
                </span>
              </div>
            </div>
          ))
        )}
      </div>
    </div>
  );
}

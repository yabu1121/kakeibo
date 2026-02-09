'use client';

import { useState, useEffect } from 'react';
import { PlusCircle } from 'lucide-react';
import { apiClient } from '@/lib/api';
import { Expense } from '@/types/expenses';
import { Category } from '@/types/category';
import ExpenseForm from '../components/ExpenseForm';
import ExpenseList from '../components/ExpenseList';
import CategoryModal from '../components/CategoryModal';
import SubscriptionManager from '../components/SubscriptionManager';

export default function NotesPage() {
  const [expenses, setExpenses] = useState<Expense[]>([]);
  const [categories, setCategories] = useState<Category[]>([]);
  const [loading, setLoading] = useState(true);
  const [showExpenseForm, setShowExpenseForm] = useState(false);
  const [showCategoryModal, setShowCategoryModal] = useState(false);
  const [subscriptions, setSubscriptions] = useState<Subscription[]>([]);
  const [categories, setCategories] = useState<Category[]>([]);
  const [loading, setLoading] = useState(true);
  const [showCategoryModal, setShowCategoryModal] = useState(false);

  useEffect(() => {
    loadData();
  }, []);

  const handleCategoryAdded = (updatedCategories: Category[]) => {
    setCategories(updatedCategories);
  };

  if (loading) {
    return (
      <div className="flex items-center justify-center min-h-[50vh]">
        <div className="animate-spin rounded-full h-16 w-16 border-t-4 border-b-4 border-indigo-600"></div>
      </div>
    );
  }

  useEffect(() => {
    loadData();
  }, []);

  const loadData = async () => {
    try {
      setLoading(true);
      const [expensesData, categoriesData] = await Promise.all([
        apiClient.getExpenses(),
        apiClient.getCategories(),
      ]);
      setExpenses(expensesData);
      setCategories(categoriesData);
    } catch (error) {
      console.error('データの読み込みに失敗しました:', error);
    } finally {
      setLoading(false);
    }
  };

  const handleExpenseAdded = async () => {
    // Reload expenses to get the updated list
    const expensesData = await apiClient.getExpenses();
    setExpenses(expensesData);
    setShowExpenseForm(false);
  };

  const handleCategoryAdded = (updatedCategories: Category[]) => {
    setCategories(updatedCategories);
  };

  if (loading) {
    return (
      <div className="flex items-center justify-center min-h-[50vh]">
        <div className="animate-spin rounded-full h-16 w-16 border-t-4 border-b-4 border-indigo-600"></div>
      </div>
    );
  }

  return (
    <div className="bg-white rounded-2xl shadow-lg p-8">
      <div className="flex justify-between items-center mb-6">
        <h2 className="text-2xl font-bold text-gray-800">消費記録</h2>
        <button
          onClick={() => setShowExpenseForm(!showExpenseForm)}
          className="flex items-center gap-2 bg-indigo-600 text-white px-6 py-3 rounded-xl hover:bg-indigo-700 transition-colors shadow-md"
        >
          <PlusCircle size={20} />
          {showExpenseForm ? 'キャンセル' : '新規追加'}
        </button>
      </div>

      {showExpenseForm && (
        <ExpenseForm
          categories={categories}
          onExpenseAdded={handleExpenseAdded}
          onOpenCategoryModal={() => setShowCategoryModal(true)}
        />
      )}

      <ExpenseList expenses={expenses} categories={categories} />

      {showCategoryModal && (
        <CategoryModal
          onClose={() => setShowCategoryModal(false)}
          onCategoryAdded={handleCategoryAdded}
        />
      )}

      <SubscriptionManager
        initialSubscriptions={subscriptions}
        categories={categories}
        onOpenCategoryModal={() => setShowCategoryModal(true)}
      />

      {showCategoryModal && (
        <CategoryModal
          onClose={() => setShowCategoryModal(false)}
          onCategoryAdded={handleCategoryAdded}
        />
      )}
    </div>
  );
}

'use client';

import { useState, useEffect } from 'react';
import { apiClient } from '@/lib/api';
import { Expense } from '@/types/expenses';
import { Category } from '@/types/category';
import { Subscription } from '@/types/subscription';
import DashboardSummary from './components/DashboardSummary';

export default function Home() {
  const [expenses, setExpenses] = useState<Expense[]>([]);
  const [categories, setCategories] = useState<Category[]>([]);
  const [subscriptions, setSubscriptions] = useState<Subscription[]>([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    loadData();
  }, []);

  const loadData = async () => {
    try {
      setLoading(true);
      const [expensesData, categoriesData, subscriptionsData] = await Promise.all([
        apiClient.getExpenses(),
        apiClient.getCategories(),
        apiClient.getSubscriptions(),
      ]);
      setExpenses(expensesData);
      setCategories(categoriesData);
      setSubscriptions(subscriptionsData);
    } catch (error) {
      console.error('データの読み込みに失敗しました:', error);
    } finally {
      setLoading(false);
    }
  };

  if (loading) {
    return (
      <div className="flex items-center justify-center min-h-[50vh]">
        <div className="animate-spin rounded-full h-16 w-16 border-t-4 border-b-4 border-indigo-600"></div>
      </div>
    );
  }

  return (
    <>
      <DashboardSummary
        expenses={expenses}
        subscriptions={subscriptions}
        categories={categories}
      />
    </>
  );
}

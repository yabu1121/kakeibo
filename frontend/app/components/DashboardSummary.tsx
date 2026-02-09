'use client';

import { TrendingDown, CreditCard, BarChart3 } from 'lucide-react';
import { Expense } from '@/types/expenses';
import { Subscription } from '@/types/subscription';
import { Category } from '@/types/category';

interface DashboardSummaryProps {
  expenses: Expense[];
  subscriptions: Subscription[];
  categories: Category[];
}

export default function DashboardSummary({ expenses, subscriptions, categories }: DashboardSummaryProps) {
  const totalExpenses = expenses.reduce((sum, exp) => sum + exp.amount, 0);

  return (
    <div className="grid grid-cols-1 md:grid-cols-3 gap-6 mb-12">
      <div className="bg-white rounded-2xl shadow-lg p-6 border border-gray-100 hover:shadow-xl transition-shadow">
        <div className="flex items-center justify-between mb-4">
          <h3 className="text-gray-600 font-medium">総支出 (全期間)</h3>
          <TrendingDown className="text-red-500" size={24} />
        </div>
        <p className="text-3xl font-bold text-gray-800">¥{totalExpenses.toLocaleString()}</p>
      </div>

      <div className="bg-white rounded-2xl shadow-lg p-6 border border-gray-100 hover:shadow-xl transition-shadow">
        <div className="flex items-center justify-between mb-4">
          <h3 className="text-gray-600 font-medium">サブスク数</h3>
          <CreditCard className="text-blue-500" size={24} />
        </div>
        <p className="text-3xl font-bold text-gray-800">{subscriptions.length}件</p>
      </div>

      <div className="bg-white rounded-2xl shadow-lg p-6 border border-gray-100 hover:shadow-xl transition-shadow">
        <div className="flex items-center justify-between mb-4">
          <h3 className="text-gray-600 font-medium">カテゴリー</h3>
          <BarChart3 className="text-green-500" size={24} />
        </div>
        <p className="text-3xl font-bold text-gray-800">{categories.length}種類</p>
      </div>
    </div>
  );
}

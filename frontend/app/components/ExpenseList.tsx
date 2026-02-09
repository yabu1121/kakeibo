'use client';

import { Expense } from '@/types/expenses';
import { Category } from '@/types/category';

interface ExpenseListProps {
  expenses: Expense[];
  categories: Category[];
}

export default function ExpenseList({ expenses, categories }: ExpenseListProps) {
  if (expenses.length === 0) {
    return <p className="text-center text-gray-500 py-8">まだ消費記録がありません</p>;
  }

  return (
    <div className="space-y-4">
      {expenses.map((expense) => (
        <div key={expense.id} className="flex justify-between items-center p-4 bg-gray-50 rounded-xl hover:bg-gray-100 transition-colors">
          <div>
            <p className="font-semibold text-gray-800">¥{expense.amount.toLocaleString()}</p>
            <p className="text-sm text-gray-500">{new Date(expense.date).toLocaleDateString('ja-JP')}</p>
          </div>
          <div className="text-right">
            <span className="inline-block px-3 py-1 bg-indigo-100 text-indigo-700 rounded-full text-sm font-medium">
              {expense.category?.name || categories.find((c) => c.id === expense.category_id)?.name || 'カテゴリー不明'}
            </span>
          </div>
        </div>
      ))}
    </div>
  );
}

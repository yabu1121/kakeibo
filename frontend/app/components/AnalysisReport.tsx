'use client';

import { useState, useMemo, useEffect } from 'react';
import { ChevronLeft, ChevronRight } from 'lucide-react';
import { apiClient } from '@/lib/api';
import { Expense } from '@/types/expenses';
import { Category } from '@/types/category';

interface AnalysisReportProps {
  initialCategories: Category[];
}

export default function AnalysisReport({ initialCategories }: AnalysisReportProps) {
  const [analysisType, setAnalysisType] = useState<'day' | 'week' | 'month' | 'year' | 'category'>('month');
  const today = new Date();
  const todayStr = today.toISOString().split('T')[0];

  const [selectedDate, setSelectedDate] = useState(todayStr);
  const [selectedWeekStart, setSelectedWeekStart] = useState(getMonday(new Date(todayStr)).toISOString().split('T')[0]);
  const [selectedYear, setSelectedYear] = useState(today.getFullYear());
  const [selectedMonth, setSelectedMonth] = useState(today.getMonth() + 1);
  const [selectedCategoryId, setSelectedCategoryId] = useState<number>(0);

  const [analysisExpenses, setAnalysisExpenses] = useState<Expense[]>([]);
  const [categories, setCategories] = useState<Category[]>(initialCategories);

  useEffect(() => {
    // If props categories are updated or empty, maybe re-fetch or sync?
    // For now assuming passed categories are up to date or we fetch if empty
    if (initialCategories.length > 0) {
      setCategories(initialCategories);
      if (selectedCategoryId === 0) setSelectedCategoryId(initialCategories[0].id);
    }
  }, [initialCategories]);

  useEffect(() => {
    loadAnalysisData();
  }, [analysisType, selectedDate, selectedWeekStart, selectedYear, selectedMonth, selectedCategoryId]);

  const loadAnalysisData = async () => {
    try {
      let data: Expense[] = [];
      switch (analysisType) {
        case 'day':
          data = await apiClient.getExpensesByDay(selectedDate);
          break;
        case 'week':
          const startDate = new Date(selectedWeekStart);
          const endDate = new Date(startDate);
          endDate.setDate(endDate.getDate() + 6);
          data = await apiClient.getExpensesByWeek(
            startDate.toISOString().split('T')[0],
            endDate.toISOString().split('T')[0]
          );
          break;
        case 'month':
          data = await apiClient.getExpensesByMonth(selectedYear, selectedMonth);
          break;
        case 'year':
          data = await apiClient.getExpensesByYear(selectedYear);
          break;
        case 'category':
          if (selectedCategoryId !== 0) {
            data = await apiClient.getExpensesByCategory(selectedCategoryId);
          }
          break;
      }
      setAnalysisExpenses(data);
    } catch (error) {
      console.error('分析データの読み込みに失敗しました:', error);
      setAnalysisExpenses([]);
    }
  };

  const moveWeek = (amount: number) => {
    const current = new Date(selectedWeekStart);
    current.setDate(current.getDate() + (amount * 7));
    setSelectedWeekStart(current.toISOString().split('T')[0]);
  };

  const analysisTotal = analysisExpenses.reduce((sum, exp) => sum + exp.amount, 0);

  const categoryBreakdown = useMemo(() => {
    const breakdown = new Map<number, number>();
    analysisExpenses.forEach(exp => {
      const current = breakdown.get(exp.category_id) || 0;
      breakdown.set(exp.category_id, current + exp.amount);
    });

    const total = Array.from(breakdown.values()).reduce((sum, val) => sum + val, 0);

    return Array.from(breakdown.entries()).map(([catId, amount]) => {
      const category = categories.find(c => c.id === catId);
      return {
        name: category?.name || '不明',
        amount,
        percentage: total > 0 ? (amount / total) * 100 : 0,
        color: getRandomColor(catId),
      };
    }).sort((a, b) => b.amount - a.amount);
  }, [analysisExpenses, categories]);

  return (
    <div className="bg-white rounded-2xl shadow-lg p-8">
      <div className="mb-8">
        <h2 className="text-2xl font-bold text-gray-800 mb-4">分析レポート</h2>
        {/* 分析タイプ切り替え */}
        <div className="flex flex-wrap gap-2 mb-6">
          {[
            { type: 'day', label: '日次' },
            { type: 'week', label: '週次' },
            { type: 'month', label: '月次' },
            { type: 'year', label: '年次' },
            { type: 'category', label: 'カテゴリー' },
          ].map((item) => (
            <button
              key={item.type}
              onClick={() => setAnalysisType(item.type as any)}
              className={`px-4 py-2 rounded-full text-sm font-medium transition-colors ${analysisType === item.type
                ? 'bg-indigo-600 text-white'
                : 'bg-gray-100 text-gray-600 hover:bg-gray-200'
                }`}
            >
              {item.label}
            </button>
          ))}
        </div>

        {/* 条件選択エリア */}
        <div className="flex items-center gap-4 bg-gray-50 p-4 rounded-xl overflow-x-auto">
          {analysisType === 'day' && (
            <input
              type="date"
              value={selectedDate}
              onChange={(e) => setSelectedDate(e.target.value)}
              className="px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-indigo-500"
              aria-label="日付選択"
            />
          )}
          {analysisType === 'week' && (
            <div className="flex items-center gap-2">
              <button onClick={() => moveWeek(-1)} className="p-2 hover:bg-gray-200 rounded-full" aria-label="前の週">
                <ChevronLeft size={20} />
              </button>
              <input
                type="date"
                value={selectedWeekStart}
                onChange={(e) => setSelectedWeekStart(getMonday(new Date(e.target.value)).toISOString().split('T')[0])}
                className="px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-indigo-500"
                aria-label="週選択"
              />
              <span className="text-gray-500 text-sm whitespace-nowrap">の週</span>
              <button onClick={() => moveWeek(1)} className="p-2 hover:bg-gray-200 rounded-full" aria-label="次の週">
                <ChevronRight size={20} />
              </button>
            </div>
          )}
          {analysisType === 'month' && (
            <div className="flex gap-2">
              <select
                value={selectedYear}
                onChange={(e) => setSelectedYear(parseInt(e.target.value))}
                className="px-4 py-2 border border-gray-300 rounded-lg"
                aria-label="年選択"
              >
                {Array.from({ length: 5 }, (_, i) => today.getFullYear() - 2 + i).map(year => (
                  <option key={year} value={year}>{year}年</option>
                ))}
              </select>
              <select
                value={selectedMonth}
                onChange={(e) => setSelectedMonth(parseInt(e.target.value))}
                className="px-4 py-2 border border-gray-300 rounded-lg"
                aria-label="月選択"
              >
                {Array.from({ length: 12 }, (_, i) => i + 1).map(month => (
                  <option key={month} value={month}>{month}月</option>
                ))}
              </select>
            </div>
          )}
          {analysisType === 'year' && (
            <select
              value={selectedYear}
              onChange={(e) => setSelectedYear(parseInt(e.target.value))}
              className="px-4 py-2 border border-gray-300 rounded-lg"
              aria-label="年選択"
            >
              {Array.from({ length: 5 }, (_, i) => today.getFullYear() - 2 + i).map(year => (
                <option key={year} value={year}>{year}年</option>
              ))}
            </select>
          )}
          {analysisType === 'category' && (
            <select
              value={selectedCategoryId}
              onChange={(e) => setSelectedCategoryId(parseInt(e.target.value))}
              className="px-4 py-2 border border-gray-300 rounded-lg"
              aria-label="カテゴリー選択"
            >
              {categories.map(cat => (
                <option key={cat.id} value={cat.id}>{cat.name}</option>
              ))}
            </select>
          )}
        </div>
      </div>

      <div className="mb-8 p-6 bg-indigo-50 rounded-xl text-center">
        <h3 className="text-gray-600 mb-2">総支出</h3>
        <p className="text-4xl font-bold text-indigo-700">¥{analysisTotal.toLocaleString()}</p>
      </div>

      {/* カテゴリー別内訳 (カテゴリー分析以外で表示) */}
      {analysisType !== 'category' && (
        <div className="space-y-6 mb-12">
          <h3 className="text-xl font-bold text-gray-800 mb-4">カテゴリー別内訳</h3>
          {categoryBreakdown.length === 0 ? (
            <p className="text-center text-gray-500 py-4">データがありません</p>
          ) : (
            categoryBreakdown.map((item, index) => (
              <div key={index} className="space-y-2">
                <div className="flex justify-between text-sm font-medium">
                  <span className="text-gray-700">{item.name}</span>
                  <span className="text-gray-900">¥{item.amount.toLocaleString()} ({item.percentage.toFixed(1)}%)</span>
                </div>
                <div className="w-full bg-gray-200 rounded-full h-2.5 overflow-hidden">
                  <div className="h-2.5 rounded-full" style={{ width: `${item.percentage}%`, backgroundColor: item.color }}></div>
                </div>
              </div>
            ))
          )}
        </div>
      )}

      {/* 履歴リスト */}
      <div>
        <h3 className="text-xl font-bold text-gray-800 mb-4">履歴 ({analysisExpenses.length}件)</h3>
        <div className="overflow-x-auto">
          <table className="min-w-full text-left text-sm">
            <thead>
              <tr className="border-b border-gray-200">
                <th className="py-3 px-4 font-semibold text-gray-600">日付</th>
                <th className="py-3 px-4 font-semibold text-gray-600">カテゴリー</th>
                <th className="py-3 px-4 font-semibold text-gray-600 text-right">金額</th>
              </tr>
            </thead>
            <tbody>
              {analysisExpenses.map((expense) => (
                <tr key={expense.id} className="border-b border-gray-100 hover:bg-gray-50">
                  <td className="py-3 px-4 text-gray-600">
                    {new Date(expense.date).toLocaleDateString('ja-JP')}
                  </td>
                  <td className="py-3 px-4">
                    <span className="inline-block px-2 py-1 rounded text-xs font-medium bg-gray-100 text-gray-800">
                      {expense.category?.name || '不明'}
                    </span>
                  </td>
                  <td className="py-3 px-4 text-right font-medium text-gray-900">
                    ¥{expense.amount.toLocaleString()}
                  </td>
                </tr>
              ))}
            </tbody>
          </table>
        </div>
      </div>
    </div>
  );
}

function getMonday(d: Date) {
  d = new Date(d);
  const day = d.getDay();
  const diff = d.getDate() - day + (day === 0 ? -6 : 1);
  return new Date(d.setDate(diff));
}

function getRandomColor(id: number) {
  const colors = [
    '#4F46E5', '#EC4899', '#10B981', '#F59E0B',
    '#8B5CF6', '#EF4444', '#3B82F6', '#6366F1'
  ];
  return colors[id % colors.length];
}

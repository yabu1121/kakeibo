'use client';

import { useState, useEffect } from 'react';
import { apiClient } from '@/lib/api';
import { Category } from '@/types/category';
import AnalysisReport from '../components/AnalysisReport';

export default function ReportPage() {
  const [categories, setCategories] = useState<Category[]>([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    loadCategories();
  }, []);

  const loadCategories = async () => {
    try {
      setLoading(true);
      const categoriesData = await apiClient.getCategories();
      setCategories(categoriesData);
    } catch (error) {
      console.error('カテゴリーの読み込みに失敗しました:', error);
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

  return <AnalysisReport initialCategories={categories} />;
}

import { Category } from "@/types/category";
import { Expense } from "@/types/expenses";
import { Subscription } from "@/types/subscription";
import { PublicUtility } from "@/types/public_utility";
import { User } from "@/types/user";

const API_BASE_URL = process.env.NEXT_PUBLIC_API_URL || 'http://localhost:8080/api';

// API Client
class ApiClient {
  private baseUrl: string;

  constructor(baseUrl: string) {
    this.baseUrl = baseUrl;
  }

  private async request<T>(endpoint: string, options?: RequestInit): Promise<T> {
    const response = await fetch(`${this.baseUrl}${endpoint}`, {
      ...options,
      headers: {
        'Content-Type': 'application/json',
        ...options?.headers,
      },
    });

    if (!response.ok) {
      throw new Error(`API Error: ${response.statusText}`);
    }

    return response.json();
  }

  // User APIs
  async getUsers(): Promise<User[]> {
    return this.request<User[]>('/users');
  }

  async createUser(user: Omit<User, 'id' | 'created_at' | 'updated_at'>): Promise<User> {
    return this.request<User>('/users', {
      method: 'POST',
      body: JSON.stringify(user),
    });
  }

  // Category APIs
  async getCategories(): Promise<Category[]> {
    return this.request<Category[]>('/categories');
  }

  async createCategory(name: string): Promise<Category> {
    return this.request<Category>('/categories', {
      method: 'POST',
      body: JSON.stringify({ name }),
    });
  }

  // Expense APIs
  async getExpenses(): Promise<Expense[]> {
    return this.request<Expense[]>('/expenses');
  }

  async createExpense(expense: Omit<Expense, 'id' | 'created_at' | 'updated_at' | 'category'>): Promise<Expense> {
    return this.request<Expense>('/expenses', {
      method: 'POST',
      body: JSON.stringify(expense),
    });
  }

  async getExpensesByDay(date: string): Promise<Expense[]> {
    return this.request<Expense[]>(`/expenses/day?date=${date}`);
  }

  async getExpensesByWeek(startDate: string, endDate: string): Promise<Expense[]> {
    return this.request<Expense[]>(`/expenses/week?start_date=${startDate}&end_date=${endDate}`);
  }

  async getExpensesByMonth(year: number, month: number): Promise<Expense[]> {
    return this.request<Expense[]>(`/expenses/month?year=${year}&month=${month}`);
  }

  async getExpensesByYear(year: number): Promise<Expense[]> {
    return this.request<Expense[]>(`/expenses/year?year=${year}`);
  }

  async getExpensesByCategory(categoryId: number): Promise<Expense[]> {
    return this.request<Expense[]>(`/expenses/category?category_id=${categoryId}`);
  }

  // Subscription APIs
  async getSubscriptions(): Promise<Subscription[]> {
    return this.request<Subscription[]>('/subscriptions');
  }

  async createSubscription(subscription: Omit<Subscription, 'id' | 'created_at' | 'updated_at' | 'category'>): Promise<Subscription> {
    return this.request<Subscription>('/subscriptions', {
      method: 'POST',
      body: JSON.stringify(subscription),
    });
  }

  // Public utility APIs
  async getPublicUtilities(params?: { user_id?: number; category_id?: number }): Promise<PublicUtility[]> {
    const query = new URLSearchParams();
    if (params?.user_id !== undefined) query.set('user_id', String(params.user_id));
    if (params?.category_id !== undefined) query.set('category_id', String(params.category_id));
    const qs = query.toString();
    const endpoint = qs ? `/public-utilities?${qs}` : '/public-utilities';
    return this.request<PublicUtility[]>(endpoint);
  }

  async createPublicUtility(utility: Omit<PublicUtility, 'id' | 'created_at' | 'updated_at' | 'category'>): Promise<PublicUtility> {
    return this.request<PublicUtility>('/public-utilities', {
      method: 'POST',
      body: JSON.stringify(utility),
    });
  }

  async updatePublicUtility(id: number, utility: Partial<Omit<PublicUtility, 'id' | 'created_at' | 'updated_at' | 'category'>>): Promise<PublicUtility> {
    return this.request<PublicUtility>(`/public-utilities/${id}`, {
      method: 'PUT',
      body: JSON.stringify(utility),
    });
  }

  async deletePublicUtility(id: number): Promise<void> {
    await this.request<void>(`/public-utilities/${id}`, {
      method: 'DELETE',
    });
  }
}

export const apiClient = new ApiClient(API_BASE_URL);

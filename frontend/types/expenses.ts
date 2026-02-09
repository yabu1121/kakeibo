import { Category } from "./category";

export type Expense = {
  id: number;
  user_id: number;
  amount: number;
  date: string;
  memo: string;
  category_id: number;
  category?: Category;
  created_at: string;
  updated_at: string;
}

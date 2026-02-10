import { Category } from "./category";

export type PublicUtility = {
  id: number;
  user_id: number;
  amount: number;
  date: string;
  category_id: number;
  memo: string;
  category?: Category;
  created_at: string;
  updated_at: string;
}
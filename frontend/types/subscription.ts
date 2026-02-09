import { Category } from "./category";

export type Subscription = {
  id: number;
  user_id: number;
  product_name: string;
  category_id: number;
  amount: number;
  frequency: string;
  category?: Category;
  created_at: string;
  updated_at: string;
}

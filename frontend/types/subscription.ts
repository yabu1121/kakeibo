import { Category } from "./category";

export type Subscription = {
  id: number;
  user_id: number;
  product_name: string;
  category_id: number;
  frequency: string;
  category?: Category;
  created_at: string;
  updated_at: string;
}

import Address from "./address";
import StatusLog from "./statuslog";

interface Order {
  id?: string;
  user_id: string;
  address: Address;
  date_time: Date;
  price: number;
  number_of_rooms: number;
  number_of_bathrooms: number;
  pollution: number;
  area: number;
  comment: string;
  status?: string;
  status_logs?: StatusLog[];
  services: string[];
  required_workers: number;
  workers?: string[];
  created_at?: Date;
  updated_at?: Date;
}

export default Order;
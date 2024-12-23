import Consumable from "./consumable";

interface Service {
    id: string;
    name: string;
    price: number;
    workers_quantity: number;
    description: string;
    consumables: Consumable[];
    created_at: string;
    updated_at?: string
  }

export default Service;
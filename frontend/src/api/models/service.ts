import Consumable from "./consumable";

interface Service {
    id: string;
    name: string;
    price: number;
    workers_quantity: number;
    description: string;
    consumables: Consumable[];
    created_at: Date;
    updated_at?: Date
  }

export default Service;
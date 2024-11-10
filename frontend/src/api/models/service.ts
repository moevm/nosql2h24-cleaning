import Consumable from "./consumable";

interface Service {
    id: string;
    name: string;
    price: number;
    workersQuantity: number;
    description: string;
    consumables: Consumable[];
    createdAt: Date;
    updatedAt?: Date
  }

export default Service;
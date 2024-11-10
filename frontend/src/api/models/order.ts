import Address from "./address";
import StatusLog from "./statuslog";

interface Order {
  id: string;
  userId: string;
  address: Address;
  dateTime: Date;
  price: number;
  numberOfRooms: number;
  numberOfBathrooms: number;
  pollution: number;
  area: number;
  comment: string;
  statusLogs: StatusLog[];
  services: string[];
  requiredWorkers: number;
  workers: string[];
  createdAt: Date;
  updatedAt?: Date;
}

export default Order;
interface Address {
  id: string;
  city: string;
  street: string;
  building: string;
  entrance: string;
  floor: string;
  door_number: string;
  created_at: Date;
  updated_at?: Date;
}

export default Address;
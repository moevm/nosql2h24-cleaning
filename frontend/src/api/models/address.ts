interface Address {
  id: string;
  city: string;
  street: string;
  building: string;
  entrance: string;
  floor: string;
  doorNumber: string;
  createdAt: Date;
  updatedAt?: Date;
}

export default Address;
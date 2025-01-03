import Address from "./address";

export type UserLoginData = {
  email: string;
  password: string;
}

export interface UserInfo {
  name: string;
  surname: string;
  patronymic?: string;
  phone_number: string;
}

export type User = {
  id: string;
  email: string;
  password: string;
  name: string;
  surname: string;
  patronymic: string;
  phone_number: string;
  addresses: Address[];
  user_type: string;
  created_at: Date;
  updated_at: Date;
  orders_count: number;
}

export type UserRegisterData = {
  email: string,
  name: string,
  password: string,
  patronymic: string,
  phone_number: string,
  surname: string
}

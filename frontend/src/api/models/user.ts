import Address from "./address";

export interface UserCredentials {
  email: string;
  password?: string;
}

export interface UserInfo {
  name: string;
  surname: string;
  patronymic?: string;
  phoneNumber: string;
}

export interface User {
  id?: string;
  email: string;
  password?: string;
  name: string;
  surname: string;
  patronymic?: string;
  phoneNumber: string;
  addresses?: Address[];
  userType: string;
  createdAt?: Date;
  updatedAt?: Date;
}

export interface NewUser {
  email: string;
  password?: string;
  name: string;
  surname: string;
  patronymic?: string;
  phoneNumber: string;
}

import axios from 'axios';
import Address from "./models/address";
import { User } from "./models/user"
import { 
  getListOfUsersPath,
  createWorkerUserPath,
  getUserInfoPath,
  updateUserPath,
  deleteUserPath,
  getClientAddressesPath,
  createNewAddressPath,
  getClientAddressPath,
  updateClientAddressPath,
  deleteClientAddressPath
} from "./endpoint";

const baseURL = 'http://localhost:8081/api';

export function getListOfUsers(): Promise<User[]> {
  return axios.get<User[]>(baseURL + getListOfUsersPath)
  .then((response) => {
    return response.data
  })
  .catch((error) => {
    console.error('Error HTTP:', error);
    throw error;
  })
}

export function createWorkerUser(newWorker: User): Promise<{ id: string }> {
  return axios.post(baseURL + createWorkerUserPath, newWorker)
  .then((response) => {
    return response.data
  })
  .catch((error) => {
    console.error('Error HTTP:', error);
    throw error;
  })
}

export function getUserInfo(id: number): Promise<User> {
  return axios.get<User>(baseURL + getUserInfoPath(id))
  .then((response) => {
    return response.data
  })
  .catch((error) => {
    console.error('Error HTTP:', error);
    throw error;
  })
}

export function updateUser(id: number, newUserData: User): Promise<any> {
  return axios.put<User>(baseURL + updateUserPath(id), newUserData)
  .then((response) => {
    return response.data
  })
  .catch((error) => {
    console.error('Error HTTP:', error);
    throw error;
  })
}

export function deleteUser(id: number): Promise<any> {
  return axios.delete(baseURL + deleteUserPath(id))
  .then((response) => {
    return response.data
  })
  .catch((error) => {
    console.error('Error HTTP:', error);
    throw error;
  })
}

export function getClientAddress(id: number, address_id: number): Promise<Address> {
  return axios.get<Address>(baseURL + getClientAddressPath(id, address_id))
  .then((response) => {
    return response.data
  })
  .catch((error) => {
    console.error('Error HTTP:', error);
    throw error;
  })
}

export function updateClientAddress(id: number, address_id: number, newAddress: Address): Promise<any> {
  return axios.put<Address>(baseURL + updateClientAddressPath(id, address_id), newAddress)
  .then((response) => {
    return response.data
  })
  .catch((error) => {
    console.error('Error HTTP:', error);
    throw error;
  })
}

export function deleteClientAddress(id: number, address_id: number): Promise<any> {
  return axios.delete(baseURL + deleteClientAddressPath(id, address_id))
  .then((response) => {
    return response.data
  })
  .catch((error) => {
    console.error('Error HTTP:', error);
    throw error;
  })
}

export function getClientAddresses(id: number): Promise<Address[]> {
  return axios.get<Address[]>(baseURL + getClientAddressesPath(id))
  .then((response) => {
    return response.data
  })
  .catch((error) => {
    console.error('Error HTTP:', error);
    throw error;
  })
}

export function createNewAddress(id: number, addressData: Address): Promise<any> {
  return axios.post(baseURL + createNewAddressPath(id), addressData)
  .then((response) => {
    return response.data
  })
  .catch((error) => {
    console.error('Error HTTP:', error);
    throw error;
  })
}
import axios from 'axios'
import {
  UserRegisterData,
  UserLoginData,
  User
} from './models/user'
import Address from './models/address'
import {
  getRegisterPath,
  getLoginPath,
  getUserInfoPath,
  getUsersPath,
  updateUserPath,
  deleteUserPath,
  getClientAddressesPath,
  createNewAddressPath,
  getClientAddressPath,
  updateClientAddressPath,
  createWorkerUserPath,
  deleteClientAddressPath,
  uploadDumpsPath,
  exportDumpsPath,
  getAllOrdersPath,
  getAllServicesPath,
  updateOrdersPath,
  createOrderPath,
  getServicesPath
} from './endpoint'
import Order from './models/order'
import Service from './models/service'
import { FilterOrder } from './models/filterOrder'
import { FilterWorkers } from './models/filterWorkers'

const baseURL = 'http://localhost:8080/api'

export async function postRegister (data: UserRegisterData): Promise<void> {
  return axios.post(baseURL + getRegisterPath, data)
  .then((response) => {
    return Promise.resolve(response.data)
  })
  .catch((error) => {
    return Promise.reject(error)
  })
}

export async function postLogin (data: UserLoginData): Promise<User> {
  return axios.post(baseURL + getLoginPath, data)
  .then((response) => {
    return Promise.resolve(response.data)
  })
  .catch((error) => {
    return Promise.reject(error)
  })
}

export async function getUserInfo(id: string) : Promise<User> {
  return axios.get(baseURL + getUserInfoPath(id))
  .then((response) => {
    return Promise.resolve(response.data)
  })
  .catch((error) => {
    return Promise.reject(error)
  })
}

export async function getUsers(type: string) : Promise<User[]> {
  return axios.get(baseURL + getUsersPath, { params: { type: type } })
  .then((response) => {
    return Promise.resolve(response.data)
  })
  .catch((error) => {
    return Promise.reject(error)
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

export function updateUser(id: string, newUserData: User): Promise<any> {
  return axios.put<User>(baseURL + updateUserPath(id), newUserData)
  .then((response) => {
    return response.data
  })
  .catch((error) => {
    console.error('Error HTTP:', error);
    throw error;
  })
}

export function deleteUser(id: string): Promise<any> {
  return axios.delete(baseURL + deleteUserPath(id))
  .then((response) => {
    return response.data
  })
  .catch((error) => {
    console.error('Error HTTP:', error);
    throw error;
  })
}

export async function getClientAddresses(id: string): Promise<Address[]> {
  return axios.get(baseURL + getClientAddressesPath(id))
  .then((response) => {
    return Promise.resolve(response.data)
  })
  .catch((error) => {
    return Promise.reject(error)
  })
}

export async function createNewAddress(id: string, addressData: Address): Promise<Address> {
  return axios.post<Address>(baseURL + createNewAddressPath(id), addressData)
  .then((response) => {
    return Promise.resolve(response.data)
  })
  .catch((error) => {
    return Promise.reject(error)
  })
}

export function getClientAddress(id: string, address_id: string): Promise<Address> {
  return axios.get<Address>(baseURL + getClientAddressPath(id, address_id))
  .then((response) => {
    return response.data
  })
  .catch((error) => {
    console.error('Error HTTP:', error);
    throw error;
  })
}

export function updateClientAddress(id: string, address_id: string, newAddress: Address): Promise<any> {
  return axios.put<Address>(baseURL + updateClientAddressPath(id, address_id), newAddress)
  .then((response) => {
    return response.data
  })
  .catch((error) => {
    console.error('Error HTTP:', error);
    throw error;
  })
}

export function deleteClientAddress(id: string, address_id: string): Promise<any> {
  return axios.delete(baseURL + deleteClientAddressPath(id, address_id))
  .then((response) => {
    return response.data
  })
  .catch((error) => {
    console.error('Error HTTP:', error);
    throw error;
  })
}

export function exportDumps(): Promise<void> {
  return axios.get(baseURL + exportDumpsPath, {
    headers: {
      'Content-Type': 'application/json',
    },
    responseType: 'blob'
  })
  .then((response) => {
    const url = window.URL.createObjectURL(new Blob([response.data]));
    const link = document.createElement('a');
    link.href = url;

    link.setAttribute('download', 'dump.json');
    document.body.appendChild(link);
    link.click();
    document.body.removeChild(link);

    return Promise.resolve();
  })
  .catch((error) => {
    console.error('Ошибка при экспорте данных:', error);
    return Promise.reject(error);
  });
}

export async function uploadDumps(file: File): Promise<void> {
  const formData = new FormData();
  formData.append('file', file);

  return axios.post(baseURL + uploadDumpsPath, formData, {
    headers: {
      'Content-Type': 'multipart/form-data',
    },
  })
  .then((response) => {
    console.log('Файл успешно загружен:', response.data);
    return Promise.resolve(response.data);
  })
  .catch((error) => {
    console.error('Ошибка при загрузке файла:', error);
    return Promise.reject(error);
  });
}

export async function filterWorkers(filterWorkers: FilterWorkers): Promise<User[]> {
  return axios.get(baseURL + getUsersPath, { params: { type: 'WORKER', ...filterWorkers} })
  .then((response) => {
    return Promise.resolve(response.data)
  })
  .catch((error) => {
    return Promise.reject(error)
  })
}

export async function getAllOrders(id: string | null): Promise<Order[]> {
  const url = `${baseURL}${getAllOrdersPath}?user_id=${id}`;
  return axios.get(url)
  .then((response) => {
    return Promise.resolve(response.data)
  })
  .catch((error) => {
    return Promise.reject(error);
  })
}

export async function updateOrder(newOrder: Order): Promise<Order> {
  return axios.put<Order>(baseURL + updateOrdersPath(newOrder.id!), newOrder)
  .then((response) => {
    return response.data
  })
  .catch((error) => {
    console.error('Error update order:', error);
    throw error;
  })
  
}

export async function getAllServices(): Promise<Service[]> {
  return axios.get(baseURL + getAllServicesPath)
  .then((response) => {
    return Promise.resolve(response.data)
  })
  .catch((error) => {
    return Promise.reject(error);
  })
}

export async function filterOrder(filterData: FilterOrder): Promise<Order[]> {
  return axios.get(baseURL + getAllOrdersPath, {
    params: filterData, 
    paramsSerializer: {
      indexes: null
    }
  })
  .then((response) => {
    return Promise.resolve(response.data)
  })
  .catch((error) => {
    return Promise.reject(error);
  })
}

export async function createOrder(newOrder: Order): Promise<Order> {
  return axios.post<Order>(baseURL + createOrderPath, newOrder)
  .then((response) => {
    return Promise.resolve(response.data)
  })
  .catch((error) => {
    return Promise.reject(error)
  })
}

export async function getOrderById(id: string): Promise<Order> {
  return axios.get(baseURL + updateOrdersPath(id))
  .then((response) => {
    return Promise.resolve(response.data)
  })
  .catch((error) => {
    return Promise.reject(error)
  })
}

export async function getService (id: string): Promise<Service> {
  return axios.get(baseURL + getServicesPath(id))
  .then((response) => {
    return Promise.resolve(response.data)
  })
  .catch((error) => {
    return Promise.reject(error)
  })
}
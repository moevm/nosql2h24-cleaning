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
  deleteClientAddressPath
} from './endpoint'

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
    console.log(response.data)
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
  try {
    const response = await fetch(baseURL+getClientAddressesPath(id), {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json',
      },
    });
  
    if (!response.ok) {
      throw new Error(`Error HTTP: ${response.status}`);
    }
  
    const data = (await response.json()) as Address[];
    return data;
  } catch (error) {
    console.error('Error request:', error);
    throw error;
  }
}

export async function createNewAddress(id: string, addressData: Address): Promise<Address> {
  try {
    const response = await fetch(baseURL+createNewAddressPath(id), {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(addressData),
    });
  
    if (!response.ok) {
      throw new Error(`Error HTTP: ${response.status}`);
    }
  
    const data = await response.json();
    return data.id; 
    
  } catch (error) {
    console.error('Error request:', error);
    throw error;
  }
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

export async function filterWorkers(name: string, surname: string): Promise<User[]> {
  return axios.get(baseURL + getUsersPath, { params: { type: 'WORKER', name: name, surname: surname } })
  .then((response) => {
    return Promise.resolve(response.data)
  })
  .catch((error) => {
    return Promise.reject(error)
  })
}

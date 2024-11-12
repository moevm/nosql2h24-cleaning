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
  getClientAddressesPath,
  createNewAddressPath 
} from './endpoint'

const baseURL = 'http://localhost:8081/api'

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

// export async function getUser(id: string): Promise<void> {
//   return axios.get(baseURL + `/v1/users/${id}`)
// }

export async function getClientAddresses(id: number): Promise<Address[]> {
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

export async function createNewAddress(id: number, addressData: Address): Promise<Address> {
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
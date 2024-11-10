import Address from "./models/address";
import { getClientAddressesPath, createNewAddressPath } from "./endpoint";

const baseURL = 'http://localhost:8081/api/v1';

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
export const getClientAddressesPath = (id: number) => `/users/${id}/addresses`;
export const createNewAddressPath = (id: number) => `/users/${id}/addresses`
export const getClientAddressPath = (id: number, address_id: number) => `/users/${id}/addresses/${address_id}`
export const updateClientAddressPath = (id: number, address_id: number) => `/users/${id}/addresses/${address_id}`
export const deleteClientAddressPath = (id: number, address_id: number) => `/users/${id}/addresses/${address_id}`

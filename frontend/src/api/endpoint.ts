export const getListOfUsersPath = '/v1/users'
export const createWorkerUserPath = '/v1/users'
export const getUserInfoPath = (id: string) => `/v1/users/${id}`
export const updateUserPath = (id: string) => `/v1/users/${id}`
export const deleteUserPath = (id: string) => `/v1/users/${id}`
export const getClientAddressesPath = (id: string) => `/v1/users/${id}/addresses`;
export const createNewAddressPath = (id: string) => `/v1/users/${id}/addresses`
export const getClientAddressPath = (id: string, address_id: string) => `/v1/users/${id}/addresses/${address_id}`
export const updateClientAddressPath = (id: string, address_id: string) => `/v1/users/${id}/addresses/${address_id}`
export const deleteClientAddressPath = (id: string, address_id: string) => `/v1/users/${id}/addresses/${address_id}`

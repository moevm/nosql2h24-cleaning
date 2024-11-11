export const getListOfUsersPath = '/v1/users'
export const createWorkerUserPath = '/v1/users'
export const getUserInfoPath = (id: number) => `/v1/users/${id}`
export const updateUserPath = (id: number) => `/v1/users/${id}`
export const deleteUserPath = (id: number) => `/v1/users/${id}`
export const getClientAddressesPath = (id: number) => `/v1/users/${id}/addresses`;
export const createNewAddressPath = (id: number) => `/v1/users/${id}/addresses`
export const getClientAddressPath = (id: number, address_id: number) => `/v1/users/${id}/addresses/${address_id}`
export const updateClientAddressPath = (id: number, address_id: number) => `/v1/users/${id}/addresses/${address_id}`
export const deleteClientAddressPath = (id: number, address_id: number) => `/v1/users/${id}/addresses/${address_id}`

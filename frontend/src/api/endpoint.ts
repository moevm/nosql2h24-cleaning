export const getLoginPath = '/auth/login'
export const getLogoutPath = '/auth/logout'
export const getRegisterPath = '/auth/register'
export const getUsersPath = '/v1/users'

export const getClientAddressesPath = (id: number) => `/v1/users/${id}/addresses`
export const createNewAddressPath = (id: number) => `/v1/users/${id}/addresses`
export const getClientAddressPath = (id: number, address_id: number) => `/v1/users/${id}/addresses/${address_id}`
export const updateClientAddressPath = (id: number, address_id: number) => `/v1/users/${id}/addresses/${address_id}`
export const deleteClientAddressPath = (id: number, address_id: number) => `/v1/users/${id}/addresses/${address_id}`
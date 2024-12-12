export const getLoginPath = '/auth/login'
export const getLogoutPath = '/auth/logout'
export const getRegisterPath = '/auth/register'
export const getUsersPath = '/v1/users'
export const createWorkerUserPath = '/v1/users'
export const updateUserPath = (id: string) => `/v1/users/${id}`
export const deleteUserPath = (id: string) => `/v1/users/${id}`
export const getUserInfoPath = (id: string) => `/v1/users/${id}`
export const getClientAddressesPath = (id: string) => `/v1/users/${id}/addresses`
export const createNewAddressPath = (id: string) => `/v1/users/${id}/addresses`
export const getClientAddressPath = (id: string, address_id: string) => `/v1/users/${id}/addresses/${address_id}`
export const updateClientAddressPath = (id: string, address_id: string) => `/v1/users/${id}/addresses/${address_id}`
export const deleteClientAddressPath = (id: string, address_id: string) => `/v1/users/${id}/addresses/${address_id}`
export const uploadDumpsPath = '/v1/dumps'
export const exportDumpsPath = '/v1/dumps'
import service from '@/utils/request'

// ==================== 客户API ====================
export const createAccount = (data) => {
  return service({
    url: '/ffProduct/createAccount',
    method: 'post',
    data
  })
}

export const updateAccount = (data) => {
  return service({
    url: '/ffProduct/updateAccount',
    method: 'put',
    data
  })
}

export const deleteAccount = (data) => {
  return service({
    url: '/ffProduct/deleteAccount',
    method: 'delete',
    data
  })
}

export const getAccountList = (params) => {
  return service({
    url: '/ffProduct/getAccountList',
    method: 'get',
    params
  })
}

export const getAccountOptions = () => {
  return service({
    url: '/ffProduct/getAccountOptions',
    method: 'get'
  })
}

// ==================== 物料API ====================
export const createMaterial = (data) => {
  return service({
    url: '/ffProduct/createMaterial',
    method: 'post',
    data
  })
}

export const updateMaterial = (data) => {
  return service({
    url: '/ffProduct/updateMaterial',
    method: 'put',
    data
  })
}

export const deleteMaterial = (data) => {
  return service({
    url: '/ffProduct/deleteMaterial',
    method: 'delete',
    data
  })
}

export const getMaterialList = (params) => {
  return service({
    url: '/ffProduct/getMaterialList',
    method: 'get',
    params
  })
}

// ==================== 工序API ====================
export const createOperation = (data) => {
  return service({
    url: '/ffProduct/createOperation',
    method: 'post',
    data
  })
}

export const updateOperation = (data) => {
  return service({
    url: '/ffProduct/updateOperation',
    method: 'put',
    data
  })
}

export const deleteOperation = (data) => {
  return service({
    url: '/ffProduct/deleteOperation',
    method: 'delete',
    data
  })
}

export const getOperationList = (params) => {
  return service({
    url: '/ffProduct/getOperationList',
    method: 'get',
    params
  })
}

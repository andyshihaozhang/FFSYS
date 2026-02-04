import service from '@/utils/request'

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

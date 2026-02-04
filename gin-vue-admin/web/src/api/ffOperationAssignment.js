import service from '@/utils/request'

export const createOperationAssignment = (data) => {
  return service({
    url: '/ffProduct/createOperationAssignment',
    method: 'post',
    data
  })
}

export const updateOperationAssignment = (data) => {
  return service({
    url: '/ffProduct/updateOperationAssignment',
    method: 'put',
    data
  })
}

export const deleteOperationAssignment = (data) => {
  return service({
    url: '/ffProduct/deleteOperationAssignment',
    method: 'delete',
    data
  })
}

export const getOperationAssignmentList = (params) => {
  return service({
    url: '/ffProduct/getOperationAssignmentList',
    method: 'get',
    params
  })
}

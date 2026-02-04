import service from '@/utils/request'

export const getEmployeeList = (params) => {
  return service({
    url: '/ffProduct/getEmployeeList',
    method: 'get',
    params
  })
}

export const createEmployee = (data) => {
  return service({
    url: '/ffProduct/createEmployee',
    method: 'post',
    data
  })
}

export const updateEmployee = (data) => {
  return service({
    url: '/ffProduct/updateEmployee',
    method: 'put',
    data
  })
}

export const deleteEmployee = (data) => {
  return service({
    url: '/ffProduct/deleteEmployee',
    method: 'delete',
    data
  })
}

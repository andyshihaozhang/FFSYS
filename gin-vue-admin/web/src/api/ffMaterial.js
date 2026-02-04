import service from '@/utils/request'

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

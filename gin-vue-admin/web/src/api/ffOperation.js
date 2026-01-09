import service from '@/utils/request'

/**
 * 创建工序
 * @param {Object} data 工序信息
 * @param {number} data.productId 产品ID
 * @param {string} data.operationName 工序名称
 * @param {number} data.referencePrice 参考单价
 * @returns {Promise} 响应数据
 */
export const createOperation = (data) => {
  return service({
    url: '/ffproduction/operation',
    method: 'post',
    data
  })
}

/**
 * 更新工序信息
 * @param {Object} data 工序信息
 * @returns {Promise} 响应数据
 */
export const updateOperation = (data) => {
  return service({
    url: '/ffproduction/operation',
    method: 'put',
    data
  })
}

/**
 * 删除工序
 * @param {Object} data 包含工序ID
 * @returns {Promise} 响应数据
 */
export const deleteOperation = (data) => {
  return service({
    url: '/ffproduction/operation',
    method: 'delete',
    data
  })
}

/**
 * 批量删除工序
 * @param {Object} data 包含工序ID数组
 * @returns {Promise} 响应数据
 */
export const deleteOperationByIds = (data) => {
  return service({
    url: '/ffproduction/operationByIds',
    method: 'delete',
    data
  })
}

/**
 * 获取单个工序信息
 * @param {Object} params 查询参数
 * @returns {Promise} 响应数据
 */
export const getOperation = (params) => {
  return service({
    url: '/ffproduction/operation',
    method: 'get',
    params
  })
}

/**
 * 获取工序列表
 * @param {Object} params 查询参数
 * @param {number} params.page 页码
 * @param {number} params.pageSize 每页数量
 * @param {number} params.productId 产品ID
 * @returns {Promise} 工序列表数据
 */
export const getOperationList = (params) => {
  return service({
    url: '/ffproduction/operationList',
    method: 'get',
    params
  })
}

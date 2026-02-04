import service from '@/utils/request'

/**
 * 创建产品
 * @param {Object} data 产品信息
 * @returns {Promise} 响应数据
 */
export const createProduct = (data) => {
  return service({
    url: '/ffProduct/createProduct',
    method: 'post',
    data
  })
}

/**
 * 更新产品信息
 * @param {Object} data 产品信息
 * @returns {Promise} 响应数据
 */
export const updateProduct = (data) => {
  return service({
    url: '/ffProduct/updateProduct',
    method: 'put',
    data
  })
}

/**
 * 删除产品
 * @param {Object} data 包含产品ID
 * @returns {Promise} 响应数据
 */
export const deleteProduct = (data) => {
  return service({
    url: '/ffProduct/deleteProduct',
    method: 'delete',
    data
  })
}

/**
 * 批量删除产品
 * @param {Object} data 包含产品ID数组
 * @returns {Promise} 响应数据
 */
export const deleteProductByIds = (data) => {
  return service({
    url: '/ffProduct/deleteProductByIds',
    method: 'delete',
    data
  })
}

/**
 * 获取单个产品信息
 * @param {Object} params 查询参数
 * @returns {Promise} 响应数据
 */
export const getProduct = (params) => {
  return service({
    url: '/ffProduct/getProduct',
    method: 'get',
    params
  })
}

/**
 * 获取产品列表
 * @param {Object} params 查询参数
 * @param {number} params.page 页码
 * @param {number} params.pageSize 每页数量
 * @param {string} params.productCode 产品款号
 * @param {string} params.productName 产品名称
 * @param {string} params.productColor 产品颜色
 * @param {number} params.accountId 客户ID
 * @param {string} params.productFlag 生产标志
 * @returns {Promise} 产品列表数据
 */
export const getProductList = (params) => {
  return service({
    url: '/ffProduct/getProductList',
    method: 'get',
    params
  })
}
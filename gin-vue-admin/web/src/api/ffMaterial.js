import service from '@/utils/request'

/**
 * 创建物料
 * @param {Object} data 物料信息
 * @param {number} data.productId 产品ID
 * @param {string} data.materialName 材料名称
 * @param {string} data.materialCode 材料编码
 * @param {number} data.materialUsage 材料用量
 * @param {string} data.materialUom 材料单位
 * @param {number} data.materialPrice 材料单价
 * @param {string} data.materialSpecification 材料规格
 * @param {string} data.materialType 材料类型
 * @param {number} data.supplierId 供应商ID
 * @returns {Promise} 响应数据
 */
export const createMaterial = (data) => {
  return service({
    url: '/ffproduction/material',
    method: 'post',
    data
  })
}

/**
 * 更新物料信息
 * @param {Object} data 物料信息
 * @returns {Promise} 响应数据
 */
export const updateMaterial = (data) => {
  return service({
    url: '/ffproduction/material',
    method: 'put',
    data
  })
}

/**
 * 删除物料
 * @param {Object} data 包含物料ID
 * @returns {Promise} 响应数据
 */
export const deleteMaterial = (data) => {
  return service({
    url: '/ffproduction/material',
    method: 'delete',
    data
  })
}

/**
 * 批量删除物料
 * @param {Object} data 包含物料ID数组
 * @returns {Promise} 响应数据
 */
export const deleteMaterialByIds = (data) => {
  return service({
    url: '/ffproduction/materialByIds',
    method: 'delete',
    data
  })
}

/**
 * 获取单个物料信息
 * @param {Object} params 查询参数
 * @returns {Promise} 响应数据
 */
export const getMaterial = (params) => {
  return service({
    url: '/ffproduction/material',
    method: 'get',
    params
  })
}

/**
 * 获取物料列表
 * @param {Object} params 查询参数
 * @param {number} params.page 页码
 * @param {number} params.pageSize 每页数量
 * @param {number} params.productId 产品ID
 * @param {string} params.materialType 材料类型
 * @returns {Promise} 物料列表数据
 */
export const getMaterialList = (params) => {
  return service({
    url: '/ffproduction/materialList',
    method: 'get',
    params
  })
}

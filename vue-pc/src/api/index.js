import request from '@/utils/request'

export function getHome(params) {
    return request({
        url: '/index',
        method: 'get',
        params: params
    })
}
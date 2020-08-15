import Api from '@/services/Api'

export default {
  index (account) {
    return Api().get('orders', {
      params: account
    })
  },
  createOrder (order) {
    return Api().post('order', order)
  },
}
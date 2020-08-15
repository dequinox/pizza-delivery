import Api from '@/services/Api'

export default {
  index (account) {
    return Api().post('orders', account)
  },
  createOrder (order) {
    return Api().post('order', order)
  },
}
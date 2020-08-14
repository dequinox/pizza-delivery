import Api from '@/services/Api'

export default {
  register (credentials) {
    return Api().post('account/create', credentials, { useCredentails: true })
  },
  login (credentials) {
    return Api().post('account/login', credentials, { useCredentails: true })
  }
}
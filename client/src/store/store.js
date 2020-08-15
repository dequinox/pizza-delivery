import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

export default new Vuex.Store({
  strict: true,
  plugins: [
  ],
  state: {
    user: null,
    isUserLoggedIn: false,
    orders: []
  },
  mutations: {
    setUser (state, user) {
      state.user = user
      state.isUserLoggedIn = !!(user)
      state.orders = []
    },
    addPizza (state, pizza){
        state.orders.push(pizza)
    },
    removePizza(state, pizzaId){
        state.orders.splice(pizzaId, 1)
    },
    clearOrder(state){
        state.orders = []
    }
  },
  actions: {
    setUser ({commit}, user) {
      commit('setUser', user)
    },
    addPizza({commit}, pizza){
        commit('addPizza', pizza)
    },
    removePizza({commit}, pizzaId){
        commit('removePizza', pizzaId)
    },
    clearOrder({commit})
    {
        commit('clearOrder')
    }
  }
})
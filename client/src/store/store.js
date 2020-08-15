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
    orders: [],
    cost: 0,
  },
  mutations: {
    setUser (state, user) {
      state.user = user
      state.isUserLoggedIn = !!(user)
      state.orders = []
    },
    addPizza (state, pizza){
        state.cost = state.cost + pizza.cost
        state.orders.push(pizza)
    },
    removePizza(state, pizzaId){
        state.cost = state.cost - state.orders[pizzaId].cost
        state.orders.splice(pizzaId, 1)
    },
    setOrder(state, orders){
        state.orders = orders
        state.cost = 0
        state.orders.forEach(element => {
            state.cost = state.cost + element.cost
        });
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
    setOrder({commit}, orders)
    {
        commit('setOrder', orders)
    }
  }
})
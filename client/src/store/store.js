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
        var exists = false
        var index = 0
        var i, len
        for (i = 0, len = state.orders.length; i < len; i++){
          if (state.orders[i].id == pizza.id){
            exists = true
            index = i
            break;
          }
        }
        console.log(pizza)
        console.log(index, exists)
        if (exists) {
          state.orders[index].amount = state.orders[index].amount + 1
        }
        else {
            state.orders.push({
              id: pizza.id,
              amount: 1,
              title: pizza.title,
              cost: pizza.cost
          })
        }
    },
    removePizza(state, pizzaId){
        state.cost = state.cost - state.orders[pizzaId].cost * state.orders[pizzaId].amount
        state.orders.splice(pizzaId, 1)
    },
    setOrder(state, orders){
        state.orders = orders
        state.cost = 0
        state.orders.forEach(element => {
            state.cost = state.cost + element.cost * element.amount
        });
    },
    increaseAmountPizza(state, pizzaId){
      state.orders[pizzaId].amount = state.orders[pizzaId].amount + 1
      state.cost = state.cost + state.orders[pizzaId].cost
    },
    reduceAmountPizza(state, pizzaId){
      if (state.orders[pizzaId].amount > 1){
        state.orders[pizzaId].amount = state.orders[pizzaId].amount - 1
        state.cost = state.cost - state.orders[pizzaId].cost
      }
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
    },
    reduceAmountPizza({commit}, pizzaId){
      commit('reduceAmountPizza', pizzaId)
    },
    increaseAmountPizza({commit}, pizzaId){
      commit('increaseAmountPizza', pizzaId)
    },
  }
})
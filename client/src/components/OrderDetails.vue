<template>
    <div v-if="orders.length > 0">
        <b-card
            no-body
            style="max-width: 20rem;"
            header="Orders"
            header-tag="header"
            class="text-left"
        >
            <template v-slot:header>
                <h6 class="mb-0">Orders {{ currency }}{{ getCostRound($store.state.costInCurrency) }} + {{ currency }}10 Delivery</h6>
                <label for="currency">Currency:</label>
                    <select v-model="currency">
                    <option>$</option>
                    <option>€</option>
                </select>
            </template>
            <b-list-group flush>
                <div v-for="(item, index) in orders" :key="index">
                    <b-list-group-item>
                        <div>
                            <button class="btn" @click="reduceAmount(index)">
                                <i class="fas fa-arrow-left"></i>
                            </button>
                            {{item.amount}}
                            <button class="btn" @click="increaseAmount(index)">
                                <i class="fas fa-arrow-right" @click="increaseAmount(index)"></i>
                            </button>
                        </div>
                        <div class="item-title text-left">{{item.title}}</div>
                        <div class="close-button"><button class="btn" @click="remove(index)">✖</button></div>
                    </b-list-group-item>
                </div>
            </b-list-group>
            <!--<template v-slot:footer >
                <b-button href="#" variant="primary" @click="order()">Order</b-button>
            </template>
            -->
        </b-card>
    </div>
</template>

<script>
import {mapState} from 'vuex'

export default {
  data () {
    return {
      currency: '$'
    }
  },
  watch: {
      currency: function(val) {
          this.$store.dispatch('changeCurrency', val)
          console.log(val)
      }
  },
  computed: {
    ...mapState([
      'orders'
    ])
  },
  methods: {
      remove(id) {
        this.$store.dispatch('removePizza', id) 
      },
      order() {
        this.$router.push({
            name: 'order'
        })
      },
      reduceAmount(id) {
          this.$store.dispatch('reduceAmountPizza', id)
      },
      increaseAmount(id) {
          this.$store.dispatch('increaseAmountPizza', id)
      },
      getCostRound(val) {
          return Math.round(val)
      }
  }
}
</script>

<style scoped>
.btn {
    width: 20%;
    margin: 0px;
    padding: 0px;
}
.item-title{
    vertical-align: middle;
    float:left;
    width: 80%;
}
</style>
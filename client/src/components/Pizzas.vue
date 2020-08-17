<template>
    <div>
        <b-card-group deck class="card-group">
            <div v-for="pizza in pizzas" :key="pizza.id">
                <b-card
                    v-bind:title="pizza.title"
                    v-bind:img-src="pizza.imageUrl"
                    img-alt="Image"
                    img-top
                    class="text-left"
                    style="max-width: 15rem;"
                >
                    <b-card-text class="text-left">
                        {{pizza.description}}
                    </b-card-text>

                    <template v-slot:footer>
                        <b-button href="#" variant="primary" @click="addPizza(pizza)">Add ${{pizza.cost}}</b-button>
                    </template>
                </b-card>
            </div>
        </b-card-group>
    </div> 
</template>

<script>
import {mapState} from 'vuex'
import PizzasService from '@/services/PizzasService'
import OrderDetails from './OrderDetails'

export default {
  components: {
      OrderDetails
  },
  computed: {
    ...mapState([
      'isUserLoggedIn'
    ])
  },
  data () {
    return {
      pizzas: null
    }
  },
  async mounted () {
    this.pizzas = (await PizzasService.index()).data
  },
  methods: {
      addPizza(pizza) {
        console.log(pizza)
        this.$store.dispatch('addPizza', pizza) 
      },
    }
}
</script>

<style scoped>
.card {
    height: 500px;
}
.card-group {
    margin: 0 auto; /* Added */
    float: none; /* Added */
    margin-bottom: 10px; /* Added */
    margin-top: 10px;
    width: 70%;
}

.order-detail {
    width: 20%;
    position: fixed;
    top: 70px;
    right: 30px;
}
</style>
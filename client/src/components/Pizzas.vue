<template>
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
                    <b-button href="#" variant="primary">Order</b-button>
                </template>
            </b-card>
        </div>
    </b-card-group>
    
</template>

<script>
import {mapState} from 'vuex'
import PizzasService from '@/services/PizzasService'

export default {
  components: {
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
</style>
<template>
    <div v-if="orders.length > 0">
        <b-card
            no-body
            style="max-width: 20rem;"
            header="Orders"
            header-tag="header"
        >
            <b-list-group flush>
                <div v-for="(item, index) in orders" :key="index">
                    <b-list-group-item>
                        <div class="item-title text-left">{{item.title}}</div>
                        <div class="close-button"><button class="btn" @click="remove(index)">✖</button></div>
                    </b-list-group-item>
                </div>
            </b-list-group>
            <template v-slot:footer>
                <b-button href="#" variant="primary" @click="order()">Order</b-button>
            </template>

        </b-card>
    </div>
</template>

<script>
import {mapState} from 'vuex'

export default {
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